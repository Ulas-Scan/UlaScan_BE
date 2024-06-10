package dto

import "errors"

var (
	ErrCreateHttpRequest    = errors.New("failed to create http request")
	ErrSendsHttpRequest     = errors.New("failed to sends http request")
	ErrReadHttpResponseBody = errors.New("failed to read http response body")
	ErrParseJson            = errors.New("failed to parse response json")
	ErrNotOk                = errors.New("received non-200 response code")
)
