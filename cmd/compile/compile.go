package compile

import (
	"log"

	"github.com/oamazing/loa/utils"
	"github.com/spf13/cobra"
)

var CompileCmd = &cobra.Command{
	Use:   `compile`,
	Short: `compile app`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start test unit")
		if err := utils.Run(`go test ./...`); err != nil {
			panic(err)
		}
		log.Println("start build app")
		if err := utils.Run(`go build`); err != nil {
			panic(err)
		}
	},
}
