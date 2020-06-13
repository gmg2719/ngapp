package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"fmt"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	service:= NewNgApp()
	service.mainWin.Show()

	// test zap
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	_, close, err := zap.Open("./logs/ngapp.log")
	if err != nil {
		close()
		panic(fmt.Sprintf("Fail to create/open log file: %v", err))
	}

	config := zap.Config{
		Level:            atom,
		Development:      true,
		//DisableCaller: false,
		//DisableStacktrace: false,
		//Sampling: nil,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout", "./logs/ngapp.log"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]interface{}{"serviceName": "ngapp"},
	}

	logger, err := config.Build(zap.Hooks(func(entry zapcore.Entry) error {
		service.logEdit.Append(fmt.Sprintf("%v [%v] %v: %s", entry.Time.Format("2006-01-02 15:04:05.999"), entry.Level.CapitalString(), entry.Caller, entry.Message))
		return nil
	}))
	if err != nil {
		panic(fmt.Sprintf("Fail to initialize logger: %v", err))
	}
	defer logger.Sync()

	logger.Info("Logger initialized.")
	logger.Error("hello logger")
	app.Exec()
}
