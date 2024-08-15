/*
TODO:

    add update, add, and delete functionality
	figure out how to change item color based on done status
	styling, such as checkmarks next to completed items(maybe) or crossed out, thing next to selected item
*/
package list

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CodyBense/todo/cmd/mySql/bubbletea_queries"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styling variables
var (
    appStyle = lipgloss.NewStyle().Margin(1, 2)

    titleStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#FFFDF5")).
        Background(lipgloss.Color("#25A065")).
        Padding(0, 1)
    
    notDoneStyle = lipgloss.Color("#CC0000")
    doneStyle = lipgloss.Color("#29FF03")
)

var ItemsList []list.Item

type item struct {
	title, desc, done, id string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.done }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
        // switch msg.String() {
        // case "ctrl+c":
        //     return m, tea.Quit
        // }
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
        if msg.String() == "u" {
            m.updateItem()
            return m, nil
        }
        if msg.String() == "d" {
            m.removeItem()
            return m, nil
        }
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}

func Main() {

    ItemsList := resultsToList()
    
    delegate := list.NewDefaultDelegate()
    delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(doneStyle)
    delegate.Styles.SelectedDesc = delegate.Styles.SelectedTitle

	m := model{list: list.New(ItemsList, delegate, 0, 0)}
	m.list.Title = "TODO"
    m.list.Styles.Title = titleStyle

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}


// Helper funcs

// adds all tasks to Item array
func resultsToList() []list.Item {

    results := bubbletea_queries.List()

    for _, r := range results {
        ItemsList = append(ItemsList, item{title: r["task"], desc: r["done"], done: r["done"], id: r["id"]})
    }

    return ItemsList
}

// Updates the status of an item
func (m model) updateItem() {

    items := bubbletea_queries.List()
    currentIndex := m.list.Index()
    currentId := items[currentIndex]["id"]
    i, err := strconv.Atoi(currentId)
    if err != nil {
        fmt.Println("Error converting string to int:", err)
        return
    }
    currentStatus := items[currentIndex]["done"]

    bubbletea_queries.Update(i)

    if currentStatus == "false" {
        m.list.SetItem(currentIndex, item{title: items[currentIndex]["task"], desc: "true", done: "true", id: items[currentIndex]["id"]})
    } else {
        m.list.SetItem(currentIndex, item{title: items[currentIndex]["task"], desc: "false", done: "false", id: items[currentIndex]["id"]})
    }
}

// Deletes an item
func (m model) removeItem() {

    items := bubbletea_queries.List()
    currentIndex := m.list.Index()
    currentId := items[currentIndex]["id"]
    i, err := strconv.Atoi(currentId)
    if err != nil {
        fmt.Println("Error converting string to int:", err)
        return
    }

    bubbletea_queries.Remove(i)
    m.list.RemoveItem(currentIndex)

}
