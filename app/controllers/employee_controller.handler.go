package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nulla-vis/golang-fiber-template/app/models"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)

func CreateEmployeeHandler(ctx *fiber.Ctx) error {

	newEmployee := new(entity.Employee)

	if err := ctx.BodyParser(newEmployee); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	sqlQuery := " INSERT INTO employee(name, nip, bidang, seksi, unit_kerja, gender, birth_place, birth_date, phone, email, created_date) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

	insertData := entity.Employee{
		Name:         newEmployee.Name,
		Nip:          newEmployee.Nip,
		Bidang:       newEmployee.Bidang,
		Seksi:        newEmployee.Seksi,
		UnitKerja:    newEmployee.UnitKerja,
		Gender:       newEmployee.Gender,
		BirthPlace:   newEmployee.BirthPlace,
		BirthDate:    newEmployee.BirthDate,
		Phone:        newEmployee.Phone,
		Email:        newEmployee.Email,
		Created_date: newEmployee.Created_date,
	}

	_, err := models.InsertNewEmployeeDatabase(sqlQuery, insertData)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	responseData := globalFunction.GetMessage("emply001", nil)

	return response.SuccessResponse(ctx, responseData)
}
