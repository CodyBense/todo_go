package bubbletea_queries

import (
    "log"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func Update(idFlag int) {

    // Open Mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/todo")
    if err != nil {
        log.Fatalf("impossible to create the connection: %s", err)
    }
    defer db.Close()

    // Test Mysql connection
    pingErr := db.Ping()
    if err != nil {
        log.Fatalf("impossilbe to pint the connection: %s", pingErr)
    }

    // Get status of desired task
    var (
        done bool
    )
    rows, err := db.Query("SELECT done FROM list WHERE id = ?", idFlag)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&done)
        if err != nil {
            log.Fatal(err)
        }
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    // Conduct update
    if done == false {
        insertQuery := "UPDATE list SET done = true WHERE id = ?"
        stmt, err := db.Prepare(insertQuery)
        if err != nil {
            log.Fatalf("not able to prepare insert query: %s", err)
        }
        defer stmt.Close()

        _, err = stmt.Exec(idFlag)
        if err != nil {
            log.Fatalf("not able to execute insert query: %s", err)
        }
    } else {
        insertQuery := "UPDATE list SET done = false WHERE id = ?"
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

}
