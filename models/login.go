package models

//LoginForm model
type LoginForm struct {
	Email    string `json:"Email,omitempty"`
	Password string ` json:"Password,omitempty"`
}
