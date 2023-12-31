package infra

import (
	"database/sql"
	"english-frequency/config"
	"strconv"
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
		DBName:               config.DBName,
		User:                 config.DBUserName,
		Passwd:               config.DBPassword,
		Addr:                 config.DBAddress,
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}

	var db *sql.DB

	for r := 1; r <= config.MaxDBRetryCount; r++ {
		db, err = sql.Open("mysql", c.FormatDSN())
		if err != nil {
			logger.Error("NewDB SQL Connection Attept #" + strconv.Itoa(r))
			logger.Error("NewDB SQL Connection Error:" + err.Error())
			time.Sleep(10 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			logger.Error("NewDB SQL Connection Attept #" + strconv.Itoa(r))
			logger.Error("NewDB SQL Connection Error:" + err.Error())
			time.Sleep(10 * time.Second)
			continue
		} else {
			break
		}
	}

	if err != nil {
		logger.Error("NewDB SQL Connection Error:" + err.Error())
		return nil, err
	}

	logger.Debug("NewDB SQL Connection Succeeded!")
	return &DB{DBConnection: db}, err
}
