package entity

type Notebook struct {
	Id     string `json:"id" gorm:"primaryKey;foreignKey"`
	Title  string `json:"title"`
	Color  string `json:"color"`
	UserId string `json:"user_id" gorm:"references:user(id)"`
}
