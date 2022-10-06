// Package logging can be used for logging in json format
// TODO: format configuration
package logging

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

type LogLevel int

// Logging levels
const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
	WARN     = WARNING
	CRITICAL = FATAL
)

// Logging levels string representations
func (level LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}[level]
}

// ExtendedMessage - extended log message
type ExtendedMessage struct {
	Message string            // message
	Report  map[string]string // report map
	Auth    map[string]string // authentication map
	Error   Error             // error
	Trace   string            // traceback
}

// Message - log message
type Message struct {
	Message string // message
}

// Error - error structure
type Error struct {
	Code string // error code
	Name string // error name
}

// ExtendedLogEntry - extended log entry
type ExtendedLogEntry struct {
	Source    string          // logger name
	LogLevel  string          // log level
	Timestamp string          // time stamp
	Message   ExtendedMessage // extended message
}

// LogEntry - log entry
type LogEntry struct {
	Source    string  // logger name
	LogLevel  string  // log level
	Timestamp string  // time stamp
	Message   Message // message
}

// LogWriter - log writer
type LogWriter struct {
	name  string   // logger name
	level LogLevel // logger level
}

// Write function writes message to log
//
// Parameter(s):
//
//	bytes - log entry
//
// Returns:
//
//	The number of bytes written and any write error encountered.
func (writer LogWriter) Write(bytes []byte) (int, error) {
	data := string(bytes)
	if len(data) > 0 && data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}
	if len(data) == 0 {
		return 0, nil
	}
	for _, source := range regexp.MustCompile(`"[a-zA-Z0-9]*?":`).FindAllString(data, -1) {
		data = strings.Replace(data, source, strings.ToLower(source), -1)
	}
	data = strings.Replace(data, "\"loglevel\":", "\"log_level\":", -1)
	return fmt.Println(data)
}

// Init function initialise logger
//
// Parameter(s):
//
//	name  - logger name
//	level - logger level
func Init(name string, level LogLevel) {
	writer := new(LogWriter)
	writer.name = name
	writer.level = level
	log.SetOutput(writer)
	log.SetFlags(0)
	log.SetPrefix("")
}

// GetExtendedMessage function returns extended log message
//
// Parameter(s):
//
//	level     - message level
//	message   - message text
//	report    - report map
//	auth      - authentication map
//	errorCode - error code
//	errorName - error name
//	trace     - backtrace
//
// Returns:
//
//	extended log message
func GetExtendedMessage(level LogLevel, message string, report map[string]string, auth map[string]string,
	errorCode string, errorName string, trace string) string {
	if level < log.Default().Writer().(*LogWriter).level {
		return ""
	}
	logData, err := json.Marshal(
		ExtendedLogEntry{
			log.Default().Writer().(*LogWriter).name,
			level.String(),
			time.Now().UTC().Format("2006-01-02T15:04:05.999+0000"),
			ExtendedMessage{
				message,
				report,
				auth,
				Error{
					errorCode,
					errorName,
				},
				trace,
			},
		},
	)
	if err == nil {
		return string(logData)
	} else {
		return "getMessage: JSON error, incorrect logging"
	}
}

// GetMessage function returns log message
//
// Parameter(s):
//
//	level   - message level
//	message - message text
//
// Returns:
//
//	log message string
func GetMessage(level LogLevel, message string) string {
	if level < log.Default().Writer().(*LogWriter).level {
		return ""
	}
	logData, err := json.Marshal(
		LogEntry{
			log.Default().Writer().(*LogWriter).name,
			level.String(),
			time.Now().UTC().Format("2006-01-02T15:04:05.999+0000"),
			Message{message},
		})
	if err == nil {
		return string(logData)
	} else {
		return "getMessage: JSON error, incorrect logging"
	}
}
