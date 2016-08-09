package main

import (
  "fmt"
  "strings"

  irc "github.com/mgordon34/goirc/irc"
)

func main() {
  fmt.Println("Launching irc...")
  bot := irc.IRCBot()
  bot.Connect()

  for {
    line := bot.ReadLine()
    parts := strings.SplitN(line, " ", 4)
    if !(strings.Contains(parts[1], "JOIN") || strings.Contains(parts[1], "PART") || strings.Contains(parts[1], "QUIT")) {
      sender := strings.Split(parts[0], "!")[0][1:]
      message := parts[3][1:]
      fmt.Print(sender + ": " + message)
      if strings.HasPrefix(message, "!") {
        bot.HandleCommand(strings.TrimSpace(message[1:]))
      }
    }
  }
}
