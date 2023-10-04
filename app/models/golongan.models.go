package models

import (
	"context"
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
