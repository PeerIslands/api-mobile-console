package auth

import (
	"github.com/gin-gonic/gin"
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/infrastructure/database"
	"mongo-admin-backend/infrastructure/repository"
	"mongo-admin-backend/usecase/login"
	"mongo-admin-backend/usecase/user"
)

// AuthService interface.
type AuthService interface {
	Login(ctx *gin.Context) string
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

func (a *authService) Login(ctx *gin.Context) string {
	var credential presenter.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	userRepo := repository.NewUserMongoDB(database.Client)
	userService := user.NewService(userRepo)

	u, err := userService.GetUser(credential.Email)
	if err != nil {
		return err.Error()
	}

	err = u.ValidatePassword(credential.Password)
	if err != nil {
		return err.Error()
	}

	return a.jwtService.GenerateToken(credential.Email, true)
}
