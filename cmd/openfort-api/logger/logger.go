package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log Logger

type Logger struct {
	*zap.Logger
	Level zap.AtomicLevel
}

func Init() {
	lLev := zap.NewAtomicLevelAt(zap.DebugLevel)
	logConfig := zap.Config{
		OutputPaths: []string{"stderr"},
		Encoding:    "json",
		Level:       lLev,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		},
	}

	log, err := logConfig.Build()
	if err != nil {
		log.Fatal(fmt.Sprintf("error initializing logger: %s", err.Error()))
	}

	Log.Logger = log
	Log.Level = lLev
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

func Error(err error, fields ...zap.Field) {
	Log.Error(err.Error(), fields...)
}
