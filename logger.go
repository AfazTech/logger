package logger

import (
	"io"
	"log"
	"os"
	"strings"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL

	CONSOLE_ONLY = iota
	FILE_ONLY
	CONSOLE_AND_FILE
)

func Init(logFileName string, outputOption int) {
	var output io.Writer

	logfile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		Log(ERROR, "cant write log file: ", err.Error())
		return
	}

	switch outputOption {
	case CONSOLE_ONLY:
		output = os.Stdout
	case FILE_ONLY:
		output = logfile
	case CONSOLE_AND_FILE:
		output = io.MultiWriter(os.Stdout, logfile)
	default:
		Log(FATAL, "invalid output option")
		return
	}

	log.SetOutput(output)
}

func Log(level int, messages ...string) {
	message := strings.Join(messages, " ")
	switch level {
	case DEBUG:
		log.Println("DEBUG:", message)
	case INFO:
		log.Println("INFO:", message)
	case WARN:
		log.Println("WARN:", message)
	case ERROR:
		log.Println("ERROR:", message)
	case FATAL:
		log.Fatal("FATAL:", message)
	default:
		log.Fatal("ERROR:", "wrong log level!")
	}
}
