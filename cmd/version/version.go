package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = `0.0.1`

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "show loa version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
