package models

type Article struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	FileName   string `json:"file_name"`
	Note       string `json:"note"`
	DateCreate string `json:"date_create"`
	UserId     int64  `json:"user_id"`
}
