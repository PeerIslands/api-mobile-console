package main

import (
	"fmt"
	"io"
	"mongo-admin-backend/api/handler"
	"mongo-admin-backend/config"
	"mongo-admin-backend/infrastructure/database"
	"mongo-admin-backend/infrastructure/repository"
	"mongo-admin-backend/pkg/contextWrapper"
	"mongo-admin-backend/usecase/auth"
	"mongo-admin-backend/usecase/login"
	"mongo-admin-backend/usecase/user"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Start()
	contextWrapper.Start()
	defer contextWrapper.Cancel()
	database.ConnectDatabase()

	loginService := login.StaticLoginService()
	jwtService := login.JWTAuthService()
	authService := auth.LoginHandler(loginService, jwtService)
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.New()

	r.Use(cors.Default())

	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	handler.AddNonAuthRoutes(r, authService)
	userRepo := repository.NewUserMongoDB(database.Client)
	userService := user.NewService(userRepo)

	//user
	handler.MakeUserHandlers(r, userService)
	handler.MakeUserNoAuthHandlers(r, userService)
	handler.MakeMongoHandlers(r)
	fmt.Println("Server is running at: " + config.ENVIRONMENT.API_PORT)
	r.Run(":" + config.ENVIRONMENT.API_PORT)
}
