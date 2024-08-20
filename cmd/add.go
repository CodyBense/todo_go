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

    addCmd.Flags().StringP("taskFlag", "t", "", "Task to be added")
    addCmd.Flags().StringP("descriptionFlag", "d", "", "Description of the task")
    addCmd.Flags().StringP("statusFlag", "s", "", "What table to add it to, ie status")
}

func add(cmd *cobra.Command, args []string) {

    // Handles task flag parsing
    taskFlag, _ := cmd.Flags().GetString("taskFlag")
    descriptionFlag, _ := cmd.Flags().GetString("descriptionFlag")
    statusFlag, _ := cmd.Flags().GetString("statusFlag")
    
    mySql.Add(&taskFlag, &descriptionFlag, &statusFlag)
}
