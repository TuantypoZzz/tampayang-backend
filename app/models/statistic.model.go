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
