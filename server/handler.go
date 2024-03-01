package main

import (
	"log"
	"net"
	"path"

	"github.com/walle/targz"
)

type Command struct {
  FilePathLength byte
  FilePath string
}

func HandleRequest(conn net.Conn) {
  defer conn.Close()
  log.Printf("Accepted connection from %s", conn.RemoteAddr())
  response := byte(1)

  command := Command{}
  err := ReadCommand(conn, &command)

  if err != nil {
    log.Printf("Failed to read command: %v", err)
    response = byte(0)
  }

  log.Printf("Received command: %s", command.FilePath)

  env := GetEnvironment()

  filePath := path.Join(env.BasePath, command.FilePath)

  log.Printf("Compressing %s...", filePath)

  err = targz.Compress(filePath, path.Join(env.OutputPath, command.FilePath + ".tar.gz"))

  if err != nil {
    log.Printf("Failed to compress %s: %v", filePath, err)
    response = byte(0)
  }

  log.Printf("Compressed %s", filePath)

  _, err = conn.Write([]byte{response})

  if err != nil {
    log.Printf("Failed to write response: %v", err)
  }
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
