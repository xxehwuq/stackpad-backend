package entity

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsConfirmed bool   `json:"is_confirmed" gorm:"default:false"`
}
