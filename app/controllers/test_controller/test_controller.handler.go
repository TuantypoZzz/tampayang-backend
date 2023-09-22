package test_controller

import (
	// "golang-fiber-template/app/models"
	// globalFunction "golang-fiber-template/core/functions"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	// gorest_api "github.com/nulla-vis/golang-fiber-template/app/libs/gorest"
	elasticsearchLib "github.com/nulla-vis/golang-fiber-template/app/libs/elasticsearch"
	test_model "github.com/nulla-vis/golang-fiber-template/app/models/test_model"
	"github.com/nulla-vis/golang-fiber-template/config/constant"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/response"
	// "encoding/base64"
)

func CreateUserHandler(ctx *fiber.Ctx) error {

	var name string
	var created string
	var created_date string
	
	name = ctx.Params("name")

	if name == "error" {

		// code := "conn001"
		// errorMessage := map[string]interface{} {
		// 	"code": code,
		// 	"id": "Koneksi Error",
		// 	"en": "Connection Error",
		// }
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err001", nil))
		// panic(ctx.JSON(errorMessage))
		// panic("Something went wrong")
	}

	now := time.Now()
	created = now.Format(constant.NOW_TIME_FORMAT)
	created_date = now.Format(constant.NOW_DATE_TIME_FORMAT)
	// data := map[string]interface{}{
	// 	"name": name,
	// 	"created": created,
	// 	"created_date": created_date,
	// }

	// lastId := test_model.InsertCategory(data)

	// responseData := map[string]interface{}{
	// 	"last_id": lastId,
	// }

	// someData := map[string]interface{} {
	// 	"code": 123,
	// 	"id": "Hanya data",
	// 	"en": "Just data",
	// 	"waktu": now.Format(constant.NOW_TIME_FORMAT),
	// }

	// // responseData = globalFunction.GetMessage(someData,"")


	// return response.SuccessResponse(ctx, responseData)



	sqlQuery := "INSERT INTO category(name, created, created_date) VALUES (?,?,?)"
	insertData := test_model.InsertCategoryStruct{
		Name: name,
		Created: created,
		Created_date: created_date,
	}

	dbResult, err := test_model.InsertCategoryDatabase(sqlQuery, insertData)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	type lastIdResponse struct {
		Message string `json:"message"`
		LastId	int64	`json:"lastId"`
	}

	responseData := lastIdResponse{
		Message: "Ini string",
		LastId: dbResult,
	}

	return response.SuccessResponse(ctx, responseData)



}

func GetAllUserHandler(ctx *fiber.Ctx) error {
	// USING CONDITION IN QUERY--------------
	/**
	whereCondition := "cat.id > ?"
	whereData := []interface{}{1}

	dbResult,err := test_model.SelectAllFromCategoryWithCondition(whereCondition, whereData)
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
	*/

	//USING NO CONDITION IN QUERY--------------
	
	dbResult,err := test_model.SelectAllFromCategoryWithoutCondition()
	if err != nil {
		return response.ErrorResponse(ctx,err)
	}

	// check elasticsearch connection
	// 1. Check ES ping
	esPing := elasticsearchLib.CheckPing()
	// 2. If ping ok, create an index example

	if esPing {
		// Define the index name and properties
		indexName := "story"
		documentId := "333-0"
		// CREATE INDEX-----------------------------------------------------------
		// properties := map[string]interface{}{
		// 	"field1": map[string]interface{}{
		// 		"type": "text",
		// 		// You can add other mapping properties for field1 here if needed
		// 	},
		// 	"field2": map[string]interface{}{
		// 		"type": "integer",
		// 		// You can add other mapping properties for field2 here if needed
		// 	},
		// 	// Add mappings for other fields as needed
		// }

		// Create the Elasticsearch index
		// result, err := elasticsearchLib.CreateIndex(indexName, properties)
		// if err != nil {
		// 	// Handle the error
		// 	return response.ErrorResponse(ctx, err)
		// }
		// CREATE INDEX-----------------------------------------------------------

		// DELETE INDEX-----------------------------------------------------------
		// result, err := elasticsearchLib.DeleteIndex(esClient, indexName)
		// if err != nil {
		// 	return response.ErrorResponse(ctx, err)
		// }
		// DELETE INDEX-----------------------------------------------------------

		// INSER DATA-------------------------------------------------------------
		// documentID := "4"
		// data := map[string]interface{}{
		// 	"field1": "123",
		// 	"field2": "456",
		// }

		// // Insert the data into Elasticsearch
		// result, err := elasticsearchLib.InsertData(indexName, documentID, data)
		// if err != nil {
		// 	response.ErrorResponse(ctx, err)
		// }
		// INSER DATA-------------------------------------------------------------

		// UPDATE DATA-------------------------------------------------------------
		// documentID := "4"
		// data := map[string]interface{}{
		// 	"field1": "444",
		// 	"field2": "444",
		// }

		// update the data into Elasticsearch
		// result, err := elasticsearchLib.UpdateIndex(indexName, documentID, data)
		// if err != nil {
		// 	response.ErrorResponse(ctx, err)
		// }
		// UPDATE DATA-------------------------------------------------------------

		// EXIST INDEX-------------------------------------------------------------
		// result, err := elasticsearchLib.IndexExists(indexName)
		// if err != nil {
		// 	response.ErrorResponse(ctx, err)
		// }
		// EXIST INDEX-------------------------------------------------------------

		// EXIST DOCUMENT----------------------------------------------------------
		// result, err := elasticsearchLib.DocumentExists(indexName, documentId)
		// if err != nil {
		// 	response.ErrorResponse(ctx, err)
		// }
		// EXIST DOCUMENT----------------------------------------------------------

		// DELETE DOCUMENT----------------------------------------------------------
		// result, err := elasticsearchLib.DeleteDocument(indexName, documentId)
		// if err != nil {
		// 	response.ErrorResponse(ctx, err)
		// }
		// DELETE DOCUMENT----------------------------------------------------------

		// GET A DOCUMENT----------------------------------------------------------
		result, err := elasticsearchLib.GetDocument(indexName, documentId)
		if err != nil {
			response.ErrorResponse(ctx, err)
		}
		// GET A DOCUMENT----------------------------------------------------------
		return response.SuccessResponse(ctx, result)

	} else {
		fmt.Println("Connection Error")
	}

	return response.SuccessResponse(ctx, dbResult)
	

	// 3rd party API Call / Fetct API
	/**
	result, err := gorest_api.GorestGetAllUser()
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}
	return response.SuccessResponse(ctx, result)
	*/
}



