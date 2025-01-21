package authentication

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type IJwtService interface {
	// GenerateToken generates a new JWT token.
	//
	// Parameters:
	//   - claims: JWT claims
	//
	// Returns:
	//   - string: JWT token
	GenerateToken(claims JWTClaims) (string, error)

	// ParseToken parses a JWT token.
	//
	// Parameters:
	//   - token: JWT token
	//
	// Returns:
	//   - *JWTClaims: JWT claims
	//   - error: error
	ParseToken(token string) (*JWTClaims, error)

	// RefreshToken refreshes a JWT token.
	//
	// Parameters:
	//   - token: JWT token
	//
	// Returns:
	//   - string: JWT token
	//   - error: error
	RefreshToken(token string) (string, error)

	// ValidateToken validates a JWT token.
	//
	// Parameters:
	//   - token: JWT token
	//
	// Returns:
	//   - error: error
	ValidateToken(token string) error
}

type JWTClaims map[string]interface{}
type JWTConfig struct {
	// JWT secret
	PrivateKey string
	PublicKey  string

	// JWT expiration time
	ExpirationTime int

	// JWT refresh time
	RefreshTime int

	// JWT issuer
	Issuer string

	// JWT audience
	Audience string
}

type jwtAuth struct {
	privateKey     string
	publicKey      string
	issuer         string
	audience       string
	refreshTime    int
	expirationTime int
}

func NewJWTService(opts *JWTConfig) IJwtService {
	return &jwtAuth{
		privateKey:     opts.PrivateKey,
		publicKey:      opts.PublicKey,
		expirationTime: opts.ExpirationTime,
		refreshTime:    opts.RefreshTime,
		issuer:         opts.Issuer,
		audience:       opts.Audience,
	}
}

func (j *jwtAuth) GenerateToken(dataClaims JWTClaims) (string, error) {
	var tokenString string
	var privateKey *rsa.PrivateKey

	block, _ := pem.Decode([]byte(j.privateKey))
	if block == nil {
		return "", errors.New("invalid private rsa key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Create the token
	token := jwt.New(jwt.SigningMethodRS256)

	now := time.Now()
	exp := now.Add(time.Duration(j.expirationTime) * time.Minute).Unix()
	// Set claims

	// Set claims
	claimsMap := jwt.MapClaims{
		"iss": j.issuer,
		"aud": j.audience,
		"exp": exp,
	}

	for key, value := range dataClaims {
		claimsMap[key] = value
	}

	token.Claims = claimsMap

	// Sign the token with the private key
	tokenString, err = token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtAuth) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwt.ParseRSAPublicKeyFromPEM([]byte(j.publicKey))
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		jwtClaims := JWTClaims{}
		for key, value := range claims {
			jwtClaims[key] = value
		}
		return &jwtClaims, nil
	}

	return nil, errors.New("invalid token")
}

func (j *jwtAuth) RefreshToken(tokenString string) (string, error) {
	var privateKey *rsa.PrivateKey
	block, _ := pem.Decode([]byte(j.privateKey))
	if block == nil {
		return "", errors.New("invalid private rsa key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	// Create a new claims map and copy existing claims
	newClaims := jwt.MapClaims{
		"iss": j.issuer,
		"aud": j.audience,
		"exp": time.Now().Add(time.Duration(j.refreshTime) * time.Minute).Unix(),
	}

	for key, value := range *claims {
		if key != "exp" { // Avoid copying the old expiration time
			newClaims[key] = value
		}
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodRS256, newClaims)

	newTokenString, err := newToken.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return newTokenString, nil
}

func (j *jwtAuth) ValidateToken(tokenString string) error {
	_, err := j.ParseToken(tokenString)
	return err
}
