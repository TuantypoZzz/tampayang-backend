package models

import (
	"context"
	"errors"
	"fmt"
	"log"

	"tampayang-backend/app/models/entity"
	"tampayang-backend/core/database"
)

func InsertNewReport(ctx context.Context, data *entity.Report) error {
	db := database.GetConnectionDB()
	sqlQuery := `
		INSERT INTO reports (
			report_id, report_number, reporter_name, reporter_phone, 
			infrastructure_category_id, damage_type_id, province_id, regency_id, 
			district_id, village_id, location_detail, description, 
			urgency_level, status, latitude, longitude, created_at
		) VALUES (
			?, ?, ?, ?,
			?, ?, ?, ?,
			?, ?, ?, ?,
			?, ?, ?, ?, ?
		)`

	result, err := db.ExecContext(ctx, sqlQuery,
		data.ReportId,
		data.ReportNumber,
		data.ReporterName,
		data.ReporterPhone,
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
