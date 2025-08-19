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

// MonthlyReport returns monthly report statistics
func MonthlyReport(ctx *fiber.Ctx) error {
	// Get query parameters
	year := ctx.QueryInt("year", 0)
	month := ctx.QueryInt("month", 0)

	// Log the request
	fmt.Printf("Monthly report requested from IP: %s (year: %d, month: %d)\n", ctx.IP(), year, month)

	// Get monthly report data
	monthlyReport := models.GetMonthlyReport(year, month)

	// Log the results for debugging
	fmt.Printf("Monthly Report - Total months: %d, Total reports: %d\n",
		len(monthlyReport.MonthlyData), monthlyReport.Summary.TotalReportsLast12Months)

	return response.SuccessResponse(ctx, monthlyReport)
}

// CategoryBreakdown returns infrastructure category analytics
func CategoryBreakdown(ctx *fiber.Ctx) error {
	// Get query parameters
	startDate := ctx.Query("start_date", "")
	endDate := ctx.Query("end_date", "")

	// Validate date format if provided
	if startDate != "" {
		if _, err := time.Parse("2006-01-02", startDate); err != nil {
			return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
				"field": "start_date",
				"error": "Invalid date format. Use YYYY-MM-DD",
			}))
		}
	}

	if endDate != "" {
		if _, err := time.Parse("2006-01-02", endDate); err != nil {
			return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
				"field": "end_date",
				"error": "Invalid date format. Use YYYY-MM-DD",
			}))
		}
	}

	// Log the request
	fmt.Printf("Category breakdown requested from IP: %s (start_date: %s, end_date: %s)\n",
		ctx.IP(), startDate, endDate)

	// Get category breakdown data
	categoryBreakdown := models.GetCategoryBreakdown(startDate, endDate)

	// Log the results for debugging
	fmt.Printf("Category Breakdown - Total categories: %d, Geographic data points: %d\n",
		len(categoryBreakdown.Categories), len(categoryBreakdown.GeographicData))

	return response.SuccessResponse(ctx, categoryBreakdown)
}

// PerformanceChart returns performance metrics for dashboard charts
func PerformanceChart(ctx *fiber.Ctx) error {
	// Get query parameters
	period := ctx.Query("period", "30d")

	// Validate period parameter
	validPeriods := []string{"7d", "30d", "90d", "180d", "365d", "1y"}
	isValidPeriod := false
	for _, validPeriod := range validPeriods {
		if period == validPeriod {
			isValidPeriod = true
			break
		}
	}

	if !isValidPeriod {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "period",
			"error": "Invalid period. Supported periods: 7d, 30d, 90d, 180d, 365d, 1y",
		}))
	}

	// Log the request
	fmt.Printf("Performance chart requested from IP: %s (period: %s)\n", ctx.IP(), period)

	// Get performance chart data
	performanceChart := models.GetPerformanceChart(period)

	// Log the results for debugging
	fmt.Printf("Performance Chart - Period: %s, Resolution rates: %d, Trends: %s\n",
		performanceChart.Period, len(performanceChart.ResolutionRates), performanceChart.Trends.OverallPerformance)

	return response.SuccessResponse(ctx, performanceChart)
}
