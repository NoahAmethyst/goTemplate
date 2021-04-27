package entity

type Demo struct {
	Id     string `gorm:"primaryKey"`
	Name   string `gorm:"column:name"`
	Number int    `gorm:"column:number"`
	BaseEntity
}

func (Demo) TableName() string {
	return "demo"
}
