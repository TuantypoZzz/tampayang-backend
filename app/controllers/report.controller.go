package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"tampayang-backend/app/models"
	"tampayang-backend/app/models/entity"
	"tampayang-backend/config/constant"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/response"

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

	if err := entity.ValidateNewReport(newReport); err != nil {
		errorCode := err.Error()
		return response.ErrorResponse(ctx, globalFunction.GetMessage(errorCode, nil))
	}

	if err := models.InsertNewReport(ctx.Context(), newReport); err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err006", nil))
	}

	for i, file := range newReport.ReportImages {
		fileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))
		savePath := filepath.Join(UploadReportPhoto, fileName)

		if err := ctx.SaveFile(file, savePath); err != nil {
			return response.ErrorResponse(ctx, globalFunction.GetMessage("err007", nil))
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

	return response.SuccessResponse(ctx, newReport)
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
	reportId := ctx.Query("report_id")
	if globalFunction.IsEmpty(reportId) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err008", nil))
	}

	details, err := models.GetDetailReport(reportId)
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
	reports := new(entity.UpdateReport)
	if err := ctx.BodyParser(reports); err != nil {
		return response.ErrorResponse(ctx, err)
	}

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

	if globalFunction.IsEmpty(reports.Status) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("rpt020", nil))
	}

	if globalFunction.IsEmpty(reports.PicName) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("rpt021", nil))
	}
	if globalFunction.IsEmpty(reports.PicPhone) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("rpt022", nil))
	}

	now := time.Now()
	updated_date := now.Format(constant.NOW_DATE_TIME_FORMAT)

	updateData := entity.UpdateReport{
		Status:                  reports.Status,
		PicName:                 reports.PicName,
		PicPhone:                reports.PicPhone,
		AdminNotes:              reports.AdminNotes,
		CompletionNotes:         reports.CompletionNotes,
		EstimatedCompletionDate: reports.EstimatedCompletionDate,
		ComletedAt:              reports.ComletedAt,
		UpdatedAt:               updated_date,
	}

	updateResult, err := models.UpdateReport(reportId, updateData)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	return response.SuccessResponse(ctx, updateResult)
}
