package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func SetupLogging() {
	os.Mkdir("logs", 0755)

	if _, err := os.Stat("logs/latest.log"); err == nil {
		fStat, _ := os.Stat("logs/latest.log")
		date := fStat.ModTime().Format("2006-01-02")
		os.Rename("logs/latest.log", "logs/"+date+".log")

		original, _ := os.Open("logs/" + date + ".log")
		defer original.Close()

		compressed, _ := os.Create("logs/" + date + ".log.gz")
		defer compressed.Close()

		writer := gzip.NewWriter(compressed)
		defer writer.Close()

		io.Copy(writer, original)

		writer.Flush()

		os.Remove("logs/" + date + ".log")
	}

	os.Create("logs/latest.log")
}

func Log(message string, level string) {
	log.Printf(fmt.Sprintf("%s: %s", level, message))
	f, _ := os.OpenFile("logs/latest.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	log.SetOutput(f)
	log.Printf(fmt.Sprintf("%s: %s", level, message))
	log.SetOutput(os.Stdout)
}
