package main

import (
	"flag"
	"log"

	"github.com/tmluthfiana/simplechat/client/clientside"
	"github.com/tmluthfiana/simplechat/tui"
)

func main() {
	address := flag.String("server", "localhost:3333", "Which server to connect to")

	flag.Parse()

	client := clientside.NewClient()
	err := client.Dial(*address)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// start the client to listen for incoming message
	go client.Start()

	tui.StartUi(client)
}
