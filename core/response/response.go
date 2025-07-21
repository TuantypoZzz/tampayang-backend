package response

import (
	"bytes"
	"encoding/json"
	"strings"

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

func GetLogData(ctx *fiber.Ctx, responseData interface{}) string {
	var bodyData string

	// Periksa apakah ada body dalam request
	if len(ctx.Body()) > 0 {
		// Periksa Content-Type header
		contentType := ctx.Get(fiber.HeaderContentType)

		if strings.Contains(contentType, fiber.MIMEApplicationJSON) {
			// Jika body adalah JSON, kita coba 'compact' agar rapi di log
			var buf bytes.Buffer
			if err := json.Compact(&buf, ctx.Body()); err == nil {
				bodyData = buf.String()
			} else {
				// Jika gagal, catat sebagai string mentah
				bodyData = string(ctx.Body())
			}
		} else {
			// Jika BUKAN JSON (misalnya: form-data), catat sebagai string mentah
			bodyData = string(ctx.Body())
		}
	}

	// Siapkan map untuk data log yang terstruktur
	logMap := map[string]interface{}{
		"ipAddress":    ctx.IP(),
		"route":        ctx.Path(),
		"method":       ctx.Method(),
		"headers":      ctx.GetReqHeaders(), // Simpan sebagai map, bukan string
		"params":       ctx.AllParams(),
		"body":         bodyData,
		"query":        ctx.Queries(),
		"responseData": responseData, // Data respons dari handler
	}

	// Ubah seluruh map log menjadi satu string JSON
	logBytes, err := json.Marshal(logMap)
	if err != nil {
		// Fallback jika proses marshal gagal
		return `{"error": "failed to marshal log data"}`
	}

	return string(logBytes)
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
