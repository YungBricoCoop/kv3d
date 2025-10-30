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

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Gets the value for a given key from a labeled container",
	Long: `Gets the value associated with a key by inspecting the Docker container.
The key is used as the container name, and the value is retrieved from the container's label.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		value, err := docker.GetContainerLabelValue(key)
		if err != nil {
			log.Fatalf("Error getting value for key '%s': %v", key, err)
		}
		fmt.Printf("%s\n", value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
