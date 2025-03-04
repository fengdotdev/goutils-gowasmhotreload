package funcs

import (

	"os/exec"

)

func CommandExec(command string) error {
	// Execute the command in a shell
	cmd := exec.Command("sh", "-c", command)

	// Execute the command and capture the output
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
