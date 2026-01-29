package data

import "time"

type EmployeeModel struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:150;unique;not null"`
	Role      string `gorm:"size:50;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
