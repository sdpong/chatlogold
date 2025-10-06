package main

import (
	"log"

	"github.com/sdpong/chatlog/cmd/chatlogold"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	chatlog.Execute()
}
