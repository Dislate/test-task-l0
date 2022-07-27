package main

import (
	natsStreaming "github.com/Dislate/test-task-l0/nats-streaming"
)

func main() {
	conn, _ := natsStreaming.NatsConnect("subcriber-1")
	defer conn.Close()

	natsStreaming.Subscribe(conn, "orders")

	select {}
}
