package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/yamato0211/html-embed-server/server"
)

var (
	flagPort string
)

func init() {
	flag.StringVar(&flagPort, "port", "8080", "port")
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	addr := net.JoinHostPort("", flagPort)
	s := server.New(addr)
	return s.ListenAndServe()
}
