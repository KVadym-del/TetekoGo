package main

import (
	"TetekoGo/src"
	"flag"
	"os"
)

var (
	clientMode = flag.Bool("client", false, "run in client mode")
	serverMode = flag.Bool("server", false, "run in server mode")

	serverPort = flag.String("port", "8080", "port to listen on")
	serverIP   = flag.String("ip", "localhost", "server ip")
)

func main() {
	flag.Parse()

	if *serverMode {
		src.CreateServer(*serverPort)
	} else if *clientMode {
		serverConfig := src.ServerConfig{IP: *serverIP, Port: *serverPort}
		src.ConnectToServer(serverConfig)
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}

	os.Exit(0)
}
