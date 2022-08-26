package models

import "time"

type CosmeticInventoryItem struct {
	ID             uint      `json:"id"`
	EndUserId      int       `json:"end_user_id"`
	CosmeticItemId int       `json:"cosmetic_item_id"`
	CurrencyUsed   int64     `json:"currency_used"`
	IsEquipped     bool      `json:"is_equipped"`
	TimePurchased  time.Time `json:"time_purchased"`
}
