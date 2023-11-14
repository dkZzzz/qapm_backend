package models

type Message struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    int    `json:"type"`
	Read    bool   `json:"read"`
}

type Compare struct {
	Id        string `json:"id"`
	Created   string `json:"created"`
	Status    int    `json:"status"`
	BeforeUrl string `json:"beforeUrl"`
	AfterUrl  string `json:"afterUrl"`
	User      int    `json:"user"`
}

type Pa struct {
	Id      string `json:"id"`
	Created string `json:"created"`
	Status  int    `json:"status"`
	Url     string `json:"url"`
	User    int    `json:"user"`
}

type Optimize struct {
	Id      string `json:"id"`
	Created string `json:"created"`
	Status  int    `json:"status"`
	User    int    `json:"user"`
	TaskId  string `json:"taskId"`
}
