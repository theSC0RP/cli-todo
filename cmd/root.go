package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "cli-todo",
	Short: "CLI Todo is a CLI application to manage your tasks",
	Long:  "CLI Todo is a CLI application to manage your tasks",
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
