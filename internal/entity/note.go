package entity

type Note struct {
	Id         string `json:"id" gorm:"primaryKey"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	NotebookId string `json:"notebook_id" gorm:"foreignKey;references:notebook(id)"`
	UserId     string `json:"user_id" gorm:"foreignKey;references:user(id)"`
}
