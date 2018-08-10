package cmd

import (
	"github.com/revan730/sephiroth-tools/src"
	"github.com/spf13/cobra"
)

var CreateAssetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "Generate assets directory with example assets",
	Run: func(cmd *cobra.Command, args []string) {
		src.CreateAssets()
	},
}
