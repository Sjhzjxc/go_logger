package go_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"time"
)

func getEncoder(format string) zapcore.Encoder {
	if format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return encoderConfig
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(" 2006/01/02 - 15:04:05.000"))
}

func levelValue(level string) zapcore.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}

}

func NewLogger(config *LogConfig) (*zap.SugaredLogger, error) {
	encoder := getEncoder(config.Format)
	logLevel := zap.LevelEnablerFunc(func(logLevel zapcore.Level) bool {
		return logLevel >= levelValue(config.Level)
	})
	writer, err := GetWriteSyncer(config.Director, config.FileName, config.LinkName, config.WithConsole)
	if err != nil {
		return nil, err
	}
	tees := []zapcore.Core{
		zapcore.NewCore(encoder, writer, logLevel),
	}
	core := zapcore.NewTee(tees...)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	return logger, nil
}

func DefaultLogger() (*zap.SugaredLogger, error) {
	config := &LogConfig{
		Director:    "./logs",
		Level:       "warn",
		FileExt:     "log",
		FileName:    "server",
		LinkName:    "latest_log",
		Format:      "json",
		WithConsole: true,
	}
	return NewLogger(config)
}
