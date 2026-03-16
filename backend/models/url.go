package models

type Url struct {
	Id    uint    `gorm:"primaryKey" json:"id"`
	Url   string  `json:"url"`
	Tahun *string `json:"tahun"`
}
