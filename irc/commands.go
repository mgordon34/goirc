package irc

import (
  "fmt"
  "encoding/json"
  "strings"
  "os"
)

type Commandable interface {
  IsTriggered(keyword string) bool
  Run(bot *Bot)
}

type Command struct {
  Name         string
  Triggers     []string
}

func (command Command) IsTriggered(keyword string) bool {
  for _, trigger := range command.Triggers {
    if keyword == trigger {
      return true
    }
  }
  return false
}

type FunctionCommand struct {
  Command
  Function func(bot *Bot)
}

func (command FunctionCommand) Run (bot *Bot) {
  command.Function(bot)
}

type TextCommand struct {
  Command
  Message string
}

func (command TextCommand) Run(bot *Bot) {
  fmt.Fprintf(bot.conn, "PRIVMSG %s :%s\r\n", bot.channel, command.Message)
}

func (bot *Bot) HandleCommand(message string) {
  fmt.Println("handling command")
  for _, command := range bot.commands {
    fmt.Println("checking: " + message)
    if command.IsTriggered(strings.Split(message, " ")[0]) {
      command.Run(bot)
    }
  }
}

func (bot *Bot) InitCommands() {
  hey := FunctionCommand {
    Command: Command {
      Name: "hey",
      Triggers: []string {
        "hey",
      },
    },
    Function: bot.hey,
  }
  bot.commands = append(bot.commands, hey)
}

func (bot *Bot) hey(mybot *Bot) {
  fmt.Fprintf(bot.conn, "PRIVMSG %s :hey\r\n", bot.channel)
}

func (bot *Bot) ImportJson(file string) {
  configFile, err := os.Open(file)
  if err != nil {
    fmt.Println("eror occured")
  }

  jsonParser := json.NewDecoder(configFile)
  commands := make([]Commandable, 1)
  if err = jsonParser.Decode(&commands); err != nil {
    fmt.Println("error occured parsing")
  }
  bot.commands = append(bot.commands, commands...)
}
