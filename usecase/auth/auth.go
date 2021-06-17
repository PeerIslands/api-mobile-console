package auth

import (
	"errors"
	"mongo-admin-backend/api/presenter"

	"mongo-admin-backend/infrastructure/database"
	"mongo-admin-backend/infrastructure/repository"
	"mongo-admin-backend/usecase/login"
	"mongo-admin-backend/usecase/user"

	"github.com/gin-gonic/gin"
)

// AuthService interface.
type AuthService interface {
	Login(ctx *gin.Context) (string, *presenter.AppHTTPError)
}

type authService struct {
	loginService login.LoginService
	jwtService   login.JWTService
}

func LoginHandler(loginService login.LoginService,
	jwtService login.JWTService) AuthService {
	return &authService{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (a *authService) Login(ctx *gin.Context) (string, *presenter.AppHTTPError) {
	var credential presenter.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "", &presenter.AppHTTPError{
			Msg:  "Credentials Null",
			Code: 400,
			Err:  errors.New("Credentials Null"),
		}
	}
	userRepo := repository.NewUserMongoDB(database.Client)
	userService := user.NewService(userRepo)

	u, err := userService.GetUser(credential.Email)
	if err != nil {
		return "", &presenter.AppHTTPError{
			Msg:  "Invalid Username or password",
			Code: 400,
			Err:  errors.New("Unavailable"),
		}
	}

	err = u.ValidatePassword(credential.Password)
	if err != nil {
		return "", &presenter.AppHTTPError{
			Msg:  "Invalid Username or password",
			Code: 400,
			Err:  errors.New("Unavailable"),
		}
	}
	return a.jwtService.GenerateToken(credential.Email, true), nil
}
