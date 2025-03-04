package msg

import "fmt"

var _ MSG_TemplateInterface = &MSG_Template{}

type MSG_Template struct {
	Welcome                string
	WatchingChanges        string
	ChangeDetectedIn       string
	ExecCommand            string
	MsgError               string
	ContentOfFileChange    string
	ServerRunningAt        string
	SusscessfulCompilation string
	OutPutCommand          string
}

func (m *MSG_Template) GetWelcome() string {
	return m.Welcome
}

func (m *MSG_Template) GetWatchingChanges(folder string) string {
	return fmt.Sprintf("%s %s", m.WatchingChanges, folder)
}

func (m *MSG_Template) GetChangeDetectedIn(fileOrDir string) string {
	return fmt.Sprintf("%s %s", m.ChangeDetectedIn, fileOrDir)
}

func (m *MSG_Template) GetExecCommand(execCmd string) string {
	return fmt.Sprintf("%s %s", m.ExecCommand, execCmd)
}

func (m *MSG_Template) GetMsgError(cmdErr string) string {
	return fmt.Sprintf("%s %s", m.MsgError, cmdErr)
}

func (m *MSG_Template) GetContentOfFileChange() string {
	return m.ContentOfFileChange
}

func (m *MSG_Template) GetServerRunningAt() string {
	return m.ServerRunningAt
}

func (m *MSG_Template) GetSusscessfulCompilation(port string) string {
	return fmt.Sprintf("%s %s", m.SusscessfulCompilation, port)
}

func (m *MSG_Template) GetOutPutCommand(cmd string) string {
	return fmt.Sprintf("%s %s", m.OutPutCommand, cmd)
}
