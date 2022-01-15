package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := readTar(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, " %v\n", err)
		os.Exit(1)
	}

	if err := readZip(os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, " %v\n", err)
		os.Exit(1)
	}
}

func readTar(out io.Writer) error {
	file, _ := os.Open("fixture/README.tar")
	tr := tar.NewReader(file)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		fmt.Fprintln(os.Stderr, "Start reading file:", hdr.Name)
		b := make([]byte, 8)
		for {
			_, err := tr.Read(b)
			out.Write(b)
			if err == io.EOF {
				break
			}
		}
	}
	return nil
}

func readZip(out io.Writer) error {
	zr, err := zip.OpenReader("fixture/README.zip")
	if err != nil {
		return err
	}
	defer zr.Close()
	for _, f := range zr.File {
		fmt.Fprintln(os.Stderr, "Start reading file:", f.Name)
		r, err := f.Open()
		if err != nil {
			return err
		}
		defer r.Close()
		b := make([]byte, 8)
		for {
			_, err := r.Read(b)
			out.Write(b)
			if err == io.EOF {
				break
			}
		}
	}
	return nil
}
