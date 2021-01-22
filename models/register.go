package models

//RegisterForm model
type RegisterForm struct {
	Name       string `bson:"Name" json:"Name,omitempty"`
	Gender     string `bson:"Gender" json:"Gender,omitempty"`
	Mobile     string `bson:"Mobile" json:"Mobile,omitempty"`
	Email      string `bson:"Email" json:"Email,omitempty"`
	Password   string `bson:"Password" json:"Password,omitempty"`
	RePassword string `json:"RePassword,omitempty"`
}
