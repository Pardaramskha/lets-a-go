package models

type Todo struct {
	ID      int64  `json:"id,omitempty"`
	Label   string `json:"label,omitempty"`
	Checked bool   `json:"checked"`
}

type TodoCreate struct {
	Label   string `json:"label"`
	Checked bool   `json:"checked"`
}
