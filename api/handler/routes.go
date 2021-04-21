package handler

import (
	"github.com/gin-gonic/gin"
	"mongo-admin-backend/usecase/auth"
	"net/http"
)

func AddNonAuthRoutes(r *gin.Engine, authService auth.AuthService) {

	v1 := r.Group("/v1")
	{
		v1.GET("/health-check", func (ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
		v1.POST("/login", func(ctx *gin.Context) {
			token := authService.Login(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, nil)
			}
		})
	}
}

func AddAuthRoutes(r gin.Engine) {
	//
}