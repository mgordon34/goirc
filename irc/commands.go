package irc

import (
  "fmt"
  "encoding/json"
  "strings"
  "os"
)

type Command struct {
  Name         string
  Message      string
}

func (bot *Bot) HandleCommand(message string) {
  keyword := strings.Split(message, " ")[0]
  fmt.Println("command detected" + keyword)
  if command, ok := bot.commands[keyword]; ok {
    fmt.Println("hey there" + command.Message)
    fmt.Fprintf(bot.conn, "PRIVMSG %s :%s\r\n", bot.channel, command.Message)
    fmt.Println("done")
  }
}

func (bot *Bot) ImportJson(file string) {
  configFile, err := os.Open(file)
  if err != nil {
    fmt.Println("eror occured")
  }

  jsonParser := json.NewDecoder(configFile)
  commands := make(map[string]Command)
  if err = jsonParser.Decode(&commands); err != nil {
    fmt.Println("error occured parsing")
  }
  bot.commands = commands
}
