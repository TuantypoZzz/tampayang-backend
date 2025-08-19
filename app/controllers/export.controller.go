package controllers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"tampayang-backend/app/models"
	"tampayang-backend/app/models/entity"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/response"

	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
)

// ExportReports exports report data in various formats
func ExportReports(ctx *fiber.Ctx) error {
	// Get query parameters
	format := ctx.Query("format", "")
	startDate := ctx.Query("start_date", "")
	endDate := ctx.Query("end_date", "")
	status := ctx.Query("status", "")
	categoryID := ctx.Query("category_id", "")
	regencyID := ctx.Query("regency_id", "")

	// Validate required format parameter
	if format == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "format",
			"error": "Format parameter is required. Use csv, excel, or pdf",
		}))
	}

	// Validate format parameter
	validFormats := []string{"csv", "excel", "pdf"}
	isValidFormat := false
	for _, validFormat := range validFormats {
		if format == validFormat {
			isValidFormat = true
			break
		}
	}

	if !isValidFormat {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "format",
			"error": "Invalid format. Use csv, excel, or pdf",
		}))
	}

	// Validate date formats if provided
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

	// Log export request
	fmt.Printf("Export reports requested from IP: %s (format: %s, filters: start_date=%s, end_date=%s, status=%s, category_id=%s, regency_id=%s)\n",
		ctx.IP(), format, startDate, endDate, status, categoryID, regencyID)

	// Test basic database connectivity first
	if err := models.TestExportQuery(); err != nil {
		fmt.Printf("ERROR: Basic database test failed: %v\n", err)
		return response.ErrorResponse(ctx, globalFunction.GetMessage("database001", nil))
	}

	// Get report data
	reports, err := models.ExportReportsData(startDate, endDate, status, categoryID, regencyID)
	if err != nil {
		fmt.Printf("ERROR: Failed to export reports data: %v\n", err)
		return response.ErrorResponse(ctx, globalFunction.GetMessage("database001", nil))
	}

	// Generate filename with timestamp
	timestamp := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("tampayang-reports-%s", timestamp)

	// Generate file based on format
	switch format {
	case "csv":
		return exportReportsCSV(ctx, reports, filename)
	case "excel":
		return exportReportsExcel(ctx, reports, filename)
	case "pdf":
		return exportReportsPDF(ctx, reports, filename)
	default:
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "format",
			"error": "Unsupported format",
		}))
	}
}

// ExportStatistics exports statistical data in various formats
func ExportStatistics(ctx *fiber.Ctx) error {
	// Get query parameters
	format := ctx.Query("format", "")
	startDate := ctx.Query("start_date", "")
	endDate := ctx.Query("end_date", "")

	// Validate required format parameter
	if format == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "format",
			"error": "Format parameter is required. Use csv, excel, or pdf",
		}))
	}

	// Validate format parameter
	validFormats := []string{"csv", "excel", "pdf"}
	isValidFormat := false
	for _, validFormat := range validFormats {
		if format == validFormat {
			isValidFormat = true
			break
		}
	}

	if !isValidFormat {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "format",
			"error": "Invalid format. Use csv, excel, or pdf",
		}))
	}

	// Validate date formats if provided
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

	// Log export request
	fmt.Printf("Export statistics requested from IP: %s (format: %s, filters: start_date=%s, end_date=%s)\n",
		ctx.IP(), format, startDate, endDate)

	// Get statistics data
	statistics, err := models.ExportStatisticsData(startDate, endDate)
	if err != nil {
		fmt.Printf("ERROR: Failed to export statistics data: %v\n", err)
		return response.ErrorResponse(ctx, globalFunction.GetMessage("database001", nil))
	}

	// Generate filename with timestamp
	timestamp := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("tampayang-statistics-%s", timestamp)

	// Generate file based on format
	switch format {
	case "csv":
		return exportStatisticsCSV(ctx, statistics, filename)
	case "excel":
		return exportStatisticsExcel(ctx, statistics, filename)
	case "pdf":
		return exportStatisticsPDF(ctx, statistics, filename)
	default:
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "format",
			"error": "Unsupported format",
		}))
	}
}

// Helper function to export reports as CSV
func exportReportsCSV(ctx *fiber.Ctx, reports []entity.ExportReport, filename string) error {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write CSV header
	header := []string{
		"Report Number", "Reporter Name", "Reporter Phone", "Reporter Email",
		"Infrastructure Category", "Damage Type", "Province", "Regency",
		"District", "Village", "Location Detail", "Description",
		"Urgency Level", "Status", "Latitude", "Longitude",
		"Created At", "Updated At",
	}
	writer.Write(header)

	// Write data rows
	for _, report := range reports {
		row := []string{
			report.ReportNumber,
			report.ReporterName,
			report.ReporterPhone,
			report.ReporterEmail,
			report.InfrastructureCategoryName,
			report.DamageTypeName,
			report.ProvinceName,
			report.RegencyName,
			report.DistrictName,
			report.VillageName,
			report.LocationDetail,
			report.Description,
			report.UrgencyLevel,
			report.Status,
			fmt.Sprintf("%.6f", report.Latitude),
			fmt.Sprintf("%.6f", report.Longitude),
			report.CreatedAt,
			report.UpdatedAt,
		}
		writer.Write(row)
	}

	writer.Flush()

	// Set headers for file download
	ctx.Set("Content-Type", "text/csv")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.csv\"", filename))

	// Log successful export
	fmt.Printf("INFO: Successfully exported %d reports as CSV for IP: %s\n", len(reports), ctx.IP())

	return ctx.Send(buf.Bytes())
}

// Helper function to export reports as Excel
func exportReportsExcel(ctx *fiber.Ctx, reports []entity.ExportReport, filename string) error {
	f := excelize.NewFile()
	defer f.Close()

	sheetName := "Reports"
	f.SetSheetName("Sheet1", sheetName)

	// Set headers
	headers := []string{
		"Report Number", "Reporter Name", "Reporter Phone", "Reporter Email",
		"Infrastructure Category", "Damage Type", "Province", "Regency",
		"District", "Village", "Location Detail", "Description",
		"Urgency Level", "Status", "Latitude", "Longitude",
		"Created At", "Updated At",
	}

	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Set data rows
	for rowIndex, report := range reports {
		row := rowIndex + 2 // Start from row 2 (after header)
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), report.ReportNumber)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), report.ReporterName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), report.ReporterPhone)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), report.ReporterEmail)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), report.InfrastructureCategoryName)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), report.DamageTypeName)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), report.ProvinceName)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), report.RegencyName)
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), report.DistrictName)
		f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), report.VillageName)
		f.SetCellValue(sheetName, fmt.Sprintf("K%d", row), report.LocationDetail)
		f.SetCellValue(sheetName, fmt.Sprintf("L%d", row), report.Description)
		f.SetCellValue(sheetName, fmt.Sprintf("M%d", row), report.UrgencyLevel)
		f.SetCellValue(sheetName, fmt.Sprintf("N%d", row), report.Status)
		f.SetCellValue(sheetName, fmt.Sprintf("O%d", row), report.Latitude)
		f.SetCellValue(sheetName, fmt.Sprintf("P%d", row), report.Longitude)
		f.SetCellValue(sheetName, fmt.Sprintf("Q%d", row), report.CreatedAt)
		f.SetCellValue(sheetName, fmt.Sprintf("R%d", row), report.UpdatedAt)
	}

	// Generate Excel file
	buf, err := f.WriteToBuffer()
	if err != nil {
		return err
	}

	// Set headers for file download
	ctx.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.xlsx\"", filename))

	// Log successful export
	fmt.Printf("INFO: Successfully exported %d reports as Excel for IP: %s\n", len(reports), ctx.IP())

	return ctx.Send(buf.Bytes())
}

// Helper function to export reports as PDF
func exportReportsPDF(ctx *fiber.Ctx, reports []entity.ExportReport, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Title
	pdf.Cell(190, 10, "TAMPAYANG - Laporan Kerusakan Infrastruktur")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(190, 5, fmt.Sprintf("Tanggal Export: %s", time.Now().Format("02 January 2006")))
	pdf.Ln(5)
	pdf.Cell(190, 5, fmt.Sprintf("Total Laporan: %d", len(reports)))
	pdf.Ln(15)

	// Table headers
	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(25, 8, "No. Laporan")
	pdf.Cell(30, 8, "Pelapor")
	pdf.Cell(25, 8, "Kategori")
	pdf.Cell(25, 8, "Lokasi")
	pdf.Cell(20, 8, "Status")
	pdf.Cell(25, 8, "Tingkat Urgensi")
	pdf.Cell(30, 8, "Tanggal")
	pdf.Ln(8)

	// Table data
	pdf.SetFont("Arial", "", 7)
	for i, report := range reports {
		if i > 0 && i%35 == 0 { // Add new page every 35 rows
			pdf.AddPage()
			// Repeat headers
			pdf.SetFont("Arial", "B", 8)
			pdf.Cell(25, 8, "No. Laporan")
			pdf.Cell(30, 8, "Pelapor")
			pdf.Cell(25, 8, "Kategori")
			pdf.Cell(25, 8, "Lokasi")
			pdf.Cell(20, 8, "Status")
			pdf.Cell(25, 8, "Tingkat Urgensi")
			pdf.Cell(30, 8, "Tanggal")
			pdf.Ln(8)
			pdf.SetFont("Arial", "", 7)
		}

		// Truncate long text to fit in cells
		reportNumber := truncateString(report.ReportNumber, 12)
		reporterName := truncateString(report.ReporterName, 15)
		category := truncateString(report.InfrastructureCategoryName, 12)
		location := truncateString(report.RegencyName, 12)
		status := truncateString(report.Status, 10)
		urgency := truncateString(report.UrgencyLevel, 12)
		createdAt := truncateString(report.CreatedAt[:10], 12) // Only date part

		pdf.Cell(25, 6, reportNumber)
		pdf.Cell(30, 6, reporterName)
		pdf.Cell(25, 6, category)
		pdf.Cell(25, 6, location)
		pdf.Cell(20, 6, status)
		pdf.Cell(25, 6, urgency)
		pdf.Cell(30, 6, createdAt)
		pdf.Ln(6)
	}

	// Generate PDF
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return err
	}

	// Set headers for file download
	ctx.Set("Content-Type", "application/pdf")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.pdf\"", filename))

	// Log successful export
	fmt.Printf("INFO: Successfully exported %d reports as PDF for IP: %s\n", len(reports), ctx.IP())

	return ctx.Send(buf.Bytes())
}

// Helper function to truncate strings for PDF display
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// Helper function to export statistics as CSV
func exportStatisticsCSV(ctx *fiber.Ctx, statistics entity.ExportStatistics, filename string) error {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write Monthly Summary
	writer.Write([]string{"=== MONTHLY SUMMARY ==="})
	writer.Write([]string{"Year", "Month", "Month Name", "Total Reports", "Completed Reports", "Completion Rate (%)"})
	for _, monthly := range statistics.MonthlySummary {
		row := []string{
			strconv.Itoa(monthly.Year),
			strconv.Itoa(monthly.Month),
			monthly.MonthName,
			strconv.Itoa(monthly.TotalReports),
			strconv.Itoa(monthly.CompletedReports),
			fmt.Sprintf("%.2f", monthly.CompletionRate),
		}
		writer.Write(row)
	}

	writer.Write([]string{}) // Empty row

	// Write Category Breakdown
	writer.Write([]string{"=== CATEGORY BREAKDOWN ==="})
	writer.Write([]string{"Category Name", "Total Reports", "Completed Reports", "Completion Rate (%)"})
	for _, category := range statistics.CategoryBreakdown {
		row := []string{
			category.CategoryName,
			strconv.Itoa(category.TotalReports),
			strconv.Itoa(category.CompletedReports),
			fmt.Sprintf("%.2f", category.CompletionRate),
		}
		writer.Write(row)
	}

	writer.Write([]string{}) // Empty row

	// Write Status Summary
	writer.Write([]string{"=== STATUS SUMMARY ==="})
	writer.Write([]string{"Status", "Count", "Percentage (%)"})
	for _, status := range statistics.StatusSummary {
		row := []string{
			status.Status,
			strconv.Itoa(status.Count),
			fmt.Sprintf("%.2f", status.Percentage),
		}
		writer.Write(row)
	}

	writer.Write([]string{}) // Empty row

	// Write Urgency Level Summary
	writer.Write([]string{"=== URGENCY LEVEL SUMMARY ==="})
	writer.Write([]string{"Urgency Level", "Count", "Percentage (%)", "Avg Resolution Days"})
	for _, urgency := range statistics.UrgencyLevelSummary {
		row := []string{
			urgency.UrgencyLevel,
			strconv.Itoa(urgency.Count),
			fmt.Sprintf("%.2f", urgency.Percentage),
			fmt.Sprintf("%.2f", urgency.AvgResolutionDays),
		}
		writer.Write(row)
	}

	writer.Write([]string{}) // Empty row

	// Write Regional Summary
	writer.Write([]string{"=== REGIONAL SUMMARY ==="})
	writer.Write([]string{"Regency Name", "Total Reports", "Completed Reports", "Completion Rate (%)"})
	for _, regional := range statistics.RegionalSummary {
		row := []string{
			regional.RegencyName,
			strconv.Itoa(regional.TotalReports),
			strconv.Itoa(regional.CompletedReports),
			fmt.Sprintf("%.2f", regional.CompletionRate),
		}
		writer.Write(row)
	}

	writer.Flush()

	// Set headers for file download
	ctx.Set("Content-Type", "text/csv")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.csv\"", filename))

	// Log successful export
	fmt.Printf("INFO: Successfully exported statistics as CSV for IP: %s\n", ctx.IP())

	return ctx.Send(buf.Bytes())
}

// Helper function to export statistics as Excel
func exportStatisticsExcel(ctx *fiber.Ctx, statistics entity.ExportStatistics, filename string) error {
	f := excelize.NewFile()
	defer f.Close()

	// Monthly Summary Sheet
	monthlySheet := "Monthly Summary"
	f.NewSheet(monthlySheet)
	f.SetCellValue(monthlySheet, "A1", "Year")
	f.SetCellValue(monthlySheet, "B1", "Month")
	f.SetCellValue(monthlySheet, "C1", "Month Name")
	f.SetCellValue(monthlySheet, "D1", "Total Reports")
	f.SetCellValue(monthlySheet, "E1", "Completed Reports")
	f.SetCellValue(monthlySheet, "F1", "Completion Rate (%)")

	for i, monthly := range statistics.MonthlySummary {
		row := i + 2
		f.SetCellValue(monthlySheet, fmt.Sprintf("A%d", row), monthly.Year)
		f.SetCellValue(monthlySheet, fmt.Sprintf("B%d", row), monthly.Month)
		f.SetCellValue(monthlySheet, fmt.Sprintf("C%d", row), monthly.MonthName)
		f.SetCellValue(monthlySheet, fmt.Sprintf("D%d", row), monthly.TotalReports)
		f.SetCellValue(monthlySheet, fmt.Sprintf("E%d", row), monthly.CompletedReports)
		f.SetCellValue(monthlySheet, fmt.Sprintf("F%d", row), monthly.CompletionRate)
	}

	// Category Breakdown Sheet
	categorySheet := "Category Breakdown"
	f.NewSheet(categorySheet)
	f.SetCellValue(categorySheet, "A1", "Category Name")
	f.SetCellValue(categorySheet, "B1", "Total Reports")
	f.SetCellValue(categorySheet, "C1", "Completed Reports")
	f.SetCellValue(categorySheet, "D1", "Completion Rate (%)")

	for i, category := range statistics.CategoryBreakdown {
		row := i + 2
		f.SetCellValue(categorySheet, fmt.Sprintf("A%d", row), category.CategoryName)
		f.SetCellValue(categorySheet, fmt.Sprintf("B%d", row), category.TotalReports)
		f.SetCellValue(categorySheet, fmt.Sprintf("C%d", row), category.CompletedReports)
		f.SetCellValue(categorySheet, fmt.Sprintf("D%d", row), category.CompletionRate)
	}

	// Status Summary Sheet
	statusSheet := "Status Summary"
	f.NewSheet(statusSheet)
	f.SetCellValue(statusSheet, "A1", "Status")
	f.SetCellValue(statusSheet, "B1", "Count")
	f.SetCellValue(statusSheet, "C1", "Percentage (%)")

	for i, status := range statistics.StatusSummary {
		row := i + 2
		f.SetCellValue(statusSheet, fmt.Sprintf("A%d", row), status.Status)
		f.SetCellValue(statusSheet, fmt.Sprintf("B%d", row), status.Count)
		f.SetCellValue(statusSheet, fmt.Sprintf("C%d", row), status.Percentage)
	}

	// Urgency Level Summary Sheet
	urgencySheet := "Urgency Level Summary"
	f.NewSheet(urgencySheet)
	f.SetCellValue(urgencySheet, "A1", "Urgency Level")
	f.SetCellValue(urgencySheet, "B1", "Count")
	f.SetCellValue(urgencySheet, "C1", "Percentage (%)")
	f.SetCellValue(urgencySheet, "D1", "Avg Resolution Days")

	for i, urgency := range statistics.UrgencyLevelSummary {
		row := i + 2
		f.SetCellValue(urgencySheet, fmt.Sprintf("A%d", row), urgency.UrgencyLevel)
		f.SetCellValue(urgencySheet, fmt.Sprintf("B%d", row), urgency.Count)
		f.SetCellValue(urgencySheet, fmt.Sprintf("C%d", row), urgency.Percentage)
		f.SetCellValue(urgencySheet, fmt.Sprintf("D%d", row), urgency.AvgResolutionDays)
	}

	// Regional Summary Sheet
	regionalSheet := "Regional Summary"
	f.NewSheet(regionalSheet)
	f.SetCellValue(regionalSheet, "A1", "Regency Name")
	f.SetCellValue(regionalSheet, "B1", "Total Reports")
	f.SetCellValue(regionalSheet, "C1", "Completed Reports")
	f.SetCellValue(regionalSheet, "D1", "Completion Rate (%)")

	for i, regional := range statistics.RegionalSummary {
		row := i + 2
		f.SetCellValue(regionalSheet, fmt.Sprintf("A%d", row), regional.RegencyName)
		f.SetCellValue(regionalSheet, fmt.Sprintf("B%d", row), regional.TotalReports)
		f.SetCellValue(regionalSheet, fmt.Sprintf("C%d", row), regional.CompletedReports)
		f.SetCellValue(regionalSheet, fmt.Sprintf("D%d", row), regional.CompletionRate)
	}

	// Delete default sheet
	f.DeleteSheet("Sheet1")

	// Generate Excel file
	buf, err := f.WriteToBuffer()
	if err != nil {
		return err
	}

	// Set headers for file download
	ctx.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.xlsx\"", filename))

	// Log successful export
	fmt.Printf("INFO: Successfully exported statistics as Excel for IP: %s\n", ctx.IP())

	return ctx.Send(buf.Bytes())
}

// Helper function to export statistics as PDF
func exportStatisticsPDF(ctx *fiber.Ctx, statistics entity.ExportStatistics, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Title
	pdf.Cell(190, 10, "TAMPAYANG - Statistik Laporan")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(190, 5, fmt.Sprintf("Tanggal Export: %s", time.Now().Format("02 January 2006")))
	pdf.Ln(15)

	// Monthly Summary
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 8, "Ringkasan Bulanan")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(20, 6, "Tahun")
	pdf.Cell(20, 6, "Bulan")
	pdf.Cell(30, 6, "Total")
	pdf.Cell(30, 6, "Selesai")
	pdf.Cell(25, 6, "Rate (%)")
	pdf.Ln(6)

	pdf.SetFont("Arial", "", 8)
	for _, monthly := range statistics.MonthlySummary {
		pdf.Cell(20, 5, strconv.Itoa(monthly.Year))
		pdf.Cell(20, 5, monthly.MonthName[:3]) // Abbreviated month name
		pdf.Cell(30, 5, strconv.Itoa(monthly.TotalReports))
		pdf.Cell(30, 5, strconv.Itoa(monthly.CompletedReports))
		pdf.Cell(25, 5, fmt.Sprintf("%.1f", monthly.CompletionRate))
		pdf.Ln(5)
	}
	pdf.Ln(10)

	// Category Breakdown
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 8, "Breakdown Kategori")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(60, 6, "Kategori")
	pdf.Cell(25, 6, "Total")
	pdf.Cell(25, 6, "Selesai")
	pdf.Cell(25, 6, "Rate (%)")
	pdf.Ln(6)

	pdf.SetFont("Arial", "", 8)
	for _, category := range statistics.CategoryBreakdown {
		pdf.Cell(60, 5, truncateString(category.CategoryName, 25))
		pdf.Cell(25, 5, strconv.Itoa(category.TotalReports))
		pdf.Cell(25, 5, strconv.Itoa(category.CompletedReports))
		pdf.Cell(25, 5, fmt.Sprintf("%.1f", category.CompletionRate))
		pdf.Ln(5)
	}
	pdf.Ln(10)

	// Status Summary
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 8, "Ringkasan Status")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 8)
	pdf.Cell(40, 6, "Status")
	pdf.Cell(25, 6, "Jumlah")
	pdf.Cell(25, 6, "Persentase")
	pdf.Ln(6)

	pdf.SetFont("Arial", "", 8)
	for _, status := range statistics.StatusSummary {
		pdf.Cell(40, 5, status.Status)
		pdf.Cell(25, 5, strconv.Itoa(status.Count))
		pdf.Cell(25, 5, fmt.Sprintf("%.1f%%", status.Percentage))
		pdf.Ln(5)
	}

	// Generate PDF
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return err
	}

	// Set headers for file download
	ctx.Set("Content-Type", "application/pdf")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.pdf\"", filename))

	// Log successful export
	fmt.Printf("INFO: Successfully exported statistics as PDF for IP: %s\n", ctx.IP())

	return ctx.Send(buf.Bytes())
}
