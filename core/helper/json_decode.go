package helper

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/nulla-vis/golang-fiber-template/config"
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

// MakeReader converts a map to an io.Reader containing JSON data.
func MakeReader(data map[string]interface{}) io.Reader {
    // Serialize the data to JSON
    jsonData, _ := json.Marshal(data)
    
    // Create a reader from the JSON data
    reader := bytes.NewReader(jsonData)
    
    return reader
}