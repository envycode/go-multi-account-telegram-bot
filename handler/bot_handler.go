package handler

import (
	"github.com/go-playground/validator"
	"multi-account-telegram-bot/contract"
	"multi-account-telegram-bot/service"
	"multi-account-telegram-bot/util"
	"net/http"
)

type BotHandler struct {
	Register     service.BotRegisterService
	Deregister   service.BotDeregisterService
	RegisterFunc service.BotAddFunction
}

func (b BotHandler) Create(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	util.UseJsonFieldValidation(validate)

	req, err := contract.NewRegisterBotContractReq(r)
	if err != nil {
		util.RenderErr(w, err)
		return
	}
	if err := validate.Struct(req); err != nil {
		util.RenderErr(w, err)
		return
	}
	res, err := b.Register.Exec(r.Context(), req)
	if err != nil {
		util.RenderErr(w, err)
		return
	}
	util.RenderDefault(w, res)
}

func (b BotHandler) Delete(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	util.UseJsonFieldValidation(validate)

	req, err := contract.NewDeregisterBotContractReq(r)
	if err != nil {
		util.RenderErr(w, err)
		return
	}
	if err := validate.Struct(req); err != nil {
		util.RenderErr(w, err)
		return
	}
	res, err := b.Deregister.Exec(r.Context(), req)
	if err != nil {
		util.RenderErr(w, err)
		return
	}
	util.RenderDefault(w, res)
}

func (b BotHandler) CreateFunction(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	util.UseJsonFieldValidation(validate)

	req, err := contract.NewRegisterFunctionContractReq(r)
	if err != nil {
		util.RenderErr(w, err)
		return
	}
	if err := validate.Struct(req); err != nil {
		util.RenderErr(w, err)
		return
	}
	res, err := b.RegisterFunc.Exec(r.Context(), req)
	if err != nil {
		util.RenderErr(w, err)
		return
	}
	util.RenderDefault(w, res)
}
