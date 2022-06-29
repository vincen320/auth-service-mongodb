package userservice

type UserDataGetResponseAppService struct {
	Id         string `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	CreatedAt  int64  `json:"createdAt,omitempty"`
	LastOnline int64  `json:"lastOnline,omitempty"`
}
