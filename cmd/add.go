package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to your list",
	Long:  "If you come up with a task to do, add it to your list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fake command for \"add\"")
		fmt.Println("This is the tasks you want to add: ", args)
	},
}