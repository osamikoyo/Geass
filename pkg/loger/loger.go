package loger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zerolog.Logger
}

func New() Logger {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel)


	fileWriter := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10, // Максимальный размер файла в мегабайтах
		MaxBackups: 3,  // Максимальное количество старых файлов
		MaxAge:     28, // Максимальное количество дней хранения
		Compress:   true, // Сжатие старых файлов
	}

	fileLogger := zerolog.New(fileWriter)


	multi := zerolog.MultiLevelWriter(logger, fileLogger)
	lg := zerolog.New(multi).With().Timestamp().Caller().Logger()

	return Logger{Logger: &lg}
}