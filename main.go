// package main

// import (
// 	"flag"
// 	"fmt"
// )

// type Color string

// const (
// 	ColorBlack  Color = "\u001b[30m"
// 	ColorRed    Color = "\u001b[31m"
// 	ColorGreen  Color = "\u001b[32m"
// 	ColorYellow Color = "\u001b[33m"
// 	ColorBlue   Color = "\u001b[34m"
// 	ColorReset  Color = "\u001b[0m"
// )

// func colorize(color Color, message string) {
// 	fmt.Println(string(color), message, string(ColorReset))
// }

// func main() {
// 	useColor := flag.Bool("color", false, "display colorized output")
// 	flag.Parse()

// 	if *useColor {
// 		colorize(ColorBlue, "Hello, DigitalOcean!")
// 		return
// 	}
// 	fmt.Println("Hello, DigitalOcean!")
// }

package main

import (
	"fmt"
)

func main() {
	chn1 := make(chan int)
	chn2 := make(chan int)

	go func() {
		for {

			select {
			case value := <-chn1:
				fmt.Println(value)
			default:
				chn2 <- 11

			}

		}
	}()

	go func() {
		for {
			select {
			case value := <-chn2:
				fmt.Println(value)
			default:
				chn1 <- 22

			}

		}
	}()
	// time.Sleep(3 * time.Second)
	// fmt.Println("end")

	select {}
}
