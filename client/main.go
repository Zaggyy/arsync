package main

import (
	"log"
	"net"
	"os"
)

func main() {
  cmdLineArgs := os.Args

  if len(cmdLineArgs) < 4 {
    log.Fatalf("Usage: %s <host> <port> <folder>", cmdLineArgs[0])
  }

  host := cmdLineArgs[1]
  port := cmdLineArgs[2]
  folder := cmdLineArgs[3]

  log.Printf("Connecting to %s:%s and getting %s", host, port, folder)

  addr := net.JoinHostPort(host, port)
  conn, err := net.Dial("tcp", addr)

  if err != nil {
    log.Fatalf("Could not connect to %s: %s", addr, err)
  }

  defer conn.Close()
  log.Printf("Connected to %s", addr)
 
  // calculate the length of the folder name
  folderLen := len(folder)

  if folderLen > 255 {
    log.Fatalf("Folder name is too long: %s", folder)
  }

  // send the length of the folder name along with the folder name
  _, err = conn.Write([]byte{byte(folderLen)})

  if err != nil {
    log.Fatalf("Could not send folder length: %s", err)
  }

  _, err = conn.Write([]byte(folder))

  if err != nil {
    log.Fatalf("Could not send folder name: %s", err)
  }

  log.Printf("Sent folder name: %s", folder)
}
