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

// MonthlyReport represents monthly statistics data
type MonthlyReport struct {
	MonthlyData []MonthlyData  `json:"monthly_data"`
	Summary     MonthlySummary `json:"summary"`
}

type MonthlyData struct {
	Year              int             `json:"year"`
	Month             int             `json:"month"`
	MonthName         string          `json:"month_name"`
	TotalReports      int             `json:"total_reports"`
	StatusBreakdown   StatusBreakdown `json:"status_breakdown"`
	GrowthPercentage  float64         `json:"growth_percentage"`
	AvgResolutionDays float64         `json:"avg_resolution_days"`
}

type MonthlySummary struct {
	TotalReportsLast12Months int     `json:"total_reports_last_12_months"`
	AvgMonthlyReports        float64 `json:"avg_monthly_reports"`
	BestMonth                string  `json:"best_month"`
	WorstMonth               string  `json:"worst_month"`
}

// CategoryBreakdown represents infrastructure category analytics
type CategoryBreakdown struct {
	Categories     []CategoryData   `json:"categories"`
	Summary        CategorySummary  `json:"summary"`
	GeographicData []GeographicData `json:"geographic_data"`
}

type CategoryData struct {
	CategoryID     string           `json:"category_id"`
	CategoryName   string           `json:"category_name"`
	CategoryCode   string           `json:"category_code"`
	TotalReports   int              `json:"total_reports"`
	CompletionRate float64          `json:"completion_rate"`
	DamageTypes    []DamageTypeData `json:"damage_types"`
}

type DamageTypeData struct {
	DamageTypeID   string  `json:"damage_type_id"`
	DamageTypeName string  `json:"damage_type_name"`
	DamageTypeCode string  `json:"damage_type_code"`
	ReportCount    int     `json:"report_count"`
	Percentage     float64 `json:"percentage"`
}

type CategorySummary struct {
	MostCommonCategory    string           `json:"most_common_category"`
	MostCommonDamageType  string           `json:"most_common_damage_type"`
	HighestCompletionRate string           `json:"highest_completion_rate"`
	LowestCompletionRate  string           `json:"lowest_completion_rate"`
	CommonDamageTypes     []DamageTypeData `json:"common_damage_types"`
}

type GeographicData struct {
	RegencyID    string  `json:"regency_id"`
	RegencyName  string  `json:"regency_name"`
	CategoryName string  `json:"category_name"`
	ReportCount  int     `json:"report_count"`
	Percentage   float64 `json:"percentage"`
}

// PerformanceChart represents performance metrics for dashboard charts
type PerformanceChart struct {
	Period          string            `json:"period"`
	ResolutionRates []ResolutionRate  `json:"resolution_rates"`
	ResponseTimes   []ResponseTime    `json:"response_times"`
	CompletionTimes []CompletionTime  `json:"completion_times"`
	Trends          PerformanceTrends `json:"trends"`
	WorkloadData    []WorkloadData    `json:"workload_data"`
}

type ResolutionRate struct {
	Date           string  `json:"date"`
	ResolvedCount  int     `json:"resolved_count"`
	TotalCount     int     `json:"total_count"`
	ResolutionRate float64 `json:"resolution_rate"`
}

type ResponseTime struct {
	Date                string  `json:"date"`
	AvgResponseHours    float64 `json:"avg_response_hours"`
	MedianResponseHours float64 `json:"median_response_hours"`
}

type CompletionTime struct {
	Date                 string  `json:"date"`
	AvgCompletionDays    float64 `json:"avg_completion_days"`
	MedianCompletionDays float64 `json:"median_completion_days"`
}

type PerformanceTrends struct {
	ResolutionTrend    string `json:"resolution_trend"`    // "improving", "declining", "stable"
	ResponseTrend      string `json:"response_trend"`      // "improving", "declining", "stable"
	CompletionTrend    string `json:"completion_trend"`    // "improving", "declining", "stable"
	OverallPerformance string `json:"overall_performance"` // "excellent", "good", "fair", "poor"
}

type WorkloadData struct {
	UrgencyLevel      string  `json:"urgency_level"`
	ReportCount       int     `json:"report_count"`
	Percentage        float64 `json:"percentage"`
	AvgResolutionDays float64 `json:"avg_resolution_days"`
}

// Export-related entities
type ExportReport struct {
	ReportID                   string  `json:"report_id"`
	ReportNumber               string  `json:"report_number"`
	ReporterName               string  `json:"reporter_name"`
	ReporterPhone              string  `json:"reporter_phone"`
	ReporterEmail              string  `json:"reporter_email"`
	InfrastructureCategoryName string  `json:"infrastructure_category_name"`
	DamageTypeName             string  `json:"damage_type_name"`
	ProvinceName               string  `json:"province_name"`
	RegencyName                string  `json:"regency_name"`
	DistrictName               string  `json:"district_name"`
	VillageName                string  `json:"village_name"`
	LocationDetail             string  `json:"location_detail"`
	Description                string  `json:"description"`
	UrgencyLevel               string  `json:"urgency_level"`
	Status                     string  `json:"status"`
	Latitude                   float64 `json:"latitude"`
	Longitude                  float64 `json:"longitude"`
	CreatedAt                  string  `json:"created_at"`
	UpdatedAt                  string  `json:"updated_at"`
}

type ExportStatistics struct {
	MonthlySummary      []ExportMonthlySummary      `json:"monthly_summary"`
	CategoryBreakdown   []ExportCategoryBreakdown   `json:"category_breakdown"`
	StatusSummary       []ExportStatusSummary       `json:"status_summary"`
	UrgencyLevelSummary []ExportUrgencyLevelSummary `json:"urgency_level_summary"`
	RegionalSummary     []ExportRegionalSummary     `json:"regional_summary"`
}

type ExportMonthlySummary struct {
	Year             int     `json:"year"`
	Month            int     `json:"month"`
	MonthName        string  `json:"month_name"`
	TotalReports     int     `json:"total_reports"`
	CompletedReports int     `json:"completed_reports"`
	CompletionRate   float64 `json:"completion_rate"`
}

type ExportCategoryBreakdown struct {
	CategoryName     string  `json:"category_name"`
	TotalReports     int     `json:"total_reports"`
	CompletedReports int     `json:"completed_reports"`
	CompletionRate   float64 `json:"completion_rate"`
}

type ExportStatusSummary struct {
	Status     string  `json:"status"`
	Count      int     `json:"count"`
	Percentage float64 `json:"percentage"`
}

type ExportUrgencyLevelSummary struct {
	UrgencyLevel      string  `json:"urgency_level"`
	Count             int     `json:"count"`
	Percentage        float64 `json:"percentage"`
	AvgResolutionDays float64 `json:"avg_resolution_days"`
}

type ExportRegionalSummary struct {
	RegencyName      string  `json:"regency_name"`
	TotalReports     int     `json:"total_reports"`
	CompletedReports int     `json:"completed_reports"`
	CompletionRate   float64 `json:"completion_rate"`
}

type ReportMapRequest struct {
	ProvinceId string `query:"province_id"`
	RegencyId  string `query:"regency_id"`
	DistrictId string `query:"district_id"`
}
