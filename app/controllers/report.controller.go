package controllers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"tampayang-backend/app/models"
	"tampayang-backend/app/models/entity"
	"tampayang-backend/app/services"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	UploadReportPhoto = "./public/uploads/reportphotos"
)

func CreateReport(ctx *fiber.Ctx) error {
	latStr := ctx.FormValue("latitude", "0.0")
	lonStr := ctx.FormValue("longitude", "0.0")

	if err := os.MkdirAll(UploadReportPhoto, os.ModePerm); err != nil {
		return response.ErrorResponse(ctx, "Gagal membuat direktori upload")
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return response.ErrorResponse(ctx, "Format request tidak valid")
	}
	files := form.File["report_images"]

	newReport := &entity.Report{
		ReportId:                 uuid.New(),
		ReportNumber:             GenerateReportNumber(),
		ReporterName:             ctx.FormValue("reporter_name"),
		ReporterPhone:            ctx.FormValue("reporter_phone"),
		ReporterEmail:            ctx.FormValue("reporter_email"),
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

	if err := entity.ValidateNewReport(newReport); err != nil {
		errorCode := err.Error()
		return response.ErrorResponse(ctx, globalFunction.GetMessage(errorCode, nil))
	}

	if err := models.InsertNewReport(ctx.Context(), newReport); err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err006", nil))
	}

	for i, file := range newReport.ReportImages {
		fileName := fmt.Sprintf("%s.jpg", uuid.New().String())
		savePath := filepath.Join(UploadReportPhoto, fileName)

		if err := globalFunction.CompressAndSaveImage(file, savePath); err != nil {
			return response.ErrorResponse(ctx, "Gagal memproses dan mengompres gambar")
		}

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

	// =============================================================
	// >> PEMANGGILAN NOTIFIKASI <<
	// =============================================================
	var villageName string = "[Lokasi tidak teridentifikasi]"
	if newReport.DistrictID != "" {
		villageList := models.GetLovVillage(newReport.DistrictID)
		for _, village := range villageList {
			if village.Id == newReport.VillageID {
				villageName = village.Name
				break // Hentikan loop jika sudah ditemukan
			}
		}
	} else {
		log.Printf("WARNING: DistrictID tidak ada di laporan, tidak dapat mengambil nama desa.")
	}

	go services.SendFonnteNotification(
		newReport.ReporterName,
		newReport.ReporterPhone,
		newReport.ReportNumber,
		villageName,
	)

	if newReport.ReporterEmail != "" {
		go services.SendEmailNotification(
			newReport.ReporterName,
			newReport.ReporterEmail,
			newReport.ReportNumber,
		)
	}

	// =============================================================

	return response.SuccessResponse(ctx, newReport)
}

func CheckStatus(ctx *fiber.Ctx) error {
	reportNumber := ctx.Query("report_number")
	if globalFunction.IsEmpty(reportNumber) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err008", nil))
	}

	details, err := models.GetCheckStatus(reportNumber)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	// Cek apakah data ditemukan
	if globalFunction.IsEmpty(details.ReportNumber) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err003", nil))
	}

	return response.SuccessResponse(ctx, details)
}

func UrgencyReport(ctx *fiber.Ctx) error {
	urgentlyReport := models.GetUrgentlyReport()
	return response.SuccessResponse(ctx, urgentlyReport)
}

func ManageReport(ctx *fiber.Ctx) error {
	keyword := ctx.Query("keyword", "")
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	year := ctx.Query("year", "")
	infCategory := ctx.Query("category", "")
	status := ctx.Query("status", "")

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	ManageReport, total, err := models.GetManageReport(keyword, year, infCategory, status, page, limit)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	return response.SuccessResponse(ctx, fiber.Map{
		"data": ManageReport,
		"pagination": fiber.Map{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func DetailReport(ctx *fiber.Ctx) error {
	reportNumber := ctx.Query("report_number")
	if globalFunction.IsEmpty(reportNumber) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err008", nil))
	}

	details, err := models.GetDetailReport(reportNumber)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	// Cek apakah data ditemukan
	if globalFunction.IsEmpty(details.ReportNumber) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err003", nil))
	}

	return response.SuccessResponse(ctx, details)
}

func UpdateReport(ctx *fiber.Ctx) error {
	reportId := ctx.Params("report_id")
	fmt.Println(reportId)
	if globalFunction.IsEmpty(reportId) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err002", nil))
	}

	dbResult, err := models.GetDetailReport(reportId)
	if err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err008", nil))
	}
	if dbResult.ReportID == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err006", nil))
	}

	claims, ok := ctx.Locals("userInfo").(jwt.MapClaims)
	if !ok {
		return response.ErrorResponse(ctx, "Akses ditolak: Gagal memproses informasi user")
	}

	adminID, ok := claims["user_id"].(string)
	if !ok {
		return response.ErrorResponse(ctx, "Akses ditolak: User ID tidak ditemukan")
	}

	updateData := new(entity.UpdateReport)
	if err := ctx.BodyParser(updateData); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	if globalFunction.IsEmpty(updateData.Status) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("rpt020", nil))
	}

	if globalFunction.IsEmpty(updateData.PicName) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("rpt021", nil))
	}
	if globalFunction.IsEmpty(updateData.PicPhone) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("rpt022", nil))
	}

	err = models.UpdateReportAndLogHistory(ctx.Context(), reportId, *updateData, adminID)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	return response.SuccessResponse(ctx, fiber.Map{
		"message": "Laporan berhasil diperbarui",
	})
}

func GenerateReportNumber() string {
	currentYear := time.Now().Year()
	lastSequence, err := models.GetLastSequenceForYear(currentYear)
	if err != nil {
		return fmt.Sprintf("TMP-%d-ERR%d", currentYear, time.Now().Unix())
	}
	newSequence := lastSequence + 1
	return fmt.Sprintf("TMP-%d-%06d", currentYear, newSequence)
}
