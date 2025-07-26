package models

import (
	"context"

	"tampayang-backend/app/models/entity"
	"tampayang-backend/core/database"
)

func GetLovInfrastructureCategory() []entity.Lov {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var lov []entity.Lov

	sqlQuery := "SELECT infrastructure_category_id, name, code FROM infrastructure_categories"
	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - GetLovInfrastructureCategory, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var value entity.Lov
		err := result.Scan(
			&value.Id,
			&value.Name,
			&value.Code)
		if err != nil {
			panic("models - GetLovInfrastructureCategory, result.Scan " + err.Error())
		}
		lov = append(lov, value)
	}
	if err := result.Err(); err != nil {
		panic("models - GetLovInfrastructureCategory, result.Err " + err.Error())
	}
	return lov
}

func GetLovDamageType(infrastructureCategoryId string) []entity.Lov {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var lov []entity.Lov

	sqlQuery := "SELECT damage_type_id, name, code FROM damage_types WHERE infrastructure_category_id = ?"
	result, err := db.QueryContext(ctx, sqlQuery, infrastructureCategoryId)
	if err != nil {
		panic("models - GetLovDamageType, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var value entity.Lov
		err := result.Scan(
			&value.Id,
			&value.Name,
			&value.Code)
		if err != nil {
			panic("models - GetLovDamageType, result.Scan " + err.Error())
		}
		lov = append(lov, value)
	}
	if err := result.Err(); err != nil {
		panic("models - GetLovInfrastructureCategory, result.Err " + err.Error())
	}
	return lov
}

func GetDefaultProvince() entity.Lov {
	defaultProvinceCode := "MAL"

	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var value entity.Lov

	sqlQuery := "SELECT province_id, province_name, province_code FROM provinces WHERE province_code = ?"
	result, err := db.QueryContext(ctx, sqlQuery, defaultProvinceCode)
	if err != nil {
		panic("models - GetDefaultProvince, db.QueryContext " + err.Error())
	}
	defer result.Close()

	if result.Next() {
		err := result.Scan(
			&value.Id,
			&value.Name,
			&value.Code)

		if err != nil {
			panic("models - GetDefaultProvince, result.Scan " + err.Error())
		}
	}
	return value
}

func GetLovRegency(provinceId string) []entity.Lov {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var lov []entity.Lov

	sqlQuery := "SELECT regency_id, regency_name, regency_code FROM regencies WHERE province_id = ?"
	result, err := db.QueryContext(ctx, sqlQuery, provinceId)
	if err != nil {
		panic("models - GetLovRegency, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var value entity.Lov
		err := result.Scan(
			&value.Id,
			&value.Name,
			&value.Code)
		lov = append(lov, value)
		if err != nil {
			panic("models - GetLovRegency, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetLovInfrastructureCategory, result.Err " + err.Error())
	}
	return lov
}

func GetLovDistrict(regencyId string) []entity.Lov {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var lov []entity.Lov

	sqlQuery := "SELECT district_id, district_name, district_code FROM districts WHERE regency_id = ?"
	result, err := db.QueryContext(ctx, sqlQuery, regencyId)
	if err != nil {
		panic("models - GetLovDistrict, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var value entity.Lov
		err := result.Scan(
			&value.Id,
			&value.Name,
			&value.Code)
		lov = append(lov, value)
		if err != nil {
			panic("models - GetLovDistrict, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetLovInfrastructureCategory, result.Err " + err.Error())
	}
	return lov
}

func GetLovVillage(districtId string) []entity.Lov {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var lov []entity.Lov

	sqlQuery := "SELECT village_id, village_name, village_code FROM villages WHERE district_id = ?"
	result, err := db.QueryContext(ctx, sqlQuery, districtId)
	if err != nil {
		panic("models - GetLovVillage, db.QueryContext " + err.Error())
	}
	defer result.Close()

	for result.Next() {
		var value entity.Lov
		err := result.Scan(
			&value.Id,
			&value.Name,
			&value.Code)
		lov = append(lov, value)
		if err != nil {
			panic("models - GetLovVillage, result.Scan " + err.Error())
		}
	}
	if err := result.Err(); err != nil {
		panic("models - GetLovInfrastructureCategory, result.Err " + err.Error())
	}
	return lov
}
