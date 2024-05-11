package main

import (
	"flag"
	"log"
	"os"
	"time"
)

const SLEEP_TIME = 5 * time.Second

var (
	address     = flag.String("address", "localhost:1337", "The address of the Arsync server")
	folder      = flag.String("folder", "", "The folder to prepare")
	username    = flag.String("username", "admin", "The username for the Arsync server")
	password    = flag.String("password", "password", "The password for the Arsync server")
	ftpUsername = flag.String("ftp-username", "", "The username for the FTP server")
	ftpPassword = flag.String("ftp-password", "", "The password for the FTP server")
	command     = flag.String("command", "prepare", "The command to execute. Available commands: prepare, list")
)

func FatalLogWithSleep(message string, sleep time.Duration) {
	log.Printf(message)
	time.Sleep(sleep)
	os.Exit(1)
}

func main() {
	flag.Parse()

	switch *command {
	case "prepare":
		ExecutePrepare(PrepareRequest{
			address:     *address,
			folder:      *folder,
			username:    *username,
			password:    *password,
			ftpUsername: *ftpUsername,
			ftpPassword: *ftpPassword,
		})
		break
	case "list":
		ExecuteList(ListRequest{
			address:  *address,
			username: *username,
			password: *password,
		})
		break
	default:
		FatalLogWithSleep("Invalid command", SLEEP_TIME)
	}
}
