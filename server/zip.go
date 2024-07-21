package main

import (
	"archive/zip"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
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
	if err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
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
	}); err != nil {
		return err
	}

	// Change permissions of the ZIP file to 775
	if err := os.Chmod(output, 0775); err != nil {
		return err
	}

	// Change ownership of the ZIP file to user zaggyy and group ftp_minecraft
	uid, gid, err := getUIDGID("zaggyy", "ftp_minecraft")
	if err != nil {
		return err
	}
	if err := os.Chown(output, uid, gid); err != nil {
		return err
	}

	return nil
}

// getUIDGID returns the user ID and group ID for a given username and group name
func getUIDGID(userName, groupName string) (int, int, error) {
	u, err := user.Lookup(userName)
	if err != nil {
		return -1, -1, err
	}
	g, err := user.LookupGroup(groupName)
	if err != nil {
		return -1, -1, err
	}
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return -1, -1, err
	}
	gid, err := strconv.Atoi(g.Gid)
	if err != nil {
		return -1, -1, err
	}
	return uid, gid, nil
}
