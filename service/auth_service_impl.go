package service

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/vincen320/auth-service-mongodb/exception"
	"github.com/vincen320/auth-service-mongodb/helper"
	"github.com/vincen320/auth-service-mongodb/model/appservice/userservice"
	"github.com/vincen320/auth-service-mongodb/model/web"
)

type AuthServiceImpl struct {
	Validator *validator.Validate
}

func NewAuthService(validator *validator.Validate) AuthService {
	return &AuthServiceImpl{
		Validator: validator,
	}
}

func (as *AuthServiceImpl) Login(authRequest web.AuthRequest) (string, error) {
	err := as.Validator.Struct(authRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}
	/*	CONTOH KIRIM DENGAN BODY
		jsonBody, err := json.Marshal(authRequest) //jsonBody dalam bentuk []byte
		if err != nil {
			panic("unable encode to json")
		}

		bodyReader := bytes.NewReader(jsonBody)

		request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/"+authRequest.Username, bodyReader)
	*/

	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/username/"+authRequest.Username, nil)
	if err != nil {
		panic("Internal Server Error") // 500 Internal Server Error
	}

	//request.Header.Set("Content-Type", "application/json") //Contoh: Cara set header saat request

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		panic("unable to get user request") // 500 Internal Server Error
	}

	switch response.StatusCode {
	case 500:
		panic("Internal Server Error") // 500 Internal Server Error
	case 404:
		panic(exception.NewNotFoundError("user not found")) // 404 User Not Found
	}
	//response.Header.Get("Nama-Header") //Contoh: Cara mengambil nilai header

	defer response.Body.Close()

	// bodii, _ := io.ReadAll(response.Body) //debugging
	// fmt.Println(string(bodii)) // debugging

	var userServiceResponse userservice.GetResponses
	err = json.NewDecoder(response.Body).Decode(&userServiceResponse)
	if err != nil {
		panic("unable converting data from user-service to json") // 500 Internal Server Error
	}

	login := helper.ComparePassword(userServiceResponse.Data.Password, authRequest.Password)
	if !login {
		panic(exception.NewUnauthorizedError("wrong password")) // 401 Unauthorized
	}
	//Buat JWT
	jwtToken := helper.GenerateJWT(userServiceResponse.Data)
	return jwtToken, nil
}

func (as *AuthServiceImpl) Refresh(tokenStr string) string {
	return helper.RefreshToken(tokenStr)
}
