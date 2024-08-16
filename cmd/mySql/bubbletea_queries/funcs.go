package bubbletea_queries

import (
	"database/sql"
	"fmt"
	"log"
)

func Add(task, state, table string) {
    // Open Mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306/List")
    if err != nil {
        log.Fatalf("impossible to create the connection: %s", err)
    }
    defer db.Close()

    // Test Mysql connection
    pingErr := db.Ping()
    if err != nil {
        log.Fatalf("impossible to ping the connection: %s", pingErr)
    }


    // Conduct insert
    insertQuery := fmt.Sprintf("INSERT INTO %s (task, state) VALUES (?, ?)", table)
    stmt, err := db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(task, state)
    if err != nil {
        log.Fatalf("not able to execute insert query: %s", err)
    }

}
