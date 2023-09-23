package globalFunction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	// "github.com/nulla-vis/golang-fiber-template/config/constant"
	"github.com/nulla-vis/golang-fiber-template/config/constant"
	"github.com/nulla-vis/golang-fiber-template/core/helper"
)

func GetMessage(code string, replacements interface{}) map[string]interface{}{
	langEn := helper.LangEn()
	langId := helper.LangId()

	mapMessage := map[string]interface{} {
		"code": code,
		"en": code,
		"id": code,
	}

	msgEn := langEn[code]
	msgId := langId[code]

	if msgEn != nil {
		mapMessage["en"] = fmt.Sprint(msgEn)
	}

	if msgId != nil {
		mapMessage["id"] = fmt.Sprint(msgId)
	}

	if replacements == nil {
		// do nothing
	} else if replacements == "" {
        //do nothing
    } else {
		replacementType := reflect.TypeOf(replacements).Kind()

		if replacementType == reflect.String {
			mapMessage["en"] = strings.Replace(fmt.Sprint(mapMessage["en"]), "%s", fmt.Sprint(replacements), 1)
			mapMessage["id"] = strings.Replace(fmt.Sprint(mapMessage["id"]), "%s", fmt.Sprint(replacements), 1)
		} else if replacementType == reflect.Slice {
			s := reflect.ValueOf(replacements)
			ret := make([]interface{}, s.Len())
			for i := 0; i < s.Len(); i++ {
				ret[i] = s.Index(i).Interface()
				mapMessage["en"] = strings.Replace(fmt.Sprint(mapMessage["en"]), "%s", fmt.Sprint(ret[i]), 1)
				mapMessage["id"] = strings.Replace(fmt.Sprint(mapMessage["id"]), "%s", fmt.Sprint(ret[i]), 1)
			}
		}
	}

	return mapMessage
}
// Convert data from database (if the value in byte slice)
func ConvertByteSlicesToStrings(data interface{}) {
    // Check if data is a slice of maps
    maps, isSliceOfMaps := data.([]map[string]interface{})
    if !isSliceOfMaps {
        // Data is not in the expected format
        return
    }

    // Iterate through the maps
    for _, m := range maps {
        for key, value := range m {
            byteSlice, isByteSlice := value.([]byte)
            if isByteSlice {
                // Convert the byte slice to a string
                m[key] = string(byteSlice)
            }
        }
    }
}

// Convert []uint8 to int64
func ConvertBytesToInt64(data []byte) (int64, error) {
    strData := string(data)
    intValue, err := strconv.ParseInt(strData, 10, 64)
    if err != nil {
        return 0, err
    }
    return intValue, nil
}

// Convert []uint8 to float64
func ConvertBytesToFloat64(data []byte) (float64, error) {
    strData := string(data)
    floatValue, err := strconv.ParseFloat(strData, 64)
    if err != nil {
        return 0.0, err
    }
    return floatValue, nil
}

func MakeAPIRequest(data map[string]interface{}) (*http.Response, error) {
    var bodyReader io.Reader

    // Type assertion for method
    method, ok := data["method"].(string)
    if !ok {
        panic("MakeAPIRequest - Method is not a string")
    }

    // Type assertion for URL
    url, ok := data["url"].(string)
    if !ok {
        panic("MakeAPIRequest - URL is not a string")
    }

    // Type assertion for headers
    headersMap, ok := data["headers"].(map[string]interface{})
    if !ok {
        panic("MakeAPIRequest - Headers is not a map")
    }

    headers := make(map[string]string)
    for key, value := range headersMap {
        // Convert headers to map[string]string
        if strValue, ok := value.(string); ok {
            headers[key] = strValue
        } else {
            panic("MakeAPIRequest - Header value is not a string")
        }
    }

    // Serialize the request body to JSON for POST requests
    if method == "POST" {
        requestBodyBytes, err := json.Marshal(data["body"])
        if err != nil {
            return nil, err
        }
        bodyReader = bytes.NewBuffer(requestBodyBytes)
    }

    // Create a context with a timeout
    timeout, ok := data["timeout"].(int)
    if !ok {
        timeout = constant.DEFAULT_TIMEOUT // Default timeout in milliseconds (adjust as needed)
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
    defer cancel()

    // Create an HTTP request with the context
    req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
    if err != nil {
        return nil, err
    }

    // Set headers
    for key, value := range headers {
        req.Header.Set(key, value)
    }

    // Perform the request
    client := &http.Client{}
    response, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    // If "json" is set to true, parse the response body as JSON
    if parseAsJSON, ok := data["json"].(bool); ok && parseAsJSON {
        response, err = parseJSONResponse(response)
        if err != nil {
            return nil, err
        }
    }

    return response, nil
}


// Helper function to parse JSON response
func parseJSONResponse(resp *http.Response) (*http.Response, error) {
    // Read the response body into a byte slice
    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    // Create a new ReadCloser from the byte slice
    resp.Body = io.NopCloser(bytes.NewReader(responseBody))

    return resp, nil
}

// IsEmpty checks if a value is empty based on its type.
func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	default:
		return false
	}
}

func IsValidEmail(email string) bool {
    // Define a regular expression pattern for a valid email address
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    
    // Compile the regular expression
    regex, err := regexp.Compile(pattern)
    if err != nil {
        return false
    }
    
    // Use the regular expression to match the email
    return regex.MatchString(email)
}






