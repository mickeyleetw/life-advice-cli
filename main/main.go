package main

import (
	"flag"
	"fmt"
	core "life-advice-cli/core"
	"strings"
	"sync"
)

func main() {

	advice := flag.String("advice", "", "Get advice: quote, joke, philosophy, motivation")
	historyFlag := flag.Bool("history", false, "Show history")
	flag.Parse()

	// initialize history
	history := core.NewHistory(10)

	Fetchers := map[string]core.AdviceFetcherInterface{
		"joke":  &core.AdviceFetcher{URL: "https://icanhazdadjoke.com/"},
		"quote": &core.AdviceFetcher{URL: "http://api.quotable.io/random"},
	}

	if *historyFlag {
		records := history.GetRecords()
		fmt.Println(records)
	}

	if *advice != "" {
		adviceTypes := strings.Split(*advice, ",")
		adviceChannel := make(chan string)
		errChannel := make(chan error)
		var selectedFetchers []core.AdviceFetcherInterface
		for _, adviceType := range adviceTypes {
			if fetcher, exists := Fetchers[strings.TrimSpace(adviceType)]; exists {
				selectedFetchers = append(selectedFetchers, fetcher)
			} else {
				fmt.Printf("Warning: unknown advice type '%s'\n", adviceType)
			}
		}

		var wg sync.WaitGroup
		wg.Add(len(selectedFetchers))
		go fetchMultipleAdvice(selectedFetchers, adviceChannel, errChannel, &wg)

		// Create a done channel to signal when all fetchers are complete
		doneChan := make(chan struct{})

		// Start a goroutine to close doneChan when all fetchers finish
		go func() {
			wg.Wait()
			close(doneChan)
		}()

		// Keep receiving messages until all fetchers are done
		for {
			select {
			case advice := <-adviceChannel:
				fmt.Println(advice)
				history.Add(advice)
			case err := <-errChannel:
				fmt.Printf("Error: %v\n", err)
			case <-doneChan:
				return
			}
		}
	}

}

func fetchMultipleAdvice(fetchers []core.AdviceFetcherInterface, ch chan string, errCh chan error, wg *sync.WaitGroup) {
	for _, fetcher := range fetchers {
		go func(f core.AdviceFetcherInterface) {
			defer wg.Done()
			f.Fetch(ch, errCh)
		}(fetcher)
	}
}
