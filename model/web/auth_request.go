package web

type AuthRequest struct {
	Username string `validate:"required,max=20,min=6" json:"username,omitempty"`
	Password string `validate:"required,max=20,min=6" json:"password,omitempty"`
}
