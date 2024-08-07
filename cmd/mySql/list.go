package mySql

import (
    "fmt"
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func List() {

    fmt.Println("# |\t\ttask\t\t|\tdone\t")

    // Open mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/todo")
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
