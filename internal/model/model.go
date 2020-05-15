package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	ID      int64  `gorm:"column:id" json:"id"`           // 实力图片
	Content string `gorm:"column:content" json:"content"` // 实力图片
	Author  string `gorm:"column:author" json:"author"`   // 实力图片
	//Powers  string `gorm:"column:powers" json:"powers"` // 实力图片
}

func (Article) TableName() string {
	return "articles"
}

type Music struct {
	ID      int64
	Content string
	Author  string
}
