package mysql

import (
	"github.com/oamazing/loa/utils"
	"github.com/spf13/cobra"
)

var MysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "mysql command",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.Run(`mysql -uroot -p`); err != nil {
			panic(err)
		}
	},
}
