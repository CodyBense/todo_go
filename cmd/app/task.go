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
