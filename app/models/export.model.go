package models

import (
	"context"
	"fmt"
	"strings"
	"tampayang-backend/app/models/entity"
	"tampayang-backend/core/database"
)

// ExportReportsData retrieves filtered report data for export
func ExportReportsData(startDate, endDate, status, categoryID, regencyID string) ([]entity.ExportReport, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var reports []entity.ExportReport
	var whereConditions []string
	var args []interface{}

	// Build WHERE conditions based on filters
	if startDate != "" && endDate != "" {
		whereConditions = append(whereConditions, "DATE(r.created_at) BETWEEN ? AND ?")
		args = append(args, startDate, endDate)
	} else if startDate != "" {
		whereConditions = append(whereConditions, "DATE(r.created_at) >= ?")
		args = append(args, startDate)
	} else if endDate != "" {
		whereConditions = append(whereConditions, "DATE(r.created_at) <= ?")
		args = append(args, endDate)
	}

	if status != "" {
		whereConditions = append(whereConditions, "LOWER(r.status) = LOWER(?)")
		args = append(args, status)
	}
	if categoryID != "" {
		whereConditions = append(whereConditions, "r.infrastructure_category_id = ?")
		args = append(args, categoryID)
	}
	if regencyID != "" {
		whereConditions = append(whereConditions, "r.regency_id = ?")
		args = append(args, regencyID)
	}

	// Build WHERE clause
	whereClause := ""
	if len(whereConditions) > 0 {
		whereClause = "WHERE " + strings.Join(whereConditions, " AND ")
	}

	sqlQuery := fmt.Sprintf(`
		SELECT
			r.report_id,
			r.report_number,
			r.reporter_name,
			r.reporter_phone,
			COALESCE(r.reporter_email, '') as reporter_email,
			COALESCE(ic.name, '') as infrastructure_category_name,
			COALESCE(dt.name, '') as damage_type_name,
			COALESCE(p.province_name, '') as province_name,
			COALESCE(reg.regency_name, '') as regency_name,
			COALESCE(d.district_name, '') as district_name,
			COALESCE(v.village_name, '') as village_name,
			COALESCE(r.location_detail, '') as location_detail,
			r.description,
			r.urgency_level,
			r.status,
			COALESCE(r.latitude, 0) as latitude,
			COALESCE(r.longitude, 0) as longitude,
			DATE_FORMAT(r.created_at, '%%Y-%%m-%%d %%H:%%i:%%s') as created_at,
			DATE_FORMAT(r.updated_at, '%%Y-%%m-%%d %%H:%%i:%%s') as updated_at
		FROM reports r
		INNER JOIN infrastructure_categories ic ON r.infrastructure_category_id = ic.infrastructure_category_id
		INNER JOIN damage_types dt ON r.damage_type_id = dt.damage_type_id
		INNER JOIN provinces p ON r.province_id = p.province_id
		INNER JOIN regencies reg ON r.regency_id = reg.regency_id
		INNER JOIN districts d ON r.district_id = d.district_id
		INNER JOIN villages v ON r.village_id = v.village_id
		%s
		ORDER BY r.report_number ASC, r.created_at DESC
	`, whereClause)

	result, err := db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query reports: %v", err)
	}
	defer result.Close()

	for result.Next() {
		var report entity.ExportReport
		err := result.Scan(
			&report.ReportID,
			&report.ReportNumber,
			&report.ReporterName,
			&report.ReporterPhone,
			&report.ReporterEmail,
			&report.InfrastructureCategoryName,
			&report.DamageTypeName,
			&report.ProvinceName,
			&report.RegencyName,
			&report.DistrictName,
			&report.VillageName,
			&report.LocationDetail,
			&report.Description,
			&report.UrgencyLevel,
			&report.Status,
			&report.Latitude,
			&report.Longitude,
			&report.CreatedAt,
			&report.UpdatedAt)

		if err != nil {
			return nil, fmt.Errorf("failed to scan report: %v", err)
		}

		reports = append(reports, report)
	}

	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("result error: %v", err)
	}

	return reports, nil
}

// ExportStatisticsData retrieves statistical data for export
func ExportStatisticsData(startDate, endDate string) (entity.ExportStatistics, error) {
	db := database.GetConnectionDB()
	defer db.Close()

	var stats entity.ExportStatistics
	var whereConditions []string
	var args []interface{}

	// Build WHERE conditions for date filtering
	if startDate != "" {
		whereConditions = append(whereConditions, "created_at >= ?")
		args = append(args, startDate)
	}
	if endDate != "" {
		whereConditions = append(whereConditions, "created_at <= ?")
		args = append(args, endDate+" 23:59:59")
	}

	whereClause := ""
	if len(whereConditions) > 0 {
		whereClause = "WHERE " + strings.Join(whereConditions, " AND ")
	}

	// Get monthly summary
	stats.MonthlySummary = getExportMonthlySummary(whereClause, args)

	// Get category breakdown
	stats.CategoryBreakdown = getExportCategoryBreakdown(whereClause, args)

	// Get status summary
	stats.StatusSummary = getExportStatusSummary(whereClause, args)

	// Get urgency level summary
	stats.UrgencyLevelSummary = getExportUrgencyLevelSummary(whereClause, args)

	// Get regional summary
	stats.RegionalSummary = getExportRegionalSummary(whereClause, args)

	return stats, nil
}

// Helper function to get monthly summary for export
func getExportMonthlySummary(whereClause string, args []interface{}) []entity.ExportMonthlySummary {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var monthlySummary []entity.ExportMonthlySummary

	sqlQuery := fmt.Sprintf(`
		SELECT 
			YEAR(created_at) as year,
			MONTH(created_at) as month,
			MONTHNAME(created_at) as month_name,
			COUNT(*) as total_reports,
			SUM(CASE WHEN status = 'selesai' THEN 1 ELSE 0 END) as completed_reports,
			ROUND((SUM(CASE WHEN status = 'selesai' THEN 1 ELSE 0 END) * 100.0) / COUNT(*), 2) as completion_rate
		FROM reports 
		%s
		GROUP BY YEAR(created_at), MONTH(created_at)
		ORDER BY year DESC, month DESC
		LIMIT 12
	`, whereClause)

	result, err := db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return monthlySummary
	}
	defer result.Close()

	for result.Next() {
		var summary entity.ExportMonthlySummary
		err := result.Scan(
			&summary.Year,
			&summary.Month,
			&summary.MonthName,
			&summary.TotalReports,
			&summary.CompletedReports,
			&summary.CompletionRate)

		if err == nil {
			monthlySummary = append(monthlySummary, summary)
		}
	}

	return monthlySummary
}

// Helper function to get category breakdown for export
func getExportCategoryBreakdown(whereClause string, args []interface{}) []entity.ExportCategoryBreakdown {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var categoryBreakdown []entity.ExportCategoryBreakdown

	sqlQuery := fmt.Sprintf(`
		SELECT 
			ic.name as category_name,
			COUNT(r.report_id) as total_reports,
			SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) as completed_reports,
			ROUND((SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) * 100.0) / COUNT(r.report_id), 2) as completion_rate
		FROM infrastructure_categories ic
		LEFT JOIN reports r ON ic.infrastructure_category_id = r.infrastructure_category_id
		%s
		GROUP BY ic.infrastructure_category_id, ic.name
		HAVING total_reports > 0
		ORDER BY total_reports DESC
	`, strings.Replace(whereClause, "created_at", "r.created_at", -1))

	result, err := db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return categoryBreakdown
	}
	defer result.Close()

	for result.Next() {
		var breakdown entity.ExportCategoryBreakdown
		err := result.Scan(
			&breakdown.CategoryName,
			&breakdown.TotalReports,
			&breakdown.CompletedReports,
			&breakdown.CompletionRate)

		if err == nil {
			categoryBreakdown = append(categoryBreakdown, breakdown)
		}
	}

	return categoryBreakdown
}

// Helper function to get status summary for export
func getExportStatusSummary(whereClause string, args []interface{}) []entity.ExportStatusSummary {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var statusSummary []entity.ExportStatusSummary

	sqlQuery := fmt.Sprintf(`
		SELECT 
			status,
			COUNT(*) as count,
			ROUND((COUNT(*) * 100.0) / (SELECT COUNT(*) FROM reports %s), 2) as percentage
		FROM reports 
		%s
		GROUP BY status
		ORDER BY count DESC
	`, whereClause, whereClause)

	// Duplicate args for the subquery
	queryArgs := append(args, args...)

	result, err := db.QueryContext(ctx, sqlQuery, queryArgs...)
	if err != nil {
		return statusSummary
	}
	defer result.Close()

	for result.Next() {
		var summary entity.ExportStatusSummary
		err := result.Scan(
			&summary.Status,
			&summary.Count,
			&summary.Percentage)

		if err == nil {
			statusSummary = append(statusSummary, summary)
		}
	}

	return statusSummary
}

// Helper function to get urgency level summary for export
func getExportUrgencyLevelSummary(whereClause string, args []interface{}) []entity.ExportUrgencyLevelSummary {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var urgencyLevelSummary []entity.ExportUrgencyLevelSummary

	sqlQuery := fmt.Sprintf(`
		SELECT 
			urgency_level,
			COUNT(*) as count,
			ROUND((COUNT(*) * 100.0) / (SELECT COUNT(*) FROM reports %s), 2) as percentage,
			AVG(CASE WHEN status = 'selesai' AND updated_at IS NOT NULL 
				THEN DATEDIFF(updated_at, created_at) ELSE NULL END) as avg_resolution_days
		FROM reports 
		%s
		GROUP BY urgency_level
		ORDER BY 
			CASE urgency_level 
				WHEN 'tinggi' THEN 1 
				WHEN 'sedang' THEN 2 
				WHEN 'rendah' THEN 3 
			END
	`, whereClause, whereClause)

	// Duplicate args for the subquery
	queryArgs := append(args, args...)

	result, err := db.QueryContext(ctx, sqlQuery, queryArgs...)
	if err != nil {
		return urgencyLevelSummary
	}
	defer result.Close()

	for result.Next() {
		var summary entity.ExportUrgencyLevelSummary
		var avgResolutionDays *float64
		err := result.Scan(
			&summary.UrgencyLevel,
			&summary.Count,
			&summary.Percentage,
			&avgResolutionDays)

		if err == nil {
			if avgResolutionDays != nil {
				summary.AvgResolutionDays = *avgResolutionDays
			}
			urgencyLevelSummary = append(urgencyLevelSummary, summary)
		}
	}

	return urgencyLevelSummary
}

// Helper function to get regional summary for export
func getExportRegionalSummary(whereClause string, args []interface{}) []entity.ExportRegionalSummary {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var regionalSummary []entity.ExportRegionalSummary

	sqlQuery := fmt.Sprintf(`
		SELECT 
			reg.regency_name,
			COUNT(r.report_id) as total_reports,
			SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) as completed_reports,
			ROUND((SUM(CASE WHEN r.status = 'selesai' THEN 1 ELSE 0 END) * 100.0) / COUNT(r.report_id), 2) as completion_rate
		FROM regencies reg
		LEFT JOIN reports r ON reg.regency_id = r.regency_id
		%s
		GROUP BY reg.regency_id, reg.regency_name
		HAVING total_reports > 0
		ORDER BY total_reports DESC
		LIMIT 20
	`, strings.Replace(whereClause, "created_at", "r.created_at", -1))

	result, err := db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return regionalSummary
	}
	defer result.Close()

	for result.Next() {
		var summary entity.ExportRegionalSummary
		err := result.Scan(
			&summary.RegencyName,
			&summary.TotalReports,
			&summary.CompletedReports,
			&summary.CompletionRate)

		if err == nil {
			regionalSummary = append(regionalSummary, summary)
		}
	}

	return regionalSummary
}

// TestExportQuery - Helper function to test the export query independently
func TestExportQuery() error {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	// Simple test query to verify basic functionality
	testQuery := `
		SELECT
			r.report_id,
			r.report_number,
			r.reporter_name,
			r.status,
			DATE_FORMAT(r.created_at, '%Y-%m-%d %H:%i:%s') as created_at
		FROM reports r
		LIMIT 1
	`

	fmt.Printf("DEBUG: Testing basic export query...\n")
	result, err := db.QueryContext(ctx, testQuery)
	if err != nil {
		fmt.Printf("ERROR: Basic test query failed: %v\n", err)
		return fmt.Errorf("basic test query failed: %v", err)
	}
	defer result.Close()

	if result.Next() {
		var reportID, reportNumber, reporterName, status, createdAt string
		err := result.Scan(&reportID, &reportNumber, &reporterName, &status, &createdAt)
		if err != nil {
			fmt.Printf("ERROR: Basic test scan failed: %v\n", err)
			return fmt.Errorf("basic test scan failed: %v", err)
		}
		fmt.Printf("SUCCESS: Basic test query returned: ID=%s, Number=%s, Name=%s, Status=%s, Created=%s\n",
			reportID, reportNumber, reporterName, status, createdAt)
	} else {
		fmt.Printf("INFO: No reports found in database\n")
	}

	return nil
}
