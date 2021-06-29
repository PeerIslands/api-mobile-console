package handler

import (
	"errors"
	"log"
	"mongo-admin-backend/api/middleware"
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/entity"
	"mongo-admin-backend/usecase/login"
	"mongo-admin-backend/usecase/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(service user.UseCase, ctx *gin.Context) {
	var input struct {
		Email       string             `json:"email"`
		Password    string             `json:"password"`
		Name        string             `json:"name"`
		AtlasParams entity.AtlasParams `json:"atlas_params,omitempty"`
	}
	err := ctx.ShouldBind(&input)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		if input.AtlasParams.PrivateKey != "" {
			user, err := service.CreateUser(input.Email, input.Password, input.Name, input.AtlasParams.PublicKey, input.AtlasParams.PrivateKey)
			if err != nil {
				status := http.StatusInternalServerError
				log.Println(err.Error())
				if err.Error() == "user already exists" {
					status = http.StatusConflict
				}
				ctx.JSON(status, gin.H{
					"error": err.Error(),
				})
			} else {
				toJ := &presenter.User{
					ID:    user.ID,
					Email: input.Email,
					Name:  input.Name,
				}
				ctx.JSON(http.StatusCreated, toJ)
			}
		} else {
			user, err := service.SignUp(input.Email, input.Password, input.Name)
			if err != nil {
				status := http.StatusInternalServerError
				log.Println(err.Error())
				if err.Error() == "user already exists" {
					status = http.StatusConflict
				}
				ctx.JSON(status, gin.H{
					"error": err.Error(),
				})
			} else {
				toJ := &presenter.User{
					ID:    user.ID,
					Email: input.Email,
					Name:  input.Name,
				}
				ctx.JSON(http.StatusCreated, toJ)
			}
		}
	}
}

func getUser(service user.UseCase, ctx *gin.Context) {
	email, err := login.GetEmail(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else if email == "" {
		ctx.JSON(http.StatusNotFound, errors.New("user not found"))
	} else {
		user, err := service.GetUser(email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			var isUserRegistered bool
			if user.AtlasParams.PrivateKey != "" {
				isUserRegistered = true
			} else {
				isUserRegistered = false
			}
			toJ := &presenter.User{
				ID:      user.ID,
				Email:   user.Email,
				Name:    user.Name,
				IsAdmin: isUserRegistered,
			}

			ctx.JSON(http.StatusOK, toJ)
		}
	}
}

func putUserCredentials(service user.UseCase, ctx *gin.Context) {
	email, err := login.GetEmail(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else if email == "" {
		ctx.JSON(http.StatusNotFound, errors.New("user not found"))
	} else {
		var input struct {
			PublicKey  string `json:"public_key"`
			PrivateKey string `json:"private_key"`
		}
		err := ctx.ShouldBind(&input)

		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		user, updateErr := service.PutAtlasCredentials(email, input.PrivateKey, input.PublicKey)
		if updateErr != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		} else {

			toJ := &presenter.User{
				ID:    user.ID,
				Email: user.Email,
				Name:  user.Name,
			}

			ctx.JSON(http.StatusOK, toJ)
		}
	}
}

func deleteUser(service user.UseCase, ctx *gin.Context) {
	email, err := login.GetEmail(ctx)
	err = service.DeleteUser(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, nil)
}

// MakeUserHandlers make url handlers.
func MakeUserHandlers(r *gin.Engine, service user.UseCase) {
	v1 := r.Group("/v1/user").Use(middleware.AuthorizeJWT())
	{
		v1.GET("/", func(ctx *gin.Context) {
			getUser(service, ctx)
		})
		v1.DELETE("/{email}", func(ctx *gin.Context) {
			deleteUser(service, ctx)
		})
		v1.PUT("/", func(ctx *gin.Context) {
			putUserCredentials(service, ctx)
		})
	}
}

// MakeUserNoAuthHandlers make url handlers.
func MakeUserNoAuthHandlers(r *gin.Engine, service user.UseCase) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, nil)
	})

	v1 := r.Group("/v1/user")
	{
		v1.POST("/", func(ctx *gin.Context) {
			createUser(service, ctx)
		})
	}
}
