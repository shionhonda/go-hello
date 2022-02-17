package bzip

import (
	"io"
	"log"
	"os/exec"
)

type Writer struct {
	cmd   exec.Cmd
	stdin io.WriteCloser
}

func NewWriter(w io.Writer) io.WriteCloser {
	cmd := exec.Cmd{Path: "/bin/bzip2", Stdout: w}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("cmd pipe: %v", err)
	}
	cmd.Start()
	if err != nil {
		log.Fatalf("cmd start: %v", err)
	}
	return &Writer{cmd, stdin}
}

func (w *Writer) Write(data []byte) (int, error) {
	return w.stdin.Write(data)
}

func (w *Writer) Close() error {
	pipeErr := w.stdin.Close()
	cmdErr := w.cmd.Wait()
	if pipeErr != nil {
		return pipeErr
	}
	if cmdErr != nil {
		return cmdErr
	}
	return nil
}
