package helper

import (
	"encoding/json"
	"github.com/nulla-vis/golang-fiber-template/config"
	"os"
	// "path/filepath"
)

func ConfigJson() map[string]interface{} {
	rootPath := config.ProjectRootPath
	reader, _ := os.Open(rootPath + "/config/config.json")
	decoder := json.NewDecoder(reader)

	var result map[string]interface{}
	decoder.Decode(&result)

	// fmt.Println(result["development"].(map[string]interface{})["database"])
	return result

}

func LangId() map[string]interface{} {
	rootPath := config.ProjectRootPath
	reader, _ := os.Open(rootPath + "/app/lang/id.json")
	decoder := json.NewDecoder(reader)

	var result map[string]interface{}
	decoder.Decode(&result)

	// fmt.Println(result["development"].(map[string]interface{})["database"])
	return result

}

func LangEn() map[string]interface{} {
	rootPath := config.ProjectRootPath
	reader, _ := os.Open(rootPath + "/app/lang/en.json")
	decoder := json.NewDecoder(reader)

	var result map[string]interface{}
	decoder.Decode(&result)

	// fmt.Println(result["development"].(map[string]interface{})["database"])
	return result

}