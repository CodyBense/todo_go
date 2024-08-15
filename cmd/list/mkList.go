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

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// From Model

type Form struct {
    help help.Model
    title textinput.Model
    description textarea.Model
}

func newDefaultForm() *Form {

    return NewForm("task name", "")

}

func NewForm(title, description string) *Form {

    form := Form{
        help: help.New(),
        title: textinput.New(),
        description: textarea.New(),
    }

    form.title.Placeholder = title
    form.description.Placeholder = description
    form.title.Focus()

    return &form
}

func (f Form) CreateTask() item {

    return item{f.title.Value(), f.description.Value(), "false", "ud"}

}

func (f Form) Init() tea.Cmd {
    
    return nil

}

func (f Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return f, tea.Quit
        case "enter":
            if f.title.Focused() {
                f.title.Blur()
                f.description.Focus()
                return f, textarea.Blink
            }
            return model.Update(f)
        }
    }

    if f.title.Focused() {
        f.title, cmd = f.title.Update(msg)
        return f, cmd
    }
    f.description, cmd = f.description.Update(msg)
    return f, cmd
}

func (f Form) View() string {
    return lipgloss.JoinVertical(
            lipgloss.Left,
            "Create a new task",
            f.title.View(),
            f.description.View())
}

// List Model
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
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
        if msg.String() == "u" {
            m.UpdateItem()
            return m, nil
        }
        if msg.String() == "d" {
            m.RemoveItem()
            return m, nil
        }
        if msg.String() == "a" {
            m.AddItem()
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

    ItemsList := ResultsToList()
    
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
