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
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PrepareRequest struct {
	address     string
	folder      string
	username    string
	password    string
	ftpUsername string
	ftpPassword string
}

func ExecutePrepare(request PrepareRequest) {
	folder := strings.TrimSpace(request.folder)

	if len(folder) == 0 {
		flag.Usage()
		sleep()
		os.Exit(1)
	}

	if len(request.ftpUsername) == 0 || len(request.ftpPassword) == 0 {
		FatalLogWithSleep("FTP username and password must be provided")
	}

	// Extract host from address
	host, _, err := net.SplitHostPort(request.address)

	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to extract host from address: %v", err))
	}

	// Create an FTP connection
	ftpAddr := net.JoinHostPort(host, "21")
	ftpConn, err := ftp.Dial(ftpAddr, ftp.DialWithTimeout(time.Second*10))
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to connect to the FTP server: %v", err))
	}
	defer ftpConn.Quit()

	// Login to the FTP server
	err = ftpConn.Login(request.ftpUsername, request.ftpPassword)
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to login to the FTP server: %v", err))
	}
	log.Printf("Successfully logged in to the FTP server")

	// Connect to the Arsync server
	conn, err := grpc.Dial(request.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to connect to Arsync server: %v", err))
	}
	defer conn.Close()

	client := arsync.NewArsyncClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := client.Prepare(ctx, &arsync.PrepareRequest{
		Path: folder,
		Auth: &arsync.AuthenticatedRequest{
			Username: request.username,
			Password: request.password,
		},
	})
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to prepare folder: %v", err))
	}

	if !response.Success {
		FatalLogWithSleep(fmt.Sprintf("Failed to prepare folder: %s", folder))
	}

	// Download the tar.gz file
	log.Printf("Downloading tar.gz file %s", folder+".tar.gz")
	archiveName := folder + ".tar.gz"

	archiveFile, err := ftpConn.Retr(archiveName)
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to download tar.gz file: %v", err))
	}
	defer archiveFile.Close()

	// Create the file
	file, err := os.Create(archiveName)
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to create file: %v", err))
	}
	defer file.Close()

	// Write the file
	_, err = io.Copy(file, archiveFile)
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to write file: %v", err))
	}

	log.Printf("Successfully downloaded tar.gz file %s", archiveName)
	sleep()
}
