package bot

type Data struct {
	Error   string
	IP      string
	Url     string
	Request interface{}
}

type Message struct {
	Service Service     `json:"service" binding:"required"`
	Data    MessageData `json:"data" binding:"required"`
}

type Service struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type MessageData struct {
	Date    string `json:"date" binding:"required"`
	Error   string `json:"error" binding:"required"`
	IP      string `json:"ip" binding:"required"`
	URL     string `json:"url" binding:"required"`
	Request string `json:"request"`
}
