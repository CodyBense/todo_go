package mySql

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func List() {
    // Load environment variable
    err := godotenv.Load()
    if err != nil {
        log.Fatalln("Error loading .env file")
    }
    connection := os.Getenv("MYSQL_CONNECTION")

    // Open mysql connection
    db, err := sql.Open("mysql", connection)
    if err != nil {
        log.Fatalf("impossible to create the connection: %s", err)
    }
    defer db.Close()

    // Test mysql connection
    pingErr := db.Ping()
    if err != nil {
        log.Fatalf("impossilbe to pint the connection: %s", pingErr)
    }
    
    // sql.Connect()

    // database variables
    var (
        task string
        description string
    )

    // To DO Header
    fmt.Println("To Do")

    // Conduct query
    rows, err := db.Query("SELECT * FROM todo")
    if err != nil {
        log.Fatalf("not able to conduct query: %s", err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&task, &description)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Task: %s\n\t%s\n", task, description)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    // To DO Header
    fmt.Printf("\n\nIn Progress\n")

    // Conduct query
    rows, err = db.Query("SELECT * FROM inProgress")
    if err != nil {
        log.Fatalf("not able to conduct query: %s", err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&task, &description)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Task: %s\n\t%s\n", task, description)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    // To DO Header
    fmt.Printf("\n\nDone\n")

    // Conduct query
    rows, err = db.Query("SELECT * FROM done")
    if err != nil {
        log.Fatalf("not able to conduct query: %s", err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&task, &description)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Task: %s\n\t%s\n", task, description)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}
