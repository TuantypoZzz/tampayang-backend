package gorest_api

import (
	"fmt"

	// globalFunction "github.com/nulla-vis/golang-fiber-template/core/functions"
	"github.com/nulla-vis/golang-fiber-template/core/helper"
)

func goRestRequest(data map[string]interface{}) map[string]interface{}{
	options := getOptions(data)
	// fmt.Println(options)
	// logger here---
	// reqGorest := globalFunction.MakeAPIRequest(options)

	return options
}

func getOptions(data map[string]interface{}) map[string]interface{}{
	head := map[string]interface{}{
		"Content-Type": "application/json",
	}

	options := map[string]interface{}{
		"method": data["method"],
		"headers": head,
		"url": data["url"],
		"json": true,
		"timeout": 3000,
	}

	return options
}


func GorestGetAllUser() {
	config := helper.ConfigJson()

	data := map[string]interface{}{
		"method": "GET",
		"url": config["gorest"],
	}

	data_res := goRestRequest(data)
	fmt.Println(data_res)

}