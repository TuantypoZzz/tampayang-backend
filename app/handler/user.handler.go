package handler

import (
	// "golang-fiber-template/app/models"
	// globalFunction "golang-fiber-template/core/functions"
	// "fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	category_model "github.com/nulla-vis/golang-fiber-template/app/models/category"
	"github.com/nulla-vis/golang-fiber-template/config/constant"
	"github.com/nulla-vis/golang-fiber-template/core/response"
	// "encoding/base64"
)

func CreateUserHandler(ctx *fiber.Ctx) error {

	name := ctx.Params("name")

	if name == "error" {

		// code := "conn001"
		// errorMessage := map[string]interface{} {
		// 	"code": code,
		// 	"id": "Koneksi Error",
		// 	"en": "Connection Error",
		// }
		return response.ErrorResponse(ctx, "sebuah error coy")
		// panic(ctx.JSON(errorMessage))
		// panic("Something went wrong")
	}

	now := time.Now()
	data := map[string]interface{}{
		"name": name,
		"created": now.Format(constant.NOW_TIME_FORMAT),
	}

	lastId := category_model.InsertCategory(data)

	responseData := map[string]interface{}{
		"last_id": lastId,
	}

	// someData := map[string]interface{} {
	// 	"code": 123,
	// 	"id": "Hanya data",
	// 	"en": "Just data",
	// 	"waktu": now.Format(constant.NOW_TIME_FORMAT),
	// }

	// // responseData = globalFunction.GetMessage(someData,"")


	return response.SuccessResponse(ctx, responseData)
}

func GetAllUserHandler(ctx *fiber.Ctx) error {
	whereCondition := "cat.id > ?"
	whereData := []interface{}{15}

	dbResult,err := category_model.SelectAllFromCategoryWithConditionWithStruct(whereCondition, whereData)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	return response.SuccessResponse(ctx, dbResult)

}