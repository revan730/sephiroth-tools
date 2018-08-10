package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/revan730/sephiroth-tools/src"
	"github.com/spf13/cobra"
)

func readKeysFromFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	items := make(map[string]string)
	for scanner.Scan() {
		items[scanner.Text()] = ""
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

var CreateAssetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "Generate assets directory with example assets",
	Run: func(cmd *cobra.Command, args []string) {
		src.CreateAssets()
	},
}

var CreateStringsCmd = &cobra.Command{
	Use:   "strings [flags] [name]",
	Short: "Generate strings resource file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0] + ".yaml"
		if description == "" && keysFile == "" {
			src.CreateStringAsset(fileName, nil)
		} else {
			data := src.StringAsset{
				Description: description,
				Items:       map[string]string{},
			}
			if keysFile != "" {
				items, err := readKeysFromFile(keysFile)
				if err != nil {
					fmt.Printf("Fatal: failed to read keys from file %s\n", keysFile)
					fmt.Println(err)
					return
				}
				data.Items = items
			}
			src.CreateStringAsset(fileName, data)
		}
	},
}
