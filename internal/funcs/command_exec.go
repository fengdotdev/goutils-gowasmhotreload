package funcs

import (
	"log"
	"os/exec"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/msg"
)

func CommandExec(command string) error {
	// Execute the command in a shell
	cmd := exec.Command("sh", "-c", command)

	// Execute the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf(msg.MSG.MsgError, "%s\n", err)
		log.Printf(msg.MSG.OutPutCommand, "%s\n", string(output))
		return err
	}

	log.Printf(msg.MSG.OutPutCommand, "%s\n", string(output))
	return nil
}
