package app

type status int

const divsor = 4

const (
    todo status = iota
    inProgress
    done
)

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
