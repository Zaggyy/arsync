package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/akamensky/argparse"
	"github.com/jlaffaye/ftp"
)

func fatalErr(format string, args ...interface{}) {
  log.Printf(format, args...)
  time.Sleep(10 * time.Second)
  os.Exit(1)
}

func main() {
	parser := argparse.NewParser("client", "Connects to a server and gets a folder")
	host := parser.String("H", "host", &argparse.Options{Required: true, Help: "Host to connect to"})
	port := parser.String("P", "port", &argparse.Options{Required: true, Help: "Port to connect to"})
	folder := parser.String("f", "folder", &argparse.Options{Required: true, Help: "Folder to get"})
	ftpUsername := parser.String("u", "username", &argparse.Options{Required: true, Help: "FTP username"})
	ftpPassword := parser.String("p", "password", &argparse.Options{Required: true, Help: "FTP password"})

	err := parser.Parse(os.Args)

	if err != nil {
		log.Fatal(parser.Usage(err))
	}

	log.Printf("Connecting to %s:%s and getting %s", *host, *port, *folder)

	addr := net.JoinHostPort(*host, *port)
	conn, err := net.Dial("tcp", addr)

	if err != nil {
    fatalErr("Could not connect to %s: %s", addr, err)
	}

	defer conn.Close()
	log.Printf("Connected to %s", addr)

	// calculate the length of the folder name
	folderLen := len(*folder)

  if folderLen == 0 {
    fatalErr("Folder name is empty")
  }

	if folderLen > 255 {
    fatalErr("Folder name is too long: %d", folderLen)
	}

  // check if the folder name is valid (no slashes)
  for i := 0; i < folderLen; i++ {
    if (*folder)[i] == '/' {
      fatalErr("No slashes allowed in folder name")
    }
  }

	// send the length of the folder name along with the folder name
	_, err = conn.Write([]byte{byte(folderLen)})

	if err != nil {
    fatalErr("Could not send folder name length: %s", err)
	}

	_, err = conn.Write([]byte(*folder))

	if err != nil {
    fatalErr("Could not send folder name: %s", err)
	}

	log.Printf("Sent folder name: %s", *folder)

	successBit, err := conn.Read([]byte{1})

	if err != nil {
    fatalErr("Could not read success bit: %s", err)
	}

	if successBit == 0 {
    fatalErr("Server did not accept folder name: %s", *folder)
	}

	log.Printf("Server accepted folder name: %s", *folder)

	ftpAddr := net.JoinHostPort(*host, "21")
	ftpConn, err := ftp.Dial(ftpAddr)

	if err != nil {
    fatalErr("Could not connect to FTP server: %s", err)
	}

	defer ftpConn.Quit()

	err = ftpConn.Login(*ftpUsername, *ftpPassword)

	if err != nil {
    fatalErr("Could not login to FTP server: %s", err)
	}

	archiveName := *folder + ".zip"

	file, err := ftpConn.Retr(archiveName)

	if err != nil {
    fatalErr("Could not retrieve file: %s", err)
	}

	defer file.Close()

	outputFile, err := os.Create(archiveName)

	if err != nil {
    fatalErr("Could not create file: %s", err)
	}

	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)

	if err != nil {
    fatalErr("Could not write file: %s", err)
	}

	log.Printf("Wrote file: %s", archiveName)
}
