package app

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Custom Item

func (t Task) FilterValue() string {
    return t.title
}

func (t Task) Title() string {
    return t.title
}

func (t Task) Description() string {
    return t.description
}

// Main model

type Model struct {
    focused     status
    lists       []list.Model
    err         error
    loaded      bool
}

func New() *Model {
    return &Model{}
}

func (m *Model) initLists(width, height int) {
    defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divsor, height)
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
            m.initLists(msg.Width, msg.Height)
            m.loaded = true
        }
    }
    var cmd tea.Cmd
    m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
    return m, cmd
}

func (m Model) View() string {
    if m.loaded {
        return lipgloss.JoinHorizontal(
            lipgloss.Left,
            m.lists[todo].View(),
            m.lists[inProgress].View(),
            m.lists[done].View(),
        )
    } else {
        return "loading..."
    }
}

func Main() {
    m := New()
    p := tea.NewProgram(m)
    if _,err := p.Run(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
