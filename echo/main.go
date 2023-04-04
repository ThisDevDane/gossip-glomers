package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var node *maelstrom.Node

func main() {
	node = maelstrom.NewNode()

	node.Handle("echo", echoHandler)

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}

func echoHandler(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	body["type"] = "echo_ok"

	return node.Reply(msg, body)
}
