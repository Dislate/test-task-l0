package main

import natsStreaming "github.com/Dislate/test-task-l0/nats-streaming"

func main() {
	natsStreaming.Publish("publisher-1", "orders", "Hello dealers")
}
