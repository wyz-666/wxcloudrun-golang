package model

import "gorm.io/gorm"

type BuyerTx struct {
	gorm.Model
	TID      string `gorm:"column:tid;type:varchar(64)" json:"tid"`
	UserID   string `gorm:"column:userId" json:"userId"`
	Project  string `gorm:"column:project" json:"project"`
	Type     string `gorm:"column:type" json:"type"`
	Price    string `gorm:"column:price" json:"price"`
	TxVolume string `gorm:"column:txVolume" json:"txVolume"`
	State    int    `gorm:"column:state" json:"state"`
}

type SellerTx struct {
	gorm.Model
	TID      string `gorm:"column:tid;type:varchar(64)" json:"tid"`
	UserID   string `gorm:"column:userId" json:"userId"`
	Project  string `gorm:"column:project" json:"project"`
	Type     string `gorm:"column:type" json:"type"`
	Price    string `gorm:"column:price" json:"price"`
	TxVolume string `gorm:"column:txVolume" json:"txVolume"`
	State    int    `gorm:"column:state" json:"state"`
}
