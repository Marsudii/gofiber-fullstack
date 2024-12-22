package config

import (
	"fmt"
	"log"
	"os"
)

// Logs struct untuk logging
type Logs struct {
	logger *log.Logger
}

// NewLogs membuat instance logger baru
func NewLogs() *Logs {
	// Pastikan direktori logs ada
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	// Buka file log atau buat jika belum ada
	file, err := os.OpenFile("logs/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// Buat logger instance
	logger := log.New(file, "", log.LstdFlags)

	return &Logs{logger: logger}
}

// Error mencatat log dengan level error
func (l *Logs) Error(format string, args ...interface{}) {
	l.logger.SetPrefix("ERROR: ")
	l.logger.Println(fmt.Sprintf(format, args...))
}

// Info mencatat log dengan level info
func (l *Logs) Info(message string) {
	l.logger.SetPrefix("INFO: ")
	l.logger.Println(message)
}
