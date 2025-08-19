package entity

import (
	"time"
)

// LocationType represents the type of location in the hierarchy
type LocationType string

const (
	LocationTypeProvince LocationType = "province"
	LocationTypeRegency  LocationType = "regency"
	LocationTypeDistrict LocationType = "district"
	LocationTypeVillage  LocationType = "village"
)

// Location represents a unified location entity for CRUD operations
type Location struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Code         string       `json:"code"`
	Type         LocationType `json:"type"`
	ParentID     *string      `json:"parent_id,omitempty"`
	Latitude     *float64     `json:"latitude,omitempty"`
	Longitude    *float64     `json:"longitude,omitempty"`
	RegencyType  *string      `json:"regency_type,omitempty"`  // For regencies: 'kabupaten' or 'kota'
	VillageType  *string      `json:"village_type,omitempty"`  // For villages: 'desa' or 'kelurahan'
	IsActive     bool         `json:"is_active"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    *time.Time   `json:"updated_at,omitempty"`
	
	// Hierarchical information
	Parent   *Location   `json:"parent,omitempty"`
	Children []Location  `json:"children,omitempty"`
}

// Province represents province entity
type Province struct {
	ProvinceID   string     `json:"province_id"`
	ProvinceName string     `json:"province_name"`
	ProvinceCode string     `json:"province_code"`
	Latitude     *float64   `json:"latitude,omitempty"`
	Longitude    *float64   `json:"longitude,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

// Regency represents regency entity
type Regency struct {
	RegencyID    string     `json:"regency_id"`
	ProvinceID   string     `json:"province_id"`
	RegencyName  string     `json:"regency_name"`
	RegencyCode  string     `json:"regency_code"`
	RegencyType  string     `json:"regency_type"` // 'kabupaten' or 'kota'
	Latitude     *float64   `json:"latitude,omitempty"`
	Longitude    *float64   `json:"longitude,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	
	// Related data
	Province *Province `json:"province,omitempty"`
}

// District represents district entity
type District struct {
	DistrictID   string     `json:"district_id"`
	RegencyID    string     `json:"regency_id"`
	DistrictName string     `json:"district_name"`
	DistrictCode string     `json:"district_code"`
	Latitude     *float64   `json:"latitude,omitempty"`
	Longitude    *float64   `json:"longitude,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	
	// Related data
	Regency  *Regency  `json:"regency,omitempty"`
	Province *Province `json:"province,omitempty"`
}

// Village represents village entity
type Village struct {
	VillageID   string     `json:"village_id"`
	DistrictID  string     `json:"district_id"`
	VillageName string     `json:"village_name"`
	VillageCode string     `json:"village_code"`
	VillageType string     `json:"village_type"` // 'desa' or 'kelurahan'
	Latitude    *float64   `json:"latitude,omitempty"`
	Longitude   *float64   `json:"longitude,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	
	// Related data
	District *District `json:"district,omitempty"`
	Regency  *Regency  `json:"regency,omitempty"`
	Province *Province `json:"province,omitempty"`
}

// CreateLocationRequest represents request for creating a location
type CreateLocationRequest struct {
	Name        string       `json:"name" validate:"required,min=2,max=100"`
	Code        string       `json:"code" validate:"required,min=2,max=10"`
	Type        LocationType `json:"type" validate:"required,oneof=province regency district village"`
	ParentID    *string      `json:"parent_id,omitempty"`
	Latitude    *float64     `json:"latitude,omitempty" validate:"omitempty,min=-90,max=90"`
	Longitude   *float64     `json:"longitude,omitempty" validate:"omitempty,min=-180,max=180"`
	RegencyType *string      `json:"regency_type,omitempty" validate:"omitempty,oneof=kabupaten kota"`
	VillageType *string      `json:"village_type,omitempty" validate:"omitempty,oneof=desa kelurahan"`
}

// UpdateLocationRequest represents request for updating a location
type UpdateLocationRequest struct {
	Name        *string  `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Code        *string  `json:"code,omitempty" validate:"omitempty,min=2,max=10"`
	ParentID    *string  `json:"parent_id,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty" validate:"omitempty,min=-90,max=90"`
	Longitude   *float64 `json:"longitude,omitempty" validate:"omitempty,min=-180,max=180"`
	RegencyType *string  `json:"regency_type,omitempty" validate:"omitempty,oneof=kabupaten kota"`
	VillageType *string  `json:"village_type,omitempty" validate:"omitempty,oneof=desa kelurahan"`
	IsActive    *bool    `json:"is_active,omitempty"`
}

// LocationListRequest represents request for listing locations
type LocationListRequest struct {
	Page       int          `json:"page" validate:"min=1"`
	Limit      int          `json:"limit" validate:"min=1,max=100"`
	Type       LocationType `json:"type,omitempty"`
	ParentID   string       `json:"parent_id,omitempty"`
	Search     string       `json:"search,omitempty"`
	IsActive   *bool        `json:"is_active,omitempty"`
	SortBy     string       `json:"sort_by,omitempty" validate:"omitempty,oneof=name code created_at"`
	SortOrder  string       `json:"sort_order,omitempty" validate:"omitempty,oneof=asc desc"`
}

// LocationListResponse represents response for listing locations
type LocationListResponse struct {
	Data       []Location         `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

// PaginationResponse represents pagination information
type PaginationResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// LocationHierarchy represents hierarchical location data
type LocationHierarchy struct {
	Province *Province  `json:"province,omitempty"`
	Regency  *Regency   `json:"regency,omitempty"`
	District *District  `json:"district,omitempty"`
	Village  *Village   `json:"village,omitempty"`
}

// LocationStats represents location statistics
type LocationStats struct {
	TotalProvinces int `json:"total_provinces"`
	TotalRegencies int `json:"total_regencies"`
	TotalDistricts int `json:"total_districts"`
	TotalVillages  int `json:"total_villages"`
}

// LocationDependency represents dependency information for deletion
type LocationDependency struct {
	HasReports   bool `json:"has_reports"`
	ReportCount  int  `json:"report_count"`
	HasChildren  bool `json:"has_children"`
	ChildrenCount int `json:"children_count"`
}
