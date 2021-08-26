package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "loa",
	Short: "a little utils",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Exec() {
	rootCmd.Execute()
}
