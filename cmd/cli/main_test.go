package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestIntegrationCLIByStdin(t *testing.T) {
	value := `5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`

	expected := `1 3 N
5 1 E
`

	sub := exec.Command("go", "run", "./main.go")

	stdin, err := sub.StdinPipe()
	if err != nil {
		t.Error(err)
		return
	}

	go func() {
		defer stdin.Close()

		io.WriteString(stdin, value)
	}()

	output, err := sub.CombinedOutput()
	if err != nil {
		t.Error(err)
		return
	}

	buf := bytes.NewBuffer(output)
	if got, expected := buf.String(), expected; strings.Compare(got, expected) != 0 {
		t.Errorf("unexpected result for integration test, got: %s, expected: %s", got, expected)
		return
	}
}

func TestIntegrationCLIByFileArgs(t *testing.T) {
	value := `5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`

	expected := `1 3 N
5 1 E
`

	f, err := ioutil.TempFile("/tmp", "mars-rover-*.txt")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.Remove(f.Name())

	f.WriteString(value)
	if err := f.Sync(); err != nil {
		t.Error(err)
		return
	}

	sub := exec.Command("go", "run", "./main.go", "--file", f.Name())

	output, err := sub.CombinedOutput()
	if err != nil {
		t.Error(err)
		return
	}

	buf := bytes.NewBuffer(output)
	if got, expected := buf.String(), expected; strings.Compare(got, expected) != 0 {
		t.Errorf("unexpected result for integration test, got: %s, expected: %s", got, expected)
		return
	}
}
