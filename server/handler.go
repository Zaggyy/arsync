package main

import (
	"log"
	"net"
)

type Command struct {
  FilePathLength byte
  FilePath string
}

func HandleRequest(conn net.Conn) {
  defer conn.Close()
  log.Printf("Accepted connection from %s", conn.RemoteAddr())

  command := Command{}
  err := ReadCommand(conn, &command)

  if err != nil {
    log.Printf("Failed to read command: %v", err)
    return
  }

  log.Printf("Received command: %s", command.FilePath)
}

func ReadCommand(conn net.Conn, command *Command) error {
  filePathLength := make([]byte, 1)
  _, err := conn.Read(filePathLength)

  if err != nil {
    log.Printf("Failed to read file path length: %v", err)
    return err
  }

  command.FilePathLength = filePathLength[0]

  filePath := make([]byte, command.FilePathLength)
  _, err = conn.Read(filePath)

  if err != nil {
    log.Printf("Failed to read file path: %v", err)
    return err
  }

  command.FilePath = string(filePath)
  return nil
}
