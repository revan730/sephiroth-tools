package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.AddCommand(CreateAssetsCmd)
}
