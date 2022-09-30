package Common

import (
	"Seller/Seller.Application/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strings"
	"time"
)

func GenerateToken(userId string, expireTime time.Time) (string, string) {
	var secret = []byte("d*&df%324sad@!!ewr")

	// Get the token instance with the Signing method
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Choose an expiration time. Shorter the better
	// Add your claims
	tokenId := uuid.New().String()
	token.Claims = Model.JWTClaimsModel{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Subject:   tokenId,
		}, Model.TokenUserInfo{
			UserId:  userId,
			TokenId: uuid.New().String(),
		},
	}

	// Sign the token with your secret key
	val, err := token.SignedString(secret)

	if err != nil {
		// On error return the error
		return err.Error(), ""
	}
	// On success return the token string
	return val, tokenId
}
func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	var secret = []byte("d*&df%324sad@!!ewr")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func GetClaimValuesFromGin(c *gin.Context) Model.TokenUserInfo {
	token := c.GetHeader("Authorization")
	Claims, _ := GetClaimsFromToken(strings.ReplaceAll(strings.ReplaceAll(token, "bearer ", ""), "Bearer ", ""))
	return Model.TokenUserInfo{
		UserId:  Claims["UserId"].(string),
		TokenId: Claims["TokenId"].(string),
	}
}
