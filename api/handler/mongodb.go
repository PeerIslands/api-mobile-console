package handler

import (
	"github.com/gin-gonic/gin"
	"mongo-admin-backend/api/middleware"
	"mongo-admin-backend/infrastructure/database"
	"mongo-admin-backend/infrastructure/repository"
	"mongo-admin-backend/pkg/crypto"
	"mongo-admin-backend/usecase/login"
	"mongo-admin-backend/usecase/mongoDb"
	"mongo-admin-backend/usecase/user"
	"net/http"
)

func getMongoDbGroups(ctx *gin.Context) {
	email, err := login.GetEmail(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	userRepo := repository.NewUserMongoDB(database.Client)
	userService := user.NewService(userRepo)

	us, err := userService.GetUser(email)
	u := crypto.Decrypt(us.AtlasParams.PublicKey, us.Key)
	p := crypto.Decrypt(us.AtlasParams.PrivateKey, us.Key)
	service := mongoDb.NewService(u, p)
	mongoGroups, err := service.Get()

	ctx.JSON(http.StatusOK, mongoGroups)
}

// MakeMongoHandlers make url handlers.
func MakeMongoHandlers(r *gin.Engine) {
	v1 := r.Group("/v1/mongodb/").Use(middleware.AuthorizeJWT())
	{
		v1.GET("/groups", func(ctx *gin.Context) {
			getMongoDbGroups(ctx)
		})
	}
}

