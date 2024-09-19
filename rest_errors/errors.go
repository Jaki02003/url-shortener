package rest_errors

import (
	"net/http"
)

var (
	// Common
	ErrCopyStruct            = NewError("failed to copy the structs", http.StatusInternalServerError)
	ErrParsingRequestBody    = NewError("failed to parse request body", http.StatusBadRequest)
	ErrParsingRequestHeaders = NewError("failed to parse request headers", http.StatusBadRequest)
	AccessForbidden          = NewError("access forbidden", http.StatusForbidden)
	ErrJSONMarshal           = NewError("failed to marshal data", http.StatusInternalServerError)
	NoLoggedInUserFound      = NewError("no logged-in user found", http.StatusUnauthorized)
	ErrParsingEndpoint       = NewError("failed to parse endpoint", http.StatusInternalServerError)
	ErrDecodingResponseBody  = NewError("failed to decode response body", http.StatusInternalServerError)

	ErrCreatingShortUrl    = NewError("failed to create the short url", http.StatusInternalServerError)
	ErrShortUrlNotProvided = NewError("no short url provided", http.StatusInternalServerError)
)
