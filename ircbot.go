package main

import ("net"
        "log"
        "bufio"
        "fmt"
        "net/textproto"
        "strings"
      )
type Bot struct{
        server string
        port string
        nick string
        user string
        channel string
        pass string
        pread, pwrite chan string
        conn net.Conn
}

func NewBot() *Bot {
        return &Bot{server: "irc.freenode.net",
                    port: "6667",
                    nick: "teeheeBot",
                    channel: "#osdg-iiith",
                    pass: "",
                    conn: nil,
                    user: "blaze"}
}
func (bot *Bot) Connect() (conn net.Conn, err error){
  conn, err = net.Dial("tcp",bot.server + ":" + bot.port)
  if err != nil{
    log.Fatal("unable to connect to IRC server ", err)
  }
  bot.conn = conn
  log.Printf("Connected to IRC server %s (%s)\n", bot.server, bot.conn.RemoteAddr())
  return bot.conn, nil
}

func (bot *Bot) WriteMessage(message string){
  fmt.Fprintf(bot.conn, "PRIVMSG %s :%s\r\n", bot.channel, message)
}
func (bot *Bot) Pong(server string){
  fmt.Fprintf(bot.conn, "PONG %s\r\n", server)
  fmt.Printf("PONG %s", server)
}
func (bot *Bot) EvaluateLine(line string){
  splitUp := strings.Split(line, ":")
  
  if len(splitUp) > 2 {
    name := strings.Split(splitUp[1], "!")
    if strings.HasPrefix(splitUp[2], "!teehee ") {
      flags := strings.Split(splitUp[2], " ")
      if flags[1] == "help" {
        bot.WriteMessage(name[0]+": TEEHEEBOT : written in Golang.")
        bot.WriteMessage(name[0]+": Current Functions: help, about")
        bot.WriteMessage(name[0]+": USAGE: '!teehee <function> [flags]'") 
      } else if flags[1] == "about" {
        bot.WriteMessage(name[0]+": Open Source Developers Group @ IIIT - H")
        bot.WriteMessage(name[0]+": Mailing List : https://groups.google.com/forum/?fromgroups#!forum/iiit-osdg")
	bot.WriteMessage(name[0]+": Blog : http://iiitosdg.wordpress.com/")
        bot.WriteMessage(name[0]+": IRC : Well, you guys are already here aren't you :P")
        bot.WriteMessage(name[0]+": GitHub : https://github.com/OSDG-IIITH/")
        bot.WriteMessage(name[0]+": Want to get a project forked under the github group? Register it at http://osdg.iiit.ac.in/github/")
        bot.WriteMessage(name[0]+": Doing GSoC this summer? Check out http://osdg.iiit.ac.in/gsoc15/") 
      } else { 
        bot.WriteMessage(name[0]+": I didn't get that, try '!teehee help' ??")
      }
      // INLINE GUIDES BE THE SHIZZ
      // Adding functionality, use the template suggested below.
      // else if flags[1] == "foobar" {
      // Do some stuff
      // }
    } else if strings.Contains(splitUp[2], "teeheeBot") {
       if strings.Contains(splitUp[1], "PRIVMSG") {
         bot.WriteMessage(name[0]+": Use '!teehee help' to get help!");
        }
      }
  }
  if strings.HasPrefix(splitUp[0], "PING") {
    bot.Pong(splitUp[1])
  }
  fmt.Printf("%q\n", splitUp)
}

func main(){
  ircbot := NewBot()
  conn, _ := ircbot.Connect()
  
  fmt.Fprintf(conn, "USER %s 8 * :%s\r\n", ircbot.nick, ircbot.nick)
  fmt.Fprintf(conn, "NICK %s\r\n", ircbot.nick)
  fmt.Fprintf(conn, "JOIN %s\r\n", ircbot.channel) 
  defer conn.Close()
  
  reader := bufio.NewReader(conn)
  tp := textproto.NewReader( reader )
  for {
        line, err := tp.ReadLine()
        if err != nil {
            break // break loop on errors    
        }
        ircbot.EvaluateLine(line)
    }
}

// Sorry I got to this so late, I've explained most things to you :P
// Genius is a bit much, don't you think :P
