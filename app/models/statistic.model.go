package models

import (
	"context"
	"fmt"

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
			COUNT(1) AS total_report, 
			COALESCE(SUM(CASE WHEN status = 'done' THEN 1 ELSE 0 END), 0) AS total_report_done,
			COALESCE(SUM(CASE WHEN status = 'in_progress' THEN 1 ELSE 0 END), 0) AS total_report_in_progress,
			COALESCE(SUM(CASE WHEN status IN ('new', 'verification', 'rejected') THEN 1 ELSE 0 END), 0) AS total_report_waiting,
			COALESCE(SUM(CASE WHEN status = 'new' THEN 1 ELSE 0 END), 0) AS total_report_new,
			COALESCE(SUM(CASE WHEN status = 'verification' THEN 1 ELSE 0 END), 0) AS total_report_verification
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

func GetWeeklyReport() []entity.WeeklyReport {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var summary []entity.WeeklyReport

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
		var weeklyReport entity.WeeklyReport
		err := result.Scan(
			&weeklyReport.Date,
			&weeklyReport.Day,
			&weeklyReport.Total)
		summary = append(summary, weeklyReport)
		if err != nil {
			panic("models - GetWeeklyReport, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetLovInfrastructureCategory, result.Err " + err.Error())
	}
	return summary
}
