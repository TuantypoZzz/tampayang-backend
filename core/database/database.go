package database

import (
	"context"
	"database/sql"
	"reflect"

	// "os"

	// "encoding/base64"
	"fmt"
	// "reflect"
	"strings"
	"time"

	"tampayang-backend/config"
	"tampayang-backend/core/helper"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionDB() *sql.DB {
	// GET CONFIG DATA
	configuration := helper.ConfigJson()

	var (
		user_name, password, database, host, port interface{}
		fullUrl                                   string
	)

	if config.GO_ENV == "development" {
		user_name = configuration["development"].(map[string]interface{})["username"]
		password = configuration["development"].(map[string]interface{})["password"]
		database = configuration["development"].(map[string]interface{})["database"]
		host = configuration["development"].(map[string]interface{})["host"]
		port = configuration["development"].(map[string]interface{})["port"]
		fullUrl = fmt.Sprint(user_name) + ":" + fmt.Sprint(password) + "@tcp(" + fmt.Sprint(host) + ":" + fmt.Sprint(port) + ")/" + fmt.Sprint(database)
	}

	if config.GO_ENV == "production" {
		user_name = configuration["production"].(map[string]interface{})["username"]
		password = configuration["production"].(map[string]interface{})["password"]
		database = configuration["production"].(map[string]interface{})["database"]
		host = configuration["production"].(map[string]interface{})["host"]
		port = configuration["production"].(map[string]interface{})["port"]
		fullUrl = fmt.Sprint(user_name) + ":" + fmt.Sprint(password) + "@tcp(" + fmt.Sprint(host) + ":" + fmt.Sprint(port) + ")/" + fmt.Sprint(database)
	}

	// SET DB VARIABLES

	// OPEN DB CONNECTION
	db, err := sql.Open("mysql", fullUrl)
	if err != nil {
		panic(err)
	}

	// SET CONNECTION POOLING
	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(500)
	// db.SetConnMaxIdleTime(5 * time.Minute)
	// db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func InserData(tblName string, data map[string]interface{}) int64 {
	// PREPARE CONNECTION AND VARIABLES
	db := GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	var data_cols strings.Builder
	var data_vals strings.Builder
	var bindings []interface{}

	for key, value := range data {
		data_cols.WriteString(key + ",")
		data_vals.WriteString("?,")
		bindings = append(bindings, value)
	}

	cols := strings.TrimSuffix(data_cols.String(), ",")
	vals := strings.TrimSuffix(data_vals.String(), ",")

	sqlQuery := "INSERT INTO " + tblName + "(" + cols + ")" + " VALUES (" + vals + ")"

	// Print query
	if config.GO_ENV == "development" {
		logQuery(sqlQuery, bindings)
	}
	result, err := db.ExecContext(ctx, sqlQuery, bindings...)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return insertId
}

func QuerySelectWitCondition(sqlQuery string, bindings []interface{}) ([]map[string]interface{}, error) {
	// PREPARE CONNECTION AND VARIABLES
	db := GetConnectionDB()
	defer db.Close()
	ctx := context.Background()

	if config.GO_ENV == "development" {
		logQuery(sqlQuery, bindings)
	}

	// Execute the query with the provided bindings
	rows, err := db.QueryContext(ctx, sqlQuery, bindings...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Get the column names dynamically
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Initialize a slice to store the result rows
	var result []map[string]interface{}

	// Create a slice of interface pointers to hold the column values
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	// Check if there are any rows in the result set
	for rows.Next() {
		// Scan the row into the values slice
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		// Create a map to store the column values for each row
		rowData := make(map[string]interface{})

		// Populate the rowData map with column names and values
		for i, colName := range columns {
			rowData[colName] = *values[i].(*interface{})
		}

		// Append the row data to the result slice
		result = append(result, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func QuerySelectWithoutCondition(query string, result interface{}) error {
	// PREPARE CONNECTION AND VARIABLES
	db := GetConnectionDB()

	defer db.Close()

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Use the ScanStruct function to scan data into the result slice
	if err := ScanStruct(rows, result); err != nil {
		return err
	}

	return nil
}

func ScanStruct(rows *sql.Rows, dest interface{}) error {
	// Ensure dest is a pointer to a slice of structs
	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("dest must be a pointer to a slice of structs")
	}

	// Get the type of the struct that each row will be scanned into
	elementType := destValue.Elem().Type().Elem()

	// Create a slice of pointers to the fields of the struct
	fieldPointers := make([]interface{}, elementType.NumField())
	for i := 0; i < elementType.NumField(); i++ {
		fieldPointers[i] = reflect.New(elementType.Field(i).Type).Interface()
	}

	// Iterate through the rows and scan each row into a new struct
	for rows.Next() {
		if err := rows.Scan(fieldPointers...); err != nil {
			return err
		}

		// Create a new instance of the result struct
		newItem := reflect.New(elementType).Elem()

		// Set the struct fields with values from pointers
		for i := 0; i < elementType.NumField(); i++ {
			newItem.Field(i).Set(reflect.ValueOf(fieldPointers[i]).Elem())
		}

		// Append the new instance to the destination slice
		destValue.Elem().Set(reflect.Append(destValue.Elem(), newItem))
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

/**
//FAILED USING STRUCT
func QuerySelect(sqlQuery string, bindings []interface{}, result interface{}) error {
    // PREPARE CONNECTION AND VARIABLES
    db := GetConnectionDB()
    defer db.Close()
    ctx := context.Background()

    if config.GO_ENV == "development" {
        logQuery(sqlQuery, bindings)
    }

    // Execute the query with the provided bindings
    rows, err := db.QueryContext(ctx, sqlQuery, bindings...)
    if err != nil {
        return err
    }
    defer rows.Close()

    // Get the column names dynamically
    columns, err := rows.Columns()
    if err != nil {
        return err
    }

    // Get the value and type of the result slice
    resultValue := reflect.ValueOf(result)
    resultType := resultValue.Type()

    // Check if result is a pointer to a slice of structs
    if resultType.Kind() != reflect.Ptr || resultType.Elem().Kind() != reflect.Slice || resultType.Elem().Elem().Kind() != reflect.Struct {
        return fmt.Errorf("result parameter must be a pointer to a slice of structs")
    }

    // Create a map to store field offsets
    fieldOffsets := make(map[string]int)

    // Populate fieldOffsets with column names and their corresponding struct field indexes
    for i := 0; i < resultType.Elem().Elem().NumField(); i++ {
        field := resultType.Elem().Elem().Field(i)
        columnName := field.Tag.Get("json")
        if columnName == "" {
            columnName = strings.ToLower(field.Name)
        }
        fieldOffsets[columnName] = i
    }

    // Create a slice to store the rows
    sliceType := resultType.Elem().Elem()
    slice := reflect.MakeSlice(resultType.Elem(), 0, 0)

    // Create a slice of interface pointers to hold the column values
    values := make([]interface{}, len(columns))
    for i := range values {
        values[i] = new(interface{})
    }

    // Loop through the rows and append struct instances to the slice
    for rows.Next() {
        // Scan the row into the values slice
        if err := rows.Scan(values...); err != nil {
            return err
        }

        // Create a new struct instance
        structPtr := reflect.New(sliceType)

        // Populate the struct fields with column values
        for i, colName := range columns {
            fieldIndex, ok := fieldOffsets[colName]
            if !ok {
                continue // Column not found in struct
            }

            // Dereference the struct pointer and set the field value
            field := structPtr.Elem().Field(fieldIndex)
            value := *values[i].(*interface{})
            fieldValue := reflect.ValueOf(value)

            if field.Type().Kind() == fieldValue.Type().Kind() {
                field.Set(fieldValue)
            }
        }

        // Append the struct to the result slice
        slice = reflect.Append(slice, structPtr.Elem())
    }

    // Set the result slice to the provided result pointer
    resultValue.Elem().Set(slice)

    if err := rows.Err(); err != nil {
        return err
    }

    return nil
}
*/

func logQuery(query string, args []interface{}) {
	// Replace question marks with actual bindings in the query
	for _, arg := range args {
		query = strings.Replace(query, "?", fmt.Sprintf("%v", arg), 1)
	}
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), query, string(colorReset))
	fmt.Println("")
	fmt.Println("")
}

// func QuerySelect(sqlQuery string, bindings []interface{}) ([]map[string]interface{}, error) {
// 	// PREPARE CONNECTION AND VARIABLES
// 	db := GetConnectionDB()
// 	defer db.Close()
// 	ctx := context.Background()

// if config.GO_ENV == "development" {
//     logQuery(sqlQuery, bindings)
// }

// 	// Execute the query with the provided bindings
//     rows, err := db.QueryContext(ctx, sqlQuery, bindings...)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     // Initialize a slice to store the result rows
//     var result []map[string]interface{}

//     // Check if there are any rows in the result set
//     if rows.Next() {
//         // Get the column names dynamically
//         columns, err := rows.Columns()
//         if err != nil {
//             return nil, err
//         }

// 		// Get column data type information
// 		columnTypes, err := rows.ColumnTypes()
// 		if err != nil {
// 			return nil, err
// 		}

//         // Iterate through the result set
//         for rows.Next() {
//             // Create a map to store the column values for each row
//             rowData := make(map[string]interface{})

//             // Create a slice of interface{} to hold the column values
//             values := make([]interface{}, len(columns))
//             for i := range values {
//                 values[i] = new(interface{})
//             }

//             // Scan the row into the values slice
//             if err := rows.Scan(values...); err != nil {
//                 return nil, err
//             }

//             // Populate the rowData map with column names and values
//             for i, colName := range columns {
// 				// Determine the data type of the column
//             	colType := columnTypes[i]
// 				rowData[colName] = *values[i].(*interface{})
// 				fmt.Println(colType)

// 				// Decode columns based on their data type
// 				if colType.DatabaseTypeName() == "BLOB" || colType.DatabaseTypeName() == "VARBINARY" {
// 					// Decode base64-encoded values to strings for BLOB or VARBINARY columns
// 					decodedValue, err := decodeBase64Value(values[i])
// 					if err != nil {
// 						return nil, err
// 					}
// 					rowData[colName] = string(decodedValue)
// 					// fmt.Println(rowData[colName])
// 				} else {
// 					// Use the value as is for other column types
// 					rowData[colName] = *values[i].(*interface{})
// 				}
//             }

//             // Append the row data to the result slice
//             result = append(result, rowData)
//         }
//     }

//     if err := rows.Err(); err != nil {
//         return nil, err
//     }

//     return result, nil
// }

/**
func decodeBase64Value(value interface{}) ([]byte, error) {
    // Use reflection to assert the value to []byte and then decode it
    val := reflect.ValueOf(value)
    if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Slice && val.Elem().Type().Elem().Kind() == reflect.Uint8 {
        encodedBytes := val.Elem().Interface().([]byte)
        decodedValue := make([]byte, base64.StdEncoding.DecodedLen(len(encodedBytes)))
        n, err := base64.StdEncoding.Decode(decodedValue, encodedBytes)
        if err != nil {
            return nil, err
        }
        return decodedValue[:n], nil
    }
    return nil, fmt.Errorf("unsupported value type for decoding")
}
*/

// this is the blue print for query--------
/**
func selectAllFromTable(ctx context.Context, db *sql.DB, tableName string, columns []string, where string, bindings []interface{}) ([]map[string]interface{}, error) {
    // Initialize the SQL query
    sqlQuery := "SELECT "

    // Build the SELECT clause with specified columns
    if len(columns) > 0 {
        sqlQuery += strings.Join(columns, ", ")
    } else {
        sqlQuery += "*"
    }

    // Add the table name
    sqlQuery += " FROM " + tableName

    // Add the WHERE clause if provided
    if where != "" {
        sqlQuery += " WHERE " + where
    }

    // Execute the query with the provided bindings and context
    rows, err := db.QueryContext(ctx, sqlQuery, bindings...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Initialize a slice to store the result rows
    var result []map[string]interface{}

    // Check if there are any rows in the result set
    if rows.Next() {
        // Get the column names dynamically
        columns, err := rows.Columns()
        if err != nil {
            return nil, err
        }

        // Iterate through the result set
        for rows.Next() {
            // Create a map to store the column values for each row
            rowData := make(map[string]interface{})

            // Create a slice of interface{} to hold the column values
            values := make([]interface{}, len(columns))
            for i := range values {
                values[i] = new(interface{})
            }

            // Scan the row into the values slice
            if err := rows.Scan(values...); err != nil {
                return nil, err
            }

            // Populate the rowData map with column names and values
            for i, colName := range columns {
                rowData[colName] = *values[i].(*interface{})
            }

            // Append the row data to the result slice
            result = append(result, rowData)
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return result, nil
}
*/
