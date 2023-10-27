package usecases

import (
	"errors"
	"os"
	"time"

	"github.com/arifpermadi999/core-echo-golang/dto"
	"github.com/arifpermadi999/core-echo-golang/helpers"
	"github.com/arifpermadi999/core-echo-golang/models"
	"github.com/arifpermadi999/core-echo-golang/repositories"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	UserRepository *repositories.UserRepository
}
type AuthUsecaseContract interface {
	Login(dto *dto.AuthDto) (*models.ResponseAuth, error)
	Register(dto *dto.RegisterUser) (*models.User, error)
}

func ProviderAuthUsecase(ar *repositories.UserRepository) AuthUsecase {
	return AuthUsecase{UserRepository: ar}
}

func (au *AuthUsecase) Login(dto *dto.AuthDto) (*models.ResponseAuth, error) {

	helpers.GetEnv()
	hashedSecret := os.Getenv("HASHED_SECRET")
	user := new(models.User)
	user.Email = dto.Email
	user, err := au.UserRepository.FindUser(user)
	responseAuth := new(models.ResponseAuth)
	if err != nil {
		return responseAuth, err
	}
	hasPassword, err := helpers.CheckPasswordHas(dto.Password, user.Password)
	if err != nil {
		return responseAuth, err
	}
	if !hasPassword {
		err := errors.New("username dan password salah")
		return responseAuth, err
	}
	jwtClaims := &models.JwtCustomClaims{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role, StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// encode it as before, but with your created type
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = jwtClaims
	tokenString, err := t.SignedString([]byte(hashedSecret))

	if err != nil {
		return responseAuth, err
	}
	responseAuth.Name = user.Name
	responseAuth.Role = user.Role
	responseAuth.Token = tokenString

	return responseAuth, nil
}

func (au *AuthUsecase) Register(dto *dto.RegisterUser) (*models.User, error) {

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := new(models.User)
	if dto.Password != dto.Cpassword {
		err := errors.New("password and confirm password not match")
		return user, err
	}
	user.Email = dto.Email
	user.Name = dto.Name
	user.Password = string(hashedPassword)
	user.Role = dto.Role

	data, err := au.UserRepository.SaveOrUpdateUser(user)
	if err != nil {
		return data, err
	}
	return data, nil
}
