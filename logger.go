package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	_ = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL

	_ = iota
	CONSOLE_ONLY
	FILE_ONLY
	CONSOLE_AND_FILE
)

var (
	outputOption int    = CONSOLE_ONLY
	logFileName  string = "app.log"
	logfile      *os.File
	output       io.Writer
	once         sync.Once
)

func initLogFile() {
	once.Do(func() {
		var err error
		logfile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("can't write log file:", err.Error())
		}
	})
}

func SetOutput(option int) {
	outputOption = option
	switch outputOption {
	case CONSOLE_ONLY:
		output = os.Stdout
	case FILE_ONLY:
		initLogFile()
		output = logfile
	case CONSOLE_AND_FILE:
		initLogFile()
		output = io.MultiWriter(os.Stdout, logfile)
	default:
		logging(FATAL, "invalid output option")
		return
	}
	log.SetOutput(output)
}
func SetLogFile(filePath string) {

	var err error
	logFileName = filePath
	logfile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("can't write log file:", err.Error())
		return
	}
	SetOutput(outputOption)
}

func logging(level int, messages ...string) {
	message := strings.Join(messages, " ")
	switch level {
	case DEBUG:
		message = "DEBUG: " + message
	case INFO:
		message = "INFO: " + message
	case WARN:
		message = "WARN: " + message
	case ERROR:
		message = "ERROR: " + message
	case FATAL:
		message = "FATAL: " + message
		log.Fatalln(message)
	}

	log.Println(message)
}

func Debug(messages ...string) {
	logging(DEBUG, messages...)
}

func Info(messages ...string) {
	logging(INFO, messages...)
}

func Warn(messages ...string) {
	logging(WARN, messages...)
}

func Error(messages ...string) {
	logging(ERROR, messages...)
}

func Fatal(messages ...string) {
	logging(FATAL, messages...)
}

func loggingf(level int, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	logging(level, message)
}

func Debugf(format string, args ...interface{}) {
	loggingf(DEBUG, format, args...)
}

func Infof(format string, args ...interface{}) {
	loggingf(INFO, format, args...)
}

func Warnf(format string, args ...interface{}) {
	loggingf(WARN, format, args...)
}

func Errorf(format string, args ...interface{}) {
	loggingf(ERROR, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	loggingf(FATAL, format, args...)
}
