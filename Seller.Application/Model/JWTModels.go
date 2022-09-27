package Model

import "github.com/golang-jwt/jwt/v4"

type TokenUserInfo struct {
	UserId  string
	TokenId string
}
type JWTClaimsModel struct {
	*jwt.RegisteredClaims
	TokenUserInfo
}
