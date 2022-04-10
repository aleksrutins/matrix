package log

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/theckman/yacspin"
)

func Info(msg string) {
	fmt.Printf("\x1b[36minfo \x1b[0m%s\n", msg)
}

func Error(msg string) {
	fmt.Printf("\x1b[1;31merr \x1b[0;31m%s\x1b[0m\n", msg)
}

func Build(config, target string, command []string) (string, error) {
	header := fmt.Sprintf(" \x1b[36mbuild \x1b[0;32mconfiguration\x1b[0m %s \x1b[0;32mtarget\x1b[0m %s", config, target)
	s, err := yacspin.New(yacspin.Config{
		Frequency: 100 * time.Millisecond,
		CharSet:   yacspin.CharSets[14],
		Message:   header,
	})
	if err != nil {
		Error(err.Error())
	}
	s.Start()
	defer s.Stop()
	cmd := exec.Command(command[0], command[1:]...)
	if err != nil {
		return "", err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	err = cmd.Run()
	var outBuf []byte
	stdout.Read(outBuf)

	return string(outBuf), err
}
