package util

import (
	"encoding/json"
	"multi-account-telegram-bot/contract"
	"net/http"
)

func RenderDefault(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		RenderErr(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func RenderErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	msg := contract.BaseApiContract{Message: err.Error()}
	res, err := json.Marshal(msg)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write(res)
}
