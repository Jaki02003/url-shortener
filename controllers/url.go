package controllers

import (
	"go-redis-url-shortener/rest_errors"
	"go-redis-url-shortener/services"
	"go-redis-url-shortener/types"
	"go-redis-url-shortener/utils/log"
	"go-redis-url-shortener/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UrlController struct {
	urlService services.IUrlService
}

// NewUrlController will initialize the controllers
func NewUrlController(mailService services.IUrlService) *UrlController {
	return &UrlController{
		urlService: mailService,
	}
}

func (urlCtr *UrlController) CreateShortUrl(c echo.Context) error {
	var req types.URLCreationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(response.GenerateErrorResponseBody(rest_errors.ErrParsingRequestBody))
	}

	resp, err := urlCtr.urlService.CreateShortUrl(req)
	if err != nil {
		log.Error(err)
		return c.JSON(response.GenerateErrorResponseBody(rest_errors.ErrCreatingShortUrl))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"short_url": resp,
	})
}

func (urlCtr *UrlController) ReturnLongUrl(c echo.Context) error {
	shortUrl := c.Param("short-url")

	if shortUrl == "" {
		return c.JSON(response.GenerateErrorResponseBody(rest_errors.ErrShortUrlNotProvided))
	}

	resp, err := urlCtr.urlService.ReturnLongUrl(shortUrl)
	if err != nil {
		log.Error(err)
		return c.JSON(response.GenerateErrorResponseBody(rest_errors.ErrCreatingShortUrl))
	}

	return c.JSON(http.StatusOK, types.LongUrlResponse{
		LongURL:  resp,
		ShortURL: shortUrl,
	})
}

func (urlCtr *UrlController) RedirectToActualUrl(c echo.Context) error {
	shortUrl := c.Param("short-url")
	resp, err := urlCtr.urlService.ReturnLongUrl(shortUrl)
	if err != nil {
		log.Error(err)
		return c.JSON(response.GenerateErrorResponseBody(rest_errors.ErrCreatingShortUrl))
	}

	return c.Redirect(http.StatusFound, resp)
}
