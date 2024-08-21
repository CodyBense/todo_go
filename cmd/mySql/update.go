package mySql

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Update(taskFlag, statusFlag, updateFlag *string) {
    // Load environment vairable
    err := godotenv.Load()
    if err != nil {
        log.Fatalln("Error loading .env file")
    }
    connection := os.Getenv("MYSQL_CONNECTION")

    // Open Mysql connection
    db, err := sql.Open("mysql", connection)
    if err != nil {
        log.Fatalf("impossible to create the connection: %s", err)
    }
    defer db.Close()

    // Test Mysql connection
    pingErr := db.Ping()
    if err != nil {
        log.Fatalf("impossilbe to pint the connection: %s", pingErr)
    }

    var description string

    // Get description
    rows, err := db.Query(fmt.Sprintf("SELECT description FROM %s WHERE task = ?", *statusFlag), taskFlag) 
    if err != nil {
        log.Fatalf("not able to prepare select (update) query: %s", err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&description)
        if err != nil {
            log.Fatal(err)
        }
    }

    // Conduct delete (update)
    deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE task = ?", *statusFlag) 
    stmt, err := db.Prepare(deleteQuery)
    if err != nil {
        log.Fatalf("not able to prepare delete (update) query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(taskFlag)
    if err != nil {
        log.Fatalf("not able to execute delete (update) query: %s", err)
    }
    
    insertQuery := fmt.Sprintf("INSERT INTO %s (task, description) VALUES (?,?)", *updateFlag)
    stmt, err = db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert (update) query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(taskFlag, description)
    if err != nil {
        log.Fatalf("not able to execute insert (update) query: %s", err)
    }
}
