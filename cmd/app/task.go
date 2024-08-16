package app

type Task struct {
    status      status
    title       string
    description string
}

func NewTask(status status, title, description string) Task {
    return Task{status: status, title: title, description: description}
}

func (t *Task) Next() {
    if t.status == done {
        t.status = todo
    } else {
        t.status++
    }
}

func (t *Task) Prev() {
    if t.status == todo {
        t.status = done
    } else {
        t.status--
    }
}

// Implement list.Item interface
func (t Task) FilterValue() string {
    return t.title
}

func (t Task) Title() string {
    return t.title
}

func (t Task) Description() string {
    return t.description
}

type Item struct {
    status      status
    title       string
    description string
}

func (i Item) Status() status {
    return i.status
}

func (i Item) Title() string {
    return i.title
}

func (i Item) Description() string {
    return i.description
}

