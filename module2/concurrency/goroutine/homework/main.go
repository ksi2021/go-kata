package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// t := time.Now()
	// rand.Seed(t.UnixNano())

	// go parseURL("https://youtube.com/")
	// parseURL("https://example.com/")

	// fmt.Printf("Parsing completed. Time Elapsed : %f seconds \n", time.Since(t).Seconds())

	// message := make(chan string)
	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		message <- fmt.Sprintf("%d", i)
	// 		time.Sleep(time.Millisecond * 150)
	// 	}

	// 	close(message)
	// }()

	// for {
	// 	msg, open := <-message

	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	// message := make(chan string, 2)
	// message <- "hello "
	// message <- "world"

	// fmt.Println(<-message)
	// fmt.Println(<-message)

	// urls := []string{
	// 	"https://youtube.com/",
	// 	"https://google.com/",
	// 	"https://github.com/",
	// 	"https://microsoft.com/",
	// }

	// var wg sync.WaitGroup

	// for _, url := range urls {
	// 	wg.Add(1)

	// 	go func(url string) {
	// 		doHTTP(url)
	// 		wg.Done()
	// 	}(url)
	// }
	// wg.Wait()

	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			message1 <- "Прошло пол секунды"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			message2 <- "Прошло 2 секунды"
		}
	}()

	for {
		// fmt.Println(<-message1)
		// fmt.Println(<-message2)

		select {
		case msg := <-message1:
			fmt.Println(msg)

		case msg := <-message2:
			fmt.Println(msg)

		}
	}

}

func doHTTP(url string) {
	t := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to get <%s> : %s", url, err.Error())
	}

	defer resp.Body.Close()

	fmt.Printf("<%s> - Status Code [%d] - Latency %d ms \n", url, resp.StatusCode, time.Since(t).Milliseconds())
}

func parseURL(url string) {
	for i := 0; i < 5; i++ {
		latency := rand.Intn(500) + 500

		time.Sleep(time.Duration(latency) * time.Millisecond)

		fmt.Printf("Parsing <%s> - Step %d - Latency %d \n", url, i+1, latency)
	}
}
