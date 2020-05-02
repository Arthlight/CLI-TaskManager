package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "task do",
	Short: "mark a task as done",
	Long:  "If you have completed a task, you can check it off",
	Run: func(cmd *cobra.Command, args []string) {

	},
}