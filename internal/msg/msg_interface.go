package msg

type MSG_TemplateInterface interface {
	GetWelcome( ) string
	GetWatchingChanges(folder string) string
	GetChangeDetectedIn(fileOrDir string) string
	GetExecCommand(execCmd string) string
	GetMsgError(cmdErr string) string
	GetContentOfFileChange() string
	GetServerRunningAt() string
	GetSusscessfulCompilation(port string) string
	GetOutPutCommand(cmd string) string
}



type MSG_Template_LogInterface interface {
	GetWelcome() string
	GetWatchingChanges() string
	GetChangeDetectedIn() string
	GetExecCommand() string
	GetMsgError() string
	GetContentOfFileChange() string
	GetServerRunningAt() string
	GetSusscessfulCompilation() string
	GetOutPutCommand() string
}