package jwtx

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	AccessSecret string
	AccessExpire time.Duration
}

func (j *JWT) Gen(u User) (string, error) {
	now := time.Now()
	claims := Claims{
		User: u,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  []string{""},
			ExpiresAt: jwt.NewNumericDate(now.Add(j.AccessExpire)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.AccessSecret))
}

func (j *JWT) Verify(token string) (*Claims, error) {
	if tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.AccessSecret), nil
	}); err != nil {
		return nil, err
	} else {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		} else {
			return nil, err
		}
	}
}
