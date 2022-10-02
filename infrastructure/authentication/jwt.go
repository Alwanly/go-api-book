package authentication

import (
	"books-api/infrastructure/config"
	"books-api/infrastructure/utils"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type IJwtAuth interface {
	Sign(userId string, username string) (tokenString string, err error)
	JwtAuthMidldleware() gin.HandlerFunc
}

type JwtAuthImpl struct {
	Secret []byte
}
type JwtPayload struct {
	jwt.StandardClaims
	Username string `json:"username,omitempty"`
}

func ConstructJwtAuth() (*JwtAuthImpl, error) {
	secretKey := []byte(config.GlobalConfig.Authentication.JwtSecret)

	if secretKey == nil {
		return nil, fmt.Errorf("JWT Secret null ")
	}
	return &JwtAuthImpl{Secret: secretKey}, nil
}

func (u *JwtAuthImpl) Sign(userId string, username string) (tokenString string, err error) {

	var jwtPayload = make(jwt.MapClaims)
	jwtPayload["aud"] = "DEVORIA"
	jwtPayload["exp"] = time.Now().Add(time.Hour * 24).Unix()
	jwtPayload["iat"] = time.Now().Unix()
	jwtPayload["iss"] = "DEVORIA"
	jwtPayload["jti"] = userId
	jwtPayload["username"] = username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	return token.SignedString(u.Secret)
}

func (u *JwtAuthImpl) JwtAuthMidldleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			utils.WriteAbortResponse(ctx, utils.WrapperReponse(http.StatusUnauthorized, "Missing header", nil))
			return
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			utils.WriteAbortResponse(ctx, utils.WrapperReponse(http.StatusUnauthorized, "Authorization is not header", nil))
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		jwtPayload := &JwtPayload{}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return u.Secret, nil
		})

		if err != nil {
			utils.WriteAbortResponse(ctx, utils.WrapperReponse(http.StatusUnauthorized, "Missing header", err))
			return
		}

		mapClaims := token.Claims.(jwt.MapClaims)

		jwtPayload.ExpiresAt = int64(mapClaims["exp"].(float64))
		jwtPayload.IssuedAt = int64(mapClaims["iat"].(float64))
		jwtPayload.Username = mapClaims["username"].(string)
		ctx.Set("JWT_AUTHENTICATED", true)
		ctx.Set("JWT_PAYLOAD", &jwtPayload)
		ctx.Next()
	}
}
