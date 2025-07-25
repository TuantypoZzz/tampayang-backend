package entity

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"regexp"
	"strings"
	"time"

	"tampayang-backend/core/database"
	globalFunction "tampayang-backend/core/functions"

	"github.com/google/uuid"
)

type Report struct {
	ReportId                 uuid.UUID               `json:"report_id"`
	ReportNumber             string                  `json:"report_number"`
	ReporterName             string                  `json:"reporter_name"`
	ReporterPhone            string                  `json:"reporter_phone"`
	ReporterEmail            string                  `json:"reporter_email"`
	InfrastructureCategoryId string                  `json:"infrastructure_category_id"`
	DamageTypeID             string                  `json:"damage_type_id"`
	ProviceID                string                  `json:"province_id"`
	RegencyID                string                  `json:"regency_id"`
	DistrictID               string                  `json:"district_id"`
	VillageID                string                  `json:"village_id"`
	LocationDetails          string                  `json:"location_detail"`
	Description              string                  `json:"description"`
	UrgencyLevel             string                  `json:"urgency_level"`
	Status                   string                  `json:"status"`
	Latitude                 *string                 `json:"latitude"`
	Longitude                *string                 `json:"longitude"`
	CreatedAt                time.Time               `json:"created_at"`
	UpdatedAt                time.Time               `json:"updated_at"`
	ReportImages             []*multipart.FileHeader `json:"-"`
}

type CheckStatus struct {
	ReportNumber               string `json:"report_number"`
	ReporterName               string `json:"reporter_name"`
	InfrastructureCategoryName string `json:"infrastructure_category_name"`
	DistrictName               string `json:"district_name"`
	VillageName                string `json:"village_name"`
	Status                     string `json:"status"`
	AdminNotes                 string `json:"admin_notes"`
	CreatedAt                  string `json:"created_at"`
}

type ValidationRule struct {
	Value      string
	IsRequired bool
	EmptyCode  string

	TableName    string
	ColumnName   string
	NotFoundCode string
}

type ReportPhoto struct {
	ReportPhotoID    uuid.UUID `json:"report_photo_id"`
	ReportID         uuid.UUID `json:"report_id"`
	Filename         string    `json:"filename"`
	OriginalFilename string    `json:"original_filename"`
	FilePath         string    `json:"file_path"`
	FileSize         int64     `json:"file_size"`
	MimeType         string    `json:"mime_type"`
	IsMain           bool      `json:"is_main"`
	UploadedAt       string    `json:"uploaded_at"`
}

type UrgencyReportRequest struct {
	ReportNumber   string `json:"report_number"`
	DamageTypeName string `json:"damage_type_name"`
	VillageName    string `json:"village_name"`
	UrgencyLevel   string `json:"urgency_level"`
}

type ManageReport struct {
	ReportNumber               string `json:"report_number"`
	CreatedAt                  string `json:"created_at"`
	ReporterName               string `json:"reporter_name"`
	InfrastructureCategoryName string `json:"infrastructure_category_name"`
	VillageName                string `json:"village_name"`
	Status                     string `json:"status"`
}

type DetailReport struct {
	ReportID                   string        `json:"report_id"`
	ReportNumber               string        `json:"report_number"`
	ReporterName               string        `json:"reporter_name"`
	ReporterPhone              string        `json:"reporter_phone"`
	InfrastructureCategoryName string        `json:"infrastructure_category_name"`
	DamageTypeName             string        `json:"damage_type_name"`
	DistrictName               string        `json:"district_name"`
	VillageName                string        `json:"village_name"`
	Description                string        `json:"description"`
	Status                     string        `json:"status"`
	UrgencyLevel               string        `json:"urgency_level"`
	Latitude                   float64       `json:"latitude"`
	Longitude                  float64       `json:"longitude"`
	CreatedAt                  string        `json:"created_at"`
	Photos                     []ReportPhoto `json:"photos"`
}

type CustomDate struct {
	time.Time
}

// UnmarshalJSON adalah metode kustom yang akan dipanggil oleh BodyParser
// untuk mengonversi string JSON menjadi tipe CustomDate.
func (cd *CustomDate) UnmarshalJSON(b []byte) (err error) {
	// Membersihkan tanda kutip dari string JSON (misal: "2025-12-12" -> 2025-12-12)
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		// Jika nilainya null atau string kosong, jangan lakukan apa-apa.
		return
	}

	// Coba parse dengan format "YYYY-MM-DD".
	// "2006-01-02" adalah layout standar Go untuk format ini.
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		// Jika gagal, coba parse dengan format RFC3339 sebagai fallback.
		t, err = time.Parse(time.RFC3339, s)
		if err != nil {
			return err
		}
	}
	cd.Time = t
	return
}

// UpdateReport adalah struct untuk menerima data update dari admin.
type UpdateReport struct {
	Status                  string      `json:"status"`
	PicName                 *string     `json:"pic_name"`
	PicPhone                *string     `json:"pic_phone"`
	AdminNotes              *string     `json:"admin_notes"`
	CompletionNotes         *string     `json:"completion_notes"`
	EstimatedCompletionDate *CustomDate `json:"estimated_completion_date"`
	CompletedAt             *CustomDate `json:"completed_at"`
}

func ValidateNewReport(newReport *Report) error {
	rules := []ValidationRule{
		{Value: newReport.ReporterName, IsRequired: true, EmptyCode: "rpt001"},
		{Value: newReport.ReporterPhone, IsRequired: true, EmptyCode: "rpt002"},
		{
			Value:        newReport.InfrastructureCategoryId,
			IsRequired:   true,
			EmptyCode:    "rpt003",
			TableName:    "infrastructure_categories",
			ColumnName:   "infrastructure_category_id",
			NotFoundCode: "rpt004",
		},
		{
			Value:        newReport.DamageTypeID,
			IsRequired:   true,
			EmptyCode:    "rpt005",
			TableName:    "damage_types",
			ColumnName:   "damage_type_id",
			NotFoundCode: "rpt006",
		},
		{
			Value:        newReport.ProviceID,
			IsRequired:   true,
			EmptyCode:    "rpt007",
			TableName:    "provinces",
			ColumnName:   "province_id",
			NotFoundCode: "rpt008",
		},
		{
			Value:        newReport.RegencyID,
			IsRequired:   true,
			EmptyCode:    "rpt009",
			TableName:    "regencies",
			ColumnName:   "regency_id",
			NotFoundCode: "rpt010",
		},
		{
			Value:        newReport.DistrictID,
			IsRequired:   true,
			EmptyCode:    "rpt011",
			TableName:    "districts",
			ColumnName:   "district_id",
			NotFoundCode: "rpt012",
		},
		{
			Value:        newReport.VillageID,
			IsRequired:   true,
			EmptyCode:    "rpt013",
			TableName:    "villages",
			ColumnName:   "village_id",
			NotFoundCode: "rpt014",
		},
	}

	// --- VALIDASI GAMBAR ---
	const (
		maxFiles    = 3
		maxFileSize = 5 * 1024 * 1024 // 5 MB
	)

	// Daftar tipe MIME yang diizinkan
	allowedMimeTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}

	// 1. Validasi jumlah gambar
	if len(newReport.ReportImages) > maxFiles {
		return errors.New("rpt017")
	}

	for _, file := range newReport.ReportImages {
		// 2. Validasi ukuran file
		if file.Size > maxFileSize {
			return errors.New("rpt018")
		}

		// 3. Validasi format file
		contentType := file.Header.Get("Content-Type")
		if !allowedMimeTypes[contentType] {
			return errors.New("rpt019")
		}
	}
	// --- AKHIR VALIDASI GAMBAR ---

	phoneRegex := `^(\+62|62|0)[0-9]{9,15}$`
	if !globalFunction.IsEmpty(newReport.ReporterPhone) {
		if isValid, _ := regexp.MatchString(phoneRegex, newReport.ReporterPhone); !isValid {
			return errors.New("err005")
		}
	}

	for _, rule := range rules {
		if rule.IsRequired && globalFunction.IsEmpty(rule.Value) {
			return errors.New(rule.EmptyCode)
		}

		if !globalFunction.IsEmpty(rule.Value) && rule.TableName != "" {
			if err := idReffExists(rule.TableName, rule.ColumnName, rule.Value); err != nil {
				return errors.New(rule.NotFoundCode)
			}
		}
	}

	return nil
}

func idReffExists(tableName string, idColumn string, id string) error {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var exists bool
	queryReff := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s = ? LIMIT 1)", tableName, idColumn)

	err := db.QueryRowContext(ctx, queryReff, id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking ID in %s table: %v", tableName, err)
	}

	if !exists {
		return fmt.Errorf("invalid ID: no record found in %s for %s = %s", tableName, idColumn, id)
	}

	return nil
}
