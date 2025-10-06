package main

import (
	"log"

	"github.com/sdpong/chatlogold/cmd/chatlog"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	chatlog.Execute()
}
