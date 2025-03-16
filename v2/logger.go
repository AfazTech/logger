package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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
	outputOption int = CONSOLE_ONLY
	logFileName  string
	logfile      *os.File
	output       io.Writer
)

func initLogFile() error {
	if logFileName == "" {
		exePath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("can't determine executable path: %w", err)
		}
		logFileName = filepath.Base(exePath) + ".log"
	}
	if logfile != nil {
		return nil
	}
	var err error
	logfile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("can't write log file: %w", err)
	}
	return nil
}

func SetOutput(option int) error {
	outputOption = option
	switch outputOption {
	case CONSOLE_ONLY:
		output = os.Stdout
	case FILE_ONLY:
		if err := initLogFile(); err != nil {
			return err
		}
		output = logfile
	case CONSOLE_AND_FILE:
		if err := initLogFile(); err != nil {
			return err
		}
		output = io.MultiWriter(os.Stdout, logfile)
	default:
		return fmt.Errorf("invalid output option")
	}
	log.SetOutput(output)
	return nil
}

func SetLogFile(filePath string) error {
	if logfile != nil {
		logfile.Close()
	}
	logFileName = filePath
	return nil
}

func CloseLogFile() {
	if logfile != nil {
		logfile.Close()
	}
}

func logging(level int, messages ...interface{}) {
	message := fmt.Sprint(messages...)
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

func Debug(messages ...interface{}) {
	logging(DEBUG, messages...)
}

func Info(messages ...interface{}) {
	logging(INFO, messages...)
}

func Warn(messages ...interface{}) {
	logging(WARN, messages...)
}

func Error(messages ...interface{}) {
	logging(ERROR, messages...)
}

func Fatal(messages ...interface{}) {
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
