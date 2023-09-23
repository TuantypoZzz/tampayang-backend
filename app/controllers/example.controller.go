package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/models"
	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)

func CreateExample(ctx *fiber.Ctx) error {
	// Data from Body POST
	example := new(entity.Example)
	if err := ctx.BodyParser(example); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	// Field validation...
	// Validate name
	if globalFunction.IsEmpty(example.Name) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0001", nil))
	}

	// name must unique
	query_get := "SELECT exa.name FROM example AS exa WHERE exa.name = ? LIMIT 1"
	nameExist := models.ExampleNameIsUnique(query_get, example.Name)
	if example.Name == nameExist {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0006", nil))
	}

	// Validate created
	if globalFunction.IsEmpty(example.Created) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0002", nil))
	}

	// Validate rating
	if globalFunction.IsEmpty(example.Rating) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0003", nil))
	}

	// Validate booleandesu
	if globalFunction.IsEmpty(example.Booleandesu) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0004", nil))
	}

	// Validate created_date
	if globalFunction.IsEmpty(example.Created_date) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0005", nil))
	}
	// String sql query
	query_insert := "INSERT INTO example(name, created, rating, booleandesu, created_date) VALUES(?,?,?,?,?)"

	// Set data based on MODEL struct
	exampleData := entity.Example{
		Name:         example.Name,
		Created:      example.Created,
		Rating:       example.Rating,
		Booleandesu:  example.Booleandesu,
		Created_date: example.Created_date,
	}

	dbResult, err := models.InsertExample(query_insert, exampleData)
	if err != nil {
		response.ErrorResponse(ctx, err)
	}

	result := entity.ExampleId{
		ExampleId: dbResult,
	}

	return response.SuccessResponse(ctx, result)
}

func GetExampleById(ctx *fiber.Ctx) error {
	example_id := ctx.Params("example_id")

	// validate example_id
	int_example_id, err := strconv.Atoi(example_id)
	if err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err002", nil))
	}

	dbResult := models.GetExampleById(int_example_id)

	if globalFunction.IsEmpty(dbResult.Name) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0007", nil))
	}

	return response.SuccessResponse(ctx, dbResult)
}

func GetAllExample(ctx *fiber.Ctx) error {
	dbResult := models.GetAllExample()

	return response.SuccessResponse(ctx, dbResult)
}

func UpdateExample(ctx *fiber.Ctx) error {

	example := new(entity.ExampleWithId)
	if err := ctx.BodyParser(example); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	// Validate id
	if globalFunction.IsEmpty(example.Id) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0008", nil))
	}

	dbResult := models.GetExampleById(example.Id)
	if globalFunction.IsEmpty(dbResult.Id) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0007", nil))
	}

	// Validate name
	if globalFunction.IsEmpty(example.Name) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0001", nil))
	}

	// Validate created
	if globalFunction.IsEmpty(example.Created) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0002", nil))
	}

	// Validate rating
	if globalFunction.IsEmpty(example.Rating) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0003", nil))
	}

	// Validate booleandesu
	if globalFunction.IsEmpty(example.Booleandesu) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0004", nil))
	}

	// Validate created_date
	if globalFunction.IsEmpty(example.Created_date) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0005", nil))
	}

	// Set data based on MODEL struct
	exampleData := entity.ExampleWithId{
		Id:           example.Id,
		Name:         example.Name,
		Created:      example.Created,
		Rating:       example.Rating,
		Booleandesu:  example.Booleandesu,
		Created_date: example.Created_date,
	}

	UpdateResult := models.UpdateExample(exampleData)

	return response.SuccessResponse(ctx, UpdateResult)
}

func DeleteExample(ctx *fiber.Ctx) error {
	example_id := ctx.Params("example_id")

	// validate example_id
	int_example_id, err := strconv.Atoi(example_id)
	if err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err002", nil))
	}

	dbResult := models.GetExampleById(int_example_id)
	if globalFunction.IsEmpty(dbResult.Id) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0007", nil))
	}

	success_delete := models.DeleteExampleByID(int_example_id)

	if !success_delete {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0010", nil))
	}

	return response.SuccessResponse(ctx, globalFunction.GetMessage("ex0009", nil))
}
