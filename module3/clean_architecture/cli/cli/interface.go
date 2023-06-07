package cli

import "os"

type CLI interface {
	ViewTasks()
	Exit()
	AddTask()
	DeleteTask()
	EditTask()
	FinishTask()
}

type Cli struct {
}

func (cli *Cli) Exit() {
	os.Exit(0)
}
