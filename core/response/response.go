package response

import (
	// "fmt"
	"reflect"
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

	if reflect.ValueOf(data).IsNil() {
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

func ErrorResponse(ctx *fiber.Ctx, errData interface{}) error{

	type ErrorResponse struct {
		Status     string      `json:"status"`
		StatusCode int         `json:"statusCode"`
		Payload    interface{} `json:"payload"`
	}

	if errData == nil {
    	errData = map[string]string{}
	}

	// LOGGER HERE -----
		// Panggil function build log data
			// reqHeaders := ctx.GetReqHeaders()
			// fmt.Println(reqHeaders)

		// insert data tersebut ke dalam log
	// LOGGER END ------

    payload := map[string]interface{}{
        "errors": []map[string]interface{}{
            {
                "message": errData,
            },
        },
    }

	errorData := ErrorResponse{
		Status: "error",
		StatusCode: fiber.StatusInternalServerError,
		Payload: payload,
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(errorData)
}

