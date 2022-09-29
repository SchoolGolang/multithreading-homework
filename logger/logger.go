package logger

import (
	"log"
	"os"
	"sync"
)

type logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

var _logger *logger
var once sync.Once

func GetInstance() *logger {
	once.Do(func() {
		_logger = createLogger()
	})

	return _logger
}

func createLogger() *logger {
	return &logger{
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime),
		errorLogger:   log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime),
	}
}

func (l *logger) Info(msg string) {
	l.infoLogger.Println(msg)
}

func (l *logger) Warn(msg string) {
	l.warningLogger.Println(msg)
}

func (l *logger) Error(msg string) {
	l.errorLogger.Println(msg)
}
