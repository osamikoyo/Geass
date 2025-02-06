package loger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zerolog.Logger
}

func New(dir string) Logger {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel)


	fileWriter := &lumberjack.Logger{
		Filename:   filepath.Join(dir, "app.log"),
		MaxSize:    10, // Максимальный размер файла в мегабайтах
		MaxBackups: 3,  // Максимальное количество старых файлов
		MaxAge:     28, // Максимальное количество дней хранения
		Compress:   true, // Сжатие старых файлов
	}

	fileLogger := zerolog.New(fileWriter)


	multi := zerolog.MultiLevelWriter(logger, fileLogger)
	lg := zerolog.New(multi).With().Timestamp().Logger()

	return Logger{Logger: &lg}
}