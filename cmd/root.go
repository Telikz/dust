// Package cmd implements the root command
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dust",
	Short: "A powder game ECS simulation",
	Long:  `Dust is a powder game built with an Entity Component System architecture`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunTUI()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
