package models

import "database/sql"

type Models struct {
	EndUser EndUser
	CosmeticItem CosmeticItem
	CosmeticInventoryItem CosmeticInventoryItem
}

var DB *sql.DB

func New(dbPool *sql.DB) Models {
	DB = dbPool

	return Models{
		EndUser: EndUser{},
		CosmeticItem: CosmeticItem{},
		CosmeticInventoryItem: CosmeticInventoryItem{},
	}
}
