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
		Body  []Frequency_response_body `json:"body"`
	}
	Frequency_response_body struct {
		Word        string        `json:"word"`
		Count       int           `json:"int"`
		Translation []Translation `json:"translation"`
	}
	Translation struct {
		Wordtype string `json:"wordtype"`
		WordJp   string `json:"word_jp"`
	}
)
