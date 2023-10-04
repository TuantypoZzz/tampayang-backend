package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/models"
	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)

func CreateGolongan(ctx *fiber.Ctx) error {
	newGolongan := new(entity.Golongan)

	if err := ctx.BodyParser(newGolongan); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	insertData := entity.Golongan{
		GolonganName: newGolongan.GolonganName,
	}

	errValidasi := insertData.ValidationGolongan()
	if errValidasi != nil {
		return response.ErrorResponse(ctx, errValidasi)
	}

	// validasi nama golongan tidak boleh sama
	queryGet := "SELECT gol.golongan_name FROM golongan AS gol WHERE gol.golongan_name = ? LIMIT 1"
	isGolUniq := models.GolonganNameIsUnique(queryGet, newGolongan.GolonganName)
	if newGolongan.GolonganName == isGolUniq {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("gol004", nil))
	}

	sqlQuery := " INSERT INTO golongan(golongan_name, created_date) VALUES (?,?)"
	_, err := models.InsertNewGolonganDatabase(sqlQuery, insertData)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	responseData := globalFunction.GetMessage("success", nil)

	return response.SuccessResponse(ctx, responseData)
}
