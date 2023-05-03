package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	engine "github.com/mapepema/yolobeer"
	toml "github.com/pelletier/go-toml"
)

var (
	confFile = flag.String("cfg", "conf.toml", "Path to the TOML configuration file")
)

func main() {

	cfgBytes, err := os.ReadFile(*confFile)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}

	var conf engine.Configuration
	err = toml.Unmarshal(cfgBytes, &conf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(conf)

}
