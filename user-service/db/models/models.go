package models

import (
	"github.com/jackc/pgx/v4"
)

type Models struct {
	EndUser               EndUser
	CosmeticItem          CosmeticItem
	CosmeticInventoryItem CosmeticInventoryItem
}

var db *pgx.Conn

func New(dbPool *pgx.Conn) Models {
	db = dbPool

	return Models{
		EndUser:               EndUser{},
		CosmeticItem:          CosmeticItem{},
		CosmeticInventoryItem: CosmeticInventoryItem{},
	}
}
