package bot_manager

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"sync"
	"time"
)

type BotManager interface {
	Spawn(bot Bot) error
	Kill(id string) error
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
