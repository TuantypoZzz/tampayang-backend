package models

import (
	"context"

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
			SUM(IF(status = 'done', 1, 0)) AS total_report_done, 
			SUM(IF(status = 'in_progress', 1, 0)) AS total_report_in_progress, 
			SUM(IF(status IN ('new', 'verification', 'rejected'), 1, 0)) AS total_report_waiting, 
			SUM(IF(status = 'new', 1, 0)) AS total_report_new, 
			SUM(IF(status = 'verification', 1, 0)) AS total_report_verification
		FROM reports
		WHERE created_at BEETWEN '?' AND '?' 
	`
	result, err := db.QueryContext(ctx, sqlQuery, request.StartDate, request.EndDate)
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
    		COUNT(1) AS total
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
		GROUP BY weekly.day_date, day_of_week
		ORDER BY weekly.day_date; 
	`
	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - GetWeeklyReport, db.QueryContext " + err.Error())
	}
	defer result.Close()

	if result.Next() {
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
	return summary
}
