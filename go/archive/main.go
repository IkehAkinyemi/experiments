package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	listOfFilesNames := []string{"main.tex", "main.toc", "util/file1.txt", "util/file1.txt"}
	outFile, err := os.Create("mytestarchive.tar.gz")
	if err != nil {
		panic("Error creating archive file")
	}
	defer outFile.Close()
	err = createTarAndGz(listOfFilesNames, outFile)
	if err != nil {
		panic("Error creating archive file.")
	}

	fmt.Println("Archiving and file compression completed.")
}

func createTarAndGz(fileNames []string, buffer io.Writer) error {
	gzipWriter := gzip.NewWriter(buffer)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()
	for _, f := range fileNames {
		err := addToTar(tarWriter, f)
		if err != nil {
			return err
		}
	}
	return nil
}

func addToTar(tarWriter *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	header.Name = filename
	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return err
	}

	return nil
}