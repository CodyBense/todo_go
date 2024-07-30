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

    updateCmd.Flags().IntP("idFlag", "i", 0, "What task needs to be updated")
}

func update(cmd *cobra.Command, args []string) {

    // Handles task flag parsing
    idFlag, _ := cmd.Flags().GetInt("idFlag")

    // Conducts the update
    mySql.Update(&idFlag)
}
