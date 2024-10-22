package model

type User struct {
	UserID      string `json:"userId,omitempty"`
	Username    string `json:"username,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phonenumber,omitempty"`
}
