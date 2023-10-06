package controllers

import (
	"strconv"

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

func GetAllGolongan(ctx *fiber.Ctx) error {
	pageParam, _ := strconv.Atoi(ctx.Query("page", "1"))
	limitParam, _ := strconv.Atoi(ctx.Query("limit", "10"))

	// Pastikan "page" selalu lebih besar atau sama dengan 1
	if pageParam < 1 {
		pageParam = 1
	}

	// Memanggil fungsi GetAllEmployee dengan nilai "page" dan "limit" yang diterima
	employees, err := models.GetAllGolonganPagenation(pageParam, limitParam)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	return response.SuccessResponse(ctx, employees)
}

func GetGolonganById(ctx *fiber.Ctx) error {
	golonganId := ctx.Params("golongan_id")

	// Validasi golongan_id
	int_golongan_id, err := strconv.Atoi(golonganId)
	if err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err002", nil))
	}

	dbResult := models.GetGolonganById(int_golongan_id)

	if globalFunction.IsEmpty(dbResult.Id) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("gol005", nil))
	}

	return response.SuccessResponse(ctx, dbResult)
}
