package model

type (
	FrequenciesCountDB struct {
		ProviderId string `db:"provider_id"`
		WordName   string `db:"word"`
		Count      int    `db:"count"`
		Date       string `db:"date"`
	}
)
