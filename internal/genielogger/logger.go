package genielogger

import (
	"log"
)

type Logger interface {
	Info(msg string)
	Error(msg string, err error)
	Fatal(msg string, err error)
}

type GenieLogger struct{}

func NewGenieLogger() Logger {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	return &GenieLogger{}
}

func (l *GenieLogger) Info(msg string) {
	log.Println("[INFO]", msg)
}

func (l *GenieLogger) Error(msg string, err error) {
	log.Printf("[ERROR] %s: %v\n", msg, err)
}

func (l *GenieLogger) Fatal(msg string, err error) {
	log.Fatalf("[FATAL] %s: %v\n", msg, err)
}
