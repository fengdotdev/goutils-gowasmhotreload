package funcs

import "fmt"

func Gorun(mainPath string) error {

	cmd := fmt.Sprintf("go run %s", mainPath)

	return CommandExec(cmd)
}
