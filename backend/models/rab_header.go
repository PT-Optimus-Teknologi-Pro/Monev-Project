package models

import "time"

type RabHeader struct {
	Id          uint  `gorm:"primaryKey" json:"id"`
	RabGroupId  *uint `gorm:"index;" json:"rab_group_id"`

	AlasanCount *int  `json:"alasan_count"`
	AlasanText   *string `json:"alasan_text"`
	
	Program      *string `json:"program"`

	DataEntryId uint       `json:"data_entry_id"`
	DataEntry   *DataEntry `gorm:"foreignKey:DataEntryId" json:"data_entry,omitempty"`

	CreatedById uint  `json:"created_by_id"`
	CreatedBy   *User `gorm:"foreignKey:CreatedById" json:"created_by,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	RabDetails []RabDetail `gorm:"foreignKey:RabHeaderId" json:"rab_details,omitempty"`
}

type RealisasiHeaderResponse struct {
	Id             uint                  `json:"id"`
	ScheduleHeader interface{}           `json:"schedule"`
	Detail         []RealisasiDetail `json:"detail"`
	DetailRevisi   []RealisasiDetail `json:"detail_revisi"`
}

