package handler

import (
	"mongo-admin-backend/api/middleware"
	"mongo-admin-backend/usecase/createcluster"

	"github.com/gin-gonic/gin"
)

func MakeMongoDDClusterHandler(r *gin.Engine, service createcluster.UseCase) {
	v1 := r.Group("/v1/mongodb/cluster").Use(middleware.AuthorizeJWT())
	{
		v1.GET("/list", func(ctx *gin.Context) {
			//Call the list use case here
		})
		v1.POST("/request", func(ctx *gin.Context) {
			//Call the create use case here
		})
	}
}
