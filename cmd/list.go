package cmd

import (
	"github.com/CodyBense/todo/cmd/mySql"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command {
    Use: "list",
    Short: "List todo items",
    Long: "Shoes a list of the items on the todo list",

    Run: list,
}

func init() {
    rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {

    mySql.List()
}
