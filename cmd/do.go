package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark a task as done",
	Long:  "If you have completed a task, you can check it off",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fake command for \"do\"")
		fmt.Println("This is the tasks you want to add: ", args)

	},
}