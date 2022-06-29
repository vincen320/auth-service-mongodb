package service

import (
	"github.com/vincen320/auth-service-mongodb/model/web"
)

type AuthService interface {
	Login(authRequest web.AuthRequest) (string, error)
	Refresh(tokenStr string) string
}
