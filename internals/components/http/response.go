package handlers

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	Errors                     []Error
	MapCustomMessageValidation map[string]string
)

// NOTE: add more if you need another custom message
var (
	mapCustomMessageValidation = MapCustomMessageValidation{
		"oneof": "The input must be one of the specified options",
	}
)

type Error struct {
	Field    string   `json:"field" example:"currency"`
	Messages []string `json:"messages" example:"required"`
}

func (es Errors) Error() string {
	var buf bytes.Buffer
	for _, e := range es {
		buf.WriteString(fmt.Sprintf("%s: %s\n", e.Field, strings.Join(e.Messages, ",")))
	}
	return buf.String()
}

/*
func ErrorValidator(vErrs validate.ValidationErrors) (newError Errors) {
	for _, vErr := range vErrs {
		newError = append(newError, Error{
			Field:    vErr.Field(),
			Messages: []string{msgForTag(vErr.Tag())},
		})
	}

	return newError
}
*/

func msgForTag(tag string) string {
	if msg, ok := mapCustomMessageValidation[tag]; ok {
		return msg
	}

	return tag
}

// Response is a generic response
// errors field will only availale if response is considered as an error.
//
// When a request is not successful, the errors field can be an
// empty string; on that occasion, the reason is explained in the message field.
//
// The errors field will not present in the response if and only if the request is accepted
// and can be processed by them.
// Errors is a list of errors found at response payload
// if a call returns an error, it should be checked whether or not
// it an instance of this type.
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message" example:"OK/ERROR"`
	Data    interface{} `json:"data"`
	Errors  Errors      `json:"errors,omitempty"`
}

type ResponseErrorWithData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  Errors      `json:"errors,omitempty"`
}

func SuccessResponse(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, Response{
		Data:    data,
		Code:    code,
		Message: "OK",
	})
}

func ErrorResponse(
	c echo.Context,
	code int,
	message string,
	errors []Error,
) error {
	return c.JSON(code, Response{
		Code:    code,
		Message: message,
		Errors:  errors,
	})
}
