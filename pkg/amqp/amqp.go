package amqp

import (
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type BruteAmqpPlain struct {
	Endpoint string
	User     string
}

func (brute *BruteAmqpPlain) TryPassword(password string) (bool, error) {
	endpoint := fmt.Sprintf("amqp://%s:%s@%s", brute.User, password, brute.Endpoint)
	conn, err := amqp.Dial(endpoint)
	if errors.Is(err, amqp.ErrCredentials) {
		return false, nil
	}
	defer conn.Close()
	return true, err
}
