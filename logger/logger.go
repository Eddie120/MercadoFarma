package logger

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

type Logger struct {
	serviceName string
	writer      *Writer
}

type level string

const (
	success level = "success"
	warning level = "warning"
	danger  level = "danger"
	info    level = "info"
)

func NewLogger(service string) *Logger {
	if service == "" {
		panic("service name cannot be empty")
	}

	return &Logger{
		serviceName: service,
		writer:      &Writer{},
	}
}

func (logger *Logger) LogSuccessMessageWithContext(ctx context.Context, event string, object []Object) {
	logger.logMessage(ctx, success, event, object)
}

func (logger *Logger) LogWarningMessageWithContext(ctx context.Context, event string, object []Object) {
	logger.logMessage(ctx, warning, event, object)
}

func (logger *Logger) LogDangerMessageWithContext(ctx context.Context, event string, object []Object) {
	logger.logMessage(ctx, danger, event, object)
}

func (logger *Logger) LogInfoMessageWithContext(ctx context.Context, event string, object []Object) {
	logger.logMessage(ctx, info, event, object)
}

func (logger *Logger) logMessage(ctx context.Context, level level, event string, object []Object) {
	properties := logger.getProperties(ctx, object)

	message := map[string]interface{}{
		"id":         uuid.Must(uuid.NewRandom()).String(),
		"service":    logger.serviceName,
		"level":      level,
		"event":      event,
		"properties": properties,
		"time":       time.Now().Format(time.RFC3339Nano),
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("marshaling data failed %x", err.Error())
	}

	logger.writer.writeToExternal(ctx, data)
}

func (logger *Logger) getProperties(ctx context.Context, objects []Object) map[string]interface{} {
	properties := map[string]interface{}{
		"lambda_function_name":        os.Getenv("AWS_LAMBDA_FUNCTION_NAME"),
		"lambda_function_memory_size": os.Getenv("AWS_LAMBDA_FUNCTION_MEMORY_SIZE"),
		"lambda_function_version":     os.Getenv("AWS_LAMBDA_FUNCTION_VERSION"),
	}

	if len(objects) > 0 {
		for i := 0; i < len(objects); i++ {
			object := objects[i]
			properties[object.LogName()] = object.LogProperties()
		}
	}

	return properties
}
