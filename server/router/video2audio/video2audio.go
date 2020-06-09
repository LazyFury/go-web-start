package video2audio

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/labstack/echo"
)

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/videoToAudio"
	_ = g.Group(baseURL)
	command("hello world")
}

func command(order string) (err error) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", order)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err = cmd.Start()
	if err != nil {
		log.Fatalf("cmd start err,failed with %s \n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd wait failed with %s \n", err)
	}

	if errStderr != nil || errStdout != nil {
		log.Fatalf("failed to capture stdout or stderr \n")
	}

	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\n out:\n%s\nerr:\n%s\n", outStr, errStr)
	return
}
