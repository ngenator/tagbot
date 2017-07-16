package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "strings"
    "net/http"
    "net/url"
    "regexp"
    "io/ioutil"

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

  fmt.Println("TagBot is now running! Ctrl + c to exit.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  discord.Close()
}

func safeCommand(command string) string {
  regex, _ := regexp.Compile("[^a-zA-Z]+")
  return regex.ReplaceAllString(command, "")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == s.State.User.ID {
    return
  }
  if m.Content == "!roll" {
    return
  }

  if strings.Contains(m.Content, "!") {
    line := strings.SplitN(m.Content, " ", 2)
    command := safeCommand(line[0][1:])
    args := line[1]

    s.ChannelMessageSend(m.ChannelID, "Dynamic command recieved: [" + command + "] with args [" + args +"]")

    resp, err := http.Get("http://" + command + "/execute?args=" + url.QueryEscape(args))
    if err != nil {
      s.ChannelMessageSend(m.ChannelID, "Error: ["+ err.Error() + "]")
      return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    s.ChannelMessageSend(m.ChannelID, "Executed: " + command + ". Got: " + string(body))
  }
}
