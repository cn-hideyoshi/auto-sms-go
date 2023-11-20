package utils

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type JWTUtils struct {
	Claims jwt.MapClaims
	Method jwt.SigningMethod
	Secret ecdsa.PublicKey
}

func (j JWTUtils) Encode() (string, error) {
	token := jwt.NewWithClaims(j.Method, j.Claims)
	//secret := []byte("hideyoshi.top")
	signingString, err := token.SigningString()
	if err != nil {
		return "", err
	}
	return signingString, err
}

func (j JWTUtils) Decode(tokenStr string) (jwt.MapClaims, error) {
	secret := []byte("hideyoshi.top")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if token.Valid {
		return nil, errors.New("invalid token")
	}
	// Get the claims from the token
	claims := token.Claims.(jwt.MapClaims)
	// Print the claims
	return claims, nil
}
