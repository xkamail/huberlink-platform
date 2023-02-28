package api

import (
	"encoding/json"
	"net/http"

	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

type Format struct {
	Code    uierr.Code  `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  []any       `json:"errors"`
	Message string      `json:"message"`
}

func WriteError(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode()
}

func Write[T any](w http.ResponseWriter, d T) {
	json.NewEncoder(w).Encode()
}
