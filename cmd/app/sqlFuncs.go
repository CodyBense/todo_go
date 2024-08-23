package app

import (
	"database/sql"
	"fmt"
	"log"

	customlog "github.com/CodyBense/todo/cmd/customLog"
	"github.com/charmbracelet/bubbles/list"
)

func SqlConnect() (*sql.DB, error) {
    // Open log file and set it
    customlog.Log()

    // Open Mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(192.168.1.129:3306)/List")
    if err != nil {
        log.Fatalf("impossible to create the connection: %s", err)
    }
    return db, err
}

func SqlAdd(task, description, table string) {
    // Open log file and set it
    customlog.Log()

    // Determines table
    var tableName string
    if table == "0" {
        tableName = "todo"
    } else if table == "1" {
        tableName = "inProgress"
    } else {
        tableName = "done"
    }

    // Open and test SQL connection
    db, err := SqlConnect()

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

    db.Close()

}

func (b *Board) SqlListTodo() []list.Item{
    // Open log file and set it
    customlog.Log()

    // Open and test SQL connection
    db, err := SqlConnect()

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

    db.Close()

    return itemsList
}


func (b *Board) SqlListInProgress() []list.Item{
    // Open log file and set it
    customlog.Log()

    // Open and test SQL connection
    db, err := SqlConnect()

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

    db.Close()

    return itemsList
}


func (b *Board) SqlListDone() []list.Item{
    // Open log file and set it
    customlog.Log()

    // Open and test SQL connection
    db, err := SqlConnect()

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

    db.Close()

    return itemsList
}

func (b *Board) SqlRemove(table, task string) {
    // Open log file and set it
    customlog.Log()

    // Determines table
    var tableName string
    if table == "To Do" {
        tableName = "todo"
    } else if table == "In Progress" {
        tableName = "inProgress"
    } else {
        tableName = "done"
    }

    // Open and test SQL connection
    db, err := SqlConnect()

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

    db.Close()
}


func SqlUpdate(currentTable, task, description string) {
    // Open log file and set it
    customlog.Log()

    var nextTable, tableName string

    if currentTable == "To Do" {
        tableName = "todo"
    } else if currentTable == "In Progress" {
        tableName = "inProgress"
    } else {
        tableName = "done"
    }

    if tableName == "todo" {
        nextTable = "inProgress"
    } else if tableName == "inProgress" {
        nextTable = "done"
    } else {
        nextTable = "todo"
    }

    // Open and test SQL connection
    db, err := SqlConnect()

    // Conduct update
    deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE task = ?", tableName)
    stmt, err := db.Prepare(deleteQuery)
    if err != nil {
        log.Fatalf("not able to prepare delete (update) query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(task)
    if err != nil {
        log.Fatalf("not able to execute delete (update) query: %s", err)
    }

    insertQuery := fmt.Sprintf("INSERT INTO %s (task, description) VALUES (?,?)", nextTable)
    stmt, err = db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert (update) query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(task, description)
    if err != nil {
        log.Fatalf("not abel to execute update query: %s", err)
    }

    db.Close()
}
