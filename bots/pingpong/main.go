package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	discordBotTokenKey = "DISCORD_BOT_TOKEN"
)

var pingpong map[string]string

func printIfError(tag string, f func() error) {
	if err := f(); err != nil {
		println(tag, err.Error())
	}
}

func main() {
	token := os.Getenv(discordBotTokenKey)
	if len(token) == 0 {
		println("ENV", discordBotTokenKey, "is required")
		os.Exit(1)
	}

	f, err := os.Open("./pingpong.json")
	if err != nil {
		println("pingpong.json:", err.Error())
		os.Exit(1)
	}
	defer printIfError("f.Close", f.Close)

	if err := json.NewDecoder(f).Decode(&pingpong); err != nil {
		println("json decode:", err.Error())
		os.Exit(1)
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		println("discordgo.New:", err.Error())
		os.Exit(1)
	}

	dg.AddHandler(doPingPong(pingpong))

	err = dg.Open()
	if err != nil {
		println("dg.Open:", err.Error())
		os.Exit(1)
	}
	defer printIfError("dg.Close:", dg.Close)

	println("PingPong bot is waiting...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

// doPingPong returns discordgo.messageCreateEventHandleFunc
func doPingPong(pingpong map[string]string) func(s *discordgo.Session, m *discordgo.MessageCreate) {

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// ignore bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		for ping, pong := range pingpong {
			if m.Content == ping {
				println("ping:", ping)
				s.ChannelMessageSend(m.ChannelID, pong)
			} else if m.Content == pong {
				println("pong:", pong)
				s.ChannelMessageSend(m.ChannelID, ping)
			}
		}
	}
}
