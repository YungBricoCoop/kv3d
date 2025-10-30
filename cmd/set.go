/*
Copyright © 2025 Elwan Mayencourt <mayencourt@elwan.ch>
*/
package cmd

import (
	"fmt"
	"kvd/internal/docker"
	"log"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Sets a key-value pair using a labeled container",
	Long: `Sets a key-value pair by creating a Docker container.
The key is used as the container name, and the value is stored in a label.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		err := docker.RunContainer(key, value, 1)
		if err != nil {
			log.Fatalf("Error setting key-value pair: %v", err)
		}
		fmt.Printf("Successfully set key '%s' with value '%s'\n", key, value)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
