package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"tampayang-backend/app/models/entity"
	"tampayang-backend/core/database"
)

func GetLastSequenceForYear(year int) (int, error) {
	db := database.GetConnectionDB()

	// Query untuk mencari nomor urut (sequence) tertinggi di tahun berjalan.
	// Contoh: 'TMP-2025-000001' -> kita ambil bagian '000001' dan ubah jadi angka.
	query := `
        SELECT MAX(CAST(SUBSTRING_INDEX(report_number, '-', -1) AS UNSIGNED)) 
        FROM reports 
        WHERE report_number LIKE ?`

	// Parameter LIKE, contoh: 'TMP-2025-%'
	likeParam := fmt.Sprintf("TMP-%d-%%", year)

	var lastSequence sql.NullInt64 // Gunakan NullInt64 untuk menangani kasus jika belum ada laporan di tahun ini (hasilnya NULL).

	err := db.QueryRow(query, likeParam).Scan(&lastSequence)
	if err != nil && err != sql.ErrNoRows {
		// Jika ada error selain karena tidak ada baris, kembalikan error.
		return 0, fmt.Errorf("gagal query sequence terakhir: %w", err)
	}

	// Jika lastSequence valid (ada isinya), kembalikan nilainya.
	if lastSequence.Valid {
		return int(lastSequence.Int64), nil
	}

	// Jika tidak ada laporan di tahun ini (hasilnya NULL), kembalikan 0.
	return 0, nil
}

func InsertNewReport(ctx context.Context, data *entity.Report) error {
	db := database.GetConnectionDB()
	sqlQuery := `
		INSERT INTO reports (
			report_id, report_number, reporter_name, reporter_phone, reporter_email,
			infrastructure_category_id, damage_type_id, province_id, regency_id, 
			district_id, village_id, location_detail, description, 
			urgency_level, status, latitude, longitude, created_at
		) VALUES (
			?, ?, ?, ?, ?,
			?, ?, ?, ?,
			?, ?, ?, ?,
			?, ?, ?, ?, ?
		)`

	result, err := db.ExecContext(ctx, sqlQuery,
		data.ReportId,
		data.ReportNumber,
		data.ReporterName,
		data.ReporterPhone,
		data.ReporterEmail,
		data.InfrastructureCategoryId,
		data.DamageTypeID,
		data.ProviceID,
		data.RegencyID,
		data.DistrictID,
		data.VillageID,
		data.LocationDetails,
		data.Description,
		data.UrgencyLevel,
		data.Status,
		data.Latitude,
		data.Longitude,
		data.CreatedAt,
	)

	// Hentikan penggunaan panic! Ganti dengan return error.
	if err != nil {
		return fmt.Errorf("gagal menjalankan query insert: %w", err)
	}

	// Verifikasi apakah ada baris yang berhasil ditambahkan.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tidak ada baris yang ditambahkan ke database")
	}

	// Berhasil, tidak ada error yang dikembalikan.
	return nil
}

func InsertReportPhoto(ctx context.Context, photo *entity.ReportPhoto) error {
	db := database.GetConnectionDB()
	defer db.Close()

	query := `
		INSERT INTO report_photos 
		(report_photo_id, report_id, filename, original_filename, file_path, file_size, mime_type, is_main) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.ExecContext(ctx, query,
		photo.ReportPhotoID,
		photo.ReportID,
		photo.Filename,
		photo.OriginalFilename,
		photo.FilePath,
		photo.FileSize,
		photo.MimeType,
		photo.IsMain,
	)

	if err != nil {
		log.Printf("Error inserting report photo metadata to DB: %v", err)
		return errors.New("gagal menyimpan metadata foto ke database")
	}

	return nil
}

func GetUrgentlyReport() []entity.UrgencyReportRequest {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var reports []entity.UrgencyReportRequest

	sqlQuery := `
		SELECT 
			r.report_number,
			d.name,
			v.village_name,
			r.urgency_level
		FROM reports r
		INNER JOIN damage_types d ON d.damage_type_id = r.damage_type_id
		INNER JOIN villages v ON v.village_id = r.village_id
		WHERE r.status NOT IN ('selesai', 'batal')
		ORDER BY r.urgency_level DESC,r.report_number DESC
		LIMIT 10
		;
	`

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - GetUrgentlyReport, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var urgentlyReport entity.UrgencyReportRequest
		err := result.Scan(
			&urgentlyReport.ReportNumber,
			&urgentlyReport.DamageTypeName,
			&urgentlyReport.VillageName,
			&urgentlyReport.UrgencyLevel)
		reports = append(reports, urgentlyReport)
		if err != nil {
			panic("models - GetUrgentlyReport, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetUrgentlyReport, result.Err " + err.Error())
	}
	return reports
}

func GetManageReport(keyword, year, infCategory, status string, page int, limit int) ([]entity.ManageReport, int64, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var category []entity.ManageReport
	var total int64
	args := []interface{}{}

	query := `
	SELECT 
		r.report_number,
		DATE_FORMAT(r.created_at, '%m-%d') AS created_at,
		r.reporter_name,
		i.name,
		v.village_name,
		r.status,
		COUNT(*) OVER() as total_rows
	FROM reports r
	INNER JOIN infrastructure_categories i ON i.infrastructure_category_id = r.infrastructure_category_id
	INNER JOIN villages v ON v.village_id = r.village_id
	WHERE 1=1
	`
	if keyword != "" {
		query += ` AND (
				LOWER(r.report_number) LIKE LOWER(?) OR
				LOWER(r.reporter_name) LIKE LOWER(?) OR
				LOWER(i.name) LIKE LOWER(?) OR
				LOWER(v.village_name) LIKE LOWER(?)
			)`
		searchPattern := "%" + keyword + "%"
		args = append(args, searchPattern, searchPattern, searchPattern, searchPattern)
	}

	if year != "" {
		query += " AND YEAR(r.created_at) = ?"
		args = append(args, year)
	}

	if infCategory != "" {
		query += " AND i.infrastructure_category_id = ?"
		args = append(args, infCategory)
	}

	if status != "" {
		query += " AND LOWER(r.status) = LOWER(?)"
		args = append(args, status)
	}

	query += " ORDER BY r.report_number ASC,r.created_at DESC, r.status ASC LIMIT ? OFFSET ?"
	args = append(args, limit, (page-1)*limit)

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat entity.ManageReport

		err := rows.Scan(
			&cat.ReportNumber,
			&cat.CreatedAt,
			&cat.ReporterName,
			&cat.InfrastructureCategoryName,
			&cat.VillageName,
			&cat.Status,
			&total,
		)
		if err != nil {
			return nil, 0, err
		}

		category = append(category, cat)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return category, total, nil
}

func GetReportPhotos(reportNumber string) ([]entity.ReportPhoto, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	query := `
	SELECT 
		rp.report_photo_id,
		rp.report_id,
		rp.filename,
		rp.original_filename,
		rp.file_path,
		rp.file_size,
		rp.mime_type,
		rp.is_main,
		rp.uploaded_at
	FROM report_photos rp
	JOIN reports r ON r.report_id = rp.report_id
	WHERE r.report_number = ?
	`

	rows, err := db.QueryContext(ctx, query, reportNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []entity.ReportPhoto
	for rows.Next() {
		var photo entity.ReportPhoto
		err := rows.Scan(
			&photo.ReportPhotoID,
			&photo.ReportID,
			&photo.Filename,
			&photo.OriginalFilename,
			&photo.FilePath,
			&photo.FileSize,
			&photo.MimeType,
			&photo.IsMain,
			&photo.UploadedAt,
		)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func GetDetailReport(reportNumber string) (entity.DetailReport, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var detailReport entity.DetailReport

	query := `
	SELECT 
		r.report_id,
		r.report_number,
		r.created_at,
		r.reporter_name,
		r.reporter_phone,
		i.name AS category,
		d.name AS type,
		v.village_name,
		s.district_name,
		r.status,
		r.latitude,
		r.longitude,
		r.urgency_level,
		r.description
	FROM reports r
	JOIN infrastructure_categories i ON i.infrastructure_category_id = r.infrastructure_category_id
	JOIN damage_types d ON d.damage_type_id = r.damage_type_id
	JOIN districts s ON s.district_id = r.district_id
	JOIN villages v ON v.village_id = r.village_id
	WHERE r.report_number = ?
	`
	err := db.QueryRowContext(ctx, query, reportNumber).Scan(
		&detailReport.ReportID,
		&detailReport.ReportNumber,
		&detailReport.CreatedAt,
		&detailReport.ReporterName,
		&detailReport.ReporterPhone,
		&detailReport.InfrastructureCategoryName,
		&detailReport.DamageTypeName,
		&detailReport.VillageName,
		&detailReport.DistrictName,
		&detailReport.Status,
		&detailReport.Latitude,
		&detailReport.Longitude,
		&detailReport.UrgencyLevel,
		&detailReport.Description,
	)
	if err != nil {
		return detailReport, err
	}

	photos, err := GetReportPhotos(reportNumber)
	if err != nil && err != sql.ErrNoRows {
		return detailReport, err
	}
	if photos != nil {
		detailReport.Photos = photos
	}
	return detailReport, nil
}

func UpdateReport(reportNumber string, data entity.UpdateReport) (entity.UpdateReport, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	query := `
        UPDATE reports 
        SET 
			status = ?,
            pic_name = ?,
            pic_contact = ?,
            admin_notes = ?,
            completion_notes = ?,
            estimated_completion = ?,
            completed_at = ?,
            updated_at = ?
        WHERE report_number = ?
    `

	_, err := db.ExecContext(ctx, query,
		data.Status,
		data.PicName,
		data.PicPhone,
		data.AdminNotes,
		data.CompletionNotes,
		data.EstimatedCompletionDate,
		data.ComletedAt,
		data.UpdatedAt,
		reportNumber)

	if err != nil {
		return data, err
	}
	return data, nil
}
