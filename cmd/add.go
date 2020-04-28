package cmd

import (
	"CLI-TaskManager/database"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to your list",
	Long:  "If you come up with a task to do, add it to your list",
	Run: func(cmd *cobra.Command, args []string) {
		var sb strings.Builder
		for _,s := range args {
			sb.WriteString(s + " ")
		}
		task := strings.TrimRight(sb.String(), "\t \n")
		fmt.Println("in add.go")
		fmt.Println(args)
		database.AddTask(task)
		fmt.Println("back in add.go")
	},
}