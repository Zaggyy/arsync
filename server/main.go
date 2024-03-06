package main

import (
	"fmt"
	"log"
	"net"

	"github.com/akamensky/argparse"
)

type Env struct {
	Port       string
	BasePath   string
	OutputPath string
}

func main() {
  parser := argparse.NewParser("server", "arsync server")
  port := parser.String("p", "port", &argparse.Options{Required: false, Help: "Port to listen on", Default: "8080"})
  basePath := parser.String("p", "path", &argparse.Options{Required: false, Help: "Base path to serve", Default: "."})
  outputPath := parser.String("o", "output", &argparse.Options{Required: false, Help: "Path to your FTP directory", Default: "."})

  env := Env{
    Port: *port,
    BasePath: *basePath,
    OutputPath: *outputPath,
  }

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", env.Port))

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	defer listener.Close()

	log.Printf("Server listening on port %s", env.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go HandleRequest(conn, env)
	}
}
