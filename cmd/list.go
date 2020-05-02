package cmd

import (
	db "CLI-TaskManager/database"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "task list",
	Short: "display your task list",
	Long:  "If you have task added to your list you can see them with this command",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vacation? 🏄🏼♂")
		}
		fmt.Println("You have the following tasks:")
		for index, task := range tasks {
			fmt.Printf("%d. %s\n", index + 1, task.Value)
		}


		fmt.Println("You have the following tasks:")
	},
}