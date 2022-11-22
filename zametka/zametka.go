package zametka

type Zametka struct {
	Id         int64  `gorm:"Key" json:"id"`
	Author     string `gorm:"type:varchar(300)" json:"author"`
	Content    string `gorm:"type:text" json:"content"`
	DatePublic string `gorm:"type:date" json:"date_public"`
}
