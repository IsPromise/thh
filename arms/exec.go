package arms

import (
	"os/exec"
)

func ExecShell(command string, arg ...string) ([]byte, error) {
	return exec.Command(command, arg...).Output()
}
