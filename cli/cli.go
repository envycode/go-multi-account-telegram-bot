package cli

import (
	"fmt"
	"log"
	"multi-account-telegram-bot/config"
	"multi-account-telegram-bot/router"
	"net/http"
	"time"
)

func Run() {
	cfg := config.AppConfig()

	srv := &http.Server{
		Handler:      router.SetupRoute(),
		Addr:         fmt.Sprintf(":%v", cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("unexpected server run failed: %s", err)
	}
}
