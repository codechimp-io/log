// Package log provides a global logger with zerolog.
package log

import (
	stdlog "log"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger is the global logger.
var Logger zerolog.Logger

func init() {
	// log with nanosecond precision time
	zerolog.TimeFieldFormat = ""
	zerolog.DurationFieldInteger = true
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"
	zerolog.CallerFieldName = "producer"
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano}).With().Timestamp().Logger()	

}

// NewStd creates new standart logger with zerolog as output
func NewStd() *stdlog.Logger {
	return stdlog.New(Logger, "", 0)
}

// WithCaller adds caller name to the logs
func WithCaller(name string) {
	Logger = Logger.With().Str(zerolog.CallerFieldName, name).Logger()
}

// Infof
func Infof(format string, v ...interface{}) {
	Logger.Info().Msgf(format, v...)
}

// Debugf
func Debugf(format string, v ...interface{}) {
	Logger.Debug().Msgf(format, v...)
}

// Errorf
func Errorf(format string, v ...interface{}) {
	Logger.Error().Msgf(format, v...)
}

// Fatalf
func Fatalf(format string, v ...interface{}) {
	Logger.Fatal().Msgf(format, v...)
}

// Warnf
func Warnf(format string, v ...interface{}) {
	Logger.Warn().Msgf(format, v...)
}

// Info
func Info(entry string) {
	Logger.Info().Msg(entry)
}

// Debug
func Debug(entry string) {
	Logger.Debug().Msg(entry)
}

// Error
func Error(entry string) {
	Logger.Error().Msg(entry)
}

// Fatal
func Fatal(entry string) {
	Logger.Fatal().Msg(entry)
}

// Warn
func Warn(entry string) {
	Logger.Warn().Msgf(entry)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	Logger.Print(v...)
}
