package bruteforcer

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"sync"
)

type BruteForcer struct {
	ConcurrentLimit int
	PasswordsReader io.Reader
	TryFunc         func(input string) (bool, error)
}

func Try(forcer *BruteForcer, wg *sync.WaitGroup, sem chan int, password string, ctxCancel context.CancelFunc) {
	defer wg.Done()
	ok, _ := forcer.TryFunc(password)
	if ok {
		log.Println("Password", password, "found!")
		ctxCancel()
	}
	<-sem
}

func (forcer *BruteForcer) Start() {
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

func StartBruteforcerWithFile(tryFunc func(input string) (bool, error), limit int, filepath string) {
	passwords, err := os.Open(filepath)
	if err != nil {
		log.Fatalln(err)
	}

	brute := BruteForcer{
		ConcurrentLimit: limit,
		PasswordsReader: passwords,
		TryFunc:         tryFunc,
	}

	brute.Start()
}
