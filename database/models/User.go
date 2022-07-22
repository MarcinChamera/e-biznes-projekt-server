package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email        string
	OAuthService string `json:"o_auth_service"`
	Token        string
}
