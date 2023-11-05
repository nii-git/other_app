package main

import (
	"english-frequency/config"
	"english-frequency/handler"
	"english-frequency/infra"
	"english-frequency/model"
	"english-frequency/usecase"
	"os"

	"golang.org/x/exp/slog"

	"github.com/labstack/echo/v4"
)

// サーバー
type Server struct {
	logger  slog.Logger
	config  config.Config
	db      infra.DB
	handler handler.Handler
}

// サーバーコンストラクタ
func NewServer(logger slog.Logger, config config.Config, db infra.DB, handler handler.Handler) *Server {
	return &Server{
		logger:  logger,
		config:  config,
		db:      db,
		handler: handler,
	}
}

// サーバー起動、ルーティング登録
func (s *Server) Start() error {
	e := echo.New()
	e.GET("/", s.handler.HandlerFunc())
	e.GET("/frequency", s.handler.FrequencyHandlerFunc())
	return e.Start(s.config.ServerAddress)
}

func main() {

	// configの作成
	config, err := config.NewConfig()

	if err != nil {
		println("Server Init Error " + err.Error())
		panic(err)
	}

	// sloggerの初期化
	loglevel := slog.Level(-8)
	switch config.LogLevel {
	case "DEBUG":
		loglevel = slog.LevelDebug
	case "INFO":
		loglevel = slog.LevelInfo
	case "WARN":
		loglevel = slog.LevelWarn
	case "ERROR":
		loglevel = slog.LevelError
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: loglevel}))

	// db接続
	db, err := infra.NewDB(config, logger)

	if err != nil {
		logger.Error("Server init Error " + err.Error())
		panic(err)
	}

	// validator
	validator := model.NewValidator()

	// usecase
	usecase := usecase.NewUsecase(*logger, *db)

	// handler
	handler := handler.NewHandler(*logger, *usecase, *validator)

	server := NewServer(*logger, *config, *db, *handler)

	if err = server.Start(); err != nil {
		logger.Error("Server init Error " + err.Error())
		panic(err)
	}

}
