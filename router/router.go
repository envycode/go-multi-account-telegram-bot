package router

import (
	"github.com/gorilla/mux"
	tb "gopkg.in/tucnak/telebot.v2"
	"multi-account-telegram-bot/bot_manager"
	"multi-account-telegram-bot/handler"
	"multi-account-telegram-bot/service"
	"net/http"
	"sync"
)

func SetupRoute() *mux.Router {
	r := mux.NewRouter()
	pool := bot_manager.BotPool{
		Bots:  map[string]*tb.Bot{},
		Mutex: sync.Mutex{},
	}

	manager := bot_manager.BotManagerImpl{
		ManagedBot: &pool,
	}

	registerSvc := service.BotRegisterServiceImpl{Manager: manager}
	deregisterSvc := service.BotDeregisterServiceImpl{Manager: manager}

	h := handler.BotHandler{
		Register:   registerSvc,
		Deregister: deregisterSvc,
	}

	r.HandleFunc("/register", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/deregister", h.Delete).Methods(http.MethodDelete)

	return r
}
