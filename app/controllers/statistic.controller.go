package controllers

import (
	"fmt"
	"time"

	"tampayang-backend/app/models"
	"tampayang-backend/app/models/entity"
	"tampayang-backend/config/constant"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/response"
	"tampayang-backend/core/validation"

	"github.com/gofiber/fiber/v2"
)

func ReportSummary(ctx *fiber.Ctx) error {
	request := new(entity.ReportSummaryRequest)
	if err := ctx.QueryParser(request); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	if request.StartDate == "" || (ctx.Locals("isLogin") != nil && ctx.Locals("isLogin") == false) {
		request.StartDate = time.Now().Format(constant.NOW_TIME_FORMAT)
	}

	if request.EndDate == "" || (ctx.Locals("isLogin") != nil && ctx.Locals("isLogin") == false) {
		request.EndDate = time.Now().Format(constant.NOW_TIME_FORMAT)
	}

	err := validation.Validate.Struct(request)
	if err != nil {
		return response.ErrorResponse(ctx, fmt.Errorf("validation failed: %w", err))
	}

	if !globalFunction.IsValidDateRange(request.StartDate, request.EndDate) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("date001", nil))
	}

	summary := models.GetReportSummary(*request)

	if ctx.Locals("isLogin") != nil && ctx.Locals("isLogin") == false {
		summary = entity.ReportSummary{
			TotalReport:           summary.TotalReport,
			TotalReportDone:       summary.TotalReportDone,
			TotalReportInProgress: summary.TotalReportInProgress,
			TotalReportWaiting:    summary.TotalReportWaiting,
		}
	}

	return response.SuccessResponse(ctx, summary)
}

func ReportWeekly(ctx *fiber.Ctx) error {
	reportWeekly := models.GetReportWeekly()
	return response.SuccessResponse(ctx, reportWeekly)
}

func ReportMap(ctx *fiber.Ctx) error {
	request := new(entity.ReportMapRequest)
	if err := ctx.QueryParser(request); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	err := validation.Validate.Struct(request)
	if err != nil {
		return response.ErrorResponse(ctx, fmt.Errorf("validation failed: %w", err))
	}

	var reportMap []entity.ReportMap

	if request.DistrictId != "" {
		reportMap = models.GetReportMapVillage(request.DistrictId)
	}

	if request.RegencyId != "" {
		reportMap = models.GetReportMapDistrict(request.RegencyId)
	}

	// if request.ProvinceId != "" {
	// 	reportMap = models.GetReportMapRegency(request.ProvinceId)
	// }

	if request.RegencyId == "" && request.DistrictId == "" {
		reportMap = models.GetReportMapAllRegencies()
	}

	return response.SuccessResponse(ctx, reportMap)
}

func DashboardSummary(ctx *fiber.Ctx) error {
	// Log the dashboard summary request
	fmt.Printf("Dashboard summary requested from IP: %s\n", ctx.IP())

	summary := models.GetDashboardSummary()

	// Log the results for debugging
	fmt.Printf("Dashboard Summary - Total Laporan: %d, Tingkat Penyelesaian: %.2f%%\n",
		summary.TotalLaporan, summary.TingkatPenyelesaian)

	return response.SuccessResponse(ctx, summary)
}
