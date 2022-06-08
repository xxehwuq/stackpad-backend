package entity

type Note struct {
	Id         string `json:"id" gorm:"primaryKey"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	NotebookId string `json:"notebook_id" gorm:"references:notebook(id)"`
	UserId     string `json:"user_id" gorm:"references:user(id)"`
}
