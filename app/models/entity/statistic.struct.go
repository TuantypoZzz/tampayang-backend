package entity

type ReportSummary struct {
	TotalReport             int `json:"total_report"`
	TotalReportDone         int `json:"total_report_done"`
	TotalReportInProgress   int `json:"total_report_in_progress"`
	TotalReportWaiting      int `json:"total_report_waiting"`
	TotalReportNew          int `json:"total_report_new"`
	TotalReportVerification int `json:"total_report_verification"`
}

type ReportSummaryRequest struct {
	StartDate string `json:"start_date" validate:"dateformat"`
	EndDate   string `json:"end_date" validate:"dateformat"`
}

type ReportWeekly struct {
	Date  string `json:"date"`
	Day   string `json:"day"`
	Total int    `json:"total"`
}

type ReportMap struct {
	Id              string          `json:"id"`
	Name            string          `json:"name"`
	Total           int             `json:"total"`
	FilterKey       string          `json:"filter_key"`
	Latitude        string          `json:"latitude"`
	Longitude       string          `json:"longitude"`
	StatusBreakdown StatusBreakdown `json:"status_breakdown"`
}

type StatusBreakdown struct {
	Menunggu int `json:"menunggu"`
	Proses   int `json:"proses"`
	Selesai  int `json:"selesai"`
	Ditolak  int `json:"ditolak"`
}

type DashboardSummary struct {
	TotalLaporan        int     `json:"total_laporan"`
	TingkatPenyelesaian float64 `json:"tingkat_penyelesaian"`
}

type ReportMapRequest struct {
	ProvinceId string `query:"province_id"`
	RegencyId  string `query:"regency_id"`
	DistrictId string `query:"district_id"`
}
