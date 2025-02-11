package logger

import (
	"log"
	"os"
)

type Logger struct {
	Info *log.Logger
	Err  *log.Logger
}

func Init() *Logger {

	logger := Logger{}

	logger.Info = log.New(os.Stdout, "WordF1nder | ", int(" "[0]))
	logger.Err = log.New(os.Stdout, "[!ОШИБКА!]\nWordF1nder: ", log.Lshortfile)

	return &logger
}
