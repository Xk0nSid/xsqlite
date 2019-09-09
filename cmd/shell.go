package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Shell represents the whole interactive shell
type Shell struct {
	input  io.ReadCloser
	output io.WriteCloser
	line   int
}

// NewShell creates and returns a new shell object reference
func NewShell() *Shell {
	return &Shell{
		os.Stdin,
		os.Stdout,
		1,
	}
}

// Run executes the main REPL for interacting with database
// Returns return status code and error, if any
func (shell *Shell) Run() (int, error) {
	var returnCode int
	var err error

	fmt.Printf("xsqlite %s-%s\n", versionNumber, gitRevNumber)
	fmt.Println("To see list of commands '.help'")
	fmt.Printf("Control+D or '.exit' to quite.\n\n")
	scanner := bufio.NewScanner(shell.input)

	fmt.Printf("xsqlite:%d> ", shell.line)

	for scanner.Scan() {
		line := scanner.Text()

		shell.line++

		if line[0] == '.' {
			returnCode, err = shell.execIntCmd(line)
			if err != nil && returnCode < 0 {
				return returnCode, err
			} else {
				fmt.Fprintf(shell.output, "error: %s\n", err)
			}
		} else {
			fmt.Println("run query")
		}

		fmt.Printf("xsqlite:%d> ", shell.line)
	}

	return 0, nil
}

// execIntCmd executes internal sqlite commands
// returns return status and error if any
func (shell *Shell) execIntCmd(input string) (int, error) {
	cmd := strings.Fields(input)
	switch cmd[0] {
	case ".exit":
		return -1, errors.New("system exit initiated")
	default:
		return 0, errors.New(fmt.Sprintf("unrecognized command '%s'", cmd[0]))
	}
}
