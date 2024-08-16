package app

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
    list list.Model
    err error
}

func New() *Model {
    return &Model{}
}

func (m *Model) initList(width, height int) {
    m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
    m.list.Title = "To Do"
    m.list.SetItems([]list.Item{
        Task{status: todo, title: "buy milk", description: "strawberry milk"},
        Task{status: todo, title: "eat sushi", description: "shrimp tempura roll"},
        Task{status: todo, title: "fold laundry", description: "or get from pile"},
    })
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.initList(msg.Width, msg.Height)
    }
    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m Model) View() string {
    return m.list.View()
}

func Main() {
    m := New()
    p := tea.NewProgram(m)
    if err := p.Start(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
