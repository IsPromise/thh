package console

import (
	"thh/app/console/cmd"
	"thh/app/console/cmd/demo"
	"thh/app/console/cmd/p2p"
	"thh/app/console/cmd/tspider"
	"thh/app/console/codemake"
	"thh/app/console/one"
	"thh/app/console/shadow"
	"thh/app/migration"
	"thh/bundles/app"
	"thh/bundles/bootstrap"

	"github.com/leancodebox/goose/fileopt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thh",
	Short: "A brief description of your application",
	Long:  `thh`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !fileopt.IsExist("config.toml") {
			err := fileopt.Put([]byte(app.GetOConfig()), "./config.toml")
			if err != nil {
				panic(err)
			}
		}
		fileopt.SetBasePath("storage/")
		bootstrap.Run()
		migration.M()
	},
	// Run: runWeb,
}

func init() {
	rootCmd.AddCommand(CmdServe)
	rootCmd.AddCommand(codemake.GetCommands()...)
	rootCmd.AddCommand(demo.GetCommands()...)
	rootCmd.AddCommand(one.GetCommands()...)
	rootCmd.AddCommand(cmd.GetCommands()...)
	rootCmd.AddCommand(shadow.GetCommands()...)
	rootCmd.AddCommand(tspider.GetCommands()...)
	rootCmd.AddCommand(p2p.GetCommands()...)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
