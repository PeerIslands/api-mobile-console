package handler

import (
	"errors"
	"mongo-admin-backend/api/middleware"
	"mongo-admin-backend/config"
	"mongo-admin-backend/infrastructure/database"
	"mongo-admin-backend/infrastructure/repository"
	"mongo-admin-backend/pkg/crypto"
	"mongo-admin-backend/usecase/accesslist"
	dbaccesslist "mongo-admin-backend/usecase/dbAccesslist"

	"mongo-admin-backend/usecase/login"
	"mongo-admin-backend/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Function to get the list of Users access request and from the existing user list

func getAccessRequests(service accesslist.UseCase, ctx *gin.Context) {
	listval, _ := service.GetAllNetworkAccessList()
	ctx.JSON(http.StatusOK, &listval)
}

func getDBAccessRequests(service dbaccesslist.UseCase, ctx *gin.Context) {
	listval, _ := service.GetAllDBAccessList()
	ctx.JSON(http.StatusOK, &listval)
}

func createNetworkAccess(service accesslist.UseCase, ctx *gin.Context) {
	email, err := login.GetEmail(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	params := ctx.Request.URL.Query()

	id, isIdPresent := params["id"]

	if isIdPresent {
		accessRequest, err := service.GetOneNetworkAccessRequest(id[0])
		//GET USER KEYS
		userRepo := repository.NewUserMongoDB(database.Client)
		userService := user.NewService(userRepo)
		us, _ := userService.GetUser(email)
		u := crypto.Decrypt(us.AtlasParams.PublicKey, us.Key)
		p := crypto.Decrypt(us.AtlasParams.PrivateKey, us.Key)
		if accessRequest != nil {
			resp, errordetail, err := service.CreateNetworkAccessRequest(u, p, accessRequest)
			if err != nil {
				ctx.JSON(errordetail.Error, errordetail.ErrorCode)
			} else {
				//Update the status as closed.
				stat, _ := service.UpdateNetworkRequestStatus(u, p, id[0], config.STR_REQ_STATUS_CLOSED)
				if stat {
				}
				ctx.JSON(http.StatusOK, &resp)
			}
		} else {
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
			} else {
				//ctx.JSON(http.StatusOK, "The ID is not present")
			}
		}

		//ctx.JSON(http.StatusOK, "{status : ok}")
	} else {
		ctx.JSON(http.StatusBadRequest, errors.New("Failed"))
	}

}

func createDBAccessRequests(service dbaccesslist.UseCase, ctx *gin.Context) {
	email, err := login.GetEmail(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	params := ctx.Request.URL.Query()

	id, isIdPresent := params["id"]
	if isIdPresent {
		accessRequest, err := service.GetOneDBAccessRequest(id[0])
		//GET USER KEYS
		userRepo := repository.NewUserMongoDB(database.Client)
		userService := user.NewService(userRepo)
		us, _ := userService.GetUser(email)
		u := crypto.Decrypt(us.AtlasParams.PublicKey, us.Key)
		p := crypto.Decrypt(us.AtlasParams.PrivateKey, us.Key)
		if accessRequest != nil {
			resp, errordetail, err := service.CreateDBAccessRequest(u, p, accessRequest)
			if err != nil {
				ctx.JSON(errordetail.Error, errordetail.ErrorCode)

			} else {
				stat, _ := service.UpdateDBAccessRequestStatus(u, p, id[0], config.STR_REQ_STATUS_CLOSED)
				if stat {

				}
				ctx.JSON(http.StatusOK, &resp)
			}
		} else {
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
			} else {
				//ctx.JSON(http.StatusOK, "The ID is not present")
			}
		}

		//ctx.JSON(http.StatusOK, "{status : ok}")
	} else {
		ctx.JSON(http.StatusBadRequest, errors.New("Failed"))
	}

}

// MakeMongoHandlers make url handlers.
func MakeMongoAccessHandlers(r *gin.Engine, service accesslist.UseCase) {
	v1 := r.Group("/v1/mongodb/access").Use(middleware.AuthorizeJWT())
	{
		v1.GET("/list", func(ctx *gin.Context) {
			getAccessRequests(service, ctx)
		})
		v1.POST("/request", func(ctx *gin.Context) {
			createNetworkAccess(service, ctx)
		})

	}
}

// MakeMongoHandlers make url handlers.
func MakeMongoDBAccessHandlers(r *gin.Engine, service dbaccesslist.UseCase) {
	v1 := r.Group("/v1/mongodb/dbaccess").Use(middleware.AuthorizeJWT())
	{
		v1.GET("/list", func(ctx *gin.Context) {
			getDBAccessRequests(service, ctx)
		})
		v1.POST("/request", func(ctx *gin.Context) {
			createDBAccessRequests(service, ctx)
		})

	}
}
