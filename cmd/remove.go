package cmd

import (
	"github.com/CodyBense/todo/cmd/mySql"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
    Use: "remove",
    Short: "Remove an item to the todo list",
    Long: "",

    Run: remove,
}

func init() {
    rootCmd.AddCommand(removeCmd)

    removeCmd.Flags().StringP("taskFlag", "t", "", "What task needs to be removed")
    removeCmd.Flags().StringP("statusFlag", "s", "", "The status of the task")
}

func remove(cmd *cobra.Command, args []string) {

    // Handles task flag parsing
    taskFlag, _ := cmd.Flags().GetString("taskFlag")
    statusFlag, _ := cmd.Flags().GetString("statusFlag")

    mySql.Remove(&taskFlag, &statusFlag)
}
