package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"tampayang-backend/app/models/entity"
	"tampayang-backend/core/database"
)

func GetReportSummary(request entity.ReportSummaryRequest) entity.ReportSummary {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var summary entity.ReportSummary

	sqlQuery := `
		SELECT 
			count(1) AS total_report, 
			IFNULL(SUM(IF(status = 'selesai', 1, 0)), 0) AS total_report_done, 
			IFNULL(SUM(IF(status = 'proses', 1, 0)), 0) AS total_report_in_progress, 
			IFNULL(SUM(IF(status IN ('baru', 'verifikasi', 'ditolak'), 1, 0)), 0) AS total_report_waiting, 
			IFNULL(SUM(IF(status = 'baru', 1, 0)), 0) AS total_report_new, 
			IFNULL(SUM(IF(status = 'verifikasi', 1, 0)), 0) AS total_report_verification
		FROM reports
		WHERE created_at BETWEEN ? AND ?
	`
	result, err := db.QueryContext(ctx, sqlQuery, fmt.Sprintf("%s 00:00:00", request.StartDate), fmt.Sprintf("%s 23:59:59", request.EndDate))
	if err != nil {
		panic("models - GetReportSummary, db.QueryContext " + err.Error())
	}
	defer result.Close()

	if result.Next() {
		err := result.Scan(
			&summary.TotalReport,
			&summary.TotalReportDone,
			&summary.TotalReportInProgress,
			&summary.TotalReportWaiting,
			&summary.TotalReportNew,
			&summary.TotalReportVerification)
		if err != nil {
			panic("models - GetReportSummary, result.Scan " + err.Error())
		}
	}
	return summary
}

func GetReportWeekly() []entity.ReportWeekly {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var summary []entity.ReportWeekly

	sqlQuery := `
		SELECT
    		weekly.day_date AS date,
    		DAYNAME(weekly.day_date) AS day,
    		COUNT(r.created_at) AS total
		FROM (
			SELECT CURDATE() - INTERVAL WEEKDAY(CURDATE()) DAY AS day_date
			UNION ALL SELECT CURDATE() - INTERVAL WEEKDAY(CURDATE()) - 1 DAY
			UNION ALL SELECT CURDATE() - INTERVAL WEEKDAY(CURDATE()) - 2 DAY
			UNION ALL SELECT CURDATE() - INTERVAL WEEKDAY(CURDATE()) - 3 DAY
			UNION ALL SELECT CURDATE() - INTERVAL WEEKDAY(CURDATE()) - 4 DAY
			UNION ALL SELECT CURDATE() - INTERVAL WEEKDAY(CURDATE()) - 5 DAY
			UNION ALL SELECT CURDATE() - INTERVAL WEEKDAY(CURDATE()) - 6 DAY
		) AS weekly
		LEFT JOIN reports r ON weekly.day_date = DATE(r.created_at)
		GROUP BY weekly.day_date, day
		ORDER BY weekly.day_date; 
	`
	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - GetWeeklyReport, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var reportWeekly entity.ReportWeekly
		err := result.Scan(
			&reportWeekly.Date,
			&reportWeekly.Day,
			&reportWeekly.Total)
		summary = append(summary, reportWeekly)
		if err != nil {
			panic("models - GetWeeklyReport, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetWeeklyReport, result.Err " + err.Error())
	}
	return summary
}

func GetReportMapProvince() []entity.ReportMap {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var locations []entity.ReportMap

	sqlQuery := `
		SELECT
    		r.province_id AS id,
    		l.province_name AS name,
    		count(1) AS total,
			'province_id' AS filter_key,
    		l.latitude AS latitude,
    		l.longitude AS longitude,
			SUM(CASE WHEN r.status IN ('baru', 'verifikasi') THEN 1 ELSE 0 END) AS menunggu,
			SUM(CASE WHEN r.status = 'proses' THEN 1 ELSE 0 END) AS proses,
			SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) AS selesai,
			SUM(CASE WHEN r.status = 'ditolak' THEN 1 ELSE 0 END) AS ditolak
		FROM reports r
		INNER JOIN provinces l ON l.province_id = r.province_id
		GROUP BY r.province_id
		ORDER BY name;
	`
	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - GetReportMapProvince, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var location entity.ReportMap
		err := result.Scan(
			&location.Id,
			&location.Name,
			&location.Total,
			&location.FilterKey,
			&location.Latitude,
			&location.Longitude,
			&location.StatusBreakdown.Menunggu,
			&location.StatusBreakdown.Proses,
			&location.StatusBreakdown.Selesai,
			&location.StatusBreakdown.Ditolak)
		locations = append(locations, location)
		if err != nil {
			panic("models - GetReportMapProvince, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetReportMapProvince, result.Err " + err.Error())
	}
	return locations
}

func GetReportMapAllRegencies() []entity.ReportMap {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var locations []entity.ReportMap

	sqlQuery := `
		SELECT
    		r.regency_id AS id,
    		l.regency_name AS name,
    		count(1) AS total,
			'regency_id' AS filter_key,
    		l.latitude AS latitude,
    		l.longitude AS longitude,
			SUM(CASE WHEN r.status IN ('baru', 'verifikasi') THEN 1 ELSE 0 END) AS menunggu,
			SUM(CASE WHEN r.status = 'proses' THEN 1 ELSE 0 END) AS proses,
			SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) AS selesai,
			SUM(CASE WHEN r.status = 'ditolak' THEN 1 ELSE 0 END) AS ditolak
		FROM reports r
		INNER JOIN regencies l ON l.regency_id = r.regency_id
		GROUP BY r.regency_id
		ORDER BY name;
	`
	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - GetReportMapAllRegencies, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var location entity.ReportMap
		err := result.Scan(
			&location.Id,
			&location.Name,
			&location.Total,
			&location.FilterKey,
			&location.Latitude,
			&location.Longitude,
			&location.StatusBreakdown.Menunggu,
			&location.StatusBreakdown.Proses,
			&location.StatusBreakdown.Selesai,
			&location.StatusBreakdown.Ditolak)
		locations = append(locations, location)
		if err != nil {
			panic("models - GetReportMapAllRegencies, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetReportMapAllRegencies, result.Err " + err.Error())
	}
	return locations
}

func GetReportMapRegency(provinceId string) []entity.ReportMap {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var locations []entity.ReportMap

	sqlQuery := `
		SELECT
    		r.regency_id AS id,
    		l.regency_name AS name,
    		count(1) AS total,
			'regency_id' AS filter_key,
    		l.latitude AS latitude,
    		l.longitude AS longitude,
			SUM(CASE WHEN r.status IN ('baru', 'verifikasi') THEN 1 ELSE 0 END) AS menunggu,
			SUM(CASE WHEN r.status = 'proses' THEN 1 ELSE 0 END) AS proses,
			SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) AS selesai,
			SUM(CASE WHEN r.status = 'ditolak' THEN 1 ELSE 0 END) AS ditolak
		FROM reports r
		INNER JOIN regencies l ON l.regency_id = r.regency_id
		WHERE l.province_id = ?
		GROUP BY r.regency_id
		ORDER BY name;
	`
	result, err := db.QueryContext(ctx, sqlQuery, provinceId)
	if err != nil {
		panic("models - GetReportMapRegency, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var location entity.ReportMap
		err := result.Scan(
			&location.Id,
			&location.Name,
			&location.Total,
			&location.FilterKey,
			&location.Latitude,
			&location.Longitude,
			&location.StatusBreakdown.Menunggu,
			&location.StatusBreakdown.Proses,
			&location.StatusBreakdown.Selesai,
			&location.StatusBreakdown.Ditolak)
		locations = append(locations, location)
		if err != nil {
			panic("models - GetReportMapRegency, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetReportMapRegency, result.Err " + err.Error())
	}
	return locations
}

func GetReportMapDistrict(regencyId string) []entity.ReportMap {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var locations []entity.ReportMap

	sqlQuery := `
		SELECT
    		r.district_id AS id,
    		l.district_name AS name,
    		count(1) AS total,
			'district_id' AS filter_key,
    		l.latitude AS latitude,
    		l.longitude AS longitude,
			SUM(CASE WHEN r.status IN ('baru', 'verifikasi') THEN 1 ELSE 0 END) AS menunggu,
			SUM(CASE WHEN r.status = 'proses' THEN 1 ELSE 0 END) AS proses,
			SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) AS selesai,
			SUM(CASE WHEN r.status = 'ditolak' THEN 1 ELSE 0 END) AS ditolak
		FROM reports r
		INNER JOIN districts l ON l.district_id = r.district_id
		WHERE l.regency_id = ?
		GROUP BY r.district_id
		ORDER BY name;
	`
	result, err := db.QueryContext(ctx, sqlQuery, regencyId)
	if err != nil {
		panic("models - GetReportMapDistrict, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var location entity.ReportMap
		err := result.Scan(
			&location.Id,
			&location.Name,
			&location.Total,
			&location.FilterKey,
			&location.Latitude,
			&location.Longitude,
			&location.StatusBreakdown.Menunggu,
			&location.StatusBreakdown.Proses,
			&location.StatusBreakdown.Selesai,
			&location.StatusBreakdown.Ditolak)
		locations = append(locations, location)
		if err != nil {
			panic("models - GetReportMapDistrict, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetReportMapDistrict, result.Err " + err.Error())
	}
	return locations
}

// GetDashboardSummary returns total reports and completion rate
func GetDashboardSummary() entity.DashboardSummary {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var summary entity.DashboardSummary

	sqlQuery := `
		SELECT
			COUNT(1) AS total_laporan,
			ROUND(
				(SUM(CASE WHEN status = 'selesai' THEN 1 ELSE 0 END) * 100.0) / COUNT(1),
				2
			) AS tingkat_penyelesaian
		FROM reports
	`

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - GetDashboardSummary, db.QueryContext " + err.Error())
	}
	defer result.Close()

	if result.Next() {
		err := result.Scan(
			&summary.TotalLaporan,
			&summary.TingkatPenyelesaian)
		if err != nil {
			panic("models - GetDashboardSummary, result.Scan " + err.Error())
		}
	}

	return summary
}

// GetMonthlyReport returns monthly report statistics
func GetMonthlyReport(year, month int) entity.MonthlyReport {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var monthlyReport entity.MonthlyReport
	var monthlyData []entity.MonthlyData

	// Get last 12 months of data
	sqlQuery := `
		SELECT
			YEAR(created_at) as year,
			MONTH(created_at) as month,
			MONTHNAME(created_at) as month_name,
			COUNT(*) as total_reports,
			SUM(CASE WHEN status IN ('baru', 'verifikasi') THEN 1 ELSE 0 END) AS menunggu,
			SUM(CASE WHEN status = 'proses' THEN 1 ELSE 0 END) AS proses,
			SUM(CASE WHEN status = 'selesai' THEN 1 ELSE 0 END) AS selesai,
			SUM(CASE WHEN status = 'ditolak' THEN 1 ELSE 0 END) AS ditolak,
			AVG(
				CASE
					WHEN status = 'selesai' AND updated_at IS NOT NULL
					THEN DATEDIFF(updated_at, created_at)
					ELSE NULL
				END
			) as avg_resolution_days
		FROM reports
		WHERE created_at >= DATE_SUB(NOW(), INTERVAL 12 MONTH)
		GROUP BY YEAR(created_at), MONTH(created_at)
		ORDER BY year DESC, month DESC
		LIMIT 12
	`

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - GetMonthlyReport, db.QueryContext " + err.Error())
	}
	defer result.Close()

	var previousTotal int = 0
	for result.Next() {
		var data entity.MonthlyData
		var avgResolutionDays sql.NullFloat64

		err := result.Scan(
			&data.Year,
			&data.Month,
			&data.MonthName,
			&data.TotalReports,
			&data.StatusBreakdown.Menunggu,
			&data.StatusBreakdown.Proses,
			&data.StatusBreakdown.Selesai,
			&data.StatusBreakdown.Ditolak,
			&avgResolutionDays)

		if err != nil {
			panic("models - GetMonthlyReport, result.Scan " + err.Error())
		}

		// Set average resolution days
		if avgResolutionDays.Valid {
			data.AvgResolutionDays = avgResolutionDays.Float64
		} else {
			data.AvgResolutionDays = 0
		}

		// Calculate growth percentage
		if previousTotal > 0 {
			data.GrowthPercentage = ((float64(data.TotalReports) - float64(previousTotal)) / float64(previousTotal)) * 100
		} else {
			data.GrowthPercentage = 0
		}
		previousTotal = data.TotalReports

		monthlyData = append(monthlyData, data)
	}

	// Calculate summary
	var totalReports int
	var bestMonth, worstMonth string
	var bestCount, worstCount int = 0, 999999

	for _, data := range monthlyData {
		totalReports += data.TotalReports
		if data.TotalReports > bestCount {
			bestCount = data.TotalReports
			bestMonth = fmt.Sprintf("%s %d", data.MonthName, data.Year)
		}
		if data.TotalReports < worstCount {
			worstCount = data.TotalReports
			worstMonth = fmt.Sprintf("%s %d", data.MonthName, data.Year)
		}
	}

	monthlyReport.MonthlyData = monthlyData
	monthlyReport.Summary = entity.MonthlySummary{
		TotalReportsLast12Months: totalReports,
		AvgMonthlyReports:        float64(totalReports) / float64(len(monthlyData)),
		BestMonth:                bestMonth,
		WorstMonth:               worstMonth,
	}

	return monthlyReport
}

// GetCategoryBreakdown returns infrastructure category analytics
func GetCategoryBreakdown(startDate, endDate string) entity.CategoryBreakdown {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var categoryBreakdown entity.CategoryBreakdown
	var categories []entity.CategoryData

	// Build date filter
	dateFilter := ""
	var args []interface{}
	if startDate != "" && endDate != "" {
		dateFilter = "WHERE r.created_at BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	// Get category data with completion rates
	sqlQuery := fmt.Sprintf(`
		SELECT
			ic.infrastructure_category_id,
			ic.name as category_name,
			ic.code as category_code,
			COUNT(r.report_id) as total_reports,
			ROUND(
				(SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) * 100.0) / COUNT(r.report_id),
				2
			) as completion_rate
		FROM infrastructure_categories ic
		LEFT JOIN reports r ON ic.infrastructure_category_id = r.infrastructure_category_id %s
		GROUP BY ic.infrastructure_category_id, ic.name, ic.code
		ORDER BY total_reports DESC
	`, dateFilter)

	result, err := db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		panic("models - GetCategoryBreakdown, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var category entity.CategoryData
		err := result.Scan(
			&category.CategoryID,
			&category.CategoryName,
			&category.CategoryCode,
			&category.TotalReports,
			&category.CompletionRate)

		if err != nil {
			panic("models - GetCategoryBreakdown, result.Scan " + err.Error())
		}

		// Get damage types for this category
		category.DamageTypes = getDamageTypesForCategory(category.CategoryID, startDate, endDate)
		categories = append(categories, category)
	}

	// Get geographic data
	geographicData := getCategoryGeographicData(startDate, endDate)

	// Calculate summary
	summary := calculateCategorySummary(categories)

	categoryBreakdown.Categories = categories
	categoryBreakdown.Summary = summary
	categoryBreakdown.GeographicData = geographicData

	return categoryBreakdown
}

// Helper function to get damage types for a category
func getDamageTypesForCategory(categoryID, startDate, endDate string) []entity.DamageTypeData {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var damageTypes []entity.DamageTypeData

	// Build date filter
	dateFilter := ""
	var args []interface{}
	args = append(args, categoryID)
	if startDate != "" && endDate != "" {
		dateFilter = "AND r.created_at BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	sqlQuery := fmt.Sprintf(`
		SELECT
			dt.damage_type_id,
			dt.name as damage_type_name,
			dt.code as damage_type_code,
			COUNT(r.report_id) as report_count,
			ROUND(
				(COUNT(r.report_id) * 100.0) /
				(SELECT COUNT(*) FROM reports r2 WHERE r2.infrastructure_category_id = ? %s),
				2
			) as percentage
		FROM damage_types dt
		LEFT JOIN reports r ON dt.damage_type_id = r.damage_type_id
		WHERE dt.infrastructure_category_id = ? %s
		GROUP BY dt.damage_type_id, dt.name, dt.code
		ORDER BY report_count DESC
	`, dateFilter, dateFilter)

	// Duplicate categoryID for the subquery
	queryArgs := []interface{}{categoryID}
	if startDate != "" && endDate != "" {
		queryArgs = append(queryArgs, startDate, endDate)
	}
	queryArgs = append(queryArgs, categoryID)
	if startDate != "" && endDate != "" {
		queryArgs = append(queryArgs, startDate, endDate)
	}

	result, err := db.QueryContext(ctx, sqlQuery, queryArgs...)
	if err != nil {
		panic("models - getDamageTypesForCategory, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var damageType entity.DamageTypeData
		err := result.Scan(
			&damageType.DamageTypeID,
			&damageType.DamageTypeName,
			&damageType.DamageTypeCode,
			&damageType.ReportCount,
			&damageType.Percentage)

		if err != nil {
			panic("models - getDamageTypesForCategory, result.Scan " + err.Error())
		}

		damageTypes = append(damageTypes, damageType)
	}

	return damageTypes
}

// Helper function to get geographic data for categories
func getCategoryGeographicData(startDate, endDate string) []entity.GeographicData {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var geographicData []entity.GeographicData

	// Build date filter
	dateFilter := ""
	var args []interface{}
	if startDate != "" && endDate != "" {
		dateFilter = "WHERE r.created_at BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	sqlQuery := fmt.Sprintf(`
		SELECT
			reg.regency_id,
			reg.regency_name,
			ic.name as category_name,
			COUNT(r.report_id) as report_count,
			ROUND(
				(COUNT(r.report_id) * 100.0) /
				(SELECT COUNT(*) FROM reports r2 %s),
				2
			) as percentage
		FROM regencies reg
		LEFT JOIN reports r ON reg.regency_id = r.regency_id
		LEFT JOIN infrastructure_categories ic ON r.infrastructure_category_id = ic.infrastructure_category_id
		%s
		GROUP BY reg.regency_id, reg.regency_name, ic.name
		HAVING report_count > 0
		ORDER BY report_count DESC
		LIMIT 20
	`, dateFilter, dateFilter)

	// Duplicate args for the subquery
	queryArgs := args
	queryArgs = append(queryArgs, args...)

	result, err := db.QueryContext(ctx, sqlQuery, queryArgs...)
	if err != nil {
		panic("models - getCategoryGeographicData, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var geoData entity.GeographicData
		err := result.Scan(
			&geoData.RegencyID,
			&geoData.RegencyName,
			&geoData.CategoryName,
			&geoData.ReportCount,
			&geoData.Percentage)

		if err != nil {
			panic("models - getCategoryGeographicData, result.Scan " + err.Error())
		}

		geographicData = append(geographicData, geoData)
	}

	return geographicData
}

// Helper function to calculate category summary
func calculateCategorySummary(categories []entity.CategoryData) entity.CategorySummary {
	var summary entity.CategorySummary
	var mostCommonCategory string
	var highestCompletionCategory, lowestCompletionCategory string
	var maxReports int = 0
	var highestCompletion, lowestCompletion float64 = 0, 100
	var allDamageTypes []entity.DamageTypeData

	for _, category := range categories {
		// Find most common category
		if category.TotalReports > maxReports {
			maxReports = category.TotalReports
			mostCommonCategory = category.CategoryName
		}

		// Find highest and lowest completion rates
		if category.CompletionRate > highestCompletion {
			highestCompletion = category.CompletionRate
			highestCompletionCategory = category.CategoryName
		}
		if category.CompletionRate < lowestCompletion {
			lowestCompletion = category.CompletionRate
			lowestCompletionCategory = category.CategoryName
		}

		// Collect all damage types
		allDamageTypes = append(allDamageTypes, category.DamageTypes...)
	}

	// Find most common damage type overall
	damageTypeMap := make(map[string]int)
	var mostCommonDamageType string
	var maxDamageCount int = 0

	for _, damageType := range allDamageTypes {
		damageTypeMap[damageType.DamageTypeName] += damageType.ReportCount
		if damageTypeMap[damageType.DamageTypeName] > maxDamageCount {
			maxDamageCount = damageTypeMap[damageType.DamageTypeName]
			mostCommonDamageType = damageType.DamageTypeName
		}
	}

	// Get top 5 common damage types
	type damageTypeCount struct {
		name  string
		count int
	}
	var damageTypeCounts []damageTypeCount
	for name, count := range damageTypeMap {
		damageTypeCounts = append(damageTypeCounts, damageTypeCount{name, count})
	}

	// Sort by count (simple bubble sort for small data)
	for i := 0; i < len(damageTypeCounts)-1; i++ {
		for j := 0; j < len(damageTypeCounts)-i-1; j++ {
			if damageTypeCounts[j].count < damageTypeCounts[j+1].count {
				damageTypeCounts[j], damageTypeCounts[j+1] = damageTypeCounts[j+1], damageTypeCounts[j]
			}
		}
	}

	// Take top 5
	var commonDamageTypes []entity.DamageTypeData
	limit := 5
	if len(damageTypeCounts) < limit {
		limit = len(damageTypeCounts)
	}

	totalReports := 0
	for _, dtc := range damageTypeCounts {
		totalReports += dtc.count
	}

	for i := 0; i < limit; i++ {
		percentage := float64(damageTypeCounts[i].count) * 100.0 / float64(totalReports)
		commonDamageTypes = append(commonDamageTypes, entity.DamageTypeData{
			DamageTypeName: damageTypeCounts[i].name,
			ReportCount:    damageTypeCounts[i].count,
			Percentage:     percentage,
		})
	}

	summary.MostCommonCategory = mostCommonCategory
	summary.MostCommonDamageType = mostCommonDamageType
	summary.HighestCompletionRate = highestCompletionCategory
	summary.LowestCompletionRate = lowestCompletionCategory
	summary.CommonDamageTypes = commonDamageTypes

	return summary
}

// GetPerformanceChart returns performance metrics for dashboard charts
func GetPerformanceChart(period string) entity.PerformanceChart {
	db := database.GetConnectionDB()
	defer db.Close()

	var performanceChart entity.PerformanceChart
	performanceChart.Period = period

	// Determine date range based on period
	var dateRange string
	var groupBy string
	var dateFormat string

	switch period {
	case "7d":
		dateRange = "DATE_SUB(NOW(), INTERVAL 7 DAY)"
		groupBy = "DATE(created_at)"
		dateFormat = "%Y-%m-%d"
	case "30d":
		dateRange = "DATE_SUB(NOW(), INTERVAL 30 DAY)"
		groupBy = "DATE(created_at)"
		dateFormat = "%Y-%m-%d"
	case "90d":
		dateRange = "DATE_SUB(NOW(), INTERVAL 90 DAY)"
		groupBy = "DATE_FORMAT(created_at, '%Y-%m-%d')"
		dateFormat = "%Y-%m-%d"
	case "180d":
		dateRange = "DATE_SUB(NOW(), INTERVAL 180 DAY)"
		groupBy = "DATE_FORMAT(DATE_SUB(created_at, INTERVAL WEEKDAY(created_at) DAY), '%Y-%m-%d')"
		dateFormat = "%Y-%m-%d"
	case "365d":
		dateRange = "DATE_SUB(NOW(), INTERVAL 365 DAY)"
		groupBy = "DATE_FORMAT(created_at, '%Y-%m')"
		dateFormat = "%Y-%m"
	case "1y":
		dateRange = "DATE_SUB(NOW(), INTERVAL 1 YEAR)"
		groupBy = "DATE_FORMAT(created_at, '%Y-%m')"
		dateFormat = "%Y-%m"
	default:
		dateRange = "DATE_SUB(NOW(), INTERVAL 30 DAY)"
		groupBy = "DATE(created_at)"
		dateFormat = "%Y-%m-%d"
	}

	// Get resolution rates
	performanceChart.ResolutionRates = getResolutionRates(dateRange, groupBy, dateFormat)

	// Get response times
	performanceChart.ResponseTimes = getResponseTimes(dateRange, groupBy, dateFormat)

	// Get completion times
	performanceChart.CompletionTimes = getCompletionTimes(dateRange, groupBy, dateFormat)

	// Get workload data
	performanceChart.WorkloadData = getWorkloadData(dateRange)

	// Calculate trends
	performanceChart.Trends = calculatePerformanceTrends(
		performanceChart.ResolutionRates,
		performanceChart.ResponseTimes,
		performanceChart.CompletionTimes)

	return performanceChart
}

// Helper function to get resolution rates
func getResolutionRates(dateRange, groupBy, dateFormat string) []entity.ResolutionRate {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var resolutionRates []entity.ResolutionRate

	sqlQuery := fmt.Sprintf(`
		SELECT
			DATE_FORMAT(created_at, '%s') as date,
			COUNT(*) as total_count,
			SUM(CASE WHEN status = 'selesai' THEN 1 ELSE 0 END) as resolved_count,
			ROUND(
				(SUM(CASE WHEN status = 'selesai' THEN 1 ELSE 0 END) * 100.0) / COUNT(*),
				2
			) as resolution_rate
		FROM reports
		WHERE created_at >= %s
		GROUP BY %s
		ORDER BY created_at ASC
	`, dateFormat, dateRange, groupBy)

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - getResolutionRates, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var rate entity.ResolutionRate
		err := result.Scan(
			&rate.Date,
			&rate.TotalCount,
			&rate.ResolvedCount,
			&rate.ResolutionRate)

		if err != nil {
			panic("models - getResolutionRates, result.Scan " + err.Error())
		}

		resolutionRates = append(resolutionRates, rate)
	}

	return resolutionRates
}

// Helper function to get response times
func getResponseTimes(dateRange, groupBy, dateFormat string) []entity.ResponseTime {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var responseTimes []entity.ResponseTime

	// Simplified query without complex median calculation
	sqlQuery := fmt.Sprintf(`
		SELECT
			DATE_FORMAT(created_at, '%s') as date,
			AVG(TIMESTAMPDIFF(HOUR, created_at, updated_at)) as avg_response_hours
		FROM reports
		WHERE created_at >= %s AND updated_at IS NOT NULL
		GROUP BY %s
		ORDER BY created_at ASC
	`, dateFormat, dateRange, groupBy)

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - getResponseTimes, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var responseTime entity.ResponseTime

		err := result.Scan(
			&responseTime.Date,
			&responseTime.AvgResponseHours)

		if err != nil {
			panic("models - getResponseTimes, result.Scan " + err.Error())
		}

		// Calculate median separately for each date group
		responseTime.MedianResponseHours = calculateMedianResponseTime(responseTime.Date, dateRange, groupBy, dateFormat)

		responseTimes = append(responseTimes, responseTime)
	}

	return responseTimes
}

// Helper function to calculate median response time for a specific date
func calculateMedianResponseTime(date, dateRange, groupBy, dateFormat string) float64 {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	// Build date filter for the specific date group
	var dateFilter string
	switch {
	case groupBy == "DATE(created_at)":
		dateFilter = fmt.Sprintf("DATE(created_at) = '%s'", date)
	case strings.Contains(groupBy, "DATE_FORMAT") && strings.Contains(groupBy, "%Y-%m-%d"):
		// Handle daily grouping with DATE_FORMAT
		dateFilter = fmt.Sprintf("DATE(created_at) = '%s'", date)
	case strings.Contains(groupBy, "DATE_SUB") && strings.Contains(groupBy, "WEEKDAY"):
		// Handle weekly grouping with week start date - date format is "2025-02-10" (Monday of the week)
		dateFilter = fmt.Sprintf("DATE_FORMAT(DATE_SUB(created_at, INTERVAL WEEKDAY(created_at) DAY), '%%Y-%%m-%%d') = '%s'", date)
	case strings.Contains(groupBy, "YEARWEEK"):
		// Handle weekly grouping - date format is "2025-W32"
		if strings.Contains(date, "-W") {
			parts := strings.Split(date, "-W")
			if len(parts) == 2 {
				year := parts[0]
				week := parts[1]
				dateFilter = fmt.Sprintf("YEARWEEK(created_at, 1) = %s%s", year, week)
			} else {
				dateFilter = "YEARWEEK(created_at, 1) = YEARWEEK(NOW(), 1)"
			}
		} else {
			dateFilter = fmt.Sprintf("YEARWEEK(created_at, 1) = YEARWEEK('%s-01-01', 1)", date)
		}
	case strings.Contains(groupBy, "DATE_FORMAT"):
		// Handle monthly grouping - date format is "2025-01"
		dateFilter = fmt.Sprintf("DATE_FORMAT(created_at, '%%Y-%%m') = '%s'", date)
	default:
		dateFilter = fmt.Sprintf("DATE(created_at) = '%s'", date)
	}

	sqlQuery := fmt.Sprintf(`
		SELECT TIMESTAMPDIFF(HOUR, created_at, updated_at) as response_hours
		FROM reports
		WHERE created_at >= %s AND updated_at IS NOT NULL AND %s
		ORDER BY response_hours
	`, dateRange, dateFilter)

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		// Return 0 if error occurs in median calculation
		return 0
	}
	defer result.Close()

	var responseTimes []float64
	for result.Next() {
		var responseHours sql.NullFloat64
		err := result.Scan(&responseHours)
		if err == nil && responseHours.Valid {
			responseTimes = append(responseTimes, responseHours.Float64)
		}
	}

	// Calculate median from the sorted array
	if len(responseTimes) == 0 {
		return 0
	}

	mid := len(responseTimes) / 2
	if len(responseTimes)%2 == 0 {
		// Even number of elements - average of two middle values
		return (responseTimes[mid-1] + responseTimes[mid]) / 2
	} else {
		// Odd number of elements - middle value
		return responseTimes[mid]
	}
}

// Helper function to get completion times
func getCompletionTimes(dateRange, groupBy, dateFormat string) []entity.CompletionTime {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var completionTimes []entity.CompletionTime

	// Simplified query without complex median calculation
	sqlQuery := fmt.Sprintf(`
		SELECT
			DATE_FORMAT(created_at, '%s') as date,
			AVG(DATEDIFF(updated_at, created_at)) as avg_completion_days
		FROM reports
		WHERE created_at >= %s AND status = 'selesai' AND updated_at IS NOT NULL
		GROUP BY %s
		ORDER BY created_at ASC
	`, dateFormat, dateRange, groupBy)

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - getCompletionTimes, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var completionTime entity.CompletionTime

		err := result.Scan(
			&completionTime.Date,
			&completionTime.AvgCompletionDays)

		if err != nil {
			panic("models - getCompletionTimes, result.Scan " + err.Error())
		}

		// Calculate median separately for each date group
		completionTime.MedianCompletionDays = calculateMedianCompletionTime(completionTime.Date, dateRange, groupBy, dateFormat)

		completionTimes = append(completionTimes, completionTime)
	}

	return completionTimes
}

// Helper function to calculate median completion time for a specific date
func calculateMedianCompletionTime(date, dateRange, groupBy, dateFormat string) float64 {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	// Build date filter for the specific date group
	var dateFilter string
	switch {
	case groupBy == "DATE(created_at)":
		dateFilter = fmt.Sprintf("DATE(created_at) = '%s'", date)
	case strings.Contains(groupBy, "DATE_FORMAT") && strings.Contains(groupBy, "%Y-%m-%d"):
		// Handle daily grouping with DATE_FORMAT
		dateFilter = fmt.Sprintf("DATE(created_at) = '%s'", date)
	case strings.Contains(groupBy, "DATE_SUB") && strings.Contains(groupBy, "WEEKDAY"):
		// Handle weekly grouping with week start date - date format is "2025-02-10" (Monday of the week)
		dateFilter = fmt.Sprintf("DATE_FORMAT(DATE_SUB(created_at, INTERVAL WEEKDAY(created_at) DAY), '%%Y-%%m-%%d') = '%s'", date)
	case strings.Contains(groupBy, "YEARWEEK"):
		// Handle weekly grouping - date format is "2025-W32"
		if strings.Contains(date, "-W") {
			parts := strings.Split(date, "-W")
			if len(parts) == 2 {
				year := parts[0]
				week := parts[1]
				dateFilter = fmt.Sprintf("YEARWEEK(created_at, 1) = %s%s", year, week)
			} else {
				dateFilter = "YEARWEEK(created_at, 1) = YEARWEEK(NOW(), 1)"
			}
		} else {
			dateFilter = fmt.Sprintf("YEARWEEK(created_at, 1) = YEARWEEK('%s-01-01', 1)", date)
		}
	case strings.Contains(groupBy, "DATE_FORMAT"):
		// Handle monthly grouping - date format is "2025-01"
		dateFilter = fmt.Sprintf("DATE_FORMAT(created_at, '%%Y-%%m') = '%s'", date)
	default:
		dateFilter = fmt.Sprintf("DATE(created_at) = '%s'", date)
	}

	sqlQuery := fmt.Sprintf(`
		SELECT DATEDIFF(updated_at, created_at) as completion_days
		FROM reports
		WHERE created_at >= %s AND status = 'selesai' AND updated_at IS NOT NULL AND %s
		ORDER BY completion_days
	`, dateRange, dateFilter)

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		// Return 0 if error occurs in median calculation
		return 0
	}
	defer result.Close()

	var completionDays []float64
	for result.Next() {
		var days sql.NullFloat64
		err := result.Scan(&days)
		if err == nil && days.Valid {
			completionDays = append(completionDays, days.Float64)
		}
	}

	// Calculate median from the sorted array
	if len(completionDays) == 0 {
		return 0
	}

	mid := len(completionDays) / 2
	if len(completionDays)%2 == 0 {
		// Even number of elements - average of two middle values
		return (completionDays[mid-1] + completionDays[mid]) / 2
	} else {
		// Odd number of elements - middle value
		return completionDays[mid]
	}
}

// Helper function to get workload data
func getWorkloadData(dateRange string) []entity.WorkloadData {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var workloadData []entity.WorkloadData

	sqlQuery := fmt.Sprintf(`
		SELECT
			urgency_level,
			COUNT(*) as report_count,
			ROUND(
				(COUNT(*) * 100.0) / (SELECT COUNT(*) FROM reports WHERE created_at >= %s),
				2
			) as percentage,
			AVG(
				CASE
					WHEN status = 'selesai' AND updated_at IS NOT NULL
					THEN DATEDIFF(updated_at, created_at)
					ELSE NULL
				END
			) as avg_resolution_days
		FROM reports
		WHERE created_at >= %s
		GROUP BY urgency_level
		ORDER BY
			CASE urgency_level
				WHEN 'tinggi' THEN 1
				WHEN 'sedang' THEN 2
				WHEN 'rendah' THEN 3
			END
	`, dateRange, dateRange)

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - getWorkloadData, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var workload entity.WorkloadData
		var avgResolutionDays sql.NullFloat64

		err := result.Scan(
			&workload.UrgencyLevel,
			&workload.ReportCount,
			&workload.Percentage,
			&avgResolutionDays)

		if err != nil {
			panic("models - getWorkloadData, result.Scan " + err.Error())
		}

		if avgResolutionDays.Valid {
			workload.AvgResolutionDays = avgResolutionDays.Float64
		}

		workloadData = append(workloadData, workload)
	}

	return workloadData
}

// Helper function to calculate performance trends
func calculatePerformanceTrends(resolutionRates []entity.ResolutionRate, responseTimes []entity.ResponseTime, completionTimes []entity.CompletionTime) entity.PerformanceTrends {
	var trends entity.PerformanceTrends

	// Calculate resolution trend
	if len(resolutionRates) >= 2 {
		firstHalf := resolutionRates[:len(resolutionRates)/2]
		secondHalf := resolutionRates[len(resolutionRates)/2:]

		var firstHalfAvg, secondHalfAvg float64
		for _, rate := range firstHalf {
			firstHalfAvg += rate.ResolutionRate
		}
		firstHalfAvg /= float64(len(firstHalf))

		for _, rate := range secondHalf {
			secondHalfAvg += rate.ResolutionRate
		}
		secondHalfAvg /= float64(len(secondHalf))

		if secondHalfAvg > firstHalfAvg+5 {
			trends.ResolutionTrend = "improving"
		} else if secondHalfAvg < firstHalfAvg-5 {
			trends.ResolutionTrend = "declining"
		} else {
			trends.ResolutionTrend = "stable"
		}
	} else {
		trends.ResolutionTrend = "stable"
	}

	// Calculate response trend
	if len(responseTimes) >= 2 {
		firstHalf := responseTimes[:len(responseTimes)/2]
		secondHalf := responseTimes[len(responseTimes)/2:]

		var firstHalfAvg, secondHalfAvg float64
		for _, time := range firstHalf {
			firstHalfAvg += time.AvgResponseHours
		}
		firstHalfAvg /= float64(len(firstHalf))

		for _, time := range secondHalf {
			secondHalfAvg += time.AvgResponseHours
		}
		secondHalfAvg /= float64(len(secondHalf))

		if secondHalfAvg < firstHalfAvg-2 {
			trends.ResponseTrend = "improving"
		} else if secondHalfAvg > firstHalfAvg+2 {
			trends.ResponseTrend = "declining"
		} else {
			trends.ResponseTrend = "stable"
		}
	} else {
		trends.ResponseTrend = "stable"
	}

	// Calculate completion trend
	if len(completionTimes) >= 2 {
		firstHalf := completionTimes[:len(completionTimes)/2]
		secondHalf := completionTimes[len(completionTimes)/2:]

		var firstHalfAvg, secondHalfAvg float64
		for _, time := range firstHalf {
			firstHalfAvg += time.AvgCompletionDays
		}
		firstHalfAvg /= float64(len(firstHalf))

		for _, time := range secondHalf {
			secondHalfAvg += time.AvgCompletionDays
		}
		secondHalfAvg /= float64(len(secondHalf))

		if secondHalfAvg < firstHalfAvg-1 {
			trends.CompletionTrend = "improving"
		} else if secondHalfAvg > firstHalfAvg+1 {
			trends.CompletionTrend = "declining"
		} else {
			trends.CompletionTrend = "stable"
		}
	} else {
		trends.CompletionTrend = "stable"
	}

	// Calculate overall performance
	improvingCount := 0
	decliningCount := 0

	if trends.ResolutionTrend == "improving" {
		improvingCount++
	} else if trends.ResolutionTrend == "declining" {
		decliningCount++
	}

	if trends.ResponseTrend == "improving" {
		improvingCount++
	} else if trends.ResponseTrend == "declining" {
		decliningCount++
	}

	if trends.CompletionTrend == "improving" {
		improvingCount++
	} else if trends.CompletionTrend == "declining" {
		decliningCount++
	}

	if improvingCount >= 2 {
		trends.OverallPerformance = "excellent"
	} else if improvingCount == 1 && decliningCount == 0 {
		trends.OverallPerformance = "good"
	} else if decliningCount >= 2 {
		trends.OverallPerformance = "poor"
	} else {
		trends.OverallPerformance = "fair"
	}

	return trends
}

func GetReportMapVillage(districtId string) []entity.ReportMap {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var locations []entity.ReportMap

	sqlQuery := `
		SELECT
    		r.village_id AS id,
    		l.village_name AS name,
    		count(1) AS total,
			'' AS filter_key,
    		l.latitude AS latitude,
    		l.longitude AS longitude,
			SUM(CASE WHEN r.status IN ('baru', 'verifikasi') THEN 1 ELSE 0 END) AS menunggu,
			SUM(CASE WHEN r.status = 'proses' THEN 1 ELSE 0 END) AS proses,
			SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) AS selesai,
			SUM(CASE WHEN r.status = 'ditolak' THEN 1 ELSE 0 END) AS ditolak
		FROM reports r
		INNER JOIN villages l ON l.village_id = r.village_id
		WHERE l.district_id = ?
		GROUP BY r.village_id
		ORDER BY name;
	`
	result, err := db.QueryContext(ctx, sqlQuery, districtId)
	if err != nil {
		panic("models - GetReportMapVillage, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var location entity.ReportMap
		err := result.Scan(
			&location.Id,
			&location.Name,
			&location.Total,
			&location.FilterKey,
			&location.Latitude,
			&location.Longitude,
			&location.StatusBreakdown.Menunggu,
			&location.StatusBreakdown.Proses,
			&location.StatusBreakdown.Selesai,
			&location.StatusBreakdown.Ditolak)
		locations = append(locations, location)
		if err != nil {
			panic("models - GetReportMapVillage, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetReportMapVillage, result.Err " + err.Error())
	}
	return locations
}
