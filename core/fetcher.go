package core

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

// APIResponse is the response from the API
type APIResponse struct {
	Content string `json:"content"`
	Joke    string `json:"joke"`
}

// AdviceFetcherInterface is the interface for the AdviceFetcher
type AdviceFetcherInterface interface {
	Fetch(chan string, chan error)
}

// AdviceFetcher is the struct for the AdviceFetcher
type AdviceFetcher struct {
	URL string
}

// FetchFromAPI is the method to fetch the advice from the API
func (aF *AdviceFetcher) FetchFromAPI() (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest("GET", aF.URL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("User-Agent", "AdviceFetcher/1.0")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("API request failed")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result APIResponse

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if result.Content != "" {
		return result.Content, nil
	}

	if result.Joke != "" {
		return result.Joke, nil
	}

	return "", errors.New("no valid string value found in response")
}

// Fetch is the method to fetch the advice from the API
func (aF *AdviceFetcher) Fetch(ch chan string, errCh chan error) {
	advice, err := aF.FetchFromAPI()
	if strings.Contains(aF.URL, "joke") {
		prefix := "Joke: "
		advice = prefix + advice
	}
	if strings.Contains(aF.URL, "quot") {
		prefix := "Quote: "
		advice = prefix + advice
	}
	if err != nil {
		errCh <- err
		return
	}
	ch <- advice
}
