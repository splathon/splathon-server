package swagutils

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/splathon"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/result"
)

func Error(err error) middleware.Responder {
	code := 500
	msg := err.Error()
	if splathonErr, ok := err.(*splathon.Error); ok {
		code = splathonErr.Code
		msg = splathonErr.Message
	}
	return result.NewGetResultDefault(code).WithPayload(&models.Error{
		Code:    int64(code),
		Message: swag.String(msg),
	})
}
