package usecase

import (
	"english-frequency/infra"
	"english-frequency/model"
	"fmt"

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

func (u *Usecase) FrequencyUsecase(request model.Frequency_request) (response *model.Frequency_response, err error) {
	u.logger.Debug("Frequency usecase called")
	query := fmt.Sprintf(`
	SELECT frequency.provider_id, word.word, frequency.count, frequency.date
	FROM frequency
	LEFT JOIN word
	ON frequency.word_id = word.id
	WHERE date = "%s"
	ORDER BY frequency.count DESC
	LIMIT %d
	`, request.Date, request.Limit)

	res, err := u.db.DBConnection.Query(query)

	if err != nil {
		u.logger.Error("FrequencyUsecase QueryError: " + err.Error())
		return nil, err
	}

	var result []model.Frequencies_Get_Database

	for res.Next() {
		var r model.Frequencies_Get_Database
		err = res.Scan(&r.ProviderId, &r.WordName, &r.Count, &r.Date)
		if err != nil {
			u.logger.Error("FrequencyUsecase ScanError : " + err.Error())
			return nil, err
		}
		result = append(result, r)
	}

	// dbmodel から responsemodelに変換
	resbody := make([]model.Frequency_response_body, len(result))
	for i, v := range result {
		resbody[i] = model.Frequency_response_body{
			Word:        v.WordName,
			Count:       v.Count,
			Translation: []model.Translation{},
		}
	}

	response.Body = resbody

	fmt.Printf("%#v \n", result)
	return nil, nil
}
