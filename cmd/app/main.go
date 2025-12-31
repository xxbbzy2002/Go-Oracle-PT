package main

import (
	"log"

	"go_ora/internal/config"
	"go_ora/internal/database"
	"go_ora/internal/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := logger.Init(cfg.Log.Dir); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Close()

	logger.Log.Info("application starting")

	oracleCfg := database.OracleConfig{
		Host:     cfg.Oracle.Host,
		Port:     cfg.Oracle.Port,
		Service:  cfg.Oracle.Service,
		User:     cfg.Oracle.User,
		Password: cfg.Oracle.Password,
	}

	if err := database.Connect(oracleCfg); err != nil {
		logger.Log.Fatal("failed to connect to oracle", zap.Error(err))
	}
	defer database.Close()

	logger.Log.Info("application started successfully")

	// Your business logic here
}
