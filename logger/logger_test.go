package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestNewLogger(t *testing.T) {
	logger := NewLogger("dummy")
	assert.IsType(t, &Logger{}, logger)
}
