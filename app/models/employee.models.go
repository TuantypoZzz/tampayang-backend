package models

import (
	"context"
	"fmt"

	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	"github.com/nulla-vis/golang-fiber-template/core/database"
)

func InsertNewEmployeeDatabase(query string, data entity.Employee) (int64, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	result, err := db.ExecContext(ctx, query, data.Name,
		data.Nip,
		data.Bidang,
		data.Seksi,
		data.UnitKerja,
		data.Gender,
		data.BirthPlace,
		data.BirthDate,
		data.Phone,
		data.Email,
		data.Created_date)
	if err != nil {
		return 0, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertId, nil
}

func GetAllEmployeePagenation(page int, limit int) (entity.PaginationEmply, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	// Mengambil jumlah total data pegawai
	totalEmployees := 0
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM employee").Scan(&totalEmployees); err != nil {
		panic(err)
	}

	// Hitung offset berdasarkan halaman dan batasan
	offset := (page - 1) * limit

	// Validasi apakah halaman melebihi jumlah maksimum halaman
	maxPage := (totalEmployees + limit - 1) / limit // Menghitung jumlah halaman dengan pembulatan ke atas
	if page > maxPage {
		// Jika halaman melebihi maksimum, ubah halaman ke halaman terakhir
		page = maxPage
		offset = (page - 1) * limit
	}

	// Query SQL untuk mengambil data pegawai pada halaman tertentu
	sqlQuery := fmt.Sprintf("SELECT employee_id, name, nip, bidang, seksi, unit_kerja, gender, birth_place, birth_date, phone, email, created_date FROM employee LIMIT %d OFFSET %d", limit, offset)
	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	var employees []entity.EmployeeWithId
	for result.Next() {
		var employee entity.EmployeeWithId
		err := result.Scan(
			&employee.Id,
			&employee.Name,
			&employee.Nip,
			&employee.Bidang,
			&employee.Seksi,
			&employee.UnitKerja,
			&employee.Gender,
			&employee.BirthPlace,
			&employee.BirthDate,
			&employee.Phone,
			&employee.Email,
			&employee.Created_date)

		if err != nil {
			panic(err)
		}

		employees = append(employees, employee)
	}

	// Membuat respons dengan data yang sesuai
	response := entity.PaginationEmply{
		Data:      employees,
		TotalRow:  len(employees), // Menghitung jumlah data pada halaman ini
		TotalPage: maxPage,        // Maksimum halaman
		NoPerPage: page,           // Halaman saat ini
	}

	return response, nil
}

func GetEmployeeById(employeeId int) entity.EmployeeWithId {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var employeeResult entity.EmployeeWithId

	sqlQuery := "SELECT emply.employee_id, emply.name, emply.nip, emply.bidang, emply.seksi, emply.unit_kerja, emply.gender, emply.birth_place, emply.birth_date, emply.phone, emply.email, emply.created_date FROM employee AS emply WHERE emply.employee_id = ?"
	result, err := db.QueryContext(ctx, sqlQuery, employeeId)
	if err != nil {
		panic("models - GetEmployeeById, db.QueryContext: " + err.Error())
	}

	if result.Next() {
		err := result.Scan(
			&employeeResult.Id,
			&employeeResult.Name,
			&employeeResult.Nip,
			&employeeResult.Bidang,
			&employeeResult.Seksi,
			&employeeResult.UnitKerja,
			&employeeResult.Gender,
			&employeeResult.BirthPlace,
			&employeeResult.BirthDate,
			&employeeResult.Phone,
			&employeeResult.Email,
			&employeeResult.Created_date)
		if err != nil {
			panic("models - GetEmployeeById, result.Scan: " + err.Error())
		}
	}

	return employeeResult
}
