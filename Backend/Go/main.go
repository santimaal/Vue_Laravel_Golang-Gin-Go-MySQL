package main

import (
	"fmt"
	"log"
	"os"
	"sanvic/Config"
	"sanvic/Routes"
	"sanvic/User"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	tele "gopkg.in/telebot.v3"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()
	if err != nil {
		fmt.Println("Status:", err)
	}

	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_APITOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		bot.Handle(tele.OnText, func(c tele.Context) (err error) {
			var (
				user = c.Sender()
				text = c.Text()
			)
			fmt.Println(user.FirstName)
			if text != "/start" {
				bool := User.ComproveCode(text, user.ID)
				if bool {
					msg, _ := bot.Send(user, "User connected correctly")
					c.Send(msg)
				} else {
					msg, _ := bot.Send(user, "Telegram code is not correct")
					c.Send(msg)
				}
			} else {
				msg, _ := bot.Send(user, "Hello "+user.FirstName+"!! Send me your telegram token")
				c.Send(msg)
			}
			return err
		})
		bot.Start()
	}()
	go func() {
		r := Routes.SetupRouter()
		r.Run(":3001")
	}()
	select {}

}
