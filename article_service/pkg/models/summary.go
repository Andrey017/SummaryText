package models

type Summary struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	FileName  string `json:"file_name"`
	Type      string `json:"type"`
	Note      string `json:"note"`
	ArticleId int64  `json:"article_id"`
}
