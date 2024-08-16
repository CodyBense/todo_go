package list

import (
	"fmt"
	"strconv"

	"github.com/CodyBense/todo/cmd/mySql/bubbletea_queries"
	"github.com/charmbracelet/bubbles/list"
)

// Helper funcs

// adds all tasks to Item array
func ResultsToList() []list.Item {

    results := bubbletea_queries.List()

    for _, r := range results {
        ItemsList = append(ItemsList, item{title: r["task"], desc: r["done"], done: r["done"], id: r["id"]})
    }

    return ItemsList
}

// Updates the status of an item
func (m model) UpdateItem() {

    items := bubbletea_queries.List()
    currentIndex := m.list.Index()
    currentId := items[currentIndex]["id"]
    i, err := strconv.Atoi(currentId)
    if err != nil {
        fmt.Println("Error converting string to int:", err)
        return
    }
    currentStatus := items[currentIndex]["done"]

    bubbletea_queries.Update(i)

    if currentStatus == "false" {
        m.list.SetItem(currentIndex, item{title: items[currentIndex]["task"], desc: "true", done: "true", id: items[currentIndex]["id"]})
    } else {
        m.list.SetItem(currentIndex, item{title: items[currentIndex]["task"], desc: "false", done: "false", id: items[currentIndex]["id"]})
    }
}

// Deletes an item
func (m model) RemoveItem() {

    items := bubbletea_queries.List()
    currentIndex := m.list.Index()
    currentId := items[currentIndex]["id"]
    i, err := strconv.Atoi(currentId)
    if err != nil {
        fmt.Println("Error converting string to int:", err)
        return
    }

    bubbletea_queries.Remove(i)
    m.list.RemoveItem(currentIndex)

}

// Adds an item
func (m model) AddItem() {

    newIndex := len(m.list.Items())
    fmt.Println(newIndex) // temp added


}
