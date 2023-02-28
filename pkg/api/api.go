package api

import (
	"encoding/json"
	"log"
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
	message := err.Error()
	errs := make([]any, 0)
	var errCode uierr.Code
	if uiErr, ok := err.(uierr.Error); ok {
		log.Printf("uierr: %+v", uiErr)
		errCode = uiErr.Code
		message = uiErr.Message

		if uiErr.Details != nil {
			errs = append(errs, uiErr.Details)
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

	_ = json.NewEncoder(w).Encode(&Format{
		Success: true,
		Data:    d,
		Errors:  nil,
		Message: "Success",
	})
}
