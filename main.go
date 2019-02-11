package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = runInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	cmd := exec.Command(input)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
