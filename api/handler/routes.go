package handler

import (
	"mongo-admin-backend/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddNonAuthRoutes(r *gin.Engine, authService auth.AuthService) {

	v1 := r.Group("/v1")
	{
		v1.GET("/health-check", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
		v1.POST("/login", func(ctx *gin.Context) {
			token, err := authService.Login(ctx)
			if err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(err.Code, gin.H{
					"status": err.Msg,
				})
			}
		})
	}
}

func AddAuthRoutes(r gin.Engine) {
	//
}
