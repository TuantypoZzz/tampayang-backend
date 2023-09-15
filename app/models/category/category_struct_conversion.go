package category_model

import (
	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
)

func ConvertToGetAllUserHandlerStruct(data []map[string]interface{}) []GetAllUserHandlerStruct {
    var result []GetAllUserHandlerStruct

    for _, item := range data {
        myStruct := GetAllUserHandlerStruct{
			Id:          item["id"].(int64),
            Name:        string(item["name"].([]byte)),
            Rating:      item["rating"].(float64),
            Booleandesu: item["booleandesu"].(int64),
            Created: 	 string(item["created"].([]byte)),
        }

        result = append(result, myStruct)
    }

    return result
}

// no need to use this, instead use QuerySelectWithoutCondition from database module
func ConvertToGetAllUserHandlerStructNew(data []map[string]interface{}) []GetAllUserHandlerStruct {
    var result []GetAllUserHandlerStruct

    for _, item := range data {
        var id int64
        var rating float64
        var booleandesu int64

        // Handle 'id' conversion (int64)
        if idValue, ok := item["id"].(int64); ok {
            id = idValue
        } else if idValueBytes, ok := item["id"].([]byte); ok {
            id, _ = globalFunction.ConvertBytesToInt64(idValueBytes)
        }

        // Handle 'rating' conversion (float64)
        if ratingValue, ok := item["rating"].(float64); ok {
            rating = ratingValue
        } else if ratingValueBytes, ok := item["rating"].([]byte); ok {
            rating, _ = globalFunction.ConvertBytesToFloat64(ratingValueBytes)
        }

        // Handle 'booleandesu' conversion (int64)
        if booleandesuValue, ok := item["booleandesu"].(int64); ok {
            booleandesu = booleandesuValue
        } else if booleandesuValueBytes, ok := item["booleandesu"].([]byte); ok {
            booleandesu, _ = globalFunction.ConvertBytesToInt64(booleandesuValueBytes)
        }

        myStruct := GetAllUserHandlerStruct{
            Id:          id,
            Name:        string(item["name"].([]byte)),
            Rating:      rating,
            Booleandesu: booleandesu,
            Created:     string(item["created"].([]byte)),
        }

        result = append(result, myStruct)
    }

    return result
}




