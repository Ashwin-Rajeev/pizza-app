package cmd

import (
	"pizza-app/internal/api"
	"pizza-app/internal/data"
	"pizza-app/internal/sms"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the api server",
	Long: `pizza-app 'api' starts the server.
	`,
	Run: apiServerRun,
}

func init() {
	RootCmd.AddCommand(apiCmd)
}

func apiServerRun(*cobra.Command, []string) {
	db := mustPrepareDB()
	
	msg := sms.NewMsg()
	app := api.NewApp(
		data.New(db),
		msg,
	)
	api.Serve(app, mustReadPort(), app.InitRouter())
}
