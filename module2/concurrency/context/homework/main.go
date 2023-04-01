package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func chanWriter(ch, data chan int, out chan bool) {
Loop:
	for num := range data {
		select {
		case <-out:
			break Loop
		default:
			ch <- num
		}
	}
	close(ch)
}

func joinChannels(chs ...<-chan int) chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(chs))

		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range ch {
					mergedCh <- id
				}
			}(ch, wg)
		}

		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

func generateData() chan int {
	out := make(chan int, 1000)

	go func() {
		defer close(out)
		for {
			select {
			case _, ok := <-out:
				if !ok {
					return
				}
			case out <- rand.Intn(100):
			}
		}
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	out := generateData()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	exit := make(chan bool)

	go chanWriter(a, out, exit)
	go chanWriter(b, out, exit)
	go chanWriter(c, out, exit)

	go func() {
		t := time.Now()
		<-ctx.Done()
		exit <- true
		exit <- true
		exit <- true
		fmt.Printf("КОНТЕКСТ ЗАКОНЧЕН за - %f s", time.Since(t).Seconds())
	}()

	mainChan := joinChannels(a, b, c)

	for num := range mainChan {
		fmt.Println(num)
	}

	time.Sleep(1 * time.Second)
}
