package models

import (
	"github.com/jinzhu/gorm"
)

type ScreenShot struct {
	gorm.Model
	TicketID  uint
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Filename  string `json:"filename",gorm:"size:255"`
	UserAgent string `json:"userAgent",gorm:"size:512"`
}
