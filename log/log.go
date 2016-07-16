package log

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type Logger struct {
	Log *log.Logger
}

func NewLogger() *Logger {
	log := log.New()
	return &Logger{Log: log}
}

func (l *Logger) SetLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		l.Fatal(`Not a valid level: "%s"`, level)
	}
	l.Log.Level = lvl
}

// Debug logs a message with severity DEBUG.
func (l *Logger) Debug(format string, v ...interface{}) {
	l.Log.Debug(fmt.Sprintf(format, v...))
}

// Info logs a message with severity INFO.
func (l *Logger) Info(format string, v ...interface{}) {
	l.Log.Info(fmt.Sprintf(format, v...))
}

// Warning logs a message with severity WARNING.
func (l *Logger) Warning(format string, v ...interface{}) {
	l.Log.Warning(fmt.Sprintf(format, v...))
}

// Error logs a message with severity ERROR.
func (l *Logger) Error(format string, v ...interface{}) {
	l.Log.Error(fmt.Sprintf(format, v...))
}

// Fatal logs a message with severity ERROR followed by a call to os.Exit().
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.Log.Fatal(fmt.Sprintf(format, v...))
}