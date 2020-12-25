package service

import (
	"context"
	"errors"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"multi-account-telegram-bot/bot_manager"
	"multi-account-telegram-bot/contract"
)

type BotRegisterService interface {
	Exec(ctx context.Context, req contract.RegisterBotContractReq) (contract.RegisterBotContractRes, error)
}

type BotRegisterServiceImpl struct {
	Manager bot_manager.BotManager
}

func (b BotRegisterServiceImpl) Exec(ctx context.Context, req contract.RegisterBotContractReq) (contract.RegisterBotContractRes, error) {
	if err := b.Manager.Spawn(bot_manager.Bot{
		Id:    req.ID,
		Token: req.Token,
		Actions: map[string]func(m *tb.Message) string{
			// sample actions
			"/hello": func(m *tb.Message) string {
				return fmt.Sprintf("Hello %s! Nice to meet you", m.Sender.FirstName)
			},
		},
	}); err != nil {
		return contract.RegisterBotContractRes{}, errors.New(fmt.Sprintf("error when adding new bot: %s", err))
	}
	return contract.RegisterBotContractRes{
		BaseApiContract: contract.BaseApiContract{Message: "bot added!"},
		ID:              req.ID,
	}, nil
}
