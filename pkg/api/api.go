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
	Errors  any         `json:"errors"`
	Message string      `json:"message"`
}

type Pagination[T any] struct {
	Data       []T `json:"data"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
}

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	message := err.Error()
	var errs []interface{}
	var errCode uierr.Code
	if uiErr, ok := err.(uierr.Error); ok {
		errCode = uiErr.Code
		message = uiErr.Message

		if uiErr.Details != nil {
			errs = uiErr.Details
		}

	}
	_ = json.NewEncoder(w).Encode(&Format{
		Success: false,
		Data:    nil,
		Errors:  errs,
		Message: message,
		Code:    errCode,
	})
}

func Write[T any](w http.ResponseWriter, d T) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&Format{
		Success: true,
		Data:    d,
		Errors:  nil,
		Message: "Success",
	})
}
