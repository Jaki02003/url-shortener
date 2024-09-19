package rest_errors

import (
	"errors"
)

// ResponseCode Map for errors with http code
var ResponseCode map[string]int

func ResponseMap() map[string]int {
	if ResponseCode == nil {
		ResponseCode = make(map[string]int, 0)
	}
	return ResponseCode
}

func NewError(message string, httpCode int) error {
	_, available := ResponseMap()[message]
	if !available {
		ResponseMap()[message] = httpCode
	}
	return errors.New(message)
}
