package middleware

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vincen320/auth-service-mongodb/exception"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(reflect.TypeOf(err)) //debugging
				switch err.(type) {
				case *exception.BadRequestError: //jangan lupa dicek dalam bentuk pointer
					BadRequestErrResponse(c, err)
				case *exception.NotFoundError: //jangan lupa dicek dalam bentuk pointer
					NotFoundErrResponse(c, err)
				case *exception.UnauthorizedError: //jangan lupa dicek dalam bentuk pointer
					UnauthorizedErrResponse(c, err)
				case validator.ValidationErrors:
					ValidationErrorResponse(c, err)
				default:
					c.JSON(http.StatusInternalServerError, gin.H{
						"Status":  http.StatusInternalServerError,
						"Message": "Internal Server Error",
						"Data":    nil,
					})
				}
			}
		}()

		c.Next()
	}
}

func BadRequestErrResponse(c *gin.Context, err interface{}) {
	badRequest, _ := err.(*exception.BadRequestError) //diset ke bentuk pointer
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":  http.StatusBadRequest,
		"Message": badRequest.Error(),
		"Data":    nil,
	})
}

func NotFoundErrResponse(c *gin.Context, err interface{}) {
	notFound, _ := err.(*exception.NotFoundError) //diset ke bentuk pointer
	c.JSON(http.StatusNotFound, gin.H{
		"Status":  http.StatusNotFound,
		"Message": notFound.Error(),
		"Data":    nil,
	})
}

func UnauthorizedErrResponse(c *gin.Context, err interface{}) {
	unauthorized, _ := err.(*exception.UnauthorizedError) //diset ke bentuk pointer
	c.Header("WWW-Authenticate", "realm wrong password")  //jelaskan saja, kenapa mendapatkan error ini seperti usernya tidak boleh ke server ini, "sepertinya sih"(?)
	c.JSON(http.StatusUnauthorized, gin.H{
		"Status":  http.StatusUnauthorized,
		"Message": unauthorized.Error(),
		"Data":    nil,
	})
}

func ValidationErrorResponse(c *gin.Context, err interface{}) {
	validationErr, _ := err.(validator.ValidationErrors)
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":  http.StatusBadRequest,
		"Message": validationErr.Error(),
		"Data":    nil,
	})
}
