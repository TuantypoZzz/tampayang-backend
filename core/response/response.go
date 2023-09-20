package response

import (
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

func SuccessResponse(ctx *fiber.Ctx, data interface{}) error{

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
		// Panggil function build log data
			// reqHeaders := ctx.GetReqHeaders()
			// fmt.Println(reqHeaders)

		// insert data tersebut ke dalam log
	// LOGGER END ------
	

	successData := SuccessResponse {
		Status: "success",
		StatusCode: fiber.StatusOK,
		Payload: data,
	}

	return ctx.Status(fiber.StatusOK).JSON(successData)
}

func ErrorResponse(ctx *fiber.Ctx, errData interface{}) error {
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

    errorData := ErrorResponse{
        Status:     "error",
        StatusCode: fiber.StatusInternalServerError,
        Payload:    payload,
    }

    // Set the status and JSON response data and return ctx
    return ctx.Status(fiber.StatusInternalServerError).JSON(errorData)
}





