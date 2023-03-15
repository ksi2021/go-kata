package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

//	type Dependence struct {
//		Dependence string
//	}
type Config struct {
	AppName    string `json:"app"`
	Production bool   `json:"is_prod"`
	// Dependences []Dependence
}

func main() {
	conf := flag.String("conf", "", "a string")
	flag.Parse()
	if *conf == "" {
		panic(errors.New("Бамм"))
	}

	file, er := os.ReadFile(*conf)
	if er != nil {
		panic(errors.New("Бумм"))
	}

	var config Config
	err := json.Unmarshal(file, &config)
	if err != nil {
		panic(errors.New("Бимм"))
	}

	fmt.Printf("%#v\n", config)

}
