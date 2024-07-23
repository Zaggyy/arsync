package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
)

func RecursivelyTarGzipDirectory(source string, output string) error {
	// Create the output file
	outFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(outFile)
	defer gzipWriter.Close()

	// Create a tar writer
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// Walk the source directory
	if err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a new tar header
		header, err := tar.FileInfoHeader(info, path)
		if err != nil {
			return err
		}

		// Set the name of the file to the relative path
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}

		// Write the header to the tar archive
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// If the file is a directory, we're done
		if info.IsDir() {
			return nil
		}

		// Open the file
		fileToTar, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fileToTar.Close()

		// Copy the file to the tar archive
		_, err = io.Copy(tarWriter, fileToTar)
		return err
	}); err != nil {
		return err
	}

	// Change permissions of the tar.gz file to 775
	if err := os.Chmod(output, 0775); err != nil {
		return err
	}

	// Change ownership of the tar.gz file to user zaggyy and group ftp_minecraft
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
