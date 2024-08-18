package app

// Use for mock data for testing

func (b *Board) initLists() {
    b.cols = []column{
        newColumn(todo),
        newColumn(inProgress),
        newColumn(done),
    }
    // Init To Do
	b.cols[todo].list.Title = "To Do"
    itemsList := b.SqlListTodo()
    b.cols[todo].list.SetItems(itemsList)

	// Init in progress
	b.cols[inProgress].list.Title = "In Progress"
    itemsList = b.SqlListInProgress()
    b.cols[inProgress].list.SetItems(itemsList)

	// Init done
	b.cols[done].list.Title = "Done"
    itemsList = b.SqlListDone()
    b.cols[done].list.SetItems(itemsList)
}
