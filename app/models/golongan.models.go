package models

import (
	"context"
	"fmt"
	"time"

	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	"github.com/nulla-vis/golang-fiber-template/config/constant"
	"github.com/nulla-vis/golang-fiber-template/core/database"
)

func InsertNewGolonganDatabase(query string, data entity.Golongan) (int64, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()
	now := time.Now()
	currentTIme := now.Format(constant.NOW_DATE_TIME_FORMAT)

	result, err := db.ExecContext(ctx, query, data.GolonganName, currentTIme)
	if err != nil {
		panic("models - InsertNewGolonganDatabase, db.ExecContext: " + err.Error())
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic("models - InsertNewGolonganDatabase, result.LastInsertId: " + err.Error())
	}

	return insertId, nil
}

func GolonganNameIsUnique(query string, name string) string {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var nameExist string
	result, err := db.QueryContext(ctx, query, name)
	if err != nil {
		panic("models - GolonganNameIsUnique, db.QueryContext : " + err.Error())
	}

	if result.Next() {
		if err := result.Scan(&nameExist); err != nil {
			panic("models - GolonganNameIsUnique, result.Scan : " + err.Error())
		}
	}

	return nameExist
}

func GetAllGolonganPagenation(page int, limit int) (entity.PageGolongan, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	// Jumlah total data golongan
	var totalGolongan int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM golongan").Scan(&totalGolongan); err != nil {
		panic("models - GetAllGolonganPagenation, db.QueryRowContext: " + err.Error())
	}

	// Hitung offset berdasarkan halaman dan batasan
	offset := (page - 1) * limit

	// Validasi apakah halaman melebihi jumlah maksimum halaman
	maxPage := (totalGolongan + limit - 1) / limit // Menghitung jumlah halaman dengan pembulatan ke atas
	if page > maxPage {
		// Jika halaman melebihi maksimum, ubah halaman ke halaman terakhir
		page = maxPage
		offset = (page - 1) * limit
	}

	// Query get data golongan
	sqlQuery := fmt.Sprintf("SELECT golongan_id, golongan_name, created_date FROM golongan LIMIT %d OFFSET %d", limit, offset)
	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic("models - GetAllGolonganPagenation, db.QueryContext: " + err.Error())
	}

	var golongan []entity.GolonganWithId
	for result.Next() {
		var gol entity.GolonganWithId
		err := result.Scan(
			&gol.Id,
			&gol.GolonganName,
			&gol.Created_date)

		if err != nil {
			panic("models - GetAllGolonganPagenation, result.Scan: " + err.Error())
		}

		golongan = append(golongan, gol)
	}

	// Membuat respons dengan data yang sesuai
	response := entity.PageGolongan{
		Data:      golongan,
		TotalRow:  len(golongan), // Menghitung jumlah data pada halaman ini
		TotalPage: maxPage,       // Maksimum halaman
		NoPerPage: page,          // Halaman saat ini
	}

	return response, nil
}

func GetGolonganById(golonganId int) entity.GolonganWithId {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var golonganResult entity.GolonganWithId

	sqlQuery := "SELECT gol.golongan_id, gol.golongan_name, gol.created_date FROM golongan AS gol WHERE gol.golongan_id = ?"
	result, err := db.QueryContext(ctx, sqlQuery, golonganId)
	if err != nil {
		panic("models - GetGolonganById, db.QueryContext: " + err.Error())
	}

	if result.Next() {
		err := result.Scan(
			&golonganResult.Id,
			&golonganResult.GolonganName,
			&golonganResult.Created_date)
		if err != nil {
			panic("models - GetGolonganById, result.Scan: " + err.Error())
		}
	}

	return golonganResult
}
