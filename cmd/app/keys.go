package app

import "github.com/charmbracelet/bubbles/key"

// Return keybinds to be shown in mini help
func (k keyMap) ShortHelp() []key.Binding {
    return []key.Binding{k.Help, k.Quit}
}

// Return keybinds for extended help
func (k keyMap) FullHelp() [][]key.Binding {
    return [][]key.Binding {
        {k.Up, k.Down, k.Left, k.Right},
        {k.Help, k.Quit},
    }
}

type keyMap struct {
    New     key.Binding
    Edit    key.Binding
    Delete  key.Binding
    Up      key.Binding
    Down    key.Binding
    Right   key.Binding
    Left    key.Binding
    Enter   key.Binding
    Help    key.Binding
    Quit    key.Binding
    Back    key.Binding
}

var keys = keyMap{
    New:    key.NewBinding(
            key.WithKeys("n"),
            key.WithHelp("n", "new"),
    ),
    Edit:   key.NewBinding(
            key.WithKeys("e"),
            key.WithHelp("e", "edit"),
    ),
    Delete: key.NewBinding(
            key.WithKeys("d"),
            key.WithHelp("d", "delete"),
    ),
    Up:     key.NewBinding(
            key.WithKeys("up", "k"),
            key.WithHelp("↑/k", "move up"),
    ),
    Down:   key.NewBinding(
            key.WithKeys("down", "j"),
            key.WithHelp("↓/j", "move down"),
    ),
    Right:  key.NewBinding(
            key.WithKeys("right", "l"),
            key.WithHelp("→/l", "move right"),
    ),
    Left:   key.NewBinding(
            key.WithKeys("left", "h"),
            key.WithHelp("←/h", "move left"),
    ),
    Enter:  key.NewBinding(
            key.WithKeys("enter"),
            key.WithHelp("enter", "enter"),
    ),
    Help:   key.NewBinding(
            key.WithKeys("?"),
            key.WithHelp("?", "help"),
    ),
    Quit:   key.NewBinding(
            key.WithKeys("ctrl+c"),
            key.WithHelp("ctrl+c", "quit"),
    ),
    Back:   key.NewBinding(
            key.WithKeys("esc"),
            key.WithHelp("esc", "back"),
    ),
}
