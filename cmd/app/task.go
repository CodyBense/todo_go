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

func (t *Task) Next() {
    if t.status == done {
        t.status = todo
    } else {
        t.status++
    }
}
