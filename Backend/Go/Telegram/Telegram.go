package Telegram

import (
	"fmt"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func Test() {

	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_APITOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	go func() {
		bot.Handle(tele.OnText, func(c tele.Context) error {
			fmt.Println("eyy")
			// All the text messages that weren't
			// captured by existing handlers.

			var (
				user = c.Sender()
				text = c.Text()
			)

			// Use full-fledged bot's functions
			// only if you need a result:
			msg, err := bot.Send(user, text)
			if err != nil {
				return err
			}

			// Instead, prefer a context short-hand:
			return c.Send(msg)
		})
	}()
	bot.Start()
}
