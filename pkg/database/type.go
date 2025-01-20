package database

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type (
	ContextTransaction string

	OptFunc func(opts *DBServiceOpts) error
)

const (
	ContextName = "Components.Database"
	PingTimeout = 10 * time.Second
)

// DBServiceOpts represents the options for configuring the database service.
type DBServiceOpts struct {
	// Debug enables debug mode.
	Debug bool
	// Logger is the logger.
	Logger *zap.Logger

	// PostgresSQL database connection strings (DSNs)
	postgresUri *string
	// Maximum number of open connections to the database. Default is 10.
	PostgresMaxOpenConnections int
	// Maximum number of idle connections to the database. Default is 5.
	PostgresMaxIdleConnections int

	// Application Name (for tracing)
	ApplicationName *string
}

// DBService represents the database service.
type DBService struct {
	Gorm *gorm.DB
}

// IDBService represents the interface for the database service.
type IDBService interface {
	// ---- Postgres

	// PingPostgres pings the Postgres database to check if it's available.
	//
	// Returns:
	//   - bool: true if the database is available, false otherwise.
	Ping() bool

	// BeginTransaction starts a new transaction and returns a new context with the transaction attached.
	//
	// Parameters:
	//   - c: context
	//
	// Returns:
	//   - context.Context: new context with the transaction attached
	//   - *gorm.DB: transaction
	BeginTransaction(c context.Context) (context.Context, *gorm.DB)

	// GetTransaction returns the transaction attached to the context.
	//
	// Parameters:
	//   - c: context
	//
	// Returns:
	//   - *gorm.DB: transaction
	GetTransaction(c context.Context) *gorm.DB

	// RollbackTransaction rolls back the transaction attached to the context.
	//
	// Parameters:
	//   - c: context
	//
	// Returns:
	//   - *gorm.DB: transaction
	RollbackTransaction(c context.Context) *gorm.DB

	// CommitTransaction commits the transaction attached to the context.
	//
	// Parameters:
	//   - c: context
	//
	// Returns:
	//   - *gorm.DB: transaction
	CommitTransaction(c context.Context) *gorm.DB

	// SetUpdateLockType sets the lock type to UPDATE.
	//
	// Parameters:
	//   - c: context
	//
	// Returns:
	//   - context.Context: new context with the lock type attached
	SetUpdateLockType(c context.Context) context.Context

	// SetShareLockType sets the lock type to SHARE.
	//
	// Parameters:
	//   - c: context
	//
	// Returns:
	//   - context.Context: new context with the lock type attached
	SetShareLockType(c context.Context) context.Context

	// GetLockType returns the lock type attached to the context.
	//
	// Parameters:
	//   - c: context
	//
	// Returns:
	//   - *string: lock type
	GetLockType(c context.Context) *string

	// Defer defers the transaction attached to the context.
	//
	// Parameters:
	//   - c: context
	Defer(c context.Context)

	// Close closes the database connection.
	Close() error
}

// set config for uri
func SetPostgresUri(uri string) OptFunc {
	return func(opts *DBServiceOpts) error {
		opts.postgresUri = &uri
		return nil
	}
}
