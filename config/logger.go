package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   log.Logger
	info    log.Logger
	warning log.Logger
	err     log.Logger
	writer  io.Writer
}

func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)
	return &Logger{
		debug:   *log.New(writer, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		info:    *log.New(writer, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warning: *log.New(writer, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		err:     *log.New(writer, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		writer:  writer,
	}
}

// Create Non-Formatted loggers
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Formatted loggers
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
