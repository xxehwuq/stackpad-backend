package entity

type Note struct {
	Id           string   `json:"id" gorm:"primaryKey"`
	Title        string   `json:"title" gorm:"default:Нова нотатка"`
	Text         string   `json:"text"`
	IsBookmarked bool     `json:"is_bookmarked" gorm:"default:false"`
	NotebookId   string   `json:"notebook_id"`
	UserId       string   `json:"user_id"`
	Notebook     Notebook `gorm:"foreignKey:notebook_id"`
	User         User     `gorm:"foreignKey:user_id"`
}
