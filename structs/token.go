package structs

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	ClaimTypeUser  claimType = "user"
	ClaimTypeAdmin claimType = "admin"
)

type claimType string

type Token struct {
	AccessToken  string                 `json:"access_token" bson:"access_token"`
	RefreshToken string                 `json:"refresh_token" bson:"refresh_token"`
	ClaimType    claimType              `json:"-" bson:"claim_type"`
	OtherInfo    map[string]interface{} `json:"-" bson:"other_info"`
}

type UserTokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"userId"`
	Status UserStatus
}
