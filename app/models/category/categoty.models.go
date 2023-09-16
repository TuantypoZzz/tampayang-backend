package category_model

import (
	"github.com/nulla-vis/golang-fiber-template/config/tables"
	"github.com/nulla-vis/golang-fiber-template/core/database"
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
)

/**
------------------
| IMPORTANT NOTE |
------------------

1. Better to query using STRUCT
2. Add struct in {modelname}_struct.go based on what field you want to query
3. If query using condition and bindings, add data conversion in {modelname}_struct_conversion.go
4. Query with condition call database.QuerySelectWitCondition
5. Just query select without condition call database.QuerySelectWithoutCondition
6. If not using STRUCT, always convert data from database using globalFunction.ConvertByteSlicesToStrings

---------------
| MODEL RULES |
---------------
1. Inside function, declare sqlQuery
2. If have additional LOGIC, if can process it before query to database or after it
3. Make as less database process as possible
3. Call function from database module based on query

---------------
| ERRORS |
---------------
1. sql: expected x destination arguments in Scan, not y -> check the query and struct KEYS


*/

func InsertCategory(data map[string]interface{}) int64{
	lastId := database.InserData("category", data)
	return lastId
}

//Query WITH struct
func SelectAllFromCategoryWithCondition(where string, bindings []interface{}) ([]GetAllUserHandlerStruct, error) {
    // Initialize the SQL query
    sqlQuery := "SELECT cat.id, cat.name, cat.rating, cat.booleandesu, cat.created, cat.created_date FROM " + tables.Category + " AS cat"

    // Add the WHERE clause if provided
    if where != "" {
        sqlQuery += " WHERE " + where
    }

    dbResult, err := database.QuerySelectWitCondition(sqlQuery, bindings)
    if err != nil {
        panic(err)
    }

    result := ConvertToGetAllUserHandlerStruct(dbResult)

	return result, nil
}

//Query WITH struct
func SelectAllFromCategoryWithoutCondition() ([]GetAllUserHandlerStruct, error) {
    sqlQuery := "SELECT cat.id, cat.name, cat.rating, cat.booleandesu, cat.created, cat.created_date FROM category AS cat"

    var result []GetAllUserHandlerStruct
    
    if err := database.QuerySelectWithoutCondition(sqlQuery, &result); err != nil {
        // Handle error
        return nil, err
    }

	return result, nil
}

//Query WITHOUT struct (NOT RECOMMENDED)--------------------------------------------------------------------------------------------
func SelectAllFromCategoryWithConditionWithouStruct(where string, bindings []interface{}) ([]map[string]interface{}, error) {
    // Initialize the SQL query
    sqlQuery := "SELECT cat.id, cat.name, cat.rating, cat.booleandesu, cat.created FROM " + tables.Category + " AS cat"

    // Add the WHERE clause if provided
    if where != "" {
        sqlQuery += " WHERE " + where
    }

    result, err := database.QuerySelectWitCondition(sqlQuery, bindings)
    if err != nil {
        panic(err)
    }

    // convert data from database (if Query WITHOUT struct)
	globalFunction.ConvertByteSlicesToStrings(result)

	return result, nil
}
//Query WITHOUT struct (NOT RECOMMENDED)--------------------------------------------------------------------------------------------



