package logger

import (
	"bus_depot/internal/configs"
	"fmt"
	"io"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Глобальные логгеры
var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
)

func Init() error {
	logParams := configs.AppSettings.LogParams

	fmt.Println("LogDirectory =", logParams.LogDirectory) 

	if err := os.MkdirAll(logParams.LogDirectory, 0755); err != nil {
		return fmt.Errorf("не удалось создать директорию логов: %w", err)
	}

	// Настройка файлов логов
	lumberLogInfo := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogInfo),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAgeDays,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogError),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAgeDays,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	lumberLogWarn := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogWarn),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAgeDays,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogDebug),
		MaxSize:    logParams.MaxSizeMegabytes,
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAgeDays,
		Compress:   logParams.Compress,
		LocalTime:  logParams.LocalTime,
	}

	
	gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	// Инициализация 
	Info = log.New(gin.DefaultWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(lumberLogError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn  = log.New(lumberLogWarn, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(lumberLogDebug, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info.Println("Логгер инициализирован")

	return nil
}
