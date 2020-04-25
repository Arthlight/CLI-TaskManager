package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "display your task list",
	Long:  "If you have task added to your list you can see them with this command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fake command for \"list\"")
		fmt.Println("This is the tasks you want to add: ", args)
	},
}