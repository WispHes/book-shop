package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wisphes/book-shop/internal/models"
	"github.com/wisphes/book-shop/internal/repository"
	"time"
)

const signingKey = "ASDFghj1234%^&*ZXCVbnm"

type Authorization interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GenerateToken(ctx context.Context, email, password string) (string, error)
	ParseToken(ctx context.Context, token string) (int, error)
	IsAdmin(ctx context.Context, id int) (bool, error)
}

type AuthService struct {
	repo repository.AuthPostgres
}

func NewAuthService(repo *repository.AuthPostgres) *AuthService {
	return &AuthService{repo: *repo}
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) CreateUser(ctx context.Context, user models.User) (int, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *AuthService) GenerateToken(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUser(ctx, email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&tokenClaims{
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
			user.Id,
		},
	)

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(ctx context.Context, accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) IsAdmin(ctx context.Context, id int) (bool, error) {
	user, err := s.repo.IsAdmin(ctx, id)
	if err != nil {
		return false, err
	}
	if !user.IsAdmin {
		return false, errors.New("user is not admin")
	}
	return true, nil
}
