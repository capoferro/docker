package httputils

import (
	"net/http"

	"github.com/docker/docker/pkg/errorutils"
)

func NewHTTPRequestError(msg string, res *http.Response) error {
	return &errorutils.JSONError{
		Message: msg,
		Code:    res.StatusCode,
	}
}
