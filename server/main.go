package main

import (
	"arsync/arsync"
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"path"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var (
	port       = flag.Int("port", 1337, "The port the server will listen on")
	basePath   = flag.String("base-path", "/tmp", "The base path to search for folders")
	outputPath = flag.String("output-path", "/tmp", "The path to output the zip files")
	username   = flag.String("username", "admin", "The username for the Arsync server")
	password   = flag.String("password", "password", "The password for the Arsync server")
)

type Server struct {
	arsync.UnimplementedArsyncServer
}

func (s *Server) Prepare(ctx context.Context, in *arsync.PrepareRequest) (*arsync.PrepareResponse, error) {
	// Get the IP address of the client
	p, _ := peer.FromContext(ctx)
	ip := p.Addr.String()
	// Check if the username and password are correct
	if in.Auth.Username != *username || in.Auth.Password != *password {
		Log("Invalid username and/or password", "ERROR", ip)
		return &arsync.PrepareResponse{Success: false}, errors.New("Invalid username and/or password")
	}

	// Clean up the path
	cleanPath := strings.TrimSpace(in.Path)

	Log(fmt.Sprintf("Received request to prepare %s", cleanPath), "INFO", ip)
	preparePath := path.Join(*basePath, cleanPath)
	Log(fmt.Sprintf("Calculated path: %s", preparePath), "INFO", ip)

	// Check if the path exists
	if _, err := os.Stat(preparePath); os.IsNotExist(err) {
		Log(fmt.Sprintf("Path %s does not exist", preparePath), "ERROR", ip)
		return &arsync.PrepareResponse{Success: false}, errors.New(fmt.Sprintf("Path %s does not exist", preparePath))
	}

	// Check if the path contains any "tricky" characters such as "..", "~", "/", etc.
	if path.Clean(preparePath) != preparePath {
		Log(fmt.Sprintf("Path %s contains tricky characters", preparePath), "ERROR", ip)
		return &arsync.PrepareResponse{Success: false}, errors.New(fmt.Sprintf("Path %s contains tricky characters", preparePath))
	}

	// Check if the path has less than 3 characters
	if len(preparePath) < 3 {
		Log(fmt.Sprintf("Path %s is too short", preparePath), "ERROR", ip)
		return &arsync.PrepareResponse{Success: false}, errors.New(fmt.Sprintf("Path %s is too short", preparePath))
	}

	// Check if the path is in a subdirectory of the base path, i.e. contains "/"
	if path.Dir(preparePath) != *basePath {
		Log(fmt.Sprintf("Path %s is not in the base path", preparePath), "ERROR", ip)
		return &arsync.PrepareResponse{Success: false}, errors.New(fmt.Sprintf("Path %s is not in the base path", preparePath))
	}

	// Check if the path is a directory
	if fileInfo, err := os.Stat(preparePath); err == nil && !fileInfo.IsDir() {
		Log(fmt.Sprintf("Path %s is not a directory", preparePath), "ERROR", ip)
		return &arsync.PrepareResponse{Success: false}, errors.New(fmt.Sprintf("Path %s is not a directory", preparePath))
	}

	// Recursively prepare the folder
	Log(fmt.Sprintf("Preparing folder %s", preparePath), "INFO", ip)
	err := RecursivelyZipDirectory(preparePath, path.Join(*outputPath, cleanPath+".zip"))

	if err != nil {
		Log(fmt.Sprintf("Failed to prepare folder %s: %v", preparePath, err), "ERROR", ip)
		return &arsync.PrepareResponse{Success: false}, err
	}

	return &arsync.PrepareResponse{Success: true}, nil
}

func main() {
	flag.Parse()
	PrepareLogs()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		Log(fmt.Sprintf("Failed to start Arsync server: %v", err), "ERROR", "localhost")
	}

	server := grpc.NewServer()
	arsync.RegisterArsyncServer(server, &Server{})
	Log(fmt.Sprintf("Arsync server listening on %v", listener.Addr()), "INFO", "localhost")
	Log(fmt.Sprintf("Configured with base path: %s", *basePath), "INFO", "localhost")
	Log(fmt.Sprintf("Configured with output path: %s", *outputPath), "INFO", "localhost")
	Log(fmt.Sprintf("Configured with username: %s and password: %s", *username, *password), "INFO", "localhost")

	if err := server.Serve(listener); err != nil {
		Log(fmt.Sprintf("Failed to start Arsync server: %v", err), "ERROR", "localhost")
	}
}
