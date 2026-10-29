package usecase

import (
	"clean-go/internal/entity"
	"clean-go/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(string, string) (string, error)
	Register(string, string, string) error
}

type authUsecase struct {
	repo      repository.UserRepository
	jwtSecret string
}

func NewAuthUsecase(repo repository.UserRepository, jwtSecret string) AuthUsecase {
	return &authUsecase{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (a *authUsecase) Login(email string, password string) (string, error) {
	user, err := a.repo.GetByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")

	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return signed, nil
}

func (a *authUsecase) Register(username string, email string, password string) error {
	findUser, err := a.repo.GetByEmail(email)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if findUser != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &entity.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = a.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
