package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken create a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payLoad, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	payloadJWT := NewPayloadJWT(payLoad)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payloadJWT)

	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	payload := &PayloadJWT{}
	jwtToken, err := jwt.ParseWithClaims(token, payload, keyFunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		} else {
			return nil, ErrInvalidToken
		}
	}

	payloadClaims, ok := jwtToken.Claims.(*PayloadJWT)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payloadClaims.Payload, nil
}
