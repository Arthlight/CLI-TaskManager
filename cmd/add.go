package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "task add",
	Short: "add a task to your list",
	Long:  "If you come up with a task to do, add it to your list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}