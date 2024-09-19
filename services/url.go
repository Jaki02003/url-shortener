package services

import (
	"go-redis-url-shortener/types"
)

type IUrlService interface {
	CreateShortUrl(ucr types.URLCreationRequest) (string, error)
	ReturnLongUrl(shortUrl string) (string, error)
}
