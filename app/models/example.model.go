package models

import (
	"context"

	"github.com/nulla-vis/golang-fiber-template/app/models/entity"
	"github.com/nulla-vis/golang-fiber-template/core/database"
)

func InsertExample(query string, exampleData entity.Example) (int64, error) {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	result, err := db.ExecContext(ctx, query, exampleData.Name, exampleData.Created, exampleData.Rating, exampleData.Booleandesu, exampleData.Created_date)
	if err != nil {
		panic(err)
	}

    insertId, err := result.LastInsertId()
    if err != nil {
        panic(err)
    }

    return insertId, nil
}

func UpdateExample(query string, exampleData entity.ExampleWithId) *entity.ExampleWithId {
	db := database.GetConnectionDB()
	defer db.Close()
	// ctx := context.Background()

	_, err := db.Exec(query, exampleData.Name, exampleData.Created, exampleData.Rating, exampleData.Booleandesu, exampleData.Created_date, exampleData.Id)
	if err != nil {
		panic(err)
	}

	query = "SELECT exa.* FROM example AS exa WHERE exa.id = ?"
	result := db.QueryRow(query, exampleData.Id)

	var example entity.ExampleWithId
	err = result.Scan(
		&example.Id,
		&example.Name,
		&example.Created,
		&example.Rating,
		&example.Booleandesu,
		&example.Created_date)

	if err != nil {
		panic(err)
	}
	return &example
}

func ExampleNameIsUnique(query string, name string) string {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var nameExist string = ""
	result, err := db.QueryContext(ctx, query, name)
	if err != nil {
		panic(err)
	}

	if result.Next() {
		if err := result.Scan(&nameExist); err != nil {
			panic(err)
		}
	}

	return nameExist
}

func GetExampleById(exampleId int) entity.ExampleWithId {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var exampleResult entity.ExampleWithId

	sqlQuery := "SELECT exa.* FROM example AS exa WHERE exa.id = ?"
	result, err := db.QueryContext(ctx, sqlQuery, exampleId)
	if err != nil {
		panic(err)
	}

	if result.Next() {
		err := result.Scan(
			&exampleResult.Id,
			&exampleResult.Name,
			&exampleResult.Created,
			&exampleResult.Rating,
			&exampleResult.Booleandesu,
			&exampleResult.Created_date)

		if err != nil {
			panic(err)
		}
	}

	return exampleResult
}

func GetAllExample() []entity.ExampleWithId {
	db := database.GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var exampleResults []entity.ExampleWithId

	sqlQuery := "SELECT exa.* FROM example AS exa"

	result, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	if result.Next() {
		var example entity.ExampleWithId
		err := result.Scan(
			&example.Id,
			&example.Name,
			&example.Rating,
			&example.Booleandesu,
			&example.Created,
			&example.Created_date)

		if err != nil {
			panic(err)
		}
		return exampleResults
	}
	if err := result.Err(); err != nil {
		panic(err)
	}
	return exampleResults
}
