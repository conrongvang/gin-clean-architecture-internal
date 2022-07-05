package models

type BaseModel struct {
	Id        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt int64  `json:"deleted_at,omitempty"`
}
