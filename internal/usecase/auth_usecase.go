package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
	"github.com/zhalisher/ip-task-manager/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepo  repository.UserRepository
	jwtSecret string
	jwtExpMin int
}

func NewAuthUsecase(
	userRepo repository.UserRepository,
	jwtSecret string,
	jwtExpMin int) *authUsecase {
	return &authUsecase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		jwtExpMin: jwtExpMin,
	}
}

func (u *authUsecase) Register(ctx context.Context, email, password, name string) (*model.User, error) {
	_, err := u.userRepo.GetByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Email:    email,
		Password: string(hash),
		Name:     name,
	}
	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
func (u *authUsecase) Login(ctx context.Context, email, password string) (accessToken, refreshToken string, err error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}
	accessToken, err = generateToken(user.ID, u.jwtSecret, u.jwtExpMin)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = generateToken(user.ID, u.jwtSecret, u.jwtExpMin)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func generateToken(userID uuid.UUID, secret string, expMin int) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(time.Duration(expMin) * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
