package main

import (
	"english-frequency/handler"
	"log"

	"database/sql"

	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("err")
	}
	defer db.Close()

	e := echo.New()
	e.GET("/", handler.Handler())
	e.GET("/frequency", handler.FrequencyHandler(db))
	e.Logger.Fatal(e.Start(":1323"))

}

func connectDB() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("err")
		return nil, err
	}
	c := mysql.Config{
		DBName:    "english_frequency",
		User:      "root",
		Passwd:    "",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("err")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		log.Println("データベース接続完了")
	}

	return db, err
}
