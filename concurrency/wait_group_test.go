package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printString(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	go printString("string", &wg)

	wg.Wait()

	_ = w.Close()

	resultBytes, _ := io.ReadAll(r)
	result := string(resultBytes)

	os.Stdout = stdOut

	if !strings.Contains(result, "string") {
		t.Errorf("Error while testing pringString - no valid output")
	}
}
