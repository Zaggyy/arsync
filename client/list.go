package main

import (
	"arsync/arsync"
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"
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

	// Connect to the Arsync server
	conn, err := grpc.Dial(request.address, grpc.WithInsecure())
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to connect to Arsync server: %v", err))
	}
	defer conn.Close()

	client := arsync.NewArsyncClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.List(ctx, &arsync.ListRequest{
		Auth: &arsync.AuthenticatedRequest{
			Username: request.username,
			Password: request.password,
		},
	})
	if err != nil {
		FatalLogWithSleep(fmt.Sprintf("Failed to list files: %v", err))
	}

	fmt.Println(strings.Join(response.Files, ","))
}
