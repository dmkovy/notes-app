package notes

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserList struct {
	Id     int
	UserId int
}

type Note struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Data        []byte `json:"data"`
	Destruct    bool   `json:"destruct"`
	Done        bool   `json:"done"`
}

type NoteList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
