package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version string // Must be initialized with ldflags

var (
	logVerbose bool
)

var RootCmd = &cobra.Command{
	Use:   "sephiroth",
	Short: "Command line tools for sephiroth game engine resource management",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Generate asset(s)",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(versionCmd)
	createCmd.AddCommand(CreateAssetsCmd)
}
