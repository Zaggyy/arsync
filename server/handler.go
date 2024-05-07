package main

import (
	"archive/zip"
	"io"
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
)

type Command struct {
	FilePathLength byte
	FilePath       string
}

func HandleRequest(conn net.Conn, env Env) {
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

	filePath := path.Join(env.BasePath, command.FilePath)

	var shouldCompress bool = false
	_, err = os.Stat(filePath)

	if len(command.FilePath) < 3 || command.FilePathLength < 3 {
		log.Printf("Illegal file path: %s", filePath)
		response = byte(0)
		shouldCompress = false
	}

	if os.IsNotExist(err) {
		log.Printf("Folder %s does not exist", filePath)
		response = byte(0)
		shouldCompress = false
	}

	if path.Clean(filePath) != filePath {
		log.Printf("Illegal file path: %s", filePath)
		response = byte(0)
		shouldCompress = false
	}

	if shouldCompress {
		log.Printf("Compressing %s...", filePath)
		err = zipSource(filePath, path.Join(env.OutputPath, command.FilePath+".zip"))

		if err != nil {
			log.Printf("Failed to compress %s: %v", filePath, err)
			response = byte(0)
		} else {
			log.Printf("Compressed %s", filePath)
		}
	}

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

func zipSource(source, target string) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)

		if err != nil {
			return err
		}

		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(filepath.Dir(source), path)

		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
		}

		headerWriter, err := writer.CreateHeader(header)

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)

		if err != nil {
			return err
		}

		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}
