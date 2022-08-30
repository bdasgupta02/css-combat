package models

type Models struct {
	EndUser               EndUser
	CosmeticItem          CosmeticItem
	CosmeticInventoryItem CosmeticInventoryItem
}

func New() Models {
	return Models{
		EndUser:               EndUser{},
		CosmeticItem:          CosmeticItem{},
		CosmeticInventoryItem: CosmeticInventoryItem{},
	}
}
