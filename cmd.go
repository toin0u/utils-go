package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

// Collected struct
type Collected struct {
	W io.Writer
	S string
}

// SSHCommand struct
type SSHCommand struct {
	user, host string
}

// NewSSHCommand returns a new SSH command
func NewSSHCommand(user, host string) *SSHCommand {
	return &SSHCommand{user, host}
}

// Worker retuns the host name
func (s SSHCommand) Worker() string {
	return s.host
}

// TimeTrack displays how long it took to execute the command
func (s SSHCommand) TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf(color.GreenString(fmt.Sprintf("\nCommand executed in %s\n", elapsed)))
}

// Command execute a SSH command with give arguments
func (s SSHCommand) Command(cmd ...string) *exec.Cmd {
	args := append(
		[]string{fmt.Sprintf("%s@%s", s.user, s.host)},
		cmd...,
	)

	return exec.Command("ssh", args...)
}

// Collect creates and sends Collected structs from its parameters to the given channel
// The caller is responsble to pull Collected from the channel
func (s SSHCommand) Collect(o chan<- *Collected, w io.Writer, readCloser io.ReadCloser) {
	scanner := bufio.NewScanner(readCloser)
	// we get a data race here cause by cmd.Wait() which want to close the pipe
	// @see https://github.com/golang/go/issues/9307
	for scanner.Scan() {
		o <- &Collected{w, scanner.Text()}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
