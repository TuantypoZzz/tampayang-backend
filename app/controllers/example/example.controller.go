package example_controller

import (

	"github.com/gofiber/fiber/v2"
	example_model "github.com/nulla-vis/golang-fiber-template/app/models/example"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)


func CreateExample(ctx *fiber.Ctx) error{
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
	query := "SELECT exa.name FROM example AS exa WHERE exa.name = ? LIMIT 1"
	nameExist:= example_model.ExampleNameIsUnique(query, example.Name)
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
	sqlQuery := "INSERT INTO example(name, created, rating, booleandesu, created_date) VALUES(?,?,?,?,?)"
	
	// Set data based on MODEL struct
	exampleData := example_model.InsertExampleStruct{
		Name: example.Name,
		Created: example.Created,
		Rating: example.Rating,
		Booleandesu: example.Booleandesu,
		Created_date: example.Created_date,
	}

	dbResult, err := example_model.InsertExample(sqlQuery, exampleData)
	if err != nil {
		response.ErrorResponse(ctx, err)
	}

	result := CreateExampleResponse{
		ExampleId: dbResult,
	}

	return response.SuccessResponse(ctx, result)
}

func GetExampleById(ctx *fiber.Ctx) error {
	return nil
}

func GetAllExample(ctx *fiber.Ctx) error {
	return nil
}

func UpdateExample(ctx *fiber.Ctx) error {
	return nil
}

func DeleteExample(ctx *fiber.Ctx) error {
	return nil
}