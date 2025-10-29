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
	Short: "Set a key-value pair in a new container.",
	Long: `Creates and starts a new Docker container with a given name and a label 'value' set to the provided value.

	For example:
	kvd set my-container my-value`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		containerName := args[0]
		labelValue := args[1]
		err := docker.RunContainer(containerName, labelValue)
		if err != nil {
			log.Fatalf("Error running container: %v", err)
		}
		fmt.Printf("Container '%s' started with value '%s'\n", containerName, labelValue)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
