package irc

import (
  "bufio"
  "fmt"
  "net"
  // "strings"
)

type Bot struct {
  host      string
  port      string
  nick      string
  pass      string
  channel   string
  conn      net.Conn
  reader    *bufio.Reader
  commands  map[string]Command
}

func IRCBot() *Bot {
  bot := Bot {
    host:     "irc.twitch.tv",
    port:     "6667",
    nick:     "MightyBoosh18",
    pass:     "oauth:r981puxzad198oudmbsi7qr0fc0pf0",
    channel:  "#poujakar18",
    conn:     nil,
    reader:   nil,
    commands: nil,
  }
  bot.ImportJson("commands.json")
  return &bot
}

func (bot *Bot) Connect() {
  fmt.Println("Connecting to server...")
  bot.conn, _ = net.Dial("tcp", bot.host + ":" + bot.port)
  fmt.Fprintf(bot.conn, "PASS %s\r\n", bot.pass)
  fmt.Fprintf(bot.conn, "NICK %s\r\n", bot.nick)
  fmt.Fprintf(bot.conn, "USER %s %s %s :%s\r\n", bot.nick, bot.nick, bot.nick, bot.nick)
  fmt.Fprintf(bot.conn, "JOIN %s\r\n", bot.channel)
  fmt.Println("Connected to Server")
  bot.reader = bufio.NewReader(bot.conn)
}

func (bot *Bot) ReadLine() string {
  line, err := bot.reader.ReadString('\n')
  if err != nil {
    fmt.Println("error occurred")
  }
  return line
}
