package authentication

import (
	"encoding/base64"
	"go-codebase/infrastructure/config"
	"go-codebase/infrastructure/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			ctx.Header("WWW-Authenticate", "Basic ley")
			utils.WriteAbortResponse(ctx, utils.WrapperReponse(http.StatusUnauthorized, "Missing authorization", nil))
			return
		}

		if !strings.HasPrefix(authorizationHeader, "Basic") {
			ctx.Header("WWW-Authenticate", "Basic ley")
			utils.WriteAbortResponse(ctx, utils.WrapperReponse(http.StatusUnauthorized, "Missing authorization", nil))
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Basic ", "", -1)

		tokenBody, err := base64.StdEncoding.DecodeString(tokenString)

		if err != nil {
			ctx.Header("WWW-Authenticate", "Basic ley")
			utils.WriteAbortResponse(ctx, utils.WrapperReponse(http.StatusUnauthorized, "Missing authorization", err.Error()))
			return
		}

		parts := strings.Split(string(tokenBody), ":")

		if !SafeCompareString(parts[0], config.GlobalConfig.Authentication.BasicUsername) || !SafeCompareString(parts[1], config.GlobalConfig.Authentication.BasicPassword) {
			ctx.Header("WWW-Authenticate", "Basic ley")
			utils.WriteAbortResponse(ctx, utils.WrapperReponse(http.StatusUnauthorized, "Missing authorization", nil))
			return
		}

		ctx.Set("BASIC_AUTHENTICATED", true)
		ctx.Next()

	}
}
