package bot_manager

import (
	"errors"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"multi-account-telegram-bot/action"
	"sync"
	"time"
)

type BotManager interface {
	Spawn(bot Bot) error
	Kill(id string) error
	RegisterFunc(id, name string, action string) error
}

type Bot struct {
	Id      string
	Token   string
	Actions map[string]func(m *tb.Message) string
}

type BotPool struct {
	Bots map[string]*tb.Bot
	sync.Mutex
}

type BotManagerImpl struct {
	ManagedBot *BotPool
	Action     action.Action
}

func (bm BotManagerImpl) RegisterFunc(id, name string, action string) error {
	if err := bm.Action.Create(id, name, action); err != nil {
		return err
	}
	bm.ManagedBot.Lock()
	defer bm.ManagedBot.Unlock()
	b, ok := bm.ManagedBot.Bots[id]
	if !ok {
		return errors.New(fmt.Sprintf("bot not found %s", id))
	}
	b.Handle(fmt.Sprintf("/%s", name), func(m *tb.Message) {
		res, err := bm.Action.Execute(id, name, m, m.Payload)
		if err != nil {
			log.Println(err)
			return
		}
		if _, err := b.Send(m.Sender, res); err != nil {
			log.Println(err)
		}
	})
	go func() {
		b.Stop()
		time.Sleep(15 * time.Second)
		b.Start()
	}()

	return nil
}

func (bm BotManagerImpl) Spawn(bot Bot) error {
	b, err := tb.NewBot(tb.Settings{
		Token:  bot.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Println(err)
		return err
	}

	b.Handle("/ping", func(m *tb.Message) {
		b.Send(m.Sender, "pong")
	})

	for k, v := range bot.Actions {
		b.Handle(k, func(m *tb.Message) {
			if _, err := b.Send(m.Sender, v(m)); err != nil {
				log.Println(err)
			}
		})
	}

	bm.ManagedBot.Lock()
	defer bm.ManagedBot.Unlock()
	bm.ManagedBot.Bots[bot.Id] = b
	go b.Start()

	return nil
}

func (bm BotManagerImpl) Kill(id string) error {
	bm.ManagedBot.Lock()
	defer bm.ManagedBot.Unlock()
	b := bm.ManagedBot.Bots[id]
	b.Stop()
	delete(bm.ManagedBot.Bots, id)
	return nil
}
