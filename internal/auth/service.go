package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

var jwtKey = []byte("my_secret")

type Service interface {
	SignUpUser(newUser *users.User) (*string, error)
	LoginUser(username, password string) (*string, error)
}

type service struct {
	repository users.Repository
}

func NewService(repository users.Repository) Service {
	return &service{repository: repository}
}

func (s *service) SignUpUser(newUser *users.User) (*string, error) {
	err := s.repository.SaveUser(newUser)

	if err != nil {
		return nil, err
	}

	roles := []string{}
	token, err := s.generateToken(newUser.Username, roles)
	return &token, err
}

func (s *service) LoginUser(username, password string) (*string, error) {
	if doesExist := s.repository.CheckUserCredentials(username, password); !doesExist {
		return nil, shared.ErrUnauthorized
	}

	roles := []string{}
	token, err := s.generateToken(username, roles)
	return &token, err
}

func (s *service) generateToken(username string, roles []string) (string, error) {
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

func (s *service) parseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
