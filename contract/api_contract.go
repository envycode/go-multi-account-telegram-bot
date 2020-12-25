package contract

import (
	"encoding/json"
	"net/http"
)

type BaseApiContract struct {
	Message string `json:"message"`
}

type RegisterBotContractReq struct {
	ID    string `json:"id" validate:"required"`
	Token string `json:"token" validate:"required"`
}

func NewRegisterBotContractReq(r *http.Request) (RegisterBotContractReq, error) {
	var contract RegisterBotContractReq
	err := json.NewDecoder(r.Body).Decode(&contract)
	if err != nil {
		return RegisterBotContractReq{}, err
	}
	return contract, nil
}

type RegisterBotContractRes struct {
	BaseApiContract
	ID string `json:"id"`
}

func NewDeregisterBotContractReq(r *http.Request) (DeregisterBotContractReq, error) {
	var contract DeregisterBotContractReq
	err := json.NewDecoder(r.Body).Decode(&contract)
	if err != nil {
		return DeregisterBotContractReq{}, err
	}
	return contract, nil
}

type DeregisterBotContractReq struct {
	ID string `json:"id" validate:"required"`
}

type DeregisterBotContractRes struct {
	BaseApiContract
	ID string `json:"id"`
}
