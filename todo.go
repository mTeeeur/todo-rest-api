package todo_rest_api

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Complited   bool   `json:"complited"`
}

type ListsItem struct {
	Id     int
	UserId int
	ItemId int
}
