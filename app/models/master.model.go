package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"tampayang-backend/app/models/entity"
	"tampayang-backend/core/database"

	"github.com/google/uuid"
)

// CreateLocation creates a new location based on type
func CreateLocation(req entity.CreateLocationRequest) (*entity.Location, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	// Generate UUID for new location
	locationID := uuid.New().String()
	now := time.Now()

	var err error
	switch req.Type {
	case entity.LocationTypeProvince:
		err = createProvince(ctx, db, locationID, req, now)
	case entity.LocationTypeRegency:
		err = createRegency(ctx, db, locationID, req, now)
	case entity.LocationTypeDistrict:
		err = createDistrict(ctx, db, locationID, req, now)
	case entity.LocationTypeVillage:
		err = createVillage(ctx, db, locationID, req, now)
	default:
		return nil, fmt.Errorf("invalid location type: %s", req.Type)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create location: %v", err)
	}

	// Return the created location
	return GetLocationByID(locationID, req.Type)
}

// GetLocationByID retrieves a location by ID and type
func GetLocationByID(id string, locationType entity.LocationType) (*entity.Location, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var location entity.Location
	var err error

	switch locationType {
	case entity.LocationTypeProvince:
		err = getProvinceByID(ctx, db, id, &location)
	case entity.LocationTypeRegency:
		err = getRegencyByID(ctx, db, id, &location)
	case entity.LocationTypeDistrict:
		err = getDistrictByID(ctx, db, id, &location)
	case entity.LocationTypeVillage:
		err = getVillageByID(ctx, db, id, &location)
	default:
		return nil, fmt.Errorf("invalid location type: %s", locationType)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("location not found")
		}
		return nil, fmt.Errorf("failed to get location: %v", err)
	}

	return &location, nil
}

// GetLocations retrieves locations with pagination and filtering
func GetLocations(req entity.LocationListRequest) (*entity.LocationListResponse, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	// Set defaults
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.SortBy == "" {
		req.SortBy = "name"
	}
	if req.SortOrder == "" {
		req.SortOrder = "asc"
	}

	var locations []entity.Location
	var total int64
	var err error

	fmt.Printf("DEBUG: GetLocations called with type: '%s', page: %d, limit: %d\n", req.Type, req.Page, req.Limit)

	switch req.Type {
	case entity.LocationTypeProvince:
		fmt.Printf("DEBUG: Calling getProvinces function\n")
		locations, total, err = getProvinces(ctx, db, req)
	case entity.LocationTypeRegency:
		fmt.Printf("DEBUG: Calling getRegencies function\n")
		locations, total, err = getRegencies(ctx, db, req)
	case entity.LocationTypeDistrict:
		fmt.Printf("DEBUG: Calling getDistricts function\n")
		locations, total, err = getDistricts(ctx, db, req)
	case entity.LocationTypeVillage:
		fmt.Printf("DEBUG: Calling getVillages function\n")
		locations, total, err = getVillages(ctx, db, req)
	case "":
		// If no type specified, return all locations (mixed types)
		fmt.Printf("DEBUG: Calling getAllLocations function\n")
		locations, total, err = getAllLocations(ctx, db, req)
	default:
		fmt.Printf("ERROR: Invalid location type: %s\n", req.Type)
		return nil, fmt.Errorf("invalid location type: %s", req.Type)
	}

	if err != nil {
		fmt.Printf("ERROR: Location query failed for type '%s': %v\n", req.Type, err)
		return nil, fmt.Errorf("failed to get locations: %v", err)
	}

	fmt.Printf("DEBUG: Successfully retrieved %d locations of type '%s'\n", len(locations), req.Type)

	// Calculate pagination
	totalPages := int((total + int64(req.Limit) - 1) / int64(req.Limit))

	return &entity.LocationListResponse{
		Data: locations,
		Pagination: entity.PaginationResponse{
			Page:       req.Page,
			Limit:      req.Limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

// UpdateLocation updates a location by ID and type
func UpdateLocation(id string, locationType entity.LocationType, req entity.UpdateLocationRequest) (*entity.Location, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var err error
	switch locationType {
	case entity.LocationTypeProvince:
		err = updateProvince(ctx, db, id, req)
	case entity.LocationTypeRegency:
		err = updateRegency(ctx, db, id, req)
	case entity.LocationTypeDistrict:
		err = updateDistrict(ctx, db, id, req)
	case entity.LocationTypeVillage:
		err = updateVillage(ctx, db, id, req)
	default:
		return nil, fmt.Errorf("invalid location type: %s", locationType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to update location: %v", err)
	}

	// Return the updated location
	return GetLocationByID(id, locationType)
}

// DeleteLocation deletes a location by ID and type
func DeleteLocation(id string, locationType entity.LocationType) error {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	// Check dependencies before deletion
	dependency, err := CheckLocationDependencies(id, locationType)
	if err != nil {
		return fmt.Errorf("failed to check dependencies: %v", err)
	}

	if dependency.HasReports {
		return fmt.Errorf("cannot delete location: has %d associated reports", dependency.ReportCount)
	}

	if dependency.HasChildren {
		return fmt.Errorf("cannot delete location: has %d child locations", dependency.ChildrenCount)
	}

	var deleteErr error
	switch locationType {
	case entity.LocationTypeProvince:
		deleteErr = deleteProvince(ctx, db, id)
	case entity.LocationTypeRegency:
		deleteErr = deleteRegency(ctx, db, id)
	case entity.LocationTypeDistrict:
		deleteErr = deleteDistrict(ctx, db, id)
	case entity.LocationTypeVillage:
		deleteErr = deleteVillage(ctx, db, id)
	default:
		return fmt.Errorf("invalid location type: %s", locationType)
	}

	if deleteErr != nil {
		return fmt.Errorf("failed to delete location: %v", deleteErr)
	}

	return nil
}

// CheckLocationDependencies checks if a location has dependencies
func CheckLocationDependencies(id string, locationType entity.LocationType) (*entity.LocationDependency, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	dependency := &entity.LocationDependency{}

	// Check for reports
	var reportCount int
	reportQuery := ""
	switch locationType {
	case entity.LocationTypeProvince:
		reportQuery = "SELECT COUNT(*) FROM reports WHERE province_id = ?"
	case entity.LocationTypeRegency:
		reportQuery = "SELECT COUNT(*) FROM reports WHERE regency_id = ?"
	case entity.LocationTypeDistrict:
		reportQuery = "SELECT COUNT(*) FROM reports WHERE district_id = ?"
	case entity.LocationTypeVillage:
		reportQuery = "SELECT COUNT(*) FROM reports WHERE village_id = ?"
	}

	err := db.QueryRowContext(ctx, reportQuery, id).Scan(&reportCount)
	if err != nil {
		return nil, fmt.Errorf("failed to check report dependencies: %v", err)
	}

	dependency.ReportCount = reportCount
	dependency.HasReports = reportCount > 0

	// Check for child locations
	var childCount int
	childQuery := ""
	switch locationType {
	case entity.LocationTypeProvince:
		childQuery = "SELECT COUNT(*) FROM regencies WHERE province_id = ?"
	case entity.LocationTypeRegency:
		childQuery = "SELECT COUNT(*) FROM districts WHERE regency_id = ?"
	case entity.LocationTypeDistrict:
		childQuery = "SELECT COUNT(*) FROM villages WHERE district_id = ?"
	case entity.LocationTypeVillage:
		// Villages don't have children
		childCount = 0
	}

	if childQuery != "" {
		err = db.QueryRowContext(ctx, childQuery, id).Scan(&childCount)
		if err != nil {
			return nil, fmt.Errorf("failed to check child dependencies: %v", err)
		}
	}

	dependency.ChildrenCount = childCount
	dependency.HasChildren = childCount > 0

	return dependency, nil
}

// GetLocationStats returns statistics about locations
func GetLocationStats() (*entity.LocationStats, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	stats := &entity.LocationStats{}

	// Count provinces
	err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM provinces").Scan(&stats.TotalProvinces)
	if err != nil {
		return nil, fmt.Errorf("failed to count provinces: %v", err)
	}

	// Count regencies
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM regencies").Scan(&stats.TotalRegencies)
	if err != nil {
		return nil, fmt.Errorf("failed to count regencies: %v", err)
	}

	// Count districts
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM districts").Scan(&stats.TotalDistricts)
	if err != nil {
		return nil, fmt.Errorf("failed to count districts: %v", err)
	}

	// Count villages
	err = db.QueryRowContext(ctx, "SELECT COUNT(*) FROM villages").Scan(&stats.TotalVillages)
	if err != nil {
		return nil, fmt.Errorf("failed to count villages: %v", err)
	}

	return stats, nil
}

// Helper functions for Province operations
func createProvince(ctx context.Context, db *sql.DB, id string, req entity.CreateLocationRequest, now time.Time) error {
	query := `
		INSERT INTO provinces (province_id, province_name, province_code, latitude, longitude, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, id, req.Name, req.Code, req.Latitude, req.Longitude, now)
	return err
}

func getProvinceByID(ctx context.Context, db *sql.DB, id string, location *entity.Location) error {
	query := `
		SELECT province_id, province_name, province_code, latitude, longitude, created_at, updated_at
		FROM provinces WHERE province_id = ?
	`
	var latitude, longitude sql.NullFloat64
	var createdAtStr, updatedAtStr sql.NullString // Handle DATETIME as strings

	fmt.Printf("DEBUG: Get province by ID query: %s\n", query)
	fmt.Printf("DEBUG: Province ID parameter: %s\n", id)

	err := db.QueryRowContext(ctx, query, id).Scan(
		&location.ID, &location.Name, &location.Code, &latitude, &longitude,
		&createdAtStr, &updatedAtStr)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("DEBUG: Province not found with ID: %s\n", id)
			return fmt.Errorf("province not found")
		}
		fmt.Printf("ERROR: Failed to get province by ID: %v\n", err)
		return fmt.Errorf("failed to get province: %v", err)
	}

	location.Type = entity.LocationTypeProvince
	location.IsActive = true

	if latitude.Valid {
		location.Latitude = &latitude.Float64
	}
	if longitude.Valid {
		location.Longitude = &longitude.Float64
	}

	// Convert string datetime to time.Time for created_at
	if createdAtStr.Valid && createdAtStr.String != "" {
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
			location.CreatedAt = parsedTime
		} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", createdAtStr.String); err == nil {
			location.CreatedAt = parsedTime
		} else {
			fmt.Printf("WARN: Failed to parse province created_at: %s, error: %v\n", createdAtStr.String, err)
			location.CreatedAt = time.Now() // Fallback
		}
	} else {
		location.CreatedAt = time.Now() // Fallback for null/empty
	}

	// Convert string datetime to time.Time for updated_at
	if updatedAtStr.Valid && updatedAtStr.String != "" {
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
			location.UpdatedAt = &parsedTime
		} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", updatedAtStr.String); err == nil {
			location.UpdatedAt = &parsedTime
		} else {
			fmt.Printf("WARN: Failed to parse province updated_at: %s, error: %v\n", updatedAtStr.String, err)
		}
	}

	fmt.Printf("DEBUG: Successfully retrieved province: %s (%s)\n", location.Name, location.ID)
	return nil
}

func getProvinces(ctx context.Context, db *sql.DB, req entity.LocationListRequest) ([]entity.Location, int64, error) {
	var locations []entity.Location
	var whereConditions []string
	var args []interface{}

	// Validate and sanitize sort parameters
	validSortColumns := map[string]string{
		"name":       "province_name",
		"code":       "province_code",
		"created_at": "created_at",
	}

	sortColumn, exists := validSortColumns[req.SortBy]
	if !exists {
		sortColumn = "province_name" // Default sort column
	}

	sortOrder := "ASC"
	if strings.ToUpper(req.SortOrder) == "DESC" {
		sortOrder = "DESC"
	}

	// Build WHERE conditions
	if req.Search != "" {
		whereConditions = append(whereConditions, "(province_name LIKE ? OR province_code LIKE ?)")
		searchTerm := "%" + req.Search + "%"
		args = append(args, searchTerm, searchTerm)
	}

	whereClause := ""
	if len(whereConditions) > 0 {
		whereClause = "WHERE " + strings.Join(whereConditions, " AND ")
	}

	// Count total records with detailed error logging
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM provinces %s", whereClause)
	fmt.Printf("DEBUG: Count query: %s\n", countQuery)
	fmt.Printf("DEBUG: Count args: %v\n", args)

	var total int64
	err := db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		fmt.Printf("ERROR: Count query failed: %v\n", err)
		return nil, 0, fmt.Errorf("failed to count provinces: %v", err)
	}

	fmt.Printf("DEBUG: Total provinces found: %d\n", total)

	// Get paginated data with safe SQL construction
	offset := (req.Page - 1) * req.Limit
	dataQuery := fmt.Sprintf(`
		SELECT province_id, province_name, province_code, latitude, longitude, created_at, updated_at
		FROM provinces %s
		ORDER BY %s %s
		LIMIT ? OFFSET ?
	`, whereClause, sortColumn, sortOrder)

	dataArgs := append(args, req.Limit, offset)
	fmt.Printf("DEBUG: Data query: %s\n", dataQuery)
	fmt.Printf("DEBUG: Data args: %v\n", dataArgs)

	rows, err := db.QueryContext(ctx, dataQuery, dataArgs...)
	if err != nil {
		fmt.Printf("ERROR: Data query failed: %v\n", err)
		return nil, 0, fmt.Errorf("failed to query provinces: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var location entity.Location
		var latitude, longitude sql.NullFloat64
		var createdAtStr, updatedAtStr sql.NullString // Handle DATETIME as strings

		err := rows.Scan(&location.ID, &location.Name, &location.Code, &latitude, &longitude,
			&createdAtStr, &updatedAtStr)
		if err != nil {
			fmt.Printf("ERROR: Province row scan failed: %v\n", err)
			return nil, 0, fmt.Errorf("failed to scan province row: %v", err)
		}

		location.Type = entity.LocationTypeProvince
		location.IsActive = true

		if latitude.Valid {
			location.Latitude = &latitude.Float64
		}
		if longitude.Valid {
			location.Longitude = &longitude.Float64
		}

		// Convert string datetime to time.Time for created_at
		if createdAtStr.Valid && createdAtStr.String != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				location.CreatedAt = parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", createdAtStr.String); err == nil {
				location.CreatedAt = parsedTime
			} else {
				fmt.Printf("WARN: Failed to parse province created_at: %s, error: %v\n", createdAtStr.String, err)
				location.CreatedAt = time.Now() // Fallback
			}
		} else {
			location.CreatedAt = time.Now() // Fallback for null/empty
		}

		// Convert string datetime to time.Time for updated_at
		if updatedAtStr.Valid && updatedAtStr.String != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
				location.UpdatedAt = &parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", updatedAtStr.String); err == nil {
				location.UpdatedAt = &parsedTime
			} else {
				fmt.Printf("WARN: Failed to parse province updated_at: %s, error: %v\n", updatedAtStr.String, err)
			}
		}

		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("ERROR: Rows iteration error: %v\n", err)
		return nil, 0, fmt.Errorf("error iterating province rows: %v", err)
	}

	fmt.Printf("DEBUG: Successfully retrieved %d provinces\n", len(locations))
	return locations, total, nil
}

func updateProvince(ctx context.Context, db *sql.DB, id string, req entity.UpdateLocationRequest) error {
	var setParts []string
	var args []interface{}

	if req.Name != nil {
		setParts = append(setParts, "province_name = ?")
		args = append(args, *req.Name)
	}
	if req.Code != nil {
		setParts = append(setParts, "province_code = ?")
		args = append(args, *req.Code)
	}
	if req.Latitude != nil {
		setParts = append(setParts, "latitude = ?")
		args = append(args, *req.Latitude)
	}
	if req.Longitude != nil {
		setParts = append(setParts, "longitude = ?")
		args = append(args, *req.Longitude)
	}

	if len(setParts) == 0 {
		return fmt.Errorf("no fields to update")
	}

	setParts = append(setParts, "updated_at = ?")
	args = append(args, time.Now())
	args = append(args, id)

	query := fmt.Sprintf("UPDATE provinces SET %s WHERE province_id = ?", strings.Join(setParts, ", "))
	_, err := db.ExecContext(ctx, query, args...)
	return err
}

func deleteProvince(ctx context.Context, db *sql.DB, id string) error {
	query := "DELETE FROM provinces WHERE province_id = ?"
	_, err := db.ExecContext(ctx, query, id)
	return err
}

// Helper functions for Regency operations
func createRegency(ctx context.Context, db *sql.DB, id string, req entity.CreateLocationRequest, now time.Time) error {
	if req.ParentID == nil {
		return fmt.Errorf("parent_id (province_id) is required for regency")
	}
	if req.RegencyType == nil {
		return fmt.Errorf("regency_type is required for regency")
	}

	query := `
		INSERT INTO regencies (regency_id, province_id, regency_name, regency_code, regency_type, latitude, longitude, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, id, *req.ParentID, req.Name, req.Code, *req.RegencyType, req.Latitude, req.Longitude, now)
	return err
}

func getRegencyByID(ctx context.Context, db *sql.DB, id string, location *entity.Location) error {
	query := `
		SELECT r.regency_id, r.province_id, r.regency_name, r.regency_code, r.regency_type,
		       r.latitude, r.longitude, r.created_at, r.updated_at,
		       p.province_name, p.province_code
		FROM regencies r
		LEFT JOIN provinces p ON r.province_id = p.province_id
		WHERE r.regency_id = ?
	`
	var latitude, longitude sql.NullFloat64
	var provinceName, provinceCode sql.NullString
	var createdAtStr, updatedAtStr sql.NullString // Handle DATETIME as strings

	fmt.Printf("DEBUG: Get regency by ID query: %s\n", query)
	fmt.Printf("DEBUG: Regency ID parameter: %s\n", id)

	err := db.QueryRowContext(ctx, query, id).Scan(
		&location.ID, &location.ParentID, &location.Name, &location.Code, &location.RegencyType,
		&latitude, &longitude, &createdAtStr, &updatedAtStr,
		&provinceName, &provinceCode)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("DEBUG: Regency not found with ID: %s\n", id)
			return fmt.Errorf("regency not found")
		}
		fmt.Printf("ERROR: Failed to get regency by ID: %v\n", err)
		return fmt.Errorf("failed to get regency: %v", err)
	}

	location.Type = entity.LocationTypeRegency
	location.IsActive = true

	if latitude.Valid {
		location.Latitude = &latitude.Float64
	}
	if longitude.Valid {
		location.Longitude = &longitude.Float64
	}

	// Convert string datetime to time.Time for created_at
	if createdAtStr.Valid && createdAtStr.String != "" {
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
			location.CreatedAt = parsedTime
		} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", createdAtStr.String); err == nil {
			location.CreatedAt = parsedTime
		} else {
			fmt.Printf("WARN: Failed to parse regency created_at: %s, error: %v\n", createdAtStr.String, err)
			location.CreatedAt = time.Now() // Fallback
		}
	} else {
		location.CreatedAt = time.Now() // Fallback for null/empty
	}

	// Convert string datetime to time.Time for updated_at
	if updatedAtStr.Valid && updatedAtStr.String != "" {
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
			location.UpdatedAt = &parsedTime
		} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", updatedAtStr.String); err == nil {
			location.UpdatedAt = &parsedTime
		} else {
			fmt.Printf("WARN: Failed to parse regency updated_at: %s, error: %v\n", updatedAtStr.String, err)
		}
	}

	// Add parent information
	if provinceName.Valid && provinceCode.Valid {
		location.Parent = &entity.Location{
			ID:   *location.ParentID,
			Name: provinceName.String,
			Code: provinceCode.String,
			Type: entity.LocationTypeProvince,
		}
	}

	fmt.Printf("DEBUG: Successfully retrieved regency: %s (%s)\n", location.Name, location.ID)
	return nil
}

func getRegencies(ctx context.Context, db *sql.DB, req entity.LocationListRequest) ([]entity.Location, int64, error) {
	var locations []entity.Location
	var whereConditions []string
	var args []interface{}

	// Validate and sanitize sort parameters for regencies
	validSortColumns := map[string]string{
		"name":       "regency_name",
		"code":       "regency_code",
		"created_at": "created_at",
	}

	sortColumn, exists := validSortColumns[req.SortBy]
	if !exists {
		sortColumn = "regency_name" // Default sort column
	}

	sortOrder := "ASC"
	if strings.ToUpper(req.SortOrder) == "DESC" {
		sortOrder = "DESC"
	}

	// Build WHERE conditions
	if req.Search != "" {
		whereConditions = append(whereConditions, "(r.regency_name LIKE ? OR r.regency_code LIKE ?)")
		searchTerm := "%" + req.Search + "%"
		args = append(args, searchTerm, searchTerm)
	}
	if req.ParentID != "" {
		whereConditions = append(whereConditions, "r.province_id = ?")
		args = append(args, req.ParentID)
	}

	whereClause := ""
	if len(whereConditions) > 0 {
		whereClause = "WHERE " + strings.Join(whereConditions, " AND ")
	}

	// Count total records with detailed error logging
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM regencies r %s", whereClause)
	fmt.Printf("DEBUG: Regency count query: %s\n", countQuery)
	fmt.Printf("DEBUG: Regency count args: %v\n", args)

	var total int64
	err := db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		fmt.Printf("ERROR: Regency count query failed: %v\n", err)
		return nil, 0, fmt.Errorf("failed to count regencies: %v", err)
	}

	fmt.Printf("DEBUG: Total regencies found: %d\n", total)

	// Get paginated data with safe SQL construction
	offset := (req.Page - 1) * req.Limit
	dataQuery := fmt.Sprintf(`
		SELECT r.regency_id, r.province_id, r.regency_name, r.regency_code, r.regency_type,
		       r.latitude, r.longitude, r.created_at, r.updated_at,
		       p.province_name, p.province_code
		FROM regencies r
		LEFT JOIN provinces p ON r.province_id = p.province_id
		%s
		ORDER BY r.%s %s
		LIMIT ? OFFSET ?
	`, whereClause, sortColumn, sortOrder)

	dataArgs := append(args, req.Limit, offset)
	fmt.Printf("DEBUG: Regency data query: %s\n", dataQuery)
	fmt.Printf("DEBUG: Regency data args: %v\n", dataArgs)

	rows, err := db.QueryContext(ctx, dataQuery, dataArgs...)
	if err != nil {
		fmt.Printf("ERROR: Regency data query failed: %v\n", err)
		return nil, 0, fmt.Errorf("failed to query regencies: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var location entity.Location
		var latitude, longitude sql.NullFloat64
		var provinceName, provinceCode sql.NullString
		var createdAtStr, updatedAtStr sql.NullString // Handle DATETIME as strings

		err := rows.Scan(&location.ID, &location.ParentID, &location.Name, &location.Code, &location.RegencyType,
			&latitude, &longitude, &createdAtStr, &updatedAtStr,
			&provinceName, &provinceCode)
		if err != nil {
			fmt.Printf("ERROR: Regency row scan failed: %v\n", err)
			return nil, 0, fmt.Errorf("failed to scan regency row: %v", err)
		}

		location.Type = entity.LocationTypeRegency
		location.IsActive = true

		if latitude.Valid {
			location.Latitude = &latitude.Float64
		}
		if longitude.Valid {
			location.Longitude = &longitude.Float64
		}

		// Convert string datetime to time.Time for created_at
		if createdAtStr.Valid && createdAtStr.String != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				location.CreatedAt = parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", createdAtStr.String); err == nil {
				location.CreatedAt = parsedTime
			} else {
				fmt.Printf("WARN: Failed to parse regency created_at: %s, error: %v\n", createdAtStr.String, err)
				location.CreatedAt = time.Now() // Fallback
			}
		} else {
			location.CreatedAt = time.Now() // Fallback for null/empty
		}

		// Convert string datetime to time.Time for updated_at
		if updatedAtStr.Valid && updatedAtStr.String != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
				location.UpdatedAt = &parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", updatedAtStr.String); err == nil {
				location.UpdatedAt = &parsedTime
			} else {
				fmt.Printf("WARN: Failed to parse regency updated_at: %s, error: %v\n", updatedAtStr.String, err)
			}
		}

		// Add parent information
		if provinceName.Valid && provinceCode.Valid {
			location.Parent = &entity.Location{
				ID:   *location.ParentID,
				Name: provinceName.String,
				Code: provinceCode.String,
				Type: entity.LocationTypeProvince,
			}
		}

		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("ERROR: Regency rows iteration error: %v\n", err)
		return nil, 0, fmt.Errorf("error iterating regency rows: %v", err)
	}

	fmt.Printf("DEBUG: Successfully retrieved %d regencies\n", len(locations))
	return locations, total, nil
}

func updateRegency(ctx context.Context, db *sql.DB, id string, req entity.UpdateLocationRequest) error {
	var setParts []string
	var args []interface{}

	if req.Name != nil {
		setParts = append(setParts, "regency_name = ?")
		args = append(args, *req.Name)
	}
	if req.Code != nil {
		setParts = append(setParts, "regency_code = ?")
		args = append(args, *req.Code)
	}
	if req.ParentID != nil {
		setParts = append(setParts, "province_id = ?")
		args = append(args, *req.ParentID)
	}
	if req.RegencyType != nil {
		setParts = append(setParts, "regency_type = ?")
		args = append(args, *req.RegencyType)
	}
	if req.Latitude != nil {
		setParts = append(setParts, "latitude = ?")
		args = append(args, *req.Latitude)
	}
	if req.Longitude != nil {
		setParts = append(setParts, "longitude = ?")
		args = append(args, *req.Longitude)
	}

	if len(setParts) == 0 {
		return fmt.Errorf("no fields to update")
	}

	setParts = append(setParts, "updated_at = ?")
	args = append(args, time.Now())
	args = append(args, id)

	query := fmt.Sprintf("UPDATE regencies SET %s WHERE regency_id = ?", strings.Join(setParts, ", "))
	_, err := db.ExecContext(ctx, query, args...)
	return err
}

func deleteRegency(ctx context.Context, db *sql.DB, id string) error {
	query := "DELETE FROM regencies WHERE regency_id = ?"
	_, err := db.ExecContext(ctx, query, id)
	return err
}

// Helper functions for District operations
func createDistrict(ctx context.Context, db *sql.DB, id string, req entity.CreateLocationRequest, now time.Time) error {
	if req.ParentID == nil {
		return fmt.Errorf("parent_id (regency_id) is required for district")
	}

	query := `
		INSERT INTO districts (district_id, regency_id, district_name, district_code, latitude, longitude, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, id, *req.ParentID, req.Name, req.Code, req.Latitude, req.Longitude, now)
	return err
}

func getDistrictByID(ctx context.Context, db *sql.DB, id string, location *entity.Location) error {
	query := `
		SELECT d.district_id, d.regency_id, d.district_name, d.district_code,
		       d.latitude, d.longitude, d.created_at, d.updated_at,
		       r.regency_name, r.regency_code, r.province_id,
		       p.province_name, p.province_code
		FROM districts d
		LEFT JOIN regencies r ON d.regency_id = r.regency_id
		LEFT JOIN provinces p ON r.province_id = p.province_id
		WHERE d.district_id = ?
	`
	var updatedAt sql.NullTime
	var latitude, longitude sql.NullFloat64
	var regencyName, regencyCode, provinceID sql.NullString
	var provinceName, provinceCode sql.NullString

	err := db.QueryRowContext(ctx, query, id).Scan(
		&location.ID, &location.ParentID, &location.Name, &location.Code,
		&latitude, &longitude, &location.CreatedAt, &updatedAt,
		&regencyName, &regencyCode, &provinceID,
		&provinceName, &provinceCode)

	if err != nil {
		return err
	}

	location.Type = entity.LocationTypeDistrict
	location.IsActive = true
	if latitude.Valid {
		location.Latitude = &latitude.Float64
	}
	if longitude.Valid {
		location.Longitude = &longitude.Float64
	}
	if updatedAt.Valid {
		location.UpdatedAt = &updatedAt.Time
	}

	// Add parent information
	if regencyName.Valid && regencyCode.Valid {
		parent := &entity.Location{
			ID:   *location.ParentID,
			Name: regencyName.String,
			Code: regencyCode.String,
			Type: entity.LocationTypeRegency,
		}

		// Add grandparent (province) information
		if provinceName.Valid && provinceCode.Valid && provinceID.Valid {
			parent.Parent = &entity.Location{
				ID:   provinceID.String,
				Name: provinceName.String,
				Code: provinceCode.String,
				Type: entity.LocationTypeProvince,
			}
		}

		location.Parent = parent
	}

	return nil
}

func getDistricts(ctx context.Context, db *sql.DB, req entity.LocationListRequest) ([]entity.Location, int64, error) {
	var locations []entity.Location
	var whereConditions []string
	var args []interface{}

	// Validate and sanitize sort parameters for districts
	validSortColumns := map[string]string{
		"name":       "district_name",
		"code":       "district_code",
		"created_at": "created_at",
	}

	sortColumn, exists := validSortColumns[req.SortBy]
	if !exists {
		sortColumn = "district_name" // Default sort column
	}

	sortOrder := "ASC"
	if strings.ToUpper(req.SortOrder) == "DESC" {
		sortOrder = "DESC"
	}

	// Build WHERE conditions
	if req.Search != "" {
		whereConditions = append(whereConditions, "(d.district_name LIKE ? OR d.district_code LIKE ?)")
		searchTerm := "%" + req.Search + "%"
		args = append(args, searchTerm, searchTerm)
	}
	if req.ParentID != "" {
		whereConditions = append(whereConditions, "d.regency_id = ?")
		args = append(args, req.ParentID)
	}

	whereClause := ""
	if len(whereConditions) > 0 {
		whereClause = "WHERE " + strings.Join(whereConditions, " AND ")
	}

	// Count total records with detailed error logging
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM districts d %s", whereClause)
	fmt.Printf("DEBUG: District count query: %s\n", countQuery)
	fmt.Printf("DEBUG: District count args: %v\n", args)

	var total int64
	err := db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		fmt.Printf("ERROR: District count query failed: %v\n", err)
		return nil, 0, fmt.Errorf("failed to count districts: %v", err)
	}

	fmt.Printf("DEBUG: Total districts found: %d\n", total)

	// Get paginated data with safe SQL construction
	offset := (req.Page - 1) * req.Limit
	dataQuery := fmt.Sprintf(`
		SELECT d.district_id, d.regency_id, d.district_name, d.district_code,
		       d.latitude, d.longitude, d.created_at, d.updated_at,
		       r.regency_name, r.regency_code, r.province_id,
		       p.province_name, p.province_code
		FROM districts d
		LEFT JOIN regencies r ON d.regency_id = r.regency_id
		LEFT JOIN provinces p ON r.province_id = p.province_id
		%s
		ORDER BY d.%s %s
		LIMIT ? OFFSET ?
	`, whereClause, sortColumn, sortOrder)

	args = append(args, req.Limit, offset)
	rows, err := db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var location entity.Location
		var latitude, longitude sql.NullFloat64
		var regencyName, regencyCode, provinceID sql.NullString
		var provinceName, provinceCode sql.NullString
		var createdAtStr, updatedAtStr sql.NullString // Handle DATETIME as strings

		err := rows.Scan(&location.ID, &location.ParentID, &location.Name, &location.Code,
			&latitude, &longitude, &createdAtStr, &updatedAtStr,
			&regencyName, &regencyCode, &provinceID,
			&provinceName, &provinceCode)
		if err != nil {
			fmt.Printf("ERROR: District row scan failed: %v\n", err)
			return nil, 0, fmt.Errorf("failed to scan district row: %v", err)
		}

		location.Type = entity.LocationTypeDistrict
		location.IsActive = true

		if latitude.Valid {
			location.Latitude = &latitude.Float64
		}
		if longitude.Valid {
			location.Longitude = &longitude.Float64
		}

		// Convert string datetime to time.Time for created_at
		if createdAtStr.Valid && createdAtStr.String != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				location.CreatedAt = parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", createdAtStr.String); err == nil {
				location.CreatedAt = parsedTime
			} else {
				fmt.Printf("WARN: Failed to parse district created_at: %s, error: %v\n", createdAtStr.String, err)
				location.CreatedAt = time.Now() // Fallback
			}
		} else {
			location.CreatedAt = time.Now() // Fallback for null/empty
		}

		// Convert string datetime to time.Time for updated_at
		if updatedAtStr.Valid && updatedAtStr.String != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
				location.UpdatedAt = &parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", updatedAtStr.String); err == nil {
				location.UpdatedAt = &parsedTime
			} else {
				fmt.Printf("WARN: Failed to parse district updated_at: %s, error: %v\n", updatedAtStr.String, err)
			}
		}

		// Add parent information
		if regencyName.Valid && regencyCode.Valid {
			parent := &entity.Location{
				ID:   *location.ParentID,
				Name: regencyName.String,
				Code: regencyCode.String,
				Type: entity.LocationTypeRegency,
			}

			// Add grandparent (province) information
			if provinceName.Valid && provinceCode.Valid && provinceID.Valid {
				parent.Parent = &entity.Location{
					ID:   provinceID.String,
					Name: provinceName.String,
					Code: provinceCode.String,
					Type: entity.LocationTypeProvince,
				}
			}

			location.Parent = parent
		}

		locations = append(locations, location)
	}

	return locations, total, nil
}

func updateDistrict(ctx context.Context, db *sql.DB, id string, req entity.UpdateLocationRequest) error {
	var setParts []string
	var args []interface{}

	if req.Name != nil {
		setParts = append(setParts, "district_name = ?")
		args = append(args, *req.Name)
	}
	if req.Code != nil {
		setParts = append(setParts, "district_code = ?")
		args = append(args, *req.Code)
	}
	if req.ParentID != nil {
		setParts = append(setParts, "regency_id = ?")
		args = append(args, *req.ParentID)
	}
	if req.Latitude != nil {
		setParts = append(setParts, "latitude = ?")
		args = append(args, *req.Latitude)
	}
	if req.Longitude != nil {
		setParts = append(setParts, "longitude = ?")
		args = append(args, *req.Longitude)
	}

	if len(setParts) == 0 {
		return fmt.Errorf("no fields to update")
	}

	setParts = append(setParts, "updated_at = ?")
	args = append(args, time.Now())
	args = append(args, id)

	query := fmt.Sprintf("UPDATE districts SET %s WHERE district_id = ?", strings.Join(setParts, ", "))
	_, err := db.ExecContext(ctx, query, args...)
	return err
}

func deleteDistrict(ctx context.Context, db *sql.DB, id string) error {
	query := "DELETE FROM districts WHERE district_id = ?"
	_, err := db.ExecContext(ctx, query, id)
	return err
}

// Helper functions for Village operations
func createVillage(ctx context.Context, db *sql.DB, id string, req entity.CreateLocationRequest, now time.Time) error {
	if req.ParentID == nil {
		return fmt.Errorf("parent_id (district_id) is required for village")
	}
	if req.VillageType == nil {
		return fmt.Errorf("village_type is required for village")
	}

	query := `
		INSERT INTO villages (village_id, district_id, village_name, village_code, village_type, latitude, longitude, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, id, *req.ParentID, req.Name, req.Code, *req.VillageType, req.Latitude, req.Longitude, now)
	return err
}

func getVillageByID(ctx context.Context, db *sql.DB, id string, location *entity.Location) error {
	query := `
		SELECT v.village_id, v.district_id, v.village_name, v.village_code, v.village_type,
		       v.latitude, v.longitude, v.created_at, v.updated_at,
		       d.district_name, d.district_code, d.regency_id,
		       r.regency_name, r.regency_code, r.province_id,
		       p.province_name, p.province_code
		FROM villages v
		LEFT JOIN districts d ON v.district_id = d.district_id
		LEFT JOIN regencies r ON d.regency_id = r.regency_id
		LEFT JOIN provinces p ON r.province_id = p.province_id
		WHERE v.village_id = ?
	`
	var updatedAt sql.NullTime
	var latitude, longitude sql.NullFloat64
	var districtName, districtCode, regencyID sql.NullString
	var regencyName, regencyCode, provinceID sql.NullString
	var provinceName, provinceCode sql.NullString

	err := db.QueryRowContext(ctx, query, id).Scan(
		&location.ID, &location.ParentID, &location.Name, &location.Code, &location.VillageType,
		&latitude, &longitude, &location.CreatedAt, &updatedAt,
		&districtName, &districtCode, &regencyID,
		&regencyName, &regencyCode, &provinceID,
		&provinceName, &provinceCode)

	if err != nil {
		return err
	}

	location.Type = entity.LocationTypeVillage
	location.IsActive = true
	if latitude.Valid {
		location.Latitude = &latitude.Float64
	}
	if longitude.Valid {
		location.Longitude = &longitude.Float64
	}
	if updatedAt.Valid {
		location.UpdatedAt = &updatedAt.Time
	}

	// Build hierarchical parent information
	if districtName.Valid && districtCode.Valid {
		district := &entity.Location{
			ID:   *location.ParentID,
			Name: districtName.String,
			Code: districtCode.String,
			Type: entity.LocationTypeDistrict,
		}

		// Add regency (grandparent) information
		if regencyName.Valid && regencyCode.Valid && regencyID.Valid {
			regency := &entity.Location{
				ID:   regencyID.String,
				Name: regencyName.String,
				Code: regencyCode.String,
				Type: entity.LocationTypeRegency,
			}

			// Add province (great-grandparent) information
			if provinceName.Valid && provinceCode.Valid && provinceID.Valid {
				regency.Parent = &entity.Location{
					ID:   provinceID.String,
					Name: provinceName.String,
					Code: provinceCode.String,
					Type: entity.LocationTypeProvince,
				}
			}

			district.Parent = regency
		}

		location.Parent = district
	}

	return nil
}

func getVillages(ctx context.Context, db *sql.DB, req entity.LocationListRequest) ([]entity.Location, int64, error) {
	var locations []entity.Location
	var whereConditions []string
	var args []interface{}

	// Validate and sanitize sort parameters for villages
	validSortColumns := map[string]string{
		"name":       "village_name",
		"code":       "village_code",
		"created_at": "created_at",
	}

	sortColumn, exists := validSortColumns[req.SortBy]
	if !exists {
		sortColumn = "village_name" // Default sort column
	}

	sortOrder := "ASC"
	if strings.ToUpper(req.SortOrder) == "DESC" {
		sortOrder = "DESC"
	}

	// Build WHERE conditions
	if req.Search != "" {
		whereConditions = append(whereConditions, "(v.village_name LIKE ? OR v.village_code LIKE ?)")
		searchTerm := "%" + req.Search + "%"
		args = append(args, searchTerm, searchTerm)
	}
	if req.ParentID != "" {
		whereConditions = append(whereConditions, "v.district_id = ?")
		args = append(args, req.ParentID)
	}

	whereClause := ""
	if len(whereConditions) > 0 {
		whereClause = "WHERE " + strings.Join(whereConditions, " AND ")
	}

	// Count total records with detailed error logging
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM villages v %s", whereClause)
	fmt.Printf("DEBUG: Village count query: %s\n", countQuery)
	fmt.Printf("DEBUG: Village count args: %v\n", args)

	var total int64
	err := db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		fmt.Printf("ERROR: Village count query failed: %v\n", err)
		return nil, 0, fmt.Errorf("failed to count villages: %v", err)
	}

	fmt.Printf("DEBUG: Total villages found: %d\n", total)

	// Get paginated data with safe SQL construction
	offset := (req.Page - 1) * req.Limit
	dataQuery := fmt.Sprintf(`
		SELECT v.village_id, v.district_id, v.village_name, v.village_code, v.village_type,
		       v.latitude, v.longitude, v.created_at, v.updated_at,
		       d.district_name, d.district_code, d.regency_id,
		       r.regency_name, r.regency_code, r.province_id,
		       p.province_name, p.province_code
		FROM villages v
		LEFT JOIN districts d ON v.district_id = d.district_id
		LEFT JOIN regencies r ON d.regency_id = r.regency_id
		LEFT JOIN provinces p ON r.province_id = p.province_id
		%s
		ORDER BY v.%s %s
		LIMIT ? OFFSET ?
	`, whereClause, sortColumn, sortOrder)

	args = append(args, req.Limit, offset)
	rows, err := db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var location entity.Location
		var latitude, longitude sql.NullFloat64
		var districtName, districtCode, regencyID sql.NullString
		var regencyName, regencyCode, provinceID sql.NullString
		var provinceName, provinceCode sql.NullString
		var createdAtStr, updatedAtStr sql.NullString

		err := rows.Scan(&location.ID, &location.ParentID, &location.Name, &location.Code, &location.VillageType,
			&latitude, &longitude, &createdAtStr, &updatedAtStr,
			&districtName, &districtCode, &regencyID,
			&regencyName, &regencyCode, &provinceID,
			&provinceName, &provinceCode)
		if err != nil {
			fmt.Printf("ERROR: Village row scan failed: %v\n", err)
			return nil, 0, fmt.Errorf("failed to scan village row: %v", err)
		}

		location.Type = entity.LocationTypeVillage
		location.IsActive = true

		if latitude.Valid {
			location.Latitude = &latitude.Float64
		}
		if longitude.Valid {
			location.Longitude = &longitude.Float64
		}

		// Convert string datetime to time.Time for created_at
		if createdAtStr.Valid && createdAtStr.String != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				location.CreatedAt = parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", createdAtStr.String); err == nil {
				location.CreatedAt = parsedTime
			} else {
				fmt.Printf("WARN: Failed to parse district created_at: %s, error: %v\n", createdAtStr.String, err)
				location.CreatedAt = time.Now() // Fallback
			}
		} else {
			location.CreatedAt = time.Now() // Fallback for null/empty
		}

		// Convert string datetime to time.Time for updated_at
		if updatedAtStr.Valid && updatedAtStr.String != "" {
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
				location.UpdatedAt = &parsedTime
			} else if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", updatedAtStr.String); err == nil {
				location.UpdatedAt = &parsedTime
			} else {
				fmt.Printf("WARN: Failed to parse district updated_at: %s, error: %v\n", updatedAtStr.String, err)
			}
		}

		// Build hierarchical parent information
		if districtName.Valid && districtCode.Valid {
			district := &entity.Location{
				ID:   *location.ParentID,
				Name: districtName.String,
				Code: districtCode.String,
				Type: entity.LocationTypeDistrict,
			}

			// Add regency (grandparent) information
			if regencyName.Valid && regencyCode.Valid && regencyID.Valid {
				regency := &entity.Location{
					ID:   regencyID.String,
					Name: regencyName.String,
					Code: regencyCode.String,
					Type: entity.LocationTypeRegency,
				}

				// Add province (great-grandparent) information
				if provinceName.Valid && provinceCode.Valid && provinceID.Valid {
					regency.Parent = &entity.Location{
						ID:   provinceID.String,
						Name: provinceName.String,
						Code: provinceCode.String,
						Type: entity.LocationTypeProvince,
					}
				}

				district.Parent = regency
			}

			location.Parent = district
		}

		locations = append(locations, location)
	}

	return locations, total, nil
}

func updateVillage(ctx context.Context, db *sql.DB, id string, req entity.UpdateLocationRequest) error {
	var setParts []string
	var args []interface{}

	if req.Name != nil {
		setParts = append(setParts, "village_name = ?")
		args = append(args, *req.Name)
	}
	if req.Code != nil {
		setParts = append(setParts, "village_code = ?")
		args = append(args, *req.Code)
	}
	if req.ParentID != nil {
		setParts = append(setParts, "district_id = ?")
		args = append(args, *req.ParentID)
	}
	if req.VillageType != nil {
		setParts = append(setParts, "village_type = ?")
		args = append(args, *req.VillageType)
	}
	if req.Latitude != nil {
		setParts = append(setParts, "latitude = ?")
		args = append(args, *req.Latitude)
	}
	if req.Longitude != nil {
		setParts = append(setParts, "longitude = ?")
		args = append(args, *req.Longitude)
	}

	if len(setParts) == 0 {
		return fmt.Errorf("no fields to update")
	}

	setParts = append(setParts, "updated_at = ?")
	args = append(args, time.Now())
	args = append(args, id)

	query := fmt.Sprintf("UPDATE villages SET %s WHERE village_id = ?", strings.Join(setParts, ", "))
	_, err := db.ExecContext(ctx, query, args...)
	return err
}

func deleteVillage(ctx context.Context, db *sql.DB, id string) error {
	query := "DELETE FROM villages WHERE village_id = ?"
	_, err := db.ExecContext(ctx, query, id)
	return err
}

// Helper function to get all locations (mixed types) when no type is specified
func getAllLocations(ctx context.Context, db *sql.DB, req entity.LocationListRequest) ([]entity.Location, int64, error) {
	var locations []entity.Location

	// Validate and sanitize sort parameters for mixed locations
	validSortColumns := map[string]string{
		"name":       "name",
		"code":       "code",
		"created_at": "created_at",
	}

	sortColumn, exists := validSortColumns[req.SortBy]
	if !exists {
		sortColumn = "name" // Default sort column
	}

	sortOrder := "ASC"
	if strings.ToUpper(req.SortOrder) == "DESC" {
		sortOrder = "DESC"
	}

	// Build table-specific WHERE conditions for search
	var provinceWhere, regencyWhere, districtWhere, villageWhere string
	var allArgs []interface{} // Combined arguments for all queries

	if req.Search != "" {
		searchTerm := "%" + req.Search + "%"

		// Province-specific WHERE clause
		provinceWhere = "WHERE (province_name LIKE ? OR province_code LIKE ?)"
		allArgs = append(allArgs, searchTerm, searchTerm)

		// Regency-specific WHERE clause
		regencyWhere = "WHERE (regency_name LIKE ? OR regency_code LIKE ?)"
		allArgs = append(allArgs, searchTerm, searchTerm)

		// District-specific WHERE clause
		districtWhere = "WHERE (district_name LIKE ? OR district_code LIKE ?)"
		allArgs = append(allArgs, searchTerm, searchTerm)

		// Village-specific WHERE clause
		villageWhere = "WHERE (village_name LIKE ? OR village_code LIKE ?)"
		allArgs = append(allArgs, searchTerm, searchTerm)
	}

	// Union query with table-specific WHERE clauses
	unionQuery := fmt.Sprintf(`
		SELECT 'province' as type, province_id as id, province_name as name, province_code as code,
		       NULL as parent_id, NULL as regency_type, NULL as village_type,
		       latitude, longitude, created_at, updated_at
		FROM provinces %s
		UNION ALL
		SELECT 'regency' as type, regency_id as id, regency_name as name, regency_code as code,
		       province_id as parent_id, regency_type, NULL as village_type,
		       latitude, longitude, created_at, updated_at
		FROM regencies %s
		UNION ALL
		SELECT 'district' as type, district_id as id, district_name as name, district_code as code,
		       regency_id as parent_id, NULL as regency_type, NULL as village_type,
		       latitude, longitude, created_at, updated_at
		FROM districts %s
		UNION ALL
		SELECT 'village' as type, village_id as id, village_name as name, village_code as code,
		       district_id as parent_id, NULL as regency_type, village_type,
		       latitude, longitude, created_at, updated_at
		FROM villages %s
	`, provinceWhere, regencyWhere, districtWhere, villageWhere)

	// Count total records with detailed error logging
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) as all_locations", unionQuery)
	fmt.Printf("DEBUG: All locations count query: %s\n", countQuery)

	var total int64

	// Use the combined arguments for the count query
	fmt.Printf("DEBUG: All locations count args length: %d\n", len(allArgs))

	err := db.QueryRowContext(ctx, countQuery, allArgs...).Scan(&total)
	if err != nil {
		fmt.Printf("ERROR: All locations count query failed: %v\n", err)
		return nil, 0, fmt.Errorf("failed to count all locations: %v", err)
	}

	fmt.Printf("DEBUG: Total all locations found: %d\n", total)

	// Get paginated data with safe SQL construction
	offset := (req.Page - 1) * req.Limit
	dataQuery := fmt.Sprintf(`
		SELECT type, id, name, code, parent_id, regency_type, village_type,
		       latitude, longitude, created_at, updated_at
		FROM (%s) as all_locations
		ORDER BY %s %s
		LIMIT ? OFFSET ?
	`, unionQuery, sortColumn, sortOrder)

	// Combine search args with pagination args
	dataArgs := make([]interface{}, len(allArgs))
	copy(dataArgs, allArgs)
	dataArgs = append(dataArgs, req.Limit, offset)

	fmt.Printf("DEBUG: All locations data query: %s\n", dataQuery)
	fmt.Printf("DEBUG: All locations data args length: %d\n", len(dataArgs))

	rows, err := db.QueryContext(ctx, dataQuery, dataArgs...)
	if err != nil {
		fmt.Printf("ERROR: All locations data query failed: %v\n", err)
		return nil, 0, fmt.Errorf("failed to query all locations: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var location entity.Location
		var latitude, longitude sql.NullFloat64
		var parentID, regencyType, villageType sql.NullString
		var locationType string
		var createdAtStr, updatedAtStr sql.NullString // Handle DATETIME as strings first

		err := rows.Scan(&locationType, &location.ID, &location.Name, &location.Code,
			&parentID, &regencyType, &villageType,
			&latitude, &longitude, &createdAtStr, &updatedAtStr)
		if err != nil {
			fmt.Printf("ERROR: All locations row scan failed: %v\n", err)
			return nil, 0, fmt.Errorf("failed to scan all locations row: %v", err)
		}

		location.Type = entity.LocationType(locationType)
		location.IsActive = true

		if parentID.Valid {
			location.ParentID = &parentID.String
		}
		if regencyType.Valid {
			location.RegencyType = &regencyType.String
		}
		if villageType.Valid {
			location.VillageType = &villageType.String
		}
		if latitude.Valid {
			location.Latitude = &latitude.Float64
		}
		if longitude.Valid {
			location.Longitude = &longitude.Float64
		}

		// Convert string datetime to time.Time for created_at
		if createdAtStr.Valid && createdAtStr.String != "" {
			// Parse MySQL datetime format: "2006-01-02 15:04:05"
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtStr.String); err == nil {
				location.CreatedAt = parsedTime
			} else {
				// Try alternative format with timezone
				if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", createdAtStr.String); err == nil {
					location.CreatedAt = parsedTime
				} else {
					fmt.Printf("WARN: Failed to parse created_at: %s, error: %v\n", createdAtStr.String, err)
					location.CreatedAt = time.Now() // Fallback to current time
				}
			}
		} else {
			location.CreatedAt = time.Now() // Fallback for null/empty created_at
		}

		// Convert string datetime to time.Time for updated_at
		if updatedAtStr.Valid && updatedAtStr.String != "" {
			// Parse MySQL datetime format: "2006-01-02 15:04:05"
			if parsedTime, err := time.Parse("2006-01-02 15:04:05", updatedAtStr.String); err == nil {
				location.UpdatedAt = &parsedTime
			} else {
				// Try alternative format with timezone
				if parsedTime, err := time.Parse("2006-01-02T15:04:05Z", updatedAtStr.String); err == nil {
					location.UpdatedAt = &parsedTime
				} else {
					fmt.Printf("WARN: Failed to parse updated_at: %s, error: %v\n", updatedAtStr.String, err)
					// Leave updated_at as nil for parsing errors
				}
			}
		}
		// If updatedAtStr is not valid or empty, leave location.UpdatedAt as nil

		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("ERROR: All locations rows iteration error: %v\n", err)
		return nil, 0, fmt.Errorf("error iterating all locations rows: %v", err)
	}

	fmt.Printf("DEBUG: Successfully retrieved %d mixed locations\n", len(locations))
	return locations, total, nil
}
