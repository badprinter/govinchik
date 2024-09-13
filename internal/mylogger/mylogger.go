package mylogger

import (
	"log"
	"os"
)

type MyLogger struct {
	logInfo    *log.Logger
	logWarning *log.Logger
	logError   *log.Logger
	logPanic   *log.Logger
	logFatal   *log.Logger
}

func NewMyLogger() *MyLogger {
	return &MyLogger{
		logInfo:    log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile),
		logWarning: log.New(os.Stdout, "WARNING: ", log.LstdFlags|log.Lshortfile),
		logError:   log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile),
		logPanic:   log.New(os.Stderr, "PANIC: ", log.LstdFlags|log.Lshortfile),
		logFatal:   log.New(os.Stderr, "FATAL: ", log.LstdFlags|log.Lshortfile),
	}
}

func (ml *MyLogger) INFO(message string) {
	ml.logInfo.Println(message)
}
func (ml *MyLogger) WARNING(message string) {
	ml.logWarning.Println(message)
}
func (ml *MyLogger) ERROR(message string) {
	ml.logError.Println(message)
}
func (ml *MyLogger) PANIC(message string) {
	ml.logPanic.Panicln(message)
}
func (ml *MyLogger) FATAL(message string) {
	ml.logFatal.Fatalln(message)
}
