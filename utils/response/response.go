package response

import (
	"fmt"
	"net/http"

	"go-redis-url-shortener/rest_errors"
)

type body map[string]interface{}

func ValidationErrors(err error, entity string) (int, body) {
	message := fmt.Sprintf("failed to validate the fields of the %v", entity)
	return validationResponse(message, err)
}

func GenerateErrorResponseBody(err error) (int, body) {
	message := err.Error()
	return readFromMap(message)
}

func readFromMap(message string) (int, body) {
	httpStatus, available := rest_errors.ResponseMap()[message]
	if available {
		return httpStatus, GenerateResponseBody(message)
	}
	return http.StatusInternalServerError, GenerateResponseBody("something went wrong")
}

func GenerateResponseBody(message string) body {
	return body{
		"message": message,
	}
}

func validationResponse(message string, err error) (int, body) {
	return http.StatusBadRequest, body{
		"message":          message,
		"validation_error": err,
	}
}
