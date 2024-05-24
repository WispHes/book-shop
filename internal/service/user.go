package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wisphes/book-shop/internal/models"
	"strings"
	"time"
)

const (
	signingKey = "ASDFghj1234%^&*ZXCVbnm"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (int, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GenerateToken(ctx context.Context, email, password string) (string, error) {
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

func (s *UserService) ParseToken(ctx context.Context, accessToken string) (int, error) {
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

func (s *UserService) IsAdmin(ctx context.Context, userId int) error {
	user, err := s.repo.IsAdmin(ctx, userId)
	if err != nil {
		return err
	}
	if !user.IsAdmin {
		return errors.New("user is not admin")
	}
	return nil
}

func (s *UserService) UserIdentity(ctx context.Context, header string) (int, error) {
	if header == "" {
		return 0, errors.New("invalid token")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return 0, errors.New("invalid token")
	}
	userId, err := s.ParseToken(ctx, headerParts[1])
	if err != nil {
		return 0, errors.New("invalid token")
	}
	return userId, err
}
