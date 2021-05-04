package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mongo-admin-backend/api/middleware"
	"mongo-admin-backend/infrastructure/database"
	"mongo-admin-backend/infrastructure/repository"
	"mongo-admin-backend/pkg/crypto"
	"mongo-admin-backend/usecase/login"
	"mongo-admin-backend/usecase/mongoDbAtlas/group"
	process "mongo-admin-backend/usecase/mongoDbAtlas/process"
	"mongo-admin-backend/usecase/user"
	"net/http"
)

func getMongoDbProcess(ctx *gin.Context) {
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
		fmt.Println(val)
		us, _ := userService.GetUser(email)
		u := crypto.Decrypt(us.AtlasParams.PublicKey, us.Key)
		p := crypto.Decrypt(us.AtlasParams.PrivateKey, us.Key)
		service := process.NewService(u, p, val[0])
		mongoGroups, _ := service.Get()

		ctx.JSON(http.StatusOK, mongoGroups)
	} else {
		ctx.JSON(http.StatusForbidden, errors.New("group_id is required"))
	}

}

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
	service := group.NewService(u, p)
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
		v1.GET("/process", func(ctx *gin.Context) {
			getMongoDbProcess(ctx)
		})
	}
}
