package example_model

import (
	"context"

	"github.com/nulla-vis/golang-fiber-template/core/database"
)

func InsertExample(query string, exampleData InsertExampleStruct) (int64, error) {
    db  := database.GetConnectionDB()
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

func ExampleNameIsUnique(query string, name string)(string) {
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