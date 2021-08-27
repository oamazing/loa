package log

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DockerLogCmd = &cobra.Command{
	Use:   "log",
	Short: "show docker log",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker log")
	},
}
