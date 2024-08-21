package cmd

import (
	"github.com/CodyBense/todo/cmd/mySql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
    Use: "update",
    Short: "Update an item to the todo list",
    Long: "",

    Run: update,
}

func init() {
    rootCmd.AddCommand(updateCmd)

    updateCmd.Flags().StringP("taskFlag", "t", "", "What task needs to be updated")
    updateCmd.Flags().StringP("statusFlag", "s", "", "Currents status of task")
    updateCmd.Flags().StringP("updateFlag", "u", "", "Updates status of task")
}

func update(cmd *cobra.Command, args []string) {

    // Handles task flag parsing
    taskFlag, _ := cmd.Flags().GetString("taskFlag")
    statusFlag, _ := cmd.Flags().GetString("statusFlag")
    updateFlag, _ := cmd.Flags().GetString("updateFlag")

    // Conducts the update
    mySql.Update(&taskFlag, &statusFlag, &updateFlag)
}
