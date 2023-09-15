package user_handler

import (
	// "golang-fiber-template/app/models"
	// globalFunction "golang-fiber-template/core/functions"
	// "fmt"
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
		"created_date": now.Format(constant.NOW_DATE_TIME_FORMAT),
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
	// USING CONDITION IN QUERY--------------
	whereCondition := "cat.id > ?"
	whereData := []interface{}{1}

	dbResult,err := category_model.SelectAllFromCategoryWithCondition(whereCondition, whereData)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	// Format created_date as "DD-MM-YYYY HH:mm:ss" just before returning the response
	for i := range dbResult {
		// Parse the Created_date string into a time.Time value
		createdDate, err := time.Parse(constant.NOW_DATE_TIME_FORMAT, dbResult[i].Created_date)
		if err != nil {
			return response.ErrorResponse(ctx, err)
		}

		// Format the parsed time.Time value
		formattedDate := createdDate.Format(constant.NOW_DATE_TIME_FORMAT)

		// Update the Created_date field with the formatted string
		dbResult[i].Created_date = formattedDate
	}

	return response.SuccessResponse(ctx, dbResult)

	//USING NO CONDITION IN QUERY--------------
	// var sqlQuery = "SELECT cat.id, cat.name, cat.rating, cat.booleandesu, cat.created FROM category AS cat"
	
	// dbResult,err := category_model.SelectAllFromCategoryWithoutCondition(sqlQuery)
	// if err != nil {
	// 	return response.ErrorResponse(ctx, err)
	// }

	// return response.SuccessResponse(ctx, dbResult)
	
}



