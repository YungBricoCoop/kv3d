/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
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
	Short: "Set a key-value pair, key=container name, value=label value",
	Long: `Creates and starts a new Docker container with a given name has 'key' and a label 'value' set to the provided value.

	For example:
	kvd set my-container my-value`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		err := docker.RunContainer(key, value)
		if err != nil {
			log.Fatalf("Error setting key-value pair: %v", err)
		}
		fmt.Printf("Successfully set key '%s' with value '%s'\n", key, value)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
