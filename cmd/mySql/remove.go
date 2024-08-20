package mySql

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Remove(taskFlag, statusFlag *string) {
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

    // Conduct insert
    deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE task = ?", *statusFlag)
    stmt, err := db.Prepare(deleteQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(taskFlag)
    if err != nil {
        log.Fatalf("not able to execute insert query: %s", err)
    }
}
