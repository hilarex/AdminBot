package main

import (
	"adminbot/cmd"
	"adminbot/config"

	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

func main() {
/*
TODO:
	- Merge progress / prolab
	- add writeup
	- Add new roles when connecting
	- remove user from users.json when he leaves discord (it breaks user.mention())

	- see https://github.com/bwmarrin/discordgo/wiki/FAQ#sending-embeds 
*/	
	
	// Create json files if doesn't exist
	files := []string{"users", "challs", "boxes", "progress", "ippsec"}
	for _, f := range files{
		_, err := ioutil.ReadFile(f + ".json")
    	if err != nil{
        	ioutil.WriteFile(f + ".json", nil, 0644)
    	}	
	}

	// Discord Bot
	bot, err := discordgo.New("Bot " + config.Discord.Token)
	if err != nil {
		panic(err.Error())
	}

	// Register handlers
	bot.AddHandler(cmd.Ready)
	bot.AddHandler(cmd.CommandHandler)
	bot.AddHandler(cmd.ReactionsHandler)
	
	// Open a websocket connection to Discord and begin listening.
	err = bot.Open()
	defer bot.Close()
	if err != nil {
		fmt.Println("Could not connect to discord", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Discord bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	
	// Cleanly close down the Discord session.
	fmt.Println("Closing connection")
	bot.Close()
}

