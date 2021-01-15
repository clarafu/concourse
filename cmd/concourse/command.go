package main

import "github.com/spf13/cobra"

var (
	ConcourseCmd = &cobra.Command{
		Use:   "concourse",
		Short: "c",
		Long:  `TODO`, // XXX: CONOCURSE DESC
	}
)

func Execute() error {
	return ConcourseCmd.Execute()
}

func init() {
	ConcourseCmd.AddCommand(WebCommand)
	// ConcourseCmd.AddCommand(workerCmd)
	// ConcourseCmd.AddCommand(migrateCmd)
	// ConcourseCmd.AddCommand(quickstartCmd)
	// ConcourseCmd.AddCommand(landWorkerCmd)
	// ConcourseCmd.AddCommand(retireWorkerCmd)
	// ConcourseCmd.AddCommand(generateKeyCmd)
}

// type ConcourseCommand struct {
// 	Version func() `short:"v" long:"version" description:"Print the version of Concourse and exit"`

// 	Web     WebCommand              `command:"web"     description:"Run the web UI and build scheduler."`
// 	Worker  workercmd.WorkerCommand `command:"worker"  description:"Run and register a worker."`
// 	Migrate atccmd.Migration        `command:"migrate" description:"Run database migrations."`

// 	Quickstart QuickstartCommand `command:"quickstart" description:"Run both 'web' and 'worker' together, auto-wired. Not recommended for production."`

// LandWorker   land.LandWorkerCommand     `command:"land-worker" description:"Safely drain a worker's assignments for temporary downtime."`
// 	RetireWorker retire.RetireWorkerCommand `command:"retire-worker" description:"Safely remove a worker from the cluster permanently."`

// 	GenerateKey GenerateKeyCommand `command:"generate-key" description:"Generate RSA key for use with Concourse components."`
// }

// func (cmd ConcourseCommand) LessenRequirements(parser *flags.Parser) {
// 	cmd.Quickstart.LessenRequirements(parser.Find("quickstart"))
// 	cmd.Web.LessenRequirements(parser.Find("web"))
// 	cmd.Worker.LessenRequirements("", parser.Find("worker"))
// }
