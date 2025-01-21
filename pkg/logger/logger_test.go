package logger_test

import (
	"testing"

	"github.com/Alwanly/go-codebase/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	log := logger.NewLogger("test_service", "debug")
	assert.NotNil(t, log)
}

func TestWithID(t *testing.T) {
	log := logger.NewLogger("test_service", "debug")
	logWithID := logger.WithID(log, "test_context", "test_scope")
	assert.NotNil(t, logWithID)
}

func TestWithPrettyPrint(t *testing.T) {
	log := logger.NewLogger("test_service", "debug", logger.WithPrettyPrint())
	assert.NotNil(t, log)
}
