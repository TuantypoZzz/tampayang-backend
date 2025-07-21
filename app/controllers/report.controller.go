package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"tampayang-backend/app/models"
	"tampayang-backend/app/models/entity"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	UploadReportPhoto = "./public/uploads/reportphotos"
)

func CreateReports(ctx *fiber.Ctx) error {
	latStr := ctx.FormValue("latitude", "0.0")
	lonStr := ctx.FormValue("longitude", "0.0")

	// Pastikan direktori upload ada
	if err := os.MkdirAll(UploadReportPhoto, os.ModePerm); err != nil {
		return response.ErrorResponse(ctx, "Gagal membuat direktori upload")
	}

	// Ambil file gambar dari form-data dengan key "report_images"
	form, err := ctx.MultipartForm()
	if err != nil {
		// Tangani jika format form bukan multipart
		return response.ErrorResponse(ctx, "Format request tidak valid")
	}
	// Dapatkan semua file dari field 'report_images'
	files := form.File["report_images"]

	newReport := &entity.Report{
		ReportId:                 uuid.New(),
		ReportNumber:             globalFunction.GenerateReportNumber(),
		ReporterName:             ctx.FormValue("reporter_name"),
		ReporterPhone:            ctx.FormValue("reporter_phone"),
		InfrastructureCategoryId: ctx.FormValue("infrastructure_category_id"),
		DamageTypeID:             ctx.FormValue("damage_type_id"),
		ProviceID:                ctx.FormValue("province_id"),
		RegencyID:                ctx.FormValue("regency_id"),
		DistrictID:               ctx.FormValue("district_id"),
		VillageID:                ctx.FormValue("village_id"),
		LocationDetails:          ctx.FormValue("location_detail"),
		Description:              ctx.FormValue("description"),
		UrgencyLevel:             ctx.FormValue("urgency_level"),
		Status:                   "baru",
		Latitude:                 &latStr,
		Longitude:                &lonStr,
		CreatedAt:                time.Now(),
		ReportImages:             files,
	}

	// Validasi semua data, termasuk gambar
	if err := entity.ValidateNewReport(newReport); err != nil {
		errorCode := err.Error()
		return response.ErrorResponse(ctx, globalFunction.GetMessage(errorCode, nil))
	}

	if err := models.InsertNewReport(ctx.Context(), newReport); err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err006", nil))
	}

	for i, file := range newReport.ReportImages {
		// Buat nama file unik
		fileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))
		savePath := filepath.Join(UploadReportPhoto, fileName)

		// Simpan file fisik ke disk
		if err := ctx.SaveFile(file, savePath); err != nil {
			return response.ErrorResponse(ctx, globalFunction.GetMessage("err007", nil))
		}

		// Siapkan data metadata untuk disimpan ke DB
		photoData := &entity.ReportPhoto{
			ReportPhotoID:    uuid.New(),
			ReportID:         newReport.ReportId,
			Filename:         fileName,
			OriginalFilename: file.Filename,
			FilePath:         savePath,
			FileSize:         file.Size,
			MimeType:         file.Header.Get("Content-Type"),
			IsMain:           i == 0,
		}

		if err := models.InsertReportPhoto(ctx.Context(), photoData); err != nil {
			os.Remove(savePath)
			return response.ErrorResponse(ctx, "Gagal menyimpan metadata file gambar ke database")
		}
	}
	newReport.ReportImages = nil

	return response.SuccessResponse(ctx, newReport)

}
