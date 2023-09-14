package category_model

import (
	// "database/sql"
	// "time"
	"github.com/nulla-vis/golang-fiber-template/config/tables"
	"github.com/nulla-vis/golang-fiber-template/core/database"
    globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	// "github.com/nulla-vis/golang-fiber-template/core"
)

// var db *sql.DB = database.GetConnectionDB()

func InsertCategory(data map[string]interface{}) int64{
	lastId := database.InserData("category", data)
	return lastId
}

// Declare struct like column(s) in sqlQuery

type SelectAllFromCategoryWithConditionStruct struct {
    Id          int32       `json:"id"`
    Name        string      `json:"name"`
    Rating      float64     `json:"Rating"`
    Booleandesu bool        `json:"booleandesu"`
}

//Query WITH struct
func SelectAllFromCategoryWithConditionWithStruct(where string, bindings []interface{}) ([]GetAllUserHandlerStruct, error) {
    // Initialize the SQL query
    sqlQuery := "SELECT cat.id, cat.name, cat.rating, cat.booleandesu, cat.created FROM " + tables.Category + " AS cat"

    // Add the WHERE clause if provided
    if where != "" {
        sqlQuery += " WHERE " + where
    }

    dbResult, err := database.QuerySelect(sqlQuery, bindings)
    if err != nil {
        panic(err)
    }

    result := ConvertToGetAllUserHandlerStruct(dbResult)

	return result, nil
}


//Query WITHOUT struct
func SelectAllFromCategoryWithCondition(where string, bindings []interface{}) ([]map[string]interface{}, error) {
    // Initialize the SQL query
    sqlQuery := "SELECT cat.id, cat.name, cat.rating, cat.booleandesu, cat.created FROM " + tables.Category + " AS cat"

    // Add the WHERE clause if provided
    if where != "" {
        sqlQuery += " WHERE " + where
    }

    result, err := database.QuerySelect(sqlQuery, bindings)
    if err != nil {
        panic(err)
    }

    // convert data from database (if Query WITHOUT struct)
	globalFunction.ConvertByteSlicesToStrings(result)

	return result, nil
}





