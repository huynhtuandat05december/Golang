package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}
