package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
  env := GetEnvironment()
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

    go HandleRequest(conn)
  }
}
