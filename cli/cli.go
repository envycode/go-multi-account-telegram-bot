package cli

import (
	"fmt"
	"log"
	"multi-account-telegram-bot/config"
	"multi-account-telegram-bot/constant"
	"multi-account-telegram-bot/router"
	"net/http"
	"os"
	"time"
)

func Run() {
	os.RemoveAll(constant.InterpreterDir)
	if err := os.Mkdir(constant.InterpreterDir, 0777); err != nil {
		log.Println("can't initiate directory needed", err)
	}

	cfg := config.AppConfig()

	srv := &http.Server{
		Handler:      router.SetupRoute(),
		Addr:         fmt.Sprintf(":%v", cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("listening application on :%v", cfg.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("unexpected server run failed: %s", err)
	}
}
