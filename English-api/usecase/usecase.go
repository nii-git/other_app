package usecase

import (
	"english-frequency/infra"
	"english-frequency/model"

	"golang.org/x/exp/slog"
)

type Usecase struct {
	logger slog.Logger
	db     infra.DB
}

func NewUsecase(logger slog.Logger, db infra.DB) *Usecase {
	return &Usecase{
		logger: logger,
		db:     db,
	}
}

func (u *Usecase) FrequencyUsecase(request model.Frequency_request) {
	u.logger.Debug("Frequency usecase called")
	// query := fmt.Sprintf(`
	// SELECT *
	// FROM frequency
	// LIMIT %d
	// `, request.Limit)

	// res, _ := u.db.DBConnection.Query(query)
	// c, _ := res.Scan()
	// fmt.Printf("%#v \n", c)
}
