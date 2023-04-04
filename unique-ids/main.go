package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync/atomic"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var node *maelstrom.Node
var accumIdCounter atomic.Int64

func main() {
	accumIdCounter = atomic.Int64{}
	node = maelstrom.NewNode()

	node.Handle("generate", echoHandler)

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}

func echoHandler(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	accumIdCounter.Add(1)
	body["type"] = "generate_ok"
	body["id"] = fmt.Sprintf("%s-%d", node.ID(), accumIdCounter.Load())

	return node.Reply(msg, body)
}
