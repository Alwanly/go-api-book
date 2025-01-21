package config

type GlobalConfig struct {
	// global config
	Environment string `mapstructure:"ENV"`
	Debug       bool   `mapstructure:"DEBUG"`
	Port        int    `mapstructure:"PORT"`
	PortGrpc    int    `mapstructure:"PORT_GRPC"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`

	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceVersion string `mapstructure:"SERVICE_VERSION"`

	// Authentication
	BasicAuthUsername string `mapstructure:"BASIC_AUTH_USERNAME"`
	BasicAuthPassword string `mapstructure:"BASIC_AUTH_PASSWORD"`
	JwtIssuer         string `mapstructure:"JWT_ISSUER"`
	JwtAudience       string `mapstructure:"JWT_AUDIENCE"`
	JwtExpirationTime int    `mapstructure:"JWT_EXPIRATION"`
	JwtRefreshTime    int    `mapstructure:"JWT_REFRESH_EXPIRATION"`

	// RSA keys
	PublicKey  string `mapstructure:"PUBLIC_KEY"`
	PrivateKey string `mapstructure:"PRIVATE_KEY"`

	// Database
	PostgresURI                string `mapstructure:"POSTGRES_URI"`
	PostgresMaxOpenConnections int    `mapstructure:"POSTGRES_MAX_OPEN_CONNECTIONS"`
	PostgresMaxIdleConnections int    `mapstructure:"POSTGRES_MAX_IDLE_CONNECTIONS"`

	// Redis
	RedisURI string `mapstructure:"REDIS_URI"`
}
