package redis

import (
	"github.com/oamazing/loa/utils"
	"github.com/spf13/cobra"
)

var RedisCmd = &cobra.Command{
	Use:   "redis",
	Short: "redis cli",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.Run(`redis-cli`); err != nil {
			panic(err)
		}
	},
}
