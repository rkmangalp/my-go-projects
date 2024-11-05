package main

import (
	"fmt"

	"github.com/rkmangalp/discord-ping/bot"
	"github.com/rkmangalp/discord-ping/config"
)

func main() {
	err := config.Readconfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
