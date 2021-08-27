package cmd

import (
	"fmt"
	"os"

	"github.com/oamazing/loa/cmd/compile"
	"github.com/oamazing/loa/cmd/docker"
	"github.com/oamazing/loa/cmd/mysql"
	"github.com/oamazing/loa/cmd/redis"
	"github.com/oamazing/loa/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loa",
	Short: "loa is a test utils",
}

func init() {
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(docker.DockerCmd)
	rootCmd.AddCommand(mysql.MysqlCmd)
	rootCmd.AddCommand(redis.RedisCmd)
	rootCmd.AddCommand(compile.CompileCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
