package model

type (
	Frequency_request struct {
		Date  string `param:"date"`
		Order string `param:"order"`
		Limit int    `param:"limit"`
		Page  int    `param:"page"`
	}

	Frequency_response struct {
		Error string                    `json:"error"`
		Body  []frequency_response_body `json:"body"`
	}
	frequency_response_body struct {
		Word        string        `json:"word"`
		Count       int           `json:"int"`
		Translation []translation `json:"translation"`
	}
	translation struct {
		Wordtype string `json:"wordtype"`
		WordJp   string `json:"word_jp"`
	}
)
