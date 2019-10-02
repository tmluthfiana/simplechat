package main

import (
	"github.com/tmluthfiana/simplechat/server/serverside"
)

func main() {
	var s serverside.ChatServer
	s = serverside.NewServer()
	s.Listen(":3333")

	// start the server
	s.Start()
}
