package controllers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	sanitizeLib "github.com/nulla-vis/golang-fiber-template/app/libs/sanitize"
	validationLib "github.com/nulla-vis/golang-fiber-template/app/libs/validation"
	"github.com/nulla-vis/golang-fiber-template/app/models"
	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	"github.com/nulla-vis/golang-fiber-template/core/database"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/response"
)

func CreateEmployeeHandler(ctx *fiber.Ctx) error {

	newEmployee := new(entity.Employee)

	if err := ctx.BodyParser(newEmployee); err != nil {
		return response.ErrorResponse(ctx, err)
	}

	// Validasi No telpon
	sanitizePhoneNumber := sanitizeLib.PhoneNumber(newEmployee.Phone)
	if !validationLib.IsValidPhoneNumber(sanitizePhoneNumber) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err005", nil))
	}

	insertData := entity.Employee{
		Name:         newEmployee.Name,
		Nip:          newEmployee.Nip,
		Bidang:       newEmployee.Bidang,
		Seksi:        newEmployee.Seksi,
		UnitKerja:    newEmployee.UnitKerja,
		Gender:       newEmployee.Gender,
		BirthPlace:   newEmployee.BirthPlace,
		BirthDate:    newEmployee.BirthDate,
		Phone:        sanitizePhoneNumber,
		Email:        newEmployee.Email,
		Created_date: newEmployee.Created_date,
	}

	errValidasi := insertData.ValidationEmployee()
	if errValidasi != nil {
		return response.ErrorResponse(ctx, errValidasi)
	}

	// validasi nip tidak boleh sama
	queryGet := "SELECT emply.nip FROM employee AS emply WHERE emply.nip = ? LIMIT 1"
	isEmplyUniq := isEmployeeNipIsUnique(queryGet, newEmployee.Nip)
	if newEmployee.Nip == isEmplyUniq {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("emply008", nil))
	}

	sqlQuery := " INSERT INTO employee(name, nip, bidang, seksi, unit_kerja, gender, birth_place, birth_date, phone, email, created_date) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

	_, err := models.InsertNewEmployeeDatabase(sqlQuery, insertData)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	responseData := globalFunction.GetMessage("emply001", nil)

	return response.SuccessResponse(ctx, responseData)
}

func GetAllEmployee(ctx *fiber.Ctx) error {
	pageParam, _ := strconv.Atoi(ctx.Query("page", "1"))
	limitParam, _ := strconv.Atoi(ctx.Query("limit", "5"))

	// Pastikan "page" selalu lebih besar atau sama dengan 1
	if pageParam < 1 {
		pageParam = 1
	}

	// Memanggil fungsi GetAllEmployee dengan nilai "page" dan "limit" yang diterima
	employees, err := models.GetAllEmployeePagenation(pageParam, limitParam)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}

	return response.SuccessResponse(ctx, employees)
}

func GetEmployeeById(ctx *fiber.Ctx) error {
	employeeId := ctx.Params("employee_id")

	// Validasi employee_id
	int_employee_id, err := strconv.Atoi(employeeId)
	if err != nil {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("err002", nil))
	}

	dbResult := models.GetEmployeeById(int_employee_id)

	if globalFunction.IsEmpty(dbResult.Id) {
		return response.ErrorResponse(ctx, globalFunction.GetMessage("emply012", nil))
	}

	return response.SuccessResponse(ctx, dbResult)
}

func isEmployeeNipIsUnique(query string, nip string) string {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var nipExist string = ""
	result, err := db.QueryContext(ctx, query, nip)
	if err != nil {
		panic(err)
	}

	if result.Next() {
		if err := result.Scan(&nipExist); err != nil {
			panic(err)
		}
	}

	return nipExist
}
