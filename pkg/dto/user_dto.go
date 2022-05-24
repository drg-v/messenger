package dto

type UserDto struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
