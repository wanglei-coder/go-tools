package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	DefaultEncoderConfig = getDefaultEncoderConfig()
)

var (
	logger = NewLogger(nil, DefaultEncoderConfig, zapcore.DebugLevel, "", "")
	Debug  = logger.Debug
	Info   = logger.Info
	Error  = logger.Error
	Warn   = logger.Warn
	Panic  = logger.Panic
	Fatal  = logger.Fatal
)

func NewLogger(hook *lumberjack.Logger, config zapcore.EncoderConfig, level zapcore.Level, format, serviceName string) *zap.Logger {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	var (
		core   zapcore.Core
		logger *zap.Logger
	)

	// log file
	if hook == nil {
		core = zapcore.NewCore(
			getEncoder(format, config),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			atomicLevel,
		)
	} else {
		core = zapcore.NewCore(
			getEncoder(format, config),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)),
			atomicLevel,
		)
	}

	logger = zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	// service label
	if serviceName != "" {
		logger = logger.WithOptions(zap.Fields(zap.String("serviceName", serviceName)))
	}

	return logger
}

func NewLumberjackLogger(filename string, maxSize, maxAge, maxBackups int, compress bool) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		Compress:   compress,
	}
}

// getEncoder Log save format
func getEncoder(format string, config zapcore.EncoderConfig) zapcore.Encoder {
	if format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)
}

func getDefaultEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		// EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime:  zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeName:  zapcore.FullNameEncoder,
	}
	return config
}
