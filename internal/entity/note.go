package entity

type Note struct {
	Id         string   `json:"id" gorm:"primaryKey"`
	Title      string   `json:"title"`
	Text       string   `json:"text"`
	NotebookId string   `json:"notebook_id"`
	UserId     string   `json:"user_id"`
	Notebook   Notebook `gorm:"foreignKey:notebook_id"`
	User       User     `gorm:"foreignKey:user_id"`
}
