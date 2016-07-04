package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func get_resp_time(url string) string {
	time_start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching: %v", err)
	}
	defer resp.Body.Close()
	return fmt.Sprintf("%s: %s", url, time.Since(time_start).String())
}

func main() {
	var urls = []string{
		"http://www.google.com/",
		"http://golang.org/",
		"http://blog.golang.org/",
	}

	messages := make(chan string)
	for _, url := range urls {
		go func(url string) {
			ticker := time.NewTicker(time.Second * 2)
			go func() {
				for {
					select {
					case <-ticker.C:
						messages <- get_resp_time(url)
					}
				}
			}()
		}(url)
	}

	for msg := range messages {
		fmt.Println(msg)
	}
}
