package service

import (
	"context"
	"errors"
	"fmt"
	"multi-account-telegram-bot/bot_manager"
	"multi-account-telegram-bot/contract"
)

type BotDeregisterService interface {
	Exec(ctx context.Context, req contract.DeregisterBotContractReq) (contract.DeregisterBotContractRes, error)
}

type BotDeregisterServiceImpl struct {
	Manager bot_manager.BotManager
}

func (b BotDeregisterServiceImpl) Exec(ctx context.Context, req contract.DeregisterBotContractReq) (contract.DeregisterBotContractRes, error) {
	if err := b.Manager.Kill(req.ID); err != nil {
		return contract.DeregisterBotContractRes{}, errors.New(fmt.Sprintf("error when removing bot: %s", err))
	}
	return contract.DeregisterBotContractRes{
		BaseApiContract: contract.BaseApiContract{Message: "bot removed!"},
		ID:              req.ID,
	}, nil
}
