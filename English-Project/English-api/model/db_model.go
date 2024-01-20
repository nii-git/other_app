package model

type (
	FrequenciesCountDB struct {
		ProviderId string `db:"provider_id"`
		WordName   string `db:"word"`
		Count      int    `db:"count"`
		Date       string `db:"date"`
	}
	MstProviderDB struct {
		Id       string `json:"id"`
		SiteName string `json:"site_name"`
		Url      string `json:"url"`
	}
)
