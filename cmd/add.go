package cmd

import (
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

    var lastId, nextId int
    
    // Get last id in the list and prepare the next id
    rows, err := db.Query("SELECT MAX(id) FROM list")
    if err != nil {
        log.Fatalf("not able to conduct query: %s", err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&lastId)
        if err != nil {
            log.Fatal(err)
        }
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    nextId = lastId + 1

    // Conduct insert
    insertQuery := "INSERT INTO list (id, task, done) VALUES (?, ?, ?)"
    stmt, err := db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(nextId, taskFlag ,false)
    if err != nil {
        log.Fatalf("not able to execute insert query: %s", err)
    }
}
