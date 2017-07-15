package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "math/rand"
    "strconv"
    "time"

    "github.com/bwmarrin/discordgo"
)


func main() {

  botToken, exists := os.LookupEnv("BOT_TOKEN")
  if !exists {
    fmt.Println("No Bot Token set (Expected as environment variable BOT_TOKEN). Exiting.")
    return
  }

  discord, err := discordgo.New("Bot " + botToken)
  if err != nil {
    fmt.Println("Error creating Discord session: ", err)
    return
  }

  discord.AddHandler(messageCreate)

  err = discord.Open()
  if err != nil {
    fmt.Println("Error opening connection to discord: ", err)
    return
  }

  fmt.Println("Dice-bot is now running! Ctrl + c to exit.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  discord.Close()
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == s.State.User.ID {
    return
  }
  if m.Content == "!roll" {
    s.ChannelMessageSend(m.ChannelID, strconv.Itoa(random(1, 20)))
  }
}
