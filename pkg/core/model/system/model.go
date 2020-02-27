package system

import "time"

type Model struct {
	Id         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	CreatedBy  string
	ModifiedBy string
}
