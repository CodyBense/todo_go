package cmd

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

// Variables for database
var db *sql.DB

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
    fmt.Println("# |\t\ttask\t\t|\tdone\t")

    var err error

    // Open mysql connection
    db, err = sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/todo")
    if err != nil {
        log.Fatalf("impossible to create the connection: %s", err)
    }
    defer db.Close()

    // Test mysql connection
    pingErr := db.Ping()
    if err != nil {
        log.Fatalf("impossilbe to pint the connection: %s", pingErr)
    }

    var (
        id int
        task string
        done bool
    )

    // Conduct query
    rows, err := db.Query("SELECT * FROM list")
    if err != nil {
        log.Fatalf("not able to conduct query: %s", err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&id, &task, &done)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%d |%s\t\t|%v\n", id, task, done)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}
