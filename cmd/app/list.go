package app

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* STYLING */

var (
    columnStyle = lipgloss.NewStyle().
            Padding(1, 2)
    focusedStyle = lipgloss.NewStyle().
            Padding(1, 2).
            Border(lipgloss.RoundedBorder()).
            BorderForeground(lipgloss.Color("62"))
    helpStyle = lipgloss.NewStyle().
            Foreground(lipgloss.Color("241"))
)

/* CUSTOM ITEM */

func (t Task) FilterValue() string {
    return t.title
}

func (t Task) Title() string {
    return t.title
}

func (t Task) Description() string {
    return t.description
}

/* MAIN MODEL */

type Model struct {
    loaded      bool
    focused     status
    lists       []list.Model
    err         error
    quitting    bool
}

func New() *Model {
    return &Model{}
}

func (m *Model) MoveToNext() tea.Msg {
    selectedItem := m.lists[m.focused].SelectedItem()
    selectedTask := selectedItem.(Task)
    m.lists[selectedTask.status].RemoveItem(m.lists[m.focused].Index())
    selectedTask.Next()
    m.lists[selectedTask.status].InsertItem(len(m.lists[selectedTask.status].Items())-1, list.Item(selectedTask))
    return nil
}

func (m *Model) Next() {
    if m.focused == done {
        m.focused = todo
    } else {
        m.focused++
    }
}

func (m *Model) Prev() {
    if m.focused == todo {
        m.focused = done
    } else {
        m.focused--
    }
}

func (m *Model) initLists(width, height int) {
    defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divsor, height/2)
    defaultList.SetShowHelp(false)
    m.lists = []list.Model{defaultList, defaultList, defaultList}

    // Init To Do
    m.lists[todo].Title = "To Do"
    m.lists[todo].SetItems([]list.Item{
        Task{status: todo, title: "buy milk", description: "strawberry milk"},
        Task{status: todo, title: "eat sushi", description: "shrimp tempura roll"},
        Task{status: todo, title: "fold laundry", description: "or get from pile"},
    })
    // Init In Progress
    m.lists[inProgress].Title = "In Progress"
    m.lists[inProgress].SetItems([]list.Item{
        Task{status: inProgress, title: "write code", description: "don't worry, it's Go"},
    })
    // Init Done
    m.lists[done].Title = "Done"
    m.lists[done].SetItems([]list.Item{
        Task{status: done, title: "stay cool", description: "as a cucumber"},
    })
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        if !m.loaded {
            columnStyle.Width(msg.Width / divsor)
            focusedStyle.Width(msg.Width / divsor)
            columnStyle.Height(msg.Height - divsor)
            focusedStyle.Height(msg.Height - divsor)
            m.initLists(msg.Width, msg.Height)
            m.loaded = true
        }
    case tea.KeyMsg:
        switch msg.String(){
        case "ctrl+c", "q":
            m.quitting = true
            return m, tea.Quit
        case "left", "h":
            m.Prev()
        case "right", "l":
            m.Next()
        case "enter":
            return m, m.MoveToNext
        }
    }
    var cmd tea.Cmd
    m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
    return m, cmd
}

func (m Model) View() string {
    if m.quitting {
        return ""
    }
    if m.loaded {
        todoView := m.lists[todo].View()
        inProgressView := m.lists[inProgress].View()
        doneView := m.lists[done].View()

        switch m.focused {
        case inProgress:
            return lipgloss.JoinHorizontal(
                lipgloss.Left,
                columnStyle.Render(todoView),
                focusedStyle.Render(inProgressView),
                columnStyle.Render(doneView),
            )
        case done:
            return lipgloss.JoinHorizontal(
                lipgloss.Left,
                columnStyle.Render(todoView),
                columnStyle.Render(inProgressView),
                focusedStyle.Render(doneView),
            )
        default:
            return lipgloss.JoinHorizontal(
                lipgloss.Left,
                focusedStyle.Render(todoView),
                columnStyle.Render(inProgressView),
                columnStyle.Render(doneView),
            )
        }

    } else {
        return "loading..."
    }
}

func Main() {
    m := New()
    p := tea.NewProgram(m, tea.WithAltScreen())
    if _,err := p.Run(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
