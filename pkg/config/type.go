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
	JwtPublicKey      string `mapstructure:"JWT_PUBLIC_KEY"`
	JwtPrivateKey     string `mapstructure:"JWT_PRIVATE_KEY"`
	JwtIssuer         string `mapstructure:"JWT_ISSUER"`
	JwtAudience       string `mapstructure:"JWT_AUDIENCE"`

	// Database
	PostgresUri                string `mapstructure:"POSTGRES_URI"`
	PostgresMaxOpenConnections int    `mapstructure:"POSTGRES_MAX_OPEN_CONNECTIONS"`
	PostgresMaxIdleConnections int    `mapstructure:"POSTGRES_MAX_IDLE_CONNECTIONS"`
}
