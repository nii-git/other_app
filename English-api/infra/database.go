package infra

import (
	"database/sql"
	"english-frequency/config"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/exp/slog"
)

type DB struct {
	DBConnection *sql.DB
}

// DBコンストラクタ
func NewDB(config *config.Config, logger *slog.Logger) (*DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logger.Error("NewDB Time LoadLocation Error:" + err.Error())
		return nil, err
	}
	c := mysql.Config{
		DBName:    config.DBName,
		User:      config.DBUserName,
		Passwd:    config.DBPassword,
		Addr:      config.DBAddress,
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		logger.Error("NewDB SQL Open Error:" + err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error("NewDB SQL Ping Error:" + err.Error())
		return nil, err
	} else {
		logger.Info("DB has been connected")
	}
	return &DB{DBConnection: db}, err
}
