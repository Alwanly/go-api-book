package authentication

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

	// GetClaims returns the JWT claims.
	//
	// Parameters:
	//   - claims: JWT claims
	//
	// Returns:
	//   - error: error
	ValidateClaims(claims JWTClaims) error
}

type JWTClaims struct {
	// User ID
	UserID string `json:"userId"`
}

type JWTConfig struct {
	// JWT secret
	Secret string

	// JWT expiration time
	ExpirationTime int

	// JWT refresh time
	RefreshTime int

	// JWT issuer
	Issuer string

	// JWT audience
	Audience string
}

type jwt struct {
	secret         string
	expirationTime int
	refreshTime    int
	issuer         string
	audience       string
}

func NewJWTService(opts *JWTConfig) IJwtService {
	return &jwt{
		secret:         opts.Secret,
		expirationTime: opts.ExpirationTime,
		refreshTime:    opts.RefreshTime,
		issuer:         opts.Issuer,
		audience:       opts.Audience,
	}
}

func (j *jwt) GenerateToken(claims JWTClaims) (string, error) {
	return "", nil
}

func (j *jwt) ParseToken(token string) (*JWTClaims, error) {
	return nil, nil
}

func (j *jwt) RefreshToken(token string) (string, error) {
	return "", nil
}

func (j *jwt) ValidateToken(token string) error {
	return nil
}

func (j *jwt) ValidateClaims(claims JWTClaims) error {
	return nil
}
