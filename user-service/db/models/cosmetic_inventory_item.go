package models

import "time"

type CosmeticInventoryItem struct {
	ID             uint      `json:"id"`
	EndUserId      int       `json:"endUser"`
	CosmeticItemId int       `json:"cosmeticItem"`
	CurrencyUsed   int64     `json:"currencyUsed"`
	IsEquipped     bool      `json:"isEquipped"`
	TimePurchased  time.Time `json:"timePurchased"`
}
