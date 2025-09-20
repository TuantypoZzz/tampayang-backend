package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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

	if err != nil {
		return fmt.Errorf("gagal menjalankan query insert: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tidak ada baris yang ditambahkan ke database")
	}

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

func GetCheckStatus(reportNumber string) (entity.CheckStatus, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var checkStatus entity.CheckStatus

	query := `
	SELECT 
		r.report_number,
		r.created_at,
		r.reporter_name,
		i.name AS category,
		v.village_name,
		s.district_name,
		r.status,
		r.admin_notes
	FROM reports r
	JOIN infrastructure_categories i ON i.infrastructure_category_id = r.infrastructure_category_id
	JOIN damage_types d ON d.damage_type_id = r.damage_type_id
	JOIN districts s ON s.district_id = r.district_id
	JOIN villages v ON v.village_id = r.village_id
	WHERE r.report_number = ?
	`
	err := db.QueryRowContext(ctx, query, reportNumber).Scan(
		&checkStatus.ReportNumber,
		&checkStatus.CreatedAt,
		&checkStatus.ReporterName,
		&checkStatus.InfrastructureCategoryName,
		&checkStatus.VillageName,
		&checkStatus.DistrictName,
		&checkStatus.Status,
		&checkStatus.AdminNotes,
	)
	if err != nil {
		return checkStatus, err
	}
	return checkStatus, nil
}

func GetStatusHistory(reportNumber string) ([]entity.StatusHistory, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var histories []entity.StatusHistory

	query := `
	SELECT 
		h.previous_status,
		h.new_status,
		h.notes,
		h.created_at
	FROM report_status_history h
	JOIN reports r ON r.report_id = h.report_id
	WHERE r.report_number = ?
	ORDER BY h.created_at ASC
	`

	rows, err := db.QueryContext(ctx, query, reportNumber)
	if err != nil {
		return histories, err
	}
	defer rows.Close()

	for rows.Next() {
		var h entity.StatusHistory
		if err := rows.Scan(&h.PreviousStatus, &h.NewStatus, &h.Notes, &h.CreatedAt); err != nil {
			return histories, err
		}
		histories = append(histories, h)
	}

	return histories, nil
}

func GetVillageNameByID(villageID string) (string, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()
	var villageName string

	query := "SELECT village_name FROM villages WHERE village_id = ?"

	err := db.QueryRowContext(ctx, query, villageID).Scan(&villageName)
	if err != nil {
		return "", err
	}

	return villageName, nil
}

func GetUrgentlyReport() []entity.UrgencyReportRequest {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var reports []entity.UrgencyReportRequest

	sqlQuery := `
		SELECT 
			r.report_id,
			r.report_number,
			d.name,
			v.village_name,
			r.urgency_level
		FROM reports r
		INNER JOIN damage_types d ON d.damage_type_id = r.damage_type_id
		INNER JOIN villages v ON v.village_id = r.village_id
		WHERE r.status NOT IN ('selesai', 'batal')
		ORDER BY r.urgency_level DESC,r.report_number DESC
		LIMIT 5
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
			&urgentlyReport.ReportID,
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

func GetManageReport(keyword, infCategory, status, startDate, endDate string, page int, limit int) ([]entity.ManageReport, int64, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var category []entity.ManageReport
	var total int64
	args := []interface{}{}

	query := `
	SELECT 
		r.report_id,
		r.report_number,
		DATE_FORMAT(r.created_at, '%Y-%m-%d') AS created_at,
		r.reporter_name,
		i.name,
		v.village_name,
		d.district_name,
		r.status,
		COUNT(*) OVER() as total_rows
	FROM reports r
	INNER JOIN infrastructure_categories i ON i.infrastructure_category_id = r.infrastructure_category_id
	INNER JOIN villages v ON v.village_id = r.village_id
	INNER JOIN districts d ON d.district_id = r.district_id
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

	if startDate != "" && endDate != "" {
		query += " AND DATE(r.created_at) BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
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
			&cat.ReportID,
			&cat.ReportNumber,
			&cat.CreatedAt,
			&cat.ReporterName,
			&cat.InfrastructureCategoryName,
			&cat.VillageName,
			&cat.DistrictName,
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

func GetReportPhotos(reportID string) ([]entity.ReportPhoto, error) {
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
	WHERE rp.report_id = ?
	`

	rows, err := db.QueryContext(ctx, query, reportID)
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

func GetDetailReport(reportId string) (entity.DetailReport, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var detailReport entity.DetailReport
	var reporterEmail, description sql.NullString

	query := `
	SELECT
		r.report_id,
		r.report_number,
		r.created_at,
		r.reporter_name,
		r.reporter_phone,
		r.reporter_email,
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
	WHERE r.report_id = ?
	`
	err := db.QueryRowContext(ctx, query, reportId).Scan(
		&detailReport.ReportID,
		&detailReport.ReportNumber,
		&detailReport.CreatedAt,
		&detailReport.ReporterName,
		&detailReport.ReporterPhone,
		&reporterEmail,
		&detailReport.InfrastructureCategoryName,
		&detailReport.DamageTypeName,
		&detailReport.VillageName,
		&detailReport.DistrictName,
		&detailReport.Status,
		&detailReport.Latitude,
		&detailReport.Longitude,
		&detailReport.UrgencyLevel,
		&description,
	)
	if err != nil {
		return detailReport, err
	}

	photos, err := GetReportPhotos(reportId)
	if err != nil && err != sql.ErrNoRows {
		return detailReport, err
	}
	if photos != nil {
		detailReport.Photos = photos
	}
	return detailReport, nil
}

func UpdateReportAndLogHistory(ctx context.Context, reportID string, data entity.UpdateReport, adminID string) error {
	db := database.GetConnectionDB()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %w", err)
	}
	defer tx.Rollback()

	var currentStatus sql.NullString
	err = tx.QueryRowContext(ctx, "SELECT status FROM reports WHERE report_id = ?", reportID).Scan(&currentStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("laporan dengan ID %s tidak ditemukan", reportID)
		}
		return fmt.Errorf("gagal mendapatkan status laporan: %w", err)
	}

	updateClauses := []string{}
	args := []interface{}{}
	statusChanged := false

	if data.Status != "" && data.Status != currentStatus.String {
		updateClauses = append(updateClauses, "status = ?")
		args = append(args, data.Status)
		statusChanged = true
	}

	if data.PicName != nil {
		updateClauses = append(updateClauses, "pic_name = ?")
		args = append(args, *data.PicName)
	}
	if data.PicPhone != nil {
		updateClauses = append(updateClauses, "pic_contact = ?")
		args = append(args, *data.PicPhone)
	}
	if data.AdminNotes != nil {
		updateClauses = append(updateClauses, "admin_notes = ?")
		args = append(args, *data.AdminNotes)
	}
	if data.CompletionNotes != nil {
		updateClauses = append(updateClauses, "completion_notes = ?")
		args = append(args, *data.CompletionNotes)
	}

	// Fix: Check for zero time before adding to update
	if data.EstimatedCompletionDate != nil && !data.EstimatedCompletionDate.Time.IsZero() {
		updateClauses = append(updateClauses, "estimated_completion = ?")
		args = append(args, data.EstimatedCompletionDate.Time)
	} else if data.EstimatedCompletionDate != nil && data.EstimatedCompletionDate.Time.IsZero() {
		// If explicitly setting to zero time, use NULL instead
		updateClauses = append(updateClauses, "estimated_completion = NULL")
	}

	// Fix: Check for zero time before adding to update
	if data.CompletedAt != nil && !data.CompletedAt.Time.IsZero() {
		updateClauses = append(updateClauses, "completed_at = ?")
		args = append(args, data.CompletedAt.Time)
	} else if data.CompletedAt != nil && data.CompletedAt.Time.IsZero() {
		// If explicitly setting to zero time, use NULL instead
		updateClauses = append(updateClauses, "completed_at = NULL")
	}

	if len(updateClauses) > 0 {
		updateClauses = append(updateClauses, "assigned_to = ?", "updated_at = ?")
		args = append(args, adminID, time.Now())

		query := fmt.Sprintf("UPDATE reports SET %s WHERE report_id = ?", strings.Join(updateClauses, ", "))
		args = append(args, reportID)

		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("gagal update laporan: %w", err)
		}
	}

	if statusChanged {
		logQuery := `
			INSERT INTO report_status_history 
			(report_id, previous_status, new_status, notes, updated_by, created_at) 
			VALUES (?, ?, ?, ?, ?, NOW())`

		var notes sql.NullString
		if data.AdminNotes != nil {
			notes = sql.NullString{String: *data.AdminNotes, Valid: true}
		}

		_, err = tx.ExecContext(ctx, logQuery, reportID, currentStatus.String, data.Status, notes, adminID)
		if err != nil {
			return fmt.Errorf("gagal menyimpan riwayat status: %w", err)
		}
	}

	return tx.Commit()
}

func DeleteReportByID(reportID string) error {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	// Gunakan transaksi untuk memastikan semua data terkait dihapus secara atomik
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %w", err)
	}
	// Defer rollback untuk membatalkan transaksi jika terjadi error di tengah jalan
	defer tx.Rollback()

	// LANGKAH 1: Ambil semua path file foto yang akan dihapus
	rows, err := tx.QueryContext(ctx, "SELECT file_path FROM report_photos WHERE report_id = ?", reportID)
	if err != nil {
		return fmt.Errorf("gagal mengambil path foto laporan: %w", err)
	}
	defer rows.Close()

	var photoPaths []string
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			// Jika scan gagal, lebih baik batalkan operasi
			return fmt.Errorf("gagal membaca path foto: %w", err)
		}
		photoPaths = append(photoPaths, path)
	}
	if err = rows.Err(); err != nil {
		return fmt.Errorf("terjadi error saat iterasi path foto: %w", err)
	}

	// LANGKAH 2: Hapus referensi dari tabel-tabel di database
	// Hapus dari tabel report_photos
	_, err = tx.ExecContext(ctx, "DELETE FROM report_photos WHERE report_id = ?", reportID)
	if err != nil {
		return fmt.Errorf("gagal menghapus data foto laporan dari db: %w", err)
	}

	// Hapus dari tabel report_status_history
	_, err = tx.ExecContext(ctx, "DELETE FROM report_status_history WHERE report_id = ?", reportID)
	if err != nil {
		return fmt.Errorf("gagal menghapus riwayat status laporan: %w", err)
	}

	// Hapus dari tabel utama reports
	result, err := tx.ExecContext(ctx, "DELETE FROM reports WHERE report_id = ?", reportID)
	if err != nil {
		return fmt.Errorf("gagal menghapus laporan utama: %w", err)
	}

	// Periksa apakah ada baris yang benar-benar dihapus dari tabel utama
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("gagal memeriksa baris yang terpengaruh: %w", err)
	}

	if rowsAffected == 0 {
		// Gunakan sql.ErrNoRows untuk menandakan bahwa laporan tidak ditemukan
		return sql.ErrNoRows
	}

	// LANGKAH 3: Hapus file fisik dari folder setelah operasi DB siap di-commit
	for _, path := range photoPaths {
		if path != "" {
			err := os.Remove(path)
			if err != nil && !os.IsNotExist(err) {
				// Log error jika file gagal dihapus (dan errornya bukan karena file sudah tidak ada)
				// Kita tidak mengembalikan error agar transaksi DB tetap bisa di-commit.
				// Menghapus data di DB lebih prioritas.
				log.Printf("Peringatan: Gagal menghapus file fisik '%s': %v", path, err)
			}
		}
	}

	// LANGKAH 4: Jika semua berhasil, commit transaksi
	return tx.Commit()
}
