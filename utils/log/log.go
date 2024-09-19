package log

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var loggerEntry *logrus.Entry

func NewLogger() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	customFormatter.ForceColors = true
	logger = &logrus.Logger{
		Out:          os.Stderr,
		ExitFunc:     os.Exit,
		Formatter:    customFormatter,
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.DebugLevel,
		ReportCaller: false,
	}
	loggerEntry = logrus.NewEntry(logger)
}

func WithField(key string, value interface{}) {
	loggerEntry = loggerEntry.WithField(key, value)
}

func GetTraceID() string {
	return fmt.Sprintf("%v", loggerEntry.Data["trace_id"])
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	loggerEntry.Debug(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	loggerEntry.WithField("file", fileInfo(2)).Info(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	loggerEntry.WithField("file", fileInfo(2)).Error(args...)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	// else {
	//	slash := strings.LastIndex(file, "/")
	//	if slash >= 0 {
	//		file = file[slash+1:]
	//	}
	// }
	return fmt.Sprintf("%s:%d", file, line)
}

// DebugWithFields logs a message with fields at level Debug on the standard logger.
func DebugWithFields(l interface{}, f logrus.Fields) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields(f))
		// entry.Data["file"] = fileInfo(2)
		entry.Debug(l)
	}
}

// InfoWithFields logs a message with fields at level Debug on the standard logger.
func InfoWithFields(l interface{}, f logrus.Fields) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields(f))
		// entry.Data["file"] = fileInfo(2)
		entry.Info(l)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Warn(args...)
	}
}

// WarnWithFields logs a message with fields at level Debug on the standard logger.
func WarnWithFields(l interface{}, f logrus.Fields) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Warn(l)
	}
}

// ErrorWithFields logs a message with fields at level Debug on the standard logger.
func ErrorWithFields(l interface{}, f logrus.Fields) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(args...)
	}
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	if logger.Level >= logrus.TraceLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Trace(args...)
	}
}

// FatalWithFields logs a message with fields at level Debug on the standard logger.
func FatalWithFields(l interface{}, f logrus.Fields) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(l)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)
	}
}

// PanicWithFields logs a message with fields at level Debug on the standard logger.
func PanicWithFields(l interface{}, f logrus.Fields) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Panic(l)
	}
}
