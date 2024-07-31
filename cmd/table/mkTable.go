package table

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/go-sql-driver/mysql"
)

// Creates a base style
var baseStyle = lipgloss.NewStyle().
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("240"))

// Struct for model
type model struct {
    table table.Model
}

func (m model) Init() tea.Cmd { return nil }

// Update function, handles keyboard inputs
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "esc":
            if m.table.Focused() {
                m.table.Blur()
            } else {
                m.table.Focus()
            }
        case "q", "ctrl+c":
            return m, tea.Quit
        case "enter":
            return m, tea.Batch(
                tea.Printf("task is %s", m.table.SelectedRow()[1]),
            )
        } 
    }
    m.table, cmd = m.table.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return baseStyle.Render(m.table.View()) + "\n"
}

// Creates table and fills it
func Main() {
    columnsBt := []table.Column{
        {Title: "#", Width: 4},
        {Title: "Task", Width: 30},
        {Title: "Done", Width: 5},
    }

    // rows := []table.Row{
    //     {"1", "complete v3 of todo list", "false"},
    // }

    rowsBt := []table.Row{}

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
    rowsSql, err := db.Query("SELECT * FROM list")
    if err != nil {
        log.Fatalf("not able to conduct query: %s", err)
    }
    defer rowsSql.Close()

    for rowsSql.Next() {
        err := rowsSql.Scan(&id, &task, &done)
        if err != nil {
            log.Fatal(err)
        }
        // Maybe append to rowsBt here
        rowsBt = append(rowsBt, table.Row{fmt.Sprintf("%d", id), task, fmt.Sprintf("%v", done)})
    }
    err = rowsSql.Err()
    if err != nil {
        log.Fatal(err)
    }

    t := table.New(
        table.WithColumns(columnsBt),
        table.WithRows(rowsBt),
        table.WithFocused(true),
        table.WithHeight(7),
    )

    s := table.DefaultStyles()
    s.Header = s.Header.
        BorderStyle(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("240")).
        BorderBottom(true).
        Bold(false)
    s.Selected = s.Selected.
        Foreground(lipgloss.Color("229")).
        Background(lipgloss.Color("57")).
        Bold(true)

    m := model{t}
    if _, err := tea.NewProgram(m).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}
