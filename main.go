package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		// homedir := user.HomeDir
		username := user.Username

		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		fmt.Print(username + "@" + dir + " ~: ")
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
	args := strings.Split(input, " ")

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	homedir := user.HomeDir
	// username := user.Username

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return os.Chdir(homedir)
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
