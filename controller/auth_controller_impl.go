package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vincen320/auth-service-mongodb/exception"
	"github.com/vincen320/auth-service-mongodb/model/web"
	"github.com/vincen320/auth-service-mongodb/service"
)

type AuthControllerImpl struct {
	Service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &AuthControllerImpl{
		Service: service,
	}
}

func (ac *AuthControllerImpl) Login(c *gin.Context) {
	var authRequest web.AuthRequest
	err := c.ShouldBind(&authRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error())) //Bad Request
	}
	jwtToken, err := ac.Service.Login(authRequest) //tidak penting sebenarnya err disini, karena difungsi Implnya langsung panic jika err, jadi sebaiknya dihilangkan saja(cuman malas saja nak ngilanginny) || mungkin kedepanny biso return error bae daripada panic diservice, tetapi kalau tidak panic diservice, kt tidak tau spesifik error pada baris mana yang menyebabkan error tsb(?) >>PENJELASAN DIBAWAH
	/**
	You should assume that a panic will be immediately fatal, for the entire program, or at the very least for the current goroutine.
	Ask yourself "when this happens, should the application immediately crash?" If yes, use a panic; otherwise, use an error.
		**/
	if err != nil {
		panic(exception.NewBadRequestError("Gagal Login")) //Bad Request
	}

	c.Header("Authorization", "Bearer "+jwtToken)
	c.JSON(200, gin.H{
		"Status":  200,
		"Message": "Behasil Login",
	})
}

func (ac *AuthControllerImpl) Refresh(c *gin.Context) {
	tokenStr := c.Query("token")

	newToken := ac.Service.Refresh(tokenStr)
	c.Header("Authorization", "Bearer "+newToken)
	c.JSON(200, gin.H{
		"Status":  200,
		"Message": "Token refreshed",
	})
}
