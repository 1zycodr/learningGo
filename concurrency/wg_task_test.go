package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go updateMessage("Test", &wg)
	wg.Wait()
	require.Equal(t, "Test", msg)
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("test", &wg)
	wg.Wait()

	printMessage()
	_ = w.Close()

	resByte, _ := io.ReadAll(r)

	os.Stdout = stdOut
	_ = w.Close()

	result := string(resByte)

	if !strings.Contains(result, "test") {
		t.Errorf("Error while testing printMessage")
	}
}
