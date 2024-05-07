package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/akamensky/argparse"
)

type Env struct {
	Port       string
	BasePath   string
	OutputPath string
}

func main() {
	parser := argparse.NewParser("server", "arsync server")
	port := parser.String("l", "port", &argparse.Options{Required: false, Help: "Port to listen on", Default: "8080"})
	basePath := parser.String("p", "path", &argparse.Options{Required: true, Help: "Base path to serve", Default: "/tmp"})
	outputPath := parser.String("o", "output", &argparse.Options{Required: true, Help: "Path to your FTP directory", Default: "/tmp/output"})

	err := parser.Parse(os.Args)

	env := Env{
		Port:       *port,
		BasePath:   *basePath,
		OutputPath: *outputPath,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", env.Port))

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	defer listener.Close()

	log.Printf("Server listening on port %s", env.Port)
	SetupLogging()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go HandleRequest(conn, env)
	}
}
