package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

const (
	jpegStr string = "jpeg"
	pngStr  string = "png"
	gifStr  string = "gif"
)

func main() {
	fmtStr := parseFlag()
	if err := convert(os.Stdin, os.Stdout, fmtStr); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fmtStr, err)
		os.Exit(1)
	}
}

func parseFlag() string {
	fmtPtr := flag.String("format", "jpeg", "Output format: jpeg, png, or gif.")
	flag.Parse()
	switch *fmtPtr {
	case jpegStr:
		return jpegStr
	case pngStr:
		return pngStr
	case gifStr:
		return gifStr
	default:
		log.Fatalf("invalid format: %s", *fmtPtr)
		return ""
	}
}

func convert(in io.Reader, out io.Writer, fmtStr string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch fmtStr {
	case jpegStr:
		jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case pngStr:
		png.Encode(out, img)
	case gifStr:
		gif.Encode(out, img, &gif.Options{NumColors: 256})
	default:
		return fmt.Errorf("unsupported format: %s", fmtStr)
	}
	return nil
}
