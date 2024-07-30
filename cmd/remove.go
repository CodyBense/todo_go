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

    removeCmd.Flags().IntP("idFlag", "i", 0, "What task needs to be removed")
}

func remove(cmd *cobra.Command, args []string) {

    // Handles task flag parsing
    idFlag, _ := cmd.Flags().GetInt("idFlag")

    mySql.Remove(&idFlag)
}
