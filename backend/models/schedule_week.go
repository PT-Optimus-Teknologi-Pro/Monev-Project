package models

import "time"

type ScheduleWeek struct {
	Id uint `gorm:"primaryKey" json:"id"`

	ScheduleItemId uint          `json:"schedule_item_id"`
	ScheduleItem   *ScheduleItem `json:"schedule_item,omitempty"`

	WeekNumber int     `json:"week_number"`
	Value      float64 `json:"value"`

	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}