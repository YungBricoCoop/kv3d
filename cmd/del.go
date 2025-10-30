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

var delCmd = &cobra.Command{
	Use:   "del [key]",
	Short: "Deletes a key-value pair by removing the labeled container",
	Long: `Deletes a key-value pair by stopping and removing the Docker container.
The key is used as the container name to identify which container to remove.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		err := docker.DeleteContainer(key)
		if err != nil {
			log.Fatalf("Error deleting key '%s': %v", key, err)
		}
		fmt.Printf("Successfully deleted key '%s'\n", key)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
