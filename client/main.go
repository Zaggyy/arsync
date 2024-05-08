package main

import (
	"arsync/arsync"
	"context"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address     = flag.String("address", "localhost:1337", "The address of the Arsync server")
	folder      = flag.String("folder", "", "The folder to prepare")
	username    = flag.String("username", "admin", "The username for the Arsync server")
	password    = flag.String("password", "password", "The password for the Arsync server")
	ftpUsername = flag.String("username", "", "The username for the FTP server")
	ftpPassword = flag.String("password", "", "The password for the FTP server")
)

func main() {
	flag.Parse()

	if len(*folder) == 0 {
		flag.Usage()
		return
	}

	if len(*ftpUsername) == 0 || len(*ftpPassword) == 0 {
		log.Fatalf("FTP username and password must be provided")
	}

	if len(*username) == 0 || len(*password) == 0 {
		log.Fatalf("Username and password must be provided")
	}

	// Create an FTP connection
	ftpAddr := net.JoinHostPort(*address, "21")
	ftpConn, err := ftp.Dial(ftpAddr, ftp.DialWithTimeout(time.Second*10))
	if err != nil {
		log.Fatalf("Failed to connect to FTP server: %v", err)
	}
	defer ftpConn.Quit()

	// Login to the FTP server
	err = ftpConn.Login(*ftpUsername, *ftpPassword)
	if err != nil {
		log.Fatalf("Failed to login to FTP server: %v", err)
	}

	// Connect to the Arsync server
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Arsync server: %v", err)
	}
	defer conn.Close()

	client := arsync.NewArsyncClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := client.Prepare(ctx, &arsync.PrepareRequest{Path: *folder, Username: *username, Password: *password})
	if err != nil {
		log.Fatalf("Failed to prepare folder: %v", err)
	}

	if !response.Success {
		log.Fatalf("Failed to prepare folder: %s", *folder)
	}

	// Download the zip file
	log.Printf("Downloading zip file %s", *folder+".zip")
	archiveName := *folder + ".zip"

	archiveFile, err := ftpConn.Retr(archiveName)
	if err != nil {
		log.Fatalf("Failed to download zip file: %v", err)
	}
	defer archiveFile.Close()

	// Create the file
	file, err := os.Create(archiveName)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Copy the file
	_, err = io.Copy(file, archiveFile)
	if err != nil {
		log.Fatalf("Failed to copy file: %v", err)
	}

	log.Printf("Successfully downloaded zip file %s", archiveName)
}
