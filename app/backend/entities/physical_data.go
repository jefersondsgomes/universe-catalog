package entities

import "time"

type PhysicalData struct {
	ID          uint64    `json:"id" gorm:"primary_key:auto_increment"`
	AstroID     uint64    `json:"-" gorm:"uniqueIndex"`
	Mass        float64   `json:"mass" gorm:"default:0;not null"`
	Temperature float64   `json:"temperature" gorm:"default:0;not null"`
	Inserted    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}

func (PhysicalData) TableName() string {
	return "data"
}
