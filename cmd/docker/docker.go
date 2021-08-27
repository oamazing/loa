package docker

import (
	"github.com/oamazing/loa/cmd/docker/log"
	"github.com/spf13/cobra"
)

var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "docker command",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	DockerCmd.AddCommand(log.DockerLogCmd)
}
