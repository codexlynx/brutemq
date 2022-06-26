package amqp

import (
	"fmt"
	"github.com/streadway/amqp"
)

type BruteAmqp struct {
	Endpoint string
	User     string
}

func (brute *BruteAmqp) TryPassword(password string) (bool, error) {
	endpoint := fmt.Sprintf("amqp://%s:%s@%s", brute.User, password, brute.Endpoint)
	conn, err := amqp.Dial(endpoint)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return false, err
	}
	defer ch.Close()
	conn.ConnectionState()
	return true, nil
}
