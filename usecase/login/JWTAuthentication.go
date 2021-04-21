package login

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mongo-admin-backend/api/presenter"
	"mongo-admin-backend/config"
	"net/http"
	"time"
)

//jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getSecretKey() string {
	secret := config.JWT_SECRET_KEY
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}

// VerifyToken func will used to Verify the JWT Token while using APIS
func VerifyToken(tokenString string) (email string, err error) {
	claims := &presenter.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_SECRET_KEY, nil
	})

	if token != nil {
		return claims.Email, nil
	}
	return "", err
}

// GetEmail func will used to Verify the JWT Token while using APIS.
func GetEmail(ctx *gin.Context) (email string, err error) {
	var BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]

	claims := &presenter.Claims{}

	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_SECRET_KEY, nil
	})

	claims, _ = token.Claims.(*presenter.Claims)
	// @todo check why token.valid is wrong and fix this
	return claims.Email, nil
	//return "", errors.New("Something happened")
}
