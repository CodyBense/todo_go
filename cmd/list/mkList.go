/*
TODO:

    add update, add, and delete functionality
        fix update to refresh after
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

func resultsToList() []list.Item {

    results := bubbletea_queries.List()
    ItemsList := []list.Item{}

    for _, r := range results {
        ItemsList = append(ItemsList, item{title: r["task"], desc: r["done"], done: r["done"], id: r["id"]})
    }

    return ItemsList
}

func (m model) updateItem() {

    var currentIndex int

    items := bubbletea_queries.List()
    currentIndex = m.list.Index()
    currentId := items[currentIndex]["id"]
    i, err := strconv.Atoi(currentId)
    if err != nil {
        fmt.Println("Error converting string to int:", err)
        return
    }
    // fmt.Println(items[currentIndex]["id"])
    bubbletea_queries.Update(i)

}
