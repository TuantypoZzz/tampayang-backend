package controllers

import (
	"fmt"
	"strconv"

	"tampayang-backend/app/models"
	"tampayang-backend/app/models/entity"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// CreateLocation creates a new location
func CreateLocation(ctx *fiber.Ctx) error {
	var req entity.CreateLocationRequest

	// Parse request body
	if err := ctx.BodyParser(&req); err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "request_body",
			"error": "Invalid request format",
		}))
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "validation",
			"error": err.Error(),
		}))
	}

	// Additional validation based on location type
	if err := validateLocationTypeRequirements(req); err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "location_type",
			"error": err.Error(),
		}))
	}

	// Create location
	location, err := models.CreateLocation(req)
	if err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("database001", nil))
	}
	return response.SuccessResponse(ctx, location)
}

// GetLocation retrieves a single location by ID
func GetLocation(ctx *fiber.Ctx) error {
	// Get path parameters
	id := ctx.Params("id")
	typeParam := ctx.Query("type", "")

	if id == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "id",
			"error": "Location ID is required",
		}))
	}

	if typeParam == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "type",
			"error": "Location type is required",
		}))
	}

	// Validate location type
	locationType := entity.LocationType(typeParam)
	if !isValidLocationType(locationType) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "type",
			"error": "Invalid location type. Use: province, regency, district, village",
		}))
	}

	// Log the request
	fmt.Printf("INFO: Get location requested from IP: %s (id: %s, type: %s)\n",
		ctx.IP(), id, locationType)

	// Get location
	fmt.Printf("INFO: Calling models.GetLocationByID from IP: %s (id: %s, type: %s)\n", ctx.IP(), id, locationType)
	location, err := models.GetLocationByID(id, locationType)
	if err != nil {
		if err.Error() == "location not found" {
			fmt.Printf("WARN: Location not found from IP: %s (id: %s, type: %s)\n",
				ctx.IP(), id, locationType)
			return response.ErrorResponse(ctx, globalFunction.GetMessage("notfound001", nil))
		}
		fmt.Printf("ERROR: models.GetLocationByID failed from IP: %s - %v\n", ctx.IP(), err)
		return response.ErrorResponse(ctx, globalFunction.GetMessage("database001", nil))
	}

	fmt.Printf("INFO: Successfully retrieved location from IP: %s - %s (%s)\n", ctx.IP(), location.Name, location.ID)
	return response.SuccessResponse(ctx, location)
}

// GetLocations retrieves locations with pagination and filtering
func GetLocations(ctx *fiber.Ctx) error {
	var req entity.LocationListRequest

	// Parse query parameters
	req.Page, _ = strconv.Atoi(ctx.Query("page", "1"))
	req.Limit, _ = strconv.Atoi(ctx.Query("limit", "10"))
	req.Type = entity.LocationType(ctx.Query("type", ""))
	req.ParentID = ctx.Query("parent_id", "")
	req.Search = ctx.Query("search", "")
	req.SortBy = ctx.Query("sort_by", "name")
	req.SortOrder = ctx.Query("sort_order", "asc")

	// Parse is_active parameter
	if isActiveStr := ctx.Query("is_active", ""); isActiveStr != "" {
		if isActive, err := strconv.ParseBool(isActiveStr); err == nil {
			req.IsActive = &isActive
		}
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "validation",
			"error": err.Error(),
		}))
	}

	// Validate location type if provided
	if req.Type != "" && !isValidLocationType(req.Type) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "type",
			"error": "Invalid location type. Use: province, regency, district, village",
		}))
	}
	// Get locations
	fmt.Printf("INFO: Calling models.GetLocations from IP: %s with request: %+v\n", ctx.IP(), req)
	response_data, err := models.GetLocations(req)
	if err != nil {
		fmt.Printf("ERROR: models.GetLocations failed from IP: %s - %v\n", ctx.IP(), err)
		return response.ErrorResponse(ctx, globalFunction.GetMessage("database001", nil))
	}

	fmt.Printf("INFO: Successfully retrieved locations from IP: %s - count: %d\n", ctx.IP(), len(response_data.Data))
	return response.SuccessResponse(ctx, response_data)
}

// UpdateLocation updates a location by ID
func UpdateLocation(ctx *fiber.Ctx) error {
	// Get path parameters
	id := ctx.Params("id")
	typeParam := ctx.Query("type", "")

	if id == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "id",
			"error": "Location ID is required",
		}))
	}

	if typeParam == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "type",
			"error": "Location type is required",
		}))
	}

	// Validate location type
	locationType := entity.LocationType(typeParam)
	if !isValidLocationType(locationType) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "type",
			"error": "Invalid location type. Use: province, regency, district, village",
		}))
	}

	var req entity.UpdateLocationRequest

	// Parse request body
	if err := ctx.BodyParser(&req); err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "request_body",
			"error": "Invalid request format",
		}))
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "validation",
			"error": err.Error(),
		}))
	}

	// Update location
	location, err := models.UpdateLocation(id, locationType, req)
	if err != nil {
		if err.Error() == "location not found" {
			return response.ErrorResponse(ctx, globalFunction.GetMessage("notfound001", nil))
		}
		return response.ErrorResponse(ctx, globalFunction.GetMessage("database001", nil))
	}

	return response.SuccessResponse(ctx, location)
}

// DeleteLocation deletes a location by ID
func DeleteLocation(ctx *fiber.Ctx) error {
	// Get path parameters
	id := ctx.Params("id")
	typeParam := ctx.Query("type", "")

	if id == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "id",
			"error": "Location ID is required",
		}))
	}

	if typeParam == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "type",
			"error": "Location type is required",
		}))
	}

	// Validate location type
	locationType := entity.LocationType(typeParam)
	if !isValidLocationType(locationType) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "type",
			"error": "Invalid location type. Use: province, regency, district, village",
		}))
	}

	// Log the request
	fmt.Printf("INFO: Delete location requested from IP: %s (id: %s, type: %s)\n",
		ctx.IP(), id, locationType)

	// Delete location
	err := models.DeleteLocation(id, locationType)
	if err != nil {
		if err.Error() == "location not found" {
			return response.ErrorResponse(ctx, globalFunction.GetMessage("notfound001", nil))
		}
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "dependencies",
			"error": err.Error(),
		}))
	}

	return response.SuccessResponse(ctx, map[string]interface{}{
		"message": "Location deleted successfully",
		"id":      id,
		"type":    locationType,
	})
}

// GetLocationStats returns statistics about locations
func GetLocationStats(ctx *fiber.Ctx) error {
	// Get location statistics
	stats, err := models.GetLocationStats()
	if err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("database001", nil))
	}

	return response.SuccessResponse(ctx, stats)
}

// CheckLocationDependencies checks dependencies for a location
func CheckLocationDependencies(ctx *fiber.Ctx) error {
	// Get path parameters
	id := ctx.Params("id")
	typeParam := ctx.Query("type", "")

	if id == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "id",
			"error": "Location ID is required",
		}))
	}

	if typeParam == "" {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "type",
			"error": "Location type is required",
		}))
	}

	// Validate location type
	locationType := entity.LocationType(typeParam)
	if !isValidLocationType(locationType) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("validation001", map[string]interface{}{
			"field": "type",
			"error": "Invalid location type. Use: province, regency, district, village",
		}))
	}

	// Check dependencies
	dependency, err := models.CheckLocationDependencies(id, locationType)
	if err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("database001", nil))
	}

	return response.SuccessResponse(ctx, dependency)
}

// Helper function to validate location type requirements
func validateLocationTypeRequirements(req entity.CreateLocationRequest) error {
	switch req.Type {
	case entity.LocationTypeProvince:
		// Province doesn't need parent_id
		if req.ParentID != nil {
			return fmt.Errorf("province should not have parent_id")
		}
	case entity.LocationTypeRegency:
		// Regency needs parent_id (province_id) and regency_type
		if req.ParentID == nil {
			return fmt.Errorf("regency requires parent_id (province_id)")
		}
		if req.RegencyType == nil {
			return fmt.Errorf("regency requires regency_type (kabupaten or kota)")
		}
	case entity.LocationTypeDistrict:
		// District needs parent_id (regency_id)
		if req.ParentID == nil {
			return fmt.Errorf("district requires parent_id (regency_id)")
		}
	case entity.LocationTypeVillage:
		// Village needs parent_id (district_id) and village_type
		if req.ParentID == nil {
			return fmt.Errorf("village requires parent_id (district_id)")
		}
		if req.VillageType == nil {
			return fmt.Errorf("village requires village_type (desa or kelurahan)")
		}
	}
	return nil
}

// Helper function to validate location type
func isValidLocationType(locationType entity.LocationType) bool {
	switch locationType {
	case entity.LocationTypeProvince, entity.LocationTypeRegency, entity.LocationTypeDistrict, entity.LocationTypeVillage:
		return true
	default:
		return false
	}
}
