package controllers

import (
	"tampayang-backend/app/models"
	globalFunction "tampayang-backend/core/functions"
	"tampayang-backend/core/response"

	"github.com/gofiber/fiber/v2"
)

func InfrastructureCategory(ctx *fiber.Ctx) error {
	lov := models.GetLovInfrastructureCategory()
	return response.SuccessResponse(ctx, lov)
}

func DamageType(ctx *fiber.Ctx) error {
	infrastructureCategoryId := ctx.Query("category_id", "")

	if globalFunction.IsEmpty(infrastructureCategoryId) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("lov001", nil))
	}

	lov := models.GetLovDamageType(infrastructureCategoryId)
	return response.SuccessResponse(ctx, lov)
}

func Regency(ctx *fiber.Ctx) error {
	provinceId := ctx.Query("province_id", "")

	if globalFunction.IsEmpty(provinceId) {
		provinceId = models.GetDefaultProvince().Id
	}

	lov := models.GetLovRegency(provinceId)
	return response.SuccessResponse(ctx, lov)
}

func District(ctx *fiber.Ctx) error {
	regencyId := ctx.Query("regency_id", "")

	if globalFunction.IsEmpty(regencyId) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("lov002", nil))
	}

	lov := models.GetLovDistrict(regencyId)
	return response.SuccessResponse(ctx, lov)
}

func Village(ctx *fiber.Ctx) error {
	districtId := ctx.Query("district_id", "")

	if globalFunction.IsEmpty(districtId) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("lov003", nil))
	}

	lov := models.GetLovVillage(districtId)
	return response.SuccessResponse(ctx, lov)
}
