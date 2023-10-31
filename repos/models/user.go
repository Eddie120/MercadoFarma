package models

import "time"

type Role string

const (
	AdminRole   Role = "admin"
	ShopperRole Role = "shopper"
)

var IsValidRole = map[Role]bool{
	AdminRole:   true,
	ShopperRole: true,
}

type User struct {
	UserId       string     `json:"user_id"`
	Email        string     `json:"email"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	PhoneNumber  string     `json:"phone_number"`
	Hash         string     `json:"hash"`
	Role         Role       `json:"role"`
	Active       bool       `json:"active"`
	CreationDate *time.Time `json:"creation_date"`
	UpdateDate   *time.Time `json:"update_date"`
}
