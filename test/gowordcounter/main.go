package main

import (
	"bufio"
	"log"
	"os"
	"sync"
	"sync/atomic"

	"github.com/ahamtat/go-interview/test/gowordcounter/counter"
)

const k = 5

// echo -e 'https://golang.org\n/etc/passwd\nhttps://golang.org\nhttps://golang.org\nGo Go Go' | go run *.go

func main() {
	var total uint64

	// Make goroutines restriction by number
	var wg sync.WaitGroup
	workerCh := make(chan struct{}, k)
	defer close(workerCh)

	// Create line scanner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	// Read lines from stdin
	for scanner.Scan() {

		// Set worker flag
		workerCh <- struct{}{}
		wg.Add(1)

		go func(source string) {
			n := counter.CountGoWords(source)

			// Update total words counter
			atomic.AddUint64(&total, uint64(n))
			log.Printf("Count for %s: %d\n", source, n)

			// Remove worker flag
			<-workerCh
			wg.Done()
		}(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("error reading standard input: %s", err)
	}

	// Wait goroutines to end
	wg.Wait()

	log.Printf("Total: %d", total)
}
