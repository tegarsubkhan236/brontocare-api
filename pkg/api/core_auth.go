package api

import (
	"github.com/golang-jwt/jwt/v4"
	"hospital-api/pkg/api/helper"
	"hospital-api/pkg/repository/model"
	"os"
	"time"
)

// AuthService contains the methods of the service
type AuthService interface {
	Login(input helper.LoginInput) (string, error)
}

// AuthRepository is what lets our service do db operations without knowing anything about the implementation
type AuthRepository interface {
	CheckPasswordHash(password, hash string) bool
	HashPassword(password string) (string, error)
	ValidToken(t *jwt.Token, id string) bool
	GetUserByEmail(email string) (model.CoreUser, error)
	GetUserByUsername(username string) (model.CoreUser, error)
}

type authService struct {
	storage AuthRepository
}

func (a authService) Login(input helper.LoginInput) (string, error) {
	singleUser := model.CoreUser{}
	username, err := a.storage.GetUserByUsername(input.Identity)
	if err != nil {
		return "Error on username", err
	}

	if username.ID != 0 {
		singleUser = model.CoreUser{
			ID:         username.ID,
			Username:   username.Username,
			Email:      username.Email,
			Password:   username.Password,
			Status:     username.Status,
			Role:       username.Role,
			Permission: username.Permission,
		}
	} else {
		return "User not found", err
	}

	if singleUser.Status != 1 {
		return "User not active", err
	}

	if !a.storage.CheckPasswordHash(input.Password, singleUser.Password) {
		return "Invalid password", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = singleUser.ID
	claims["username"] = singleUser.Username
	claims["role"] = singleUser.Role
	claims["permission"] = singleUser.Permission
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t, err
}

func NewAuthService(authRepo AuthRepository) AuthService {
	return &authService{
		storage: authRepo,
	}
}
