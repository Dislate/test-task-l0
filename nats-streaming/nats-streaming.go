package natsStreaming

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

func NatsConnect(nameClient string) (stan.Conn, error) {
	return stan.Connect("test-cluster", nameClient)
}

func Subscribe(connect stan.Conn, topic string) {

	aw, _ := time.ParseDuration("5s")
	connect.Subscribe(topic, func(msg *stan.Msg) {
		msg.Ack()
		fmt.Printf("Got: %s\n", msg.Data)
	}, stan.SetManualAckMode(), stan.AckWait(aw))

}

func Publish(client string, topic string, msg string) {
	connectObj, _ := NatsConnect(client)

	for i := 1; ; i++ {
		connectObj.Publish(topic, []byte(msg+" "+strconv.Itoa(i)))
		time.Sleep(4 * time.Second)
	}
}
