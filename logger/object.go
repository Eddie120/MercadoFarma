package logger

type Object interface {
	LogName() string
	LogProperties() map[string]interface{}
}
