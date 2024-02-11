package main

type Todo struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	Checked   bool   `json:"checked"`
	CreatedAt string `json:"createdAt"`
}

// dumb data
var todos = []*Todo{
	{
		ID:        "1",
		Label:     "Test todo",
		Checked:   false,
		CreatedAt: "creation date",
	},
}

func listTodos() []*Todo {
	return todos
}
