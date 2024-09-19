package types

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type URLCreationRequest struct {
	LongURL string `json:"long_url"`
	UserId  string `json:"user_id"`
}

func (ucr *URLCreationRequest) Validate() error {
	return v.ValidateStruct(ucr,
		v.Field(&ucr.LongURL, v.Required),
		v.Field(&ucr.UserId, v.Required),
	)
}

type LongUrlResponse struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}
