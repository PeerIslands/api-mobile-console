package handler

import (
	"errors"
	"fmt"
	"mongo-admin-backend/api/middleware"
	"mongo-admin-backend/config"
	"mongo-admin-backend/infrastructure/database"
	"mongo-admin-backend/infrastructure/repository"
	"mongo-admin-backend/pkg/crypto"
	"mongo-admin-backend/usecase/login"
	"mongo-admin-backend/usecase/mongoDbAtlas/group"
	process "mongo-admin-backend/usecase/mongoDbAtlas/process"
	processMeasurements "mongo-admin-backend/usecase/mongoDbAtlas/processMeasurement"
	"mongo-admin-backend/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
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

//Function to get the measurement from Altas by processID
func getMongoDbProcessMeasurement(ctx *gin.Context) {
	email, err := login.GetEmail(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	userRepo := repository.NewUserMongoDB(database.Client)
	userService := user.NewService(userRepo)
	params := ctx.Request.URL.Query()
	//fmt.Println(params)
	groupId, isGroupIdPresent := params[config.PARAM_GROUP]
	processId, isProcessIdPresent := params[config.PARAM_PROCESS]
	if isGroupIdPresent && isProcessIdPresent {
		fmt.Println(params[config.PARAM_MEASUREMENT])
		var paramMap = make(map[string][]string)

		if val, ok := params[config.PARAM_GRANULARITY]; ok {
			paramMap[config.PARAM_GRANULARITY] = val
		}
		if val, ok := params[config.PARAM_PERIOD]; ok {
			paramMap[config.PARAM_PERIOD] = val
		}
		if val, ok := params[config.PARAM_ST_DATE]; ok {
			paramMap[config.PARAM_ST_DATE] = val
		}
		if val, ok := params[config.PARAM_END_DATE]; ok {
			paramMap[config.PARAM_END_DATE] = val
		}
		if val, ok := params[config.PARAM_MEASUREMENT]; ok {
			paramMap[config.PARAM_MEASUREMENT] = val
		}

		us, _ := userService.GetUser(email)
		u := crypto.Decrypt(us.AtlasParams.PublicKey, us.Key)
		p := crypto.Decrypt(us.AtlasParams.PrivateKey, us.Key)
		service := processMeasurements.NewService(u, p, groupId[0], processId[0], paramMap)
		mongoGroups, errorDetail, err := service.Get()
		if err != nil {
			fmt.Println("Error in calling the API 2")
			fmt.Println(err)
			fmt.Println(errorDetail)
			ctx.JSON(errorDetail.Error, errorDetail)
		} else {
			ctx.JSON(http.StatusOK, mongoGroups)
		}

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
		v1.GET("/process/measurements", func(ctx *gin.Context) {
			//Call the the function here to all the MongoDB Atlas API
			getMongoDbProcessMeasurement(ctx)
		})
	}
}
