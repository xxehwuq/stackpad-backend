package entity

type Notebook struct {
	Id     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Color  string `json:"color"`
	UserId string `json:"user_id" gorm:"foreignKey;references:user(id)"`
}
