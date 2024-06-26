package entity

type Notebook struct {
	Id     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"default:Новий зошит"`
	Color  string `json:"color"`
	UserId string `json:"user_id"`
	User   User   `gorm:"foreignKey:user_id"`
}
