package models

type Message struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    int    `json:"type"`
	Read    bool   `json:"read"`
}

type Compare struct {
	ID        string `gorm:"primaryKey;autoIncrement" json:"id"`
	Created   string `json:"created"`
	Status    int    `json:"status"`
	BeforeUrl string `json:"beforeUrl"`
	AfterUrl  string `json:"afterUrl"`
	User      int    `json:"user"`
}

type Pa struct {
	ID      string `gorm:"primaryKey;autoIncrement" json:"id"`
	Created string `json:"created"`
	Status  int    `json:"status"`
	Url     string `json:"url"`
	User    int    `json:"user"`
}

type Optimize struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	User       int    `json:"user"`
	TaskId     string `json:"taskId"`
	GptVersion string `json:"gptVersion"`
	Url        string `json:"url"`
}
