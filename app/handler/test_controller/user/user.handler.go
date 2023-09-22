package user_handler

import (
	"github.com/gofiber/fiber/v2"
	user_model "github.com/nulla-vis/golang-fiber-template/app/models/user"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)

func CreateUserHandler(ctx *fiber.Ctx) error {

	newUser := new(CreateUserHandlerStruct)

	if err := ctx.BodyParser(newUser); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	sqlQuery := "INSERT INTO user(name, age, created_date) VALUES (?,?,?)"

	insertData := user_model.InsertNewUserStruct{
		Name:         newUser.Name,
		Age:          newUser.Age,
		Created_date: newUser.Created_date,
	}

	_, err := user_model.InsertNewUserDatabase(sqlQuery, insertData)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	responseData := globalFunction.GetMessage("user001", nil)

	return response.SuccessResponse(ctx, responseData)
}

func GetAllUsersHandler(ctx *fiber.Ctx) error {
	dbResult, err := user_model.SelectAllFromUser()
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}
	return response.SuccessResponse(ctx, dbResult)
}
