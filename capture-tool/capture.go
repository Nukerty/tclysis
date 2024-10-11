package main

import (
	"bufio"
	"log"
	"net"
	"net/textproto"
	"os"
	"strings"
	"time"
)

const (
	server  = "irc.chat.twitch.tv:6667"
	channel = "#xqc" // Replace with the channel you want to capture
)

func main() {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal("Encountered a connection error : ", err)
	}
	defer conn.Close()

	file, err := os.Create("twitch_chat.log")
	if err != nil {
		log.Fatal("Could not create file : ", err)
	}
	defer file.Close()

	conn.Write([]byte("NICK justinfan1338\r\n"))
	conn.Write([]byte("JOIN " + channel + "\r\n"))

	reader := textproto.NewReader(bufio.NewReader(conn))
	writer := bufio.NewWriter(file)

	for {
		line, err := reader.ReadLine()
		if err != nil {
			log.Fatal("Could not capture : ", err)
			break;
		}
		
		username, after, found := strings.Cut(line, "!")
		if found == false {
			continue
		}
		username = username[1:]
		_, msg , _ := strings.Cut(after, ":")
		
		writer.WriteString(time.Now().String() + ":" + username + ":" + msg + "\n")
		writer.Flush()
	}
}
