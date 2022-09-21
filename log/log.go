package log

import (
	"bufio"
	"fmt"
	"os/exec"
)

func Info(msg string) {
	fmt.Printf("\x1b[36minfo \x1b[0m%s\n", msg)
}

func Error(msg string) {
	fmt.Printf("\x1b[1;31merr \x1b[0;31m%s\x1b[0m\n", msg)
}

func Build(config, target string, command []string) error {
	fmt.Printf("\x1b[36mbuild \x1b[0;32mconfiguration\x1b[0m %s \x1b[0;32mtarget\x1b[0m %s\n", config, target)

	cmd := exec.Command(command[0], command[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	reader := bufio.NewReader(stdout)
	defer stdout.Close()
	go func() {
		for {
			str, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Print(str)
		}
	}()

	err = cmd.Run()

	return err
}
