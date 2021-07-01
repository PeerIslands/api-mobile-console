package handler

import (
	"errors"
	"mongo-admin-backend/api/middleware"
	"mongo-admin-backend/infrastructure/database"
	"mongo-admin-backend/infrastructure/repository"
	"mongo-admin-backend/pkg/crypto"
	"mongo-admin-backend/usecase/dashboard"
	"mongo-admin-backend/usecase/login"
	"mongo-admin-backend/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMongoDBHome(service dashboard.UseCase, ctx *gin.Context) {
	email, err := login.GetEmail(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	userRepo := repository.NewUserMongoDB(database.Client)
	userService := user.NewService(userRepo)
	params := ctx.Request.URL.Query()
	if val, ok := params["group_id"]; ok {
		us, _ := userService.GetUser(email)
		u := crypto.Decrypt(us.AtlasParams.PublicKey, us.Key)
		p := crypto.Decrypt(us.AtlasParams.PrivateKey, us.Key)

		response, err := service.GetDashboardData(u, p, val[0])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.JSON(http.StatusForbidden, errors.New("group_id is required"))
	}

}

func MakeDashboardHandlers(r *gin.Engine, service dashboard.UseCase) {
	v1 := r.Group("/v1/mongodb").Use(middleware.AuthorizeJWT())
	{
		v1.GET("/home", func(ctx *gin.Context) {
			getMongoDBHome(service, ctx)
		})

	}
}
