package bruteforcer

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type BruteForcer struct {
	ConcurrentLimit int
	PasswordsReader io.Reader
	TryFunc         func(input string) (bool, error)
	Metadata        string
}

const version = "0.2.1"

func SendWebhook(data interface{}) {
	webhookUrl, ok := os.LookupEnv("WEBHOOK_URL")
	if ok {
		serializedData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		_, err = http.Post(webhookUrl, "application/json", bytes.NewBuffer(serializedData))
		if err != nil {
			log.Println(err)
		}
	}
}

func Try(forcer *BruteForcer, wg *sync.WaitGroup, sem chan int, password string, ctxCancel context.CancelFunc) {
	defer wg.Done()
	ok, err := forcer.TryFunc(password)
	if ok {
		log.Println("Password", password, "found!")
		SendWebhook(
			map[string]string{
				"metadata": forcer.Metadata,
				"password": password,
			},
		)
		ctxCancel()
	}
	if err != nil {
		log.Println(err)
	}
	<-sem
}

func (forcer *BruteForcer) Start() {
	log.Println("Attacking...")

	scanner := bufio.NewScanner(forcer.PasswordsReader)

	var sem = make(chan int, forcer.ConcurrentLimit)
	var wg sync.WaitGroup
	ctx, ctxCancel := context.WithCancel(context.Background())

ScannerLoop:
	for scanner.Scan() {
		select {
		default:
			sem <- 1
			wg.Add(1)
			go Try(forcer, &wg, sem, scanner.Text(), ctxCancel)

		case <-ctx.Done():
			break ScannerLoop
		}
	}

	wg.Wait()
	ctxCancel()
	close(sem)
}

func NewBruterforcerFile(tryFunc func(input string) (bool, error), limit int, metadata string, filepath string) BruteForcer {
	log.Println("brutemq", version, "/ An exotic service bruteforce tool")
	passwords, err := os.Open(filepath)
	if err != nil {
		log.Fatalln(err)
	}

	brute := BruteForcer{
		ConcurrentLimit: limit,
		PasswordsReader: passwords,
		TryFunc:         tryFunc,
		Metadata:        metadata,
	}

	return brute
}
