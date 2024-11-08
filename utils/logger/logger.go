package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger() error {
	var err error
	logger, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}
	return nil
}

func Sync() {
	if logger != nil {
		logger.Sync()
	}
}

func GetDefaultLogger() *zap.Logger {
	return logger
}

// Các hàm tiện ích cho logger
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	logger.Fatal(msg, fields...)
}
