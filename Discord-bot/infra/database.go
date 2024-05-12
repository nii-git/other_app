package infra

import (
	"database/sql"
	"discord/config"
	"discord/model"
	"fmt"

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

func (d *DB) InsertTransaction(userid string) error {
	query := fmt.Sprintf(`
	INSERT INTO transaction(user_id)
	VALUES(%s)
	`, userid)

	_, err := d.DBConnection.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetTransaction(userid string) (*model.Transaction, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM transaction
	WHERE user_id = %s
	AND deleted_at IS NULL
	`, userid)

	row := d.DBConnection.QueryRow(query)
	var p model.Transaction

	err := row.Scan(&p.UserID, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (d *DB) DeleteTransaction(userid string) error {
	query := fmt.Sprintf(`
	UPDATE transaction 
	SET deleted_at = "%s"
	WHERE user_id = "%s"
	`, time.Now().Format("2006-01-02 15:04:05"), userid)

	_, err := d.DBConnection.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetStudyTime(userid string) (*model.StudyTimes, error) {
	query := fmt.Sprintf(`
	SELECT * 
	FROM study_times
	WHERE user_id = %s
	AND deleted_at IS NULL
	`, userid)

	row := d.DBConnection.QueryRow(query)
	var p model.StudyTimes

	err := row.Scan(&p.UserID, &p.StudyTimeMinutes, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		} else {
			return nil, nil
		}

	}
	return &p, nil
}

func (d *DB) UpsertStudyTime(userid string, studyTimeMinutes int) error {
	query := fmt.Sprintf(`
	INSERT INTO study_times(user_id,study_time_minutes)
	VALUES(%s,%d)
	ON DUPLICATE KEY
	UPDATE study_time_minutes = %d
	`, userid, studyTimeMinutes, studyTimeMinutes)

	_, err := d.DBConnection.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
