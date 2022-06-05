package entity

type User struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique; not null"`
	Password    string `json:"password"`
	IsConfirmed bool   `json:"is_confirmed" gorm:"default:false"`
}
