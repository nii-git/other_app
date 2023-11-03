package model

type (
	frequency_request struct {
		Date  string `json:"date"`
		Order string `json:"order"`
		Limit int    `json:"limit"`
		Page  int    `json:"page"`
	}

	frequency_response struct {
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
