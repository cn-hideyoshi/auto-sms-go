package utils

import (
	"crypto/ecdsa"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type JWTUtils struct {
	Claims JwtClaims
	Method jwt.SigningMethod
	Secret ecdsa.PublicKey
}

var secret = []byte("blog.hideyoshi.top")

type JwtClaims struct {
	Data interface{}
	jwt.RegisteredClaims
}

func (j JWTUtils) Encode() (string, error) {
	token := jwt.NewWithClaims(j.Method, j.Claims)
	signingString, err := token.SignedString(secret)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return signingString, err
}

func (j JWTUtils) Decode(tokenStr string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
