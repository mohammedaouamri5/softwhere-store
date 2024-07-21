package models
import (
    "github.com/google/uuid" 
)


type Goods struct {
    GOODS_ID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"goods_id"`
    REF       string    `gorm:"type:varchar(100);not null" json:"ref"`
    Photo     string    `gorm:"type:varchar(255)" json:"photo"`
};  