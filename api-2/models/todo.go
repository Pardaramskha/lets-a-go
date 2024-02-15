package models

type Todo struct {
	ID      any    `json:"id"`
	Label   string `json:"label"`
	Checked bool   `json:"checked"`
}
