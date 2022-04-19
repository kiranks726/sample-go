package todos

type TodoItem struct {
	Id      string `json:"Id,omitempty"`
	Title   string `json:"Title"`
	Details string `json:"Details"`
	Message string `json:"Message"`
}
