package models

import (
	"time"
)

type Log struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	Timestamp time.Time
	Method    string
	Path      string
	Status    int
}
