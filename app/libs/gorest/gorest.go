package gorest_api

import (
	"encoding/json"
	// "errors"
	"fmt"
	"io"
	"net/http"

	// "strconv"

	globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/helper"
)

// CustomError is a custom error type that includes a data structure.
type CustomError struct {
    Data []map[string]interface{}
}

// Error implements the error interface by returning a string.
func (e *CustomError) Error() string {
    return "CustomError: " + fmt.Sprint(e.Data)
}

func goRestRequest(data map[string]interface{}) (*http.Response, error){
	options := getOptions(data)
	// fmt.Println(options)
	// ---logger here---
	reqGorest, err := globalFunction.MakeAPIRequest(options)
	if err != nil {
		return reqGorest, err
	}
	// CLOSE HTTP REQUEST
	defer reqGorest.Body.Close()

	return reqGorest, nil
}

func getOptions(data map[string]interface{}) map[string]interface{}{
	config := helper.ConfigJson()
	accessToken := fmt.Sprint(config["gorest"].(map[string]interface{})["access-token"])

	head := map[string]interface{}{
		"Accept" : "application/json",
		"Content-Type": "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	options := map[string]interface{}{
		"method": data["method"],
		"headers": head,
		"url": data["url"],
		"json": true,
		"timeout": 3000,
	}

	// Check if "body" key exists in data map
	if dataBody, ok := data["body"]; ok {
		// Add "body" key to options map
		options["body"] = dataBody
	}

	return options
}


func GorestGetAllUser() ([]map[string]interface{}, error){
	config := helper.ConfigJson()
	url := fmt.Sprint(config["gorest"].(map[string]interface{})["url"]) + "users/"
	data := map[string]interface{}{
		"method": "GET",
		"url": url,
	}

	data_res, err := goRestRequest(data)
	if err != nil {
		return nil, err
	}

    // Read the response body into a byte slice
    responseBody, err := io.ReadAll(data_res.Body)
    if err != nil {
        return nil, err
    }

	// Parse the JSON response into a map[string]interface{}
	var result []map[string]interface{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GorestGetUserDetail(userId string) (map[string]interface{}, error){
	config := helper.ConfigJson()
	url := fmt.Sprint(config["gorest"].(map[string]interface{})["url"]) + "users/" + userId
	data := map[string]interface{}{
		"method": "GET",
		"url": url,
	}

	data_res, err := goRestRequest(data)
	if err != nil {
		return nil, err
	}

    // Read the response body into a byte slice
    responseBody, err := io.ReadAll(data_res.Body)
    if err != nil {
        return nil, err
    }

	// Parse the JSON response into a map[string]interface{}
	var result map[string]interface{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GorestCreateUser() (map[string]interface{}, error){
	config := helper.ConfigJson()
	url := fmt.Sprint(config["gorest"].(map[string]interface{})["url"]) + "users"
	data := map[string]interface{}{
		"method": "POST",
		"url": url,
		"body": map[string]interface{}{
			"name": "Aan",
			"gender": "male",
			"email": "asdasd@mail12.com",
			"status": "active",
		  },
	}

	data_res, err := goRestRequest(data)
	if err != nil {
		return nil, err
	}

	fmt.Println(data_res.StatusCode)
	// handle Error
	if data_res.StatusCode >= 300 {
		var result []map[string]interface{}

		responseBody, err := io.ReadAll(data_res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(responseBody, &result)
		if err != nil{
			return nil, err
		}

		errCustom := &CustomError{
			Data: result,
		}
		return nil, errCustom
	}

    // Read the response body into a byte slice
    responseBody, err := io.ReadAll(data_res.Body)
    if err != nil {
        return nil, err
    }
	
	// Parse the JSON response into a map[string]interface{}
	var result map[string]interface{}
	err = json.Unmarshal(responseBody, &result)
	if err != nil{
		return nil, err
	}

	return result, nil
}
