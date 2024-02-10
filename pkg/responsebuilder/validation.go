package responsebuilder

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

type JSONFormatter struct{}

// NewJSONFormatter will create a new JSON formatter and register a custom tag
// name func to gin's validator
func NewJSONFormatter() *JSONFormatter {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return &JSONFormatter{}
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (JSONFormatter) Descriptive(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Message: err})
	}

	return errs
}

func (JSONFormatter) Simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		param := strcase.ToSnake(f.Field())
		errs[param] = msgReadable(err, param)
	}

	return errs
}

func msgReadable(tag string, param string) string {
	switch tag {
	case "required":
		return "Kolom " + param + " wajib diisi"
	case "email":
		return "Email tidak valid"
	}
	return ""
}

func CatchError(errDTO error, ctx *gin.Context) Response {
	if errDTO != nil {
		if errDTO == io.EOF {
			response := BuildErrorResponse("please fill the json", "", EmptyObj{})
			return response
		}
		var verr validator.ValidationErrors
		if errors.As(errDTO, &verr) {
			response := BuildErrorResponse("Failed to process request", NewJSONFormatter().Simple(verr), EmptyObj{})
			return response
		} else {
			response := BuildErrorResponse("Failed to process request", errDTO.Error(), EmptyObj{})
			return response
		}
	}
	return Response{Status: true}

}
