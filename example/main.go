package main

import (
	"errors"
	"github.com/Sjhzjxc/go_logger"
	"go.uber.org/zap"
	"log"
)

var Logger *zap.SugaredLogger

func Level1() {
	Logger.Debug("Level1")
	Logger.Info("Level1")
	Logger.Warn("Level1")
	Logger.Error("Level1")
	//Logger.Panic("Level1")
	//Logger.DPanic("Level1")
	Level2()
}
func Level2() {
	Logger.Debug("Level2")
	Logger.Info("Level2")
	Logger.Warn("Level2")
	Logger.Error("Level2")
	err := Level3()
	Logger.Error(err.Error())
	//Logger.Panic("Level2")
	//Logger.DPanic("Level2")
}

func Level3() error {
	return errors.New("Level3")
	//Logger.Panic("Level2")
	//Logger.DPanic("Level2")
}

func main() {
	logger, err := go_logger.NewLogger(&go_logger.LogConfig{
		Director:    "./logs",
		Level:       "warn",
		FileExt:     "log",
		FileName:    "server",
		LinkName:    "latest_log",
		Format:      "json",
		WithConsole: true,
	})
	if err != nil {
		log.Panicln(err.Error())
	}
	Logger = logger
	Level1()
}
