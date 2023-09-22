package example_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	example_model "github.com/nulla-vis/golang-fiber-template/app/models/example"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)

func CreateExample(ctx *fiber.Ctx) error {
	// Data from Body POST
	example := new(CreateExampleStruct)
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
	nameExist := example_model.ExampleNameIsUnique(query_get, example.Name)
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
	exampleData := example_model.InsertExampleStruct{
		Name:         example.Name,
		Created:      example.Created,
		Rating:       example.Rating,
		Booleandesu:  example.Booleandesu,
		Created_date: example.Created_date,
	}

	dbResult, err := example_model.InsertExample(query_insert, exampleData)
	if err != nil {
		response.ErrorResponse(ctx, err)
	}

	result := CreateExampleResponse{
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

	dbResult := example_model.GetExampleById(int_example_id)

	if globalFunction.IsEmpty(dbResult.Name) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0007", nil))
	}

	return response.SuccessResponse(ctx, dbResult)
}

func GetAllExample(ctx *fiber.Ctx) error {
	dbResult := example_model.GetAllExample()

	return response.SuccessResponse(ctx, dbResult)
}

func UpdateExample(ctx *fiber.Ctx) error {

	example := new(UpdateExampleStruct)
	if err := ctx.BodyParser(example); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	// Validate id
	if globalFunction.IsEmpty(example.Id) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("ex0008", nil))
	}

	dbResult := example_model.GetExampleById(example.Id)
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
	// String sql query
	query_example_update := "UPDATE example SET name = ?, created = ?, rating = ?, booleandesu = ?, created_date = ? WHERE id = ?"

	// Set data based on MODEL struct
	exampleData := example_model.UpdateExampleStruct{
		Id:           example.Id,
		Name:         example.Name,
		Created:      example.Created,
		Rating:       example.Rating,
		Booleandesu:  example.Booleandesu,
		Created_date: example.Created_date,
	}

	UpdateResult := example_model.UpdateExample(query_example_update, exampleData)

	return response.SuccessResponse(ctx, UpdateResult)
}

func DeleteExample(ctx *fiber.Ctx) error {
	return nil
}
