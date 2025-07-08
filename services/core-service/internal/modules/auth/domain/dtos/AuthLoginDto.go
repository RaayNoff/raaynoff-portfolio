package authDtos

type AuthLoginDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
