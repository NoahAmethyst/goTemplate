package entity

import "time"

/**

 */
type BaseEntity struct {
	CreatedTime time.Time `gorm:"column:created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time"`
}

func InitBaseEntity() BaseEntity {
	now := time.Now()
	return BaseEntity{
		CreatedTime: now,
		UpdatedTime: now,
	}
}
