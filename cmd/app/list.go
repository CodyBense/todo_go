package app

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* MODEL MANAGEMENT */
var models []tea.Model
const (
    model status = iota
    form
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
        case "a":
            models[model] = m // save the state of the current model
            models[form] = NewForm(m.focused)
            return models[form].Update(nil)
        }
    case Task:
        task := msg
        return m, m.lists[task.status].InsertItem(len(m.lists[task.status].Items()), task)
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

/* FORM MODEL */
type Form struct {
    focused     status
    title       textinput.Model
    description textarea.Model
}

func NewForm(focused status) *Form {
    form := &Form{focused: focused}
    form.title = textinput.New()
    form.title.Focus()
    form.description = textarea.New()
    return form
}

func (f Form) Init() tea.Cmd {
    return nil
}

func (f Form) CreateTask() tea.Msg {
    task := NewTask(f.focused, f.title.Value(), f.description.Value())
    return task
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
            } else {
                models[form] = f
                return models[model], f.CreateTask
            }
        }
    }
    if f.title.Focused() {
        f.title, cmd = f.title.Update(msg)
        return f, cmd
    } else {
        f.description, cmd = f.description.Update(msg)
        return f, cmd
    }
}

func (f Form) View() string {
    return lipgloss.JoinVertical(lipgloss.Left, f.title.View(), f.description.View())
}

func Main() {
    models = []tea.Model{New(), NewForm(todo)}
    m := models[model]
    p := tea.NewProgram(m, tea.WithAltScreen())
    if _,err := p.Run(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
