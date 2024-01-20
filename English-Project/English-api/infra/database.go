package infra

import (
	"database/sql"
	"english-frequency/config"
	"english-frequency/model"
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

func (d *DB) GetFrequencyByProvider(date string, provider string, limit int, page int) (result []model.FrequenciesCountDB, err error) {
	query := fmt.Sprintf(`
	SELECT frequency.provider_id, word.word, frequency.count, frequency.date
	FROM frequency
	LEFT JOIN word
	ON frequency.word_id = word.id
	WHERE date = "%s"
	AND provider_id = "%s"
	ORDER BY frequency.count DESC
	LIMIT %d
	OFFSET %d ;
	`, date, provider, limit, page*limit)

	res, err := d.DBConnection.Query(query)

	if err != nil {
		return nil, err
	}

	for res.Next() {
		var r model.FrequenciesCountDB
		err = res.Scan(&r.ProviderId, &r.WordName, &r.Count, &r.Date)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}
func (d *DB) GetFrequencyALL(date string, limit int, page int) (result []model.FrequenciesCountDB, err error) {
	query := fmt.Sprintf(`
	SELECT word.word, SUM(frequency.count) as count, frequency.date
	FROM frequency
	LEFT JOIN word
	ON frequency.word_id = word.id
	WHERE date = "%s"
	GROUP BY word.id
	ORDER BY SUM(frequency.count) DESC
	LIMIT %d
	OFFSET %d ;
	`, date, limit, page*limit)

	res, err := d.DBConnection.Query(query)

	if err != nil {
		return nil, err
	}

	for res.Next() {
		var r model.FrequenciesCountDB
		err = res.Scan(&r.WordName, &r.Count, &r.Date)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}

func (d *DB) ValidateProvider(provider_id string) (bool, error) {
	query := fmt.Sprintf(`
	SELECT id FROM mst_provider
	WHERE id = "%s";
	`, provider_id)

	row := d.DBConnection.QueryRow(query)
	var p string

	err := row.Scan(&p)
	if err != nil {
		// 該当レコードがない場合は sql.ErrNoRowsが帰る
		return false, err
	}
	return true, nil
}

func (d *DB) GetMstProvider() (result []model.MstProviderDB, err error) {
	query := `
	SELECT *
	FROM mst_provider;
	`

	res, err := d.DBConnection.Query(query)

	if err != nil {
		return nil, err
	}

	for res.Next() {
		var r model.MstProviderDB
		err = res.Scan(&r.Id, &r.SiteName, &r.Url)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}
