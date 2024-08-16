package app

// Use for mock data for testing

func (b *Board) initLists() {
    b.cols = []column{
        newColumn(todo),
        newColumn(inProgress),
        newColumn(done),
    }
}
