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
		panic(errors.New("arg not found"))
	}

	file, er := os.ReadFile(*conf)
	if er != nil {
		panic(er)
	}

	var config Config
	err := json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", config)

}
