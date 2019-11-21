package logger

import (
	"github.com/aseTo2016/app-footstone/pkg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var (
	Infof  = func(string, ...interface{}) {}
	Info   = func(...interface{}) {}
	Errorf = func(string, ...interface{}) {}
	Error  = func(...interface{}) {}
	Warnf  = func(string, ...interface{}) {}
	Warn   = func(...interface{}) {}
	Debugf = func(string, ...interface{}) {}
	Debug  = func(...interface{}) {}
)

func init() {
	cfg, err := loadConfigs()
	if err != nil {
		panic(err)
	}

	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	Infof = logger.Sugar().Infof
	Info = logger.Sugar().Info
	Errorf = logger.Sugar().Errorf
	Error = logger.Sugar().Error
	Warnf = logger.Sugar().Warnf
	Warn = logger.Sugar().Warn
	Debugf = logger.Sugar().Debugf
	Debug = logger.Sugar().Debug

}

func loadConfigs() (*zap.Config, error) {
	appConfigsDir := os.Getenv("app_configs_dir")
	if len(appConfigsDir) == 0 {
		curPath := pkg.GetCodeFilePath()
		appConfigsDir = filepath.Join(filepath.Dir(filepath.Dir(curPath)), "configs")
	}

	appConfigsPath := filepath.Join(appConfigsDir, "log.yaml")

	cfg := new(zap.Config)
	err := pkg.LoadYamlData(appConfigsPath, cfg)
	return cfg, err
}
