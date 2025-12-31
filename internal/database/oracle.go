package database

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
	"go_ora/internal/logger"
	"go.uber.org/zap"
)

type OracleConfig struct {
	Host     string
	Port     int
	Service  string
	User     string
	Password string
}

var DB *sql.DB

func Connect(cfg OracleConfig) error {
	connStr := fmt.Sprintf("oracle://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Service)

	var err error
	DB, err = sql.Open("oracle", connStr)
	if err != nil {
		logger.Log.Error("failed to open oracle connection", zap.Error(err))
		return err
	}

	if err = DB.Ping(); err != nil {
		logger.Log.Error("failed to ping oracle", zap.Error(err))
		return err
	}

	logger.Log.Info("oracle connected successfully",
		zap.String("host", cfg.Host),
		zap.Int("port", cfg.Port),
		zap.String("service", cfg.Service))
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
		logger.Log.Info("oracle connection closed")
	}
}
