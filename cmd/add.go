package cmd

import (
	"github.com/CodyBense/todo/cmd/mySql"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use: "add",
    Short: "Add an item to the todo list",
    Long: "",

    Run: add,
}

func init() {
    rootCmd.AddCommand(addCmd)

    addCmd.Flags().StringP("taskFlag", "t", "", "What task needs to be added")
}

func add(cmd *cobra.Command, args []string) {

    // Handles task flag parsing
    taskFlag, _ := cmd.Flags().GetString("taskFlag")

    mySql.Add(&taskFlag)
}
