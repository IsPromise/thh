package console

import (
	"github.com/leancodebox/goose/fileopt"
	"github.com/spf13/cobra"
	"thh/app/bundles/bootstrap"
	"thh/app/bundles/kernel"
	"thh/app/console/cmd"
	"thh/app/console/cmd/demo"
	"thh/app/console/cmd/p2p"
	"thh/app/console/cmd/spidercmd"
	"thh/app/console/codemake"
	"thh/app/console/shadow"
	"thh/app/migration"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "thh",
	Short: "A brief description of your application",
	Long:  `thh`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !fileopt.IsExist("config.toml") {
			err := fileopt.Put([]byte(kernel.GetOConfig()), "./config.toml")
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
	rootCmd.AddCommand(CmdServe, scheduleAction)
	rootCmd.AddCommand(codemake.GetCommands()...)
	rootCmd.AddCommand(demo.GetCommands()...)
	rootCmd.AddCommand(cmd.GetCommands()...)
	rootCmd.AddCommand(shadow.GetCommands()...)
	rootCmd.AddCommand(spidercmd.GetCommands()...)
	rootCmd.AddCommand(p2p.GetCommands()...)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
