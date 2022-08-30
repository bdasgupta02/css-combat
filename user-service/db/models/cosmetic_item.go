package models

type CosmeticItem struct {
	ID              uint              `json:"id"`
	Type            string            `json:"type"`
	Description     string            `json:"description"`
	Price           int64             `json:"price"`
	AvatarImage     string            `json:"avatarImage"`
	BannerImage     string            `json:"bannerImage"`
	EditorColorType string            `json:"editorColorType"`
	EditorColors    map[string]string `json:"editorColors"`
}
