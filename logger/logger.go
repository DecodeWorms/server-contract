package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

const Stderr = "stderr"

// Config represents the logger configuration
type Config struct {
	Name           string
	Level          zapcore.Level
	WithCaller     bool
	WithStacktrace bool
	OutputPaths    []string
	Debug          bool
}

func CreateLog(config ...Config) *zap.Logger {
	var conf Config
	if len(config) > 0 {
		conf = config[0]
	}

	cfg := zap.Config{
		Encoding:          "json",
		Level:             zap.NewAtomicLevelAt(conf.Level),
		ErrorOutputPaths:  []string{"stderr", "stdout"},
		DisableCaller:     conf.WithCaller,
		DisableStacktrace: !conf.WithStacktrace,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "msg",
			NameKey:       "logger",
			LevelKey:      "level",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			TimeKey:       "time",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			CallerKey:     "caller",
			EncodeCaller:  zapcore.ShortCallerEncoder,
			StacktraceKey: "stacktrace",
		},
	}

	if len(conf.OutputPaths) > 0 {
		cfg.OutputPaths = make([]string, 0)

		for _, p := range conf.OutputPaths {
			if len(p) == 0 {
				continue
			}

			if p != Stderr {
				// Create log files and directories
				dir := filepath.Dir(p)
				if _, err := os.Stat(dir); os.IsNotExist(err) {
					os.MkdirAll(dir, os.ModePerm)
				}
			}

			cfg.OutputPaths = append(cfg.OutputPaths, p)
		}
	}

	// Lod to stderr by default
	if len(cfg.OutputPaths) == 0 {
		cfg.OutputPaths = []string{"stderr"}
	}

	logger, err := cfg.Build()
	if err != nil {
		fmt.Printf("Couldn't create logger: %s", err.Error())
		return nil
	}

	logger, err = cfg.Build()
	if err != nil {
		fmt.Printf("Couldn't create logger: %s", err.Error())
		return nil
	}

	if len(conf.Name) > 0 {
		logger = logger.Named(conf.Name)
	}

	if conf.Debug {
		logger = logger.WithOptions(DebugCore())
	}

	return logger

}

// DebugCore returns a debug configuration to use
func DebugCore() zap.Option {
	return zap.WrapCore(
		func(zapcore.Core) zapcore.Core {
			encoder := zap.NewDevelopmentEncoderConfig()
			encoder.EncodeLevel = customCapitalLevelEncoder()
			return zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), zapcore.AddSync(os.Stderr), zapcore.DebugLevel)
		})
}

// customCapitalLevelEncoder serializes a Level to an all-caps string.
func customCapitalLevelEncoder() func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("[%s]", l.CapitalString()))
	}
}
