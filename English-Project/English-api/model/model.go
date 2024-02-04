package model

type (
	Frequency_request struct {
		Date     string `param:"date" validate:"date_validation,required"`
		Order    string `param:"order"`
		Limit    int    `param:"limit" validate:"limit_validation"`
		Page     int    `param:"page" validate:"page_validation"`
		Provider string `param:"provider"` // validationはusecaseで実施
	}

	Frequency_response struct {
		Error string                    `json:"error"`
		Body  []Frequency_response_body `json:"body"`
	}
	Frequency_response_body struct {
		Word        string        `json:"word"`
		Count       int           `json:"count"`
		ProviderId  string        `json:"provider_id"`
		Translation []Translation `json:"translation"`
	}
	Translation struct {
		Wordtype string `json:"wordtype"`
		WordJp   string `json:"word_jp"`
	}

	GetProvider_request struct {
		Limit int `param:"limit" validate:"limit_validation"`
		Page  int `param:"page" validate:"page_validation"`
	}

	GetProvider_response struct {
		Error string                      `json:"error"`
		Body  []GetProvider_response_body `json:"body"`
	}

	GetProvider_response_body struct {
		Id       string `json:"id"`
		SiteName string `json:"site_name"`
		Url      string `json:"url"`
	}
)
