package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

var jwtKey = []byte("my_secret")

type Service interface {
	SignUpUser(username, password, firstName, lastName string) error
}

type service struct {
	repository users.Repository
}

func NewService(repository users.Repository) Service {
	return &service{repository: repository}
}

func (s *service) SignUpUser(username, password, firstName, lastName string) error {
	return s.repository.SaveUser(username, password, firstName, lastName)
}

func (s *service) GenerateToken(username string, roles []string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (s *service) ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
