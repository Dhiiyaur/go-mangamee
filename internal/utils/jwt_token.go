package utils

import (
	"context"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/appengine/log"
)

const JWT_SECRET = "secret"

type CustomClaim struct {
	jwt.StandardClaims
}

func EncodeToken(jwtClaim jwt.StandardClaims) (string, error) {
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)

	token, err := claim.SignedString([]byte(JWT_SECRET))

	if err != nil {
		log.Errorf(context.Background(), "Error Token %v", err)
		return "", err
	}

	return token, nil
}

func DecodeToken(jwtToken string) (*CustomClaim, error) {
	decodedToken, err := jwt.ParseWithClaims(
		jwtToken,
		&CustomClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET), nil
		},
	)

	if err != nil {
		return nil, err
	}

	result := decodedToken.Claims.(*CustomClaim)

	return result, nil
}

func CheckToken(token string) (string, error) {

	SplitToken := strings.Split(token, " ")
	jwt := SplitToken[1]

	username, err := DecodeToken(jwt)

	// if err != nil {
	// 	return nil, err
	// }

	return username.Issuer, err

}
