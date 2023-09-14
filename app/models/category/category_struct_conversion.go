package category_model

import "fmt"

// list for converting interface to it's struct

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

		fmt.Println(myStruct)

        result = append(result, myStruct)
    }

    return result
}
