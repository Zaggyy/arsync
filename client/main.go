package main

import (
	"arsync/arsync"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const SLEEP_TIME = 5 * time.Second

var (
	address     = flag.String("address", "localhost:1337", "The address of the Arsync server")
	folder      = flag.String("folder", "", "The folder to prepare")
	username    = flag.String("username", "admin", "The username for the Arsync server")
	password    = flag.String("password", "password", "The password for the Arsync server")
	ftpUsername = flag.String("ftp-username", "", "The username for the FTP server")
	ftpPassword = flag.String("ftp-password", "", "The password for the FTP server")
)

func FatalLogWithSleep(message string, sleep time.Duration) {
	log.Printf(message)
	time.Sleep(sleep)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if len(*folder) == 0 {
		flag.Usage()
		time.Sleep(SLEEP_TIME)
		os.Exit(1)
	}

	if len(*ftpUsername) == 0 || len(*ftpPassword) == 0 {
		FatalLogWithSleep("FTP username and password must be provided", SLEEP_TIME)
	}

	if len(*username) == 0 || len(*password) == 0 {
		FatalLogWithSleep("Username and password must be provided", SLEEP_TIME)
	}

	// Create an FTP connection
  // Extract host from address
  host, _, err := net.SplitHostPort(*address)

  if err != nil {
    FatalLogWithSleep(fmt.Sprintf("Failed to extract host from address: %v", err), SLEEP_TIME)
  }

	ftpAddr := net.JoinHostPort(host, "21")
	ftpConn, err := ftp.Dial(ftpAddr, ftp.DialWithTimeout(time.Second*10))
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to connect to the FTP server: %v", err), SLEEP_TIME)
	}
	defer ftpConn.Quit()

	// Login to the FTP server
	err = ftpConn.Login(*ftpUsername, *ftpPassword)
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to login to the FTP server: %v", err), SLEEP_TIME)
	}
	log.Printf("Successfully logged in to the FTP server")

	// Connect to the Arsync server
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to connect to Arsync server: %v", err), SLEEP_TIME)
	}
	defer conn.Close()

	client := arsync.NewArsyncClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := client.Prepare(ctx, &arsync.PrepareRequest{Path: *folder, Username: *username, Password: *password})
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to prepare folder: %v", err), SLEEP_TIME)
	}

	if !response.Success {
		FatalLogWithSleep(fmt.Sprintf("Failed to prepare folder: %s", *folder), SLEEP_TIME)
	}

	// Download the zip file
	log.Printf("Downloading zip file %s", *folder+".zip")
	archiveName := *folder + ".zip"

	archiveFile, err := ftpConn.Retr(archiveName)
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to download zip file: %v", err), SLEEP_TIME)
	}
	defer archiveFile.Close()

	// Create the file
	file, err := os.Create(archiveName)
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to create file: %v", err), SLEEP_TIME)
	}
	defer file.Close()

	// Copy the file
	_, err = io.Copy(file, archiveFile)
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to copy file: %v", err), SLEEP_TIME)
	}

	log.Printf("Successfully downloaded zip file %s", archiveName)
	time.Sleep(SLEEP_TIME)
}
