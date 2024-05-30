package domain

import "github.com/golang-jwt/jwt/v5"

type Rol struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IpsId       int    `json:"ips_id"`
}

type JwtCustomClaims struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	IpsID int    `json:"ipsId"`
	jwt.RegisteredClaims
}
