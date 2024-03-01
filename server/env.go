package main

import (
	"log"
	"os"
)

type Env struct {
  Port string
  BasePath string
  OutputPath string
} 

func GetEnvironment() Env {
  port := os.Getenv("PORT")

  if port == "" {
    port = "8080"
  }

  basePath := os.Getenv("BASE_PATH")
  if basePath == "" {
    log.Fatalf("BASE_PATH environment variable is required")
  }

  outputPath := os.Getenv("OUTPUT_PATH")
  if outputPath == "" {
    log.Fatalf("OUTPUT_PATH environment variable is required")
  }

  return Env{
    Port: port,
    BasePath: basePath,
    OutputPath: outputPath,
  }
}
