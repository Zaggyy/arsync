package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func RecursivelyZipDirectory(source string, output string) error {
	// Create the output file
	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Create a new zip writer
	writer := zip.NewWriter(outFile)
	defer writer.Close()

	// Walk the source directory
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a new file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set the method to deflate
		header.Method = zip.Deflate
		// Set the name of the file to the relative path
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}

		// Check if the file is a directory
		if info.IsDir() {
			header.Name += "/"
		}

		// Create the file in the zip archive
		file, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		// If the file is a directory, we're done
		if info.IsDir() {
			return nil
		}

		// Open the file
		fileToZip, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		// Copy the file to the zip archive
		_, err = io.Copy(file, fileToZip)
		return err
	})
}
