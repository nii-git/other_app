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

func (u *Usecase) FrequencyUsecase(request model.Frequency_request) (response *model.Frequency_response, err error) {
	u.logger.Debug("Frequency usecase called")

	// Providerのvalidation
	// handlerのvalidationだとSQL操作があるのでこちらで実施

	var result []model.FrequenciesCountDB

	if request.Provider != "" {
		var isExistProvider bool
		isExistProvider, err = u.db.ValidateProvider(request.Provider)
		if err != nil || !isExistProvider {
			u.logger.Error("FrequencyUsecase Provider Validation Error: " + err.Error())
			return nil, err
		}
		result, err = u.db.GetFrequencyByProvider(request.Date, request.Provider, request.Limit, request.Page*request.Limit)
	} else {
		// providerが空の場合は全providerの合計を用いる
		result, err = u.db.GetFrequencyALL(request.Date, request.Limit, request.Page)
	}

	if err != nil {
		u.logger.Error("FrequencyUsecase DBError: " + err.Error())
		return nil, err
	}

	// dbmodel から responsemodelに変換
	resbody := make([]model.Frequency_response_body, len(result))
	for i, v := range result {
		resbody[i] = model.Frequency_response_body{
			Word:        v.WordName,
			Count:       v.Count,
			ProviderId:  v.ProviderId,
			Translation: []model.Translation{},
		}
	}
	response = new(model.Frequency_response)
	response.Body = resbody

	return response, nil
}

func (u *Usecase) GetProviderUsecase(request model.GetProvider_request) (response *model.GetProvider_response, err error) {
	u.logger.Debug("GetProvider usecase called")
	result, err := u.db.GetMstProvider(request.Limit, request.Page)

	if err != nil {
		u.logger.Error("GetProviderUsecase DBError: ", err)
		return nil, err
	}

	// dbmodel から responsemodelに変換
	resbody := make([]model.GetProvider_response_body, len(result))
	for i, v := range result {
		resbody[i] = model.GetProvider_response_body(v)
	}

	response = new(model.GetProvider_response)
	response.Body = resbody
	return response, nil
}
