package mySql

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Add( taskFlag, descriptionFlag, statusFlag *string) {

    // Open Mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/List")
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
    insertQuery := fmt.Sprintf("INSERT INTO %s (task, description) VALUES (?,?)", *statusFlag)
    stmt, err := db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(taskFlag, descriptionFlag)
    if err != nil {
        log.Fatalf("not able to execute insert query: %s", err)
    }
}
