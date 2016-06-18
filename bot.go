package main

import (
	"fmt"
	"os"
	"os/signal"
	"flag"

	"github.com/bwmarrin/discordgo"

	"github.com/sbaildon/bot/handlers"

	_ "github.com/sbaildon/bot/commands/hello"
	_ "github.com/sbaildon/bot/commands/roles"
	_ "github.com/sbaildon/bot/commands/list"
)

func main() {
	var (
		Token = flag.String("token", "", "Discord auth token")
	)
	flag.Parse()

	dg, err := discordgo.New(*Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	handlers.SetupHandlers(dg)

	err = dg.Open()
	if err != nil {
		fmt.Println("Unable to open discord connection")
	}

	fmt.Println("Connected. Maybe?")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
