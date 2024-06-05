package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

type ListRequest struct {
	address  string
	username string
	password string
}

func ExecuteList(request ListRequest) {
	if len(request.username) == 0 || len(request.password) == 0 {
		FatalLogWithSleep("Username and password must be provided")
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
	err = ftpConn.Login(request.username, request.password)
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to login to the FTP server: %v", err))
	}
	log.Printf("Successfully logged in to the FTP server,")

	// List files in the FTP server
	files, err := ftpConn.List("/worlds/")
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to list files: %v", err))
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name)
	}

	fmt.Println(strings.Join(fileNames, ","))
}
