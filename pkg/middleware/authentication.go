package middleware

import (
	"go-codebase/pkg/authentication"
	"go-codebase/pkg/utils"
	"strings"

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

type authOpts struct {
	secret         string
	expirationTime int
	refreshTime    int
	issuer         string
	audience       string

	// Basic Auth
	username string
	password string
}

type AuthConfig func(*authOpts)

type AuthToken authentication.JWTClaims

func SetJwtAuth(jwtConfig authentication.JWTConfig) AuthConfig {
	return func(o *authOpts) {
		o.secret = jwtConfig.Secret
		o.expirationTime = jwtConfig.ExpirationTime
		o.refreshTime = jwtConfig.RefreshTime
		o.issuer = jwtConfig.Issuer
		o.audience = jwtConfig.Audience
	}
}

func SetBasicAuth(basicAuthConfig authentication.BasicAuthTConfig) AuthConfig {
	return func(o *authOpts) {
		o.username = basicAuthConfig.Username
		o.password = basicAuthConfig.Password
	}
}

func NewAuthMiddleware(opts ...AuthConfig) *AuthMiddleware {
	var o authOpts
	for _, opt := range opts {
		opt(&o)
	}

	jwtOpts := &authentication.JWTConfig{
		Secret:         o.secret,
		ExpirationTime: o.expirationTime,
		RefreshTime:    o.refreshTime,
		Issuer:         o.issuer,
		Audience:       o.audience,
	}

	jwtAuth := authentication.NewJWTService(jwtOpts)

	// basic auth
	basicAuthOpts := &authentication.BasicAuthTConfig{
		Username: o.username,
		Password: o.password,
	}
	basicAuth := authentication.NewBasicAuthService(*basicAuthOpts)
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
			return utils.ResponseUnauthorized(ctx, "Bearer", "Invalid token")
		}

		// validate token
		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			return utils.ResponseUnauthorized(ctx, "Bearer", "Invalid token")
		}

		// parse token
		auth, err := a.Jwt.ParseToken(token)
		if err != nil {
			return utils.ResponseUnauthorized(ctx, "Bearer", "Invalid token")
		}

		// set claims to context
		ctx.Locals("auth", auth)

		return ctx.Next()
	}
}

func (a *AuthMiddleware) BasicAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// get auth from header
		auth := ctx.Get(fiber.HeaderAuthorization)
		if !strings.Contains(auth, "Basic") {
			return utils.ResponseUnauthorized(ctx, "Basic", "Invalid auth")
		}

		// decode auth
		username, password := a.Basic.DecodeFromHeader(auth)
		if !a.Basic.Validate(username, password) {
			return utils.ResponseUnauthorized(ctx, "Basic", "Invalid auth")
		}
		return ctx.Next()
	}
}
