package cmd

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

    var err error

    // Open Mysql connection
    db, err = sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/todo")
    if err != nil {
        log.Fatalf("impossible to create the connection: %s", err)
    }
    defer db.Close()

    // Test Mysql connection
    pingErr := db.Ping()
    if err != nil {
        log.Fatalf("impossilbe to pint the connection: %s", pingErr)
    }

    // Conduct insert
    insertQuery := "DELETE FROM list WHERE id = ?"
    stmt, err := db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(idFlag)
    if err != nil {
        log.Fatalf("not able to execute insert query: %s", err)
    }
}
