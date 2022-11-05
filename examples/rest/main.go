package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/marioarranzr/go_lnmarkets/rest"
)

func main() {
	milli := time.Now().UnixNano() / int64(time.Millisecond)

	api := rest.New(
		os.Getenv("LNMARKETS_API_KEY"),
		os.Getenv("LNMARKETS_SECRET"),
		os.Getenv("LNMARKETS_PASSPHRASE"),
		fmt.Sprint(milli),
	)

	tr, err := api.Ticker()
	if err != nil {
		log.Panicln(err)
		return
	}
	log.Printf("%+v", *tr)

	pr, err := api.Positions()
	if err != nil {
		log.Panicln(err)
		return
	}
	log.Printf("%+v", *pr)

	ur, err := api.User()
	if err != nil {
		log.Panicln(err)
		return
	}
	log.Printf("%+v", *ur)

	amr, err := api.AddMargin(100, "99c470e1-2e03-4486-a37f-1255e08178b1")
	if err != nil {
		log.Panicln(err)
		return
	}
	log.Printf("%+v", *amr)

	npr, err := api.NewPosition(rest.OrderTypeMarket, rest.OrderSideBuy, 10000, 0, 0, 0, 0, 0)
	if err != nil {
		log.Panicln(err)
		return
	}
	log.Printf("%+v", *npr)
}
