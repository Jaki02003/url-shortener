package impl

import (
	"go-redis-url-shortener/consts"
	"go-redis-url-shortener/repositories"
	"go-redis-url-shortener/services"
	"go-redis-url-shortener/shortener"
	"go-redis-url-shortener/types"
)

type urlService struct {
	cacheRepo repositories.ICacheRepository
}

func NewUrlService(cacheRepo repositories.ICacheRepository) services.IUrlService {
	return &urlService{
		cacheRepo: cacheRepo,
	}
}

func (us *urlService) CreateShortUrl(cr types.URLCreationRequest) (string, error) {
	shortUrl := shortener.GenerateShortURL(cr.LongURL, cr.UserId)
	err := us.cacheRepo.Set(shortUrl, cr.LongURL, consts.CacheDuration)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (us *urlService) ReturnLongUrl(shortUrl string) (string, error) {
	initialUrl, err := us.cacheRepo.Get(shortUrl)
	if err != nil {
		return "", err
	}
	return initialUrl, nil
}
