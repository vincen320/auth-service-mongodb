package userservice

type GetResponses struct {
	Code    int                           `json:"code,omitempty"`
	Message string                        `json:"message,omitempty"`
	Data    UserDataGetResponseAppService `json:"data,omitempty"`
}
