package logging

type Level string

const (
	Error   Level = "ERROR"
	Warning Level = "WARNING"
	Notice  Level = "NOTICE"
	Info    Level = "INFO"
	Debug   Level = "DEBUG"
)

type Logger interface {
	Log(level Level, data string)
}
