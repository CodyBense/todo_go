package app

import (
	"database/sql"
	"fmt"
	"log"

    "github.com/charmbracelet/bubbles/list"
)

/* TODO: Move to app dir
   made func (b Board)
*/

func SqlAdd(task, description, table string) {
    // Determines table
    var tableName string
    if table == "0" {
        tableName = "todo"
    } else if table == "1" {
        tableName = "inProgress"
    } else {
        tableName = "done"
    }
    // Open Mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/List")
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
    insertQuery := fmt.Sprintf("INSERT INTO %s (task, description) VALUES (?, ?)", tableName)
    stmt, err := db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(task, description)
    if err != nil {
        log.Fatalf("not able to execute insert query: %s", err)
    }

}

func (b *Board) SqlListTodo() []list.Item{
    // Open mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/List")
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
        task        string
        description string
        itemsList   []list.Item
    )

    // Conduct query

    // Select all from todo
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
        // output rows here
        itemsList = append(itemsList, NewTask(todo, task, description))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    return itemsList
}


func (b *Board) SqlListInProgress() []list.Item{
    // Open mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/List")
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
        task        string
        description string
        itemsList   []list.Item
    )

    // Conduct query

    // Select all from todo
    rows, err := db.Query("SELECT * FROM inProgress")
    if err != nil {
        log.Fatalf("not able to conduct query: %s", err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&task, &description)
        if err != nil {
            log.Fatal(err)
        }
        // output rows here
        itemsList = append(itemsList, NewTask(inProgress, task, description))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    return itemsList
}


func (b *Board) SqlListDone() []list.Item{
    // Open mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/List")
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
        task        string
        description string
        itemsList   []list.Item
    )

    // Conduct query

    // Select all from todo
    rows, err := db.Query("SELECT * FROM done")
    if err != nil {
        log.Fatalf("not able to conduct query: %s", err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&task, &description)
        if err != nil {
            log.Fatal(err)
        }
        // output rows here
        itemsList = append(itemsList, NewTask(done, task, description))
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    return itemsList
}

func (b *Board) SqlRemove(table, task string) {
    // Determines table
    var tableName string
    if table == "To Do" {
        tableName = "todo"
    } else if table == "In Progress" {
        tableName = "inProgress"
    } else {
        tableName = "done"
    }
    // Open mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/List")
    if err != nil {
        log.Fatalf("impossible to create the connection: %s", err)
    }
    defer db.Close()

    // Test mysql connection
    pingErr := db.Ping()
    if err != nil {
        log.Fatalf("impossilbe to pint the connection: %s", pingErr)
    }

    // Conduct query
    removeQuery := fmt.Sprintf("DELETE FROM %v WHERE task = ?", tableName)
    stmt, err := db.Prepare(removeQuery)
    if err != nil {
        log.Fatalf("not able to prepare remove query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(task)
    if err != nil {
        log.Fatalf("not able to execute remove query: %s", err)
    }
}


func SqlUpdate(table, task string) {

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

    // Conduct update
    insertQuery := "UPDATE list SET done = true WHERE id = ?"
    stmt, err := db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(task)
    if err != nil {
        log.Fatalf("not able to execute insert query: %s", err)
    }
}
