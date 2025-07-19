package response

import (
	"bytes"
	"encoding/json"
	"fmt"

	mylogger "tampayang-backend/core/logger"

	"github.com/gofiber/fiber/v2"
)

/*
Function berikut untuk build dan handle log data
let getLogData = (req, data) => {
    let IPFromRequest = req.headers['x-forwarded-for'] == undefined ? req.connection.remoteAddress : req.headers['x-forwarded-for'].split(',')[0].trim()
    let indexOfColon = IPFromRequest.lastIndexOf(':')
    let ipaddress = IPFromRequest.substring(indexOfColon + 1, IPFromRequest.length)

    let logdata = {
        "ipaddress": ipaddress || '',
        "route": req.originalUrl,
        "method": req.method || '',
        "headers": req.headers || '',
        "params": req.params || '',
        "body": req.body || '',
        "query": req.query || '',
        "response_data": data || ''
    }
    return logdata
}

*/

func GetLogData(ctx *fiber.Ctx, data interface{}) string {
	var queryData, reqHeadersData, bodyData, logData string
	for key, value := range ctx.Queries() {
		queryData += fmt.Sprintf("%s: %s,", key, value)
	}
	for key, value := range ctx.GetReqHeaders() {
		reqHeadersData += fmt.Sprintf("%s: %s,", key, value)
	}
	if len(ctx.Body()) != 0 {
		var buf bytes.Buffer
		if err := json.Compact(&buf, ctx.Body()); err != nil {
			panic("response - GetLogData, json.Compact: " + err.Error())
		}
		// Convert the buffer to a string
		bodyData = buf.String()

		// Unmarshal the JSON string into a map
		var jsonData map[string]interface{}
		if err := json.Unmarshal([]byte(bodyData), &jsonData); err != nil {
			panic("response - GetLogData, json.Compact: " + err.Error())
		}

		// Marshal the map back into a string without escape characters
		updatedJSONStr, err := json.Marshal(jsonData)
		if err != nil {
			panic("response - GetLogData, json.Compact: " + err.Error())
		}

		// Convert the JSON byte slice to a string
		bodyData = string(updatedJSONStr)
	} else {
		bodyData = ""
	}

	log := map[string]interface{}{
		"ipAddress":     ctx.IP(),
		"route":         ctx.Path(),
		"method":        ctx.Method(),
		"headers":       string(reqHeadersData),
		"params":        convertLoggerData(ctx.AllParams()),
		"body":          bodyData,
		"query":         string(queryData),
		"response_data": convertLoggerData(data),
	}

	for key, value := range log {
		logData += fmt.Sprintf("%s: %s,", key, value)
	}
	return string(logData)
}

func SuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	type SuccessResponse struct {
		Status     string      `json:"status"`
		StatusCode int         `json:"statusCode"`
		Payload    interface{} `json:"payload"`
	}

	if _, ok := data.(string); ok {
		// myVar is a string, and str contains its value
		message := map[string]interface{}{
			"message": "Please use globalFuntion GetMessage",
		}
		return ErrorResponse(ctx, message)
	}

	if data == nil {
		data = map[string]string{}
	}

	// LOGGER HERE -----
	logData := GetLogData(ctx, data)
	mylogger.Trace("response_success", logData)
	// LOGGER END ------

	successData := SuccessResponse{
		Status:     "success",
		StatusCode: fiber.StatusOK,
		Payload:    data,
	}

	return ctx.Status(fiber.StatusOK).JSON(successData)
}

func ErrorResponse(ctx *fiber.Ctx, errData interface{}) error {
	httpCodeErr := fiber.StatusInternalServerError
	type ErrorResponse struct {
		Status     string      `json:"status"`
		StatusCode int         `json:"statusCode"`
		Payload    interface{} `json:"payload"`
	}

	if errData == nil {
		errData = map[string]string{}
	}

	payload := make(map[string]interface{})

	codeOk := false
	enOk := false
	idOk := false

	switch v := errData.(type) {
	case string:
		// If errData is a string, use it as the message
		payload["message"] = v
	case map[string]interface{}:
		// Check for specific keys
		if code, ok := v["code"].(string); ok {
			payload["code"] = code
			if payload["code"] == "err003" {
				httpCodeErr = fiber.StatusNotFound
			}
			codeOk = true
		}
		if en, ok := v["en"].(string); ok {
			payload["en"] = en
			enOk = true
		}
		if id, ok := v["id"].(string); ok {
			payload["id"] = id
			idOk = true
		}

		if !codeOk || !enOk || !idOk {
			payload = v
		}

	case error:
		// If errData is an error, use its error message
		payload["message"] = v.Error()
	default:
		// Handle other data types if needed
		payload["message"] = "An error occurred"
	}

	// LOGGER HERE -----
	logData := GetLogData(ctx, payload)
	mylogger.Error("response_error", logData)
	// LOGGER END ------

	errorData := ErrorResponse{
		Status:     "error",
		StatusCode: fiber.StatusInternalServerError,
		Payload:    payload,
	}

	// Set the status and JSON response data and return ctx
	return ctx.Status(httpCodeErr).JSON(errorData)
}

func convertLoggerData(data interface{}) string {
	// Use a type switch to check the type of data
	var result string
	switch value := data.(type) {
	case map[string]interface{}:
		// Convert the map to a JSON string
		jsonString, _ := json.Marshal(value)
		result = string(jsonString)

	case string:
		// Value is already a string, no need to convert
		result = value
	default:
		jsonString, _ := json.Marshal(value)
		result = string(jsonString)
	}

	return result
}
