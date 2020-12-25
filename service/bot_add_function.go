package service

import (
	"context"
	"multi-account-telegram-bot/bot_manager"
	"multi-account-telegram-bot/contract"
)

type BotAddFunction interface {
	Exec(ctx context.Context, req contract.RegisterFunctionContractReq) (contract.RegisterFunctionContractRes, error)
}

type BotAddFunctionImpl struct {
	Manager bot_manager.BotManager
}

func (b BotAddFunctionImpl) Exec(ctx context.Context, req contract.RegisterFunctionContractReq) (contract.RegisterFunctionContractRes, error) {
	if err := b.Manager.RegisterFunc(req.ID, req.Name, req.Action); err != nil {
		return contract.RegisterFunctionContractRes{}, err
	}
	return contract.RegisterFunctionContractRes{
		BaseApiContract: contract.BaseApiContract{Message: "function registered!"},
	}, nil
}
