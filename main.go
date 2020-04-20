package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	typePtr := flag.String("type", "login", "Denotes what type of server to start: login, world, channel")
	configPtr := flag.String("config", "config_login.toml", "config toml file")

	flag.Parse()
	fmt.Println(*typePtr, *configPtr)
	switch *typePtr {
	case "login":
		s := newLoginServer(*configPtr)
		s.run()
	case "world":
		s := newWorldServer(*configPtr)
		s.run()
	case "channel":
		s := newChannelServer(*configPtr)
		s.run()
	default:
		log.Println("Unkown server type:", *typePtr)
	}
}
