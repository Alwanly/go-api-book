package middleware

import (
	"net/http"
	"strings"

	"github.com/Alwanly/go-codebase/pkg/authentication"
	"github.com/gofiber/fiber/v2"
)

type IAuthMiddleware interface {
	// Jwt token
	JwtAuth() fiber.Handler

	// Basic Auth
	BasicAuth() fiber.Handler
}

type AuthMiddleware struct {
	Jwt   authentication.IJwtService
	Basic authentication.IBasicAuthService
}

// mockery:ignore
type AuthConfig func(*AuthOpts)

type AuthUserData struct {
	UserID string `json:"userId"`
}

type AuthOpts struct {
	*authentication.JWTConfig
	*authentication.BasicAuthTConfig
}

const LocalTokenKey = "user"

func SetJwtAuth(jwtConfig *authentication.JWTConfig) AuthConfig {
	return func(o *AuthOpts) {
		o.JWTConfig = jwtConfig
	}
}

func SetBasicAuth(basicAuthConfig *authentication.BasicAuthTConfig) AuthConfig {
	return func(o *AuthOpts) {
		o.BasicAuthTConfig = basicAuthConfig
	}
}

func NewAuthMiddleware(opts ...AuthConfig) *AuthMiddleware {
	var o AuthOpts
	for _, opt := range opts {
		opt(&o)
	}

	jwtAuth := authentication.NewJWTService(o.JWTConfig)

	basicAuth := authentication.NewBasicAuthService(o.BasicAuthTConfig)
	return &AuthMiddleware{
		Jwt:   jwtAuth,
		Basic: basicAuth,
	}
}

func (a *AuthMiddleware) JwtAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// get token from header
		token := ctx.Get(fiber.HeaderAuthorization)
		if !strings.Contains(token, "Bearer") {
			return responseUnauthorized(ctx, "Bearer", "Invalid token")
		}

		// validate token
		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			return responseUnauthorized(ctx, "Bearer", "Invalid token")
		}

		// parse token
		auth, err := a.Jwt.ParseToken(token)
		if err != nil {
			return responseUnauthorized(ctx, "Bearer", "Invalid token")
		}

		// set claims to context
		ctx.Locals(LocalTokenKey, decodeAuthToken(*auth))

		return ctx.Next()
	}
}

func (a *AuthMiddleware) BasicAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// get auth from header
		auth := ctx.Get(fiber.HeaderAuthorization)
		if !strings.Contains(auth, "Basic") {
			return responseUnauthorized(ctx, "Basic", "Invalid auth")
		}

		// decode auth
		username, password := a.Basic.DecodeFromHeader(auth)
		if !a.Basic.Validate(username, password) {
			return responseUnauthorized(ctx, "Basic", "Invalid auth")
		}
		return ctx.Next()
	}
}

func decodeAuthToken(dataClaims authentication.JWTClaims) *AuthUserData {
	return &AuthUserData{
		UserID: dataClaims["userId"].(string),
	}
}

func responseUnauthorized(c *fiber.Ctx, _ string, message ...string) error {
	c.Set("WWW-Authenticate", "Basic realm=Restricted")
	response := fiber.Map{
		"message": message[0],
	}
	if len(message) > 1 {
		response["statusCode"] = message[1]
	}
	return c.Status(http.StatusUnauthorized).JSON(response)
}
