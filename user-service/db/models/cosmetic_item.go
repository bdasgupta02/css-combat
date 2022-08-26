package models

type CosmeticItem struct {
	ID              uint              `json:"id"`
	Type            string            `json:"type"`
	Description     string            `json:"description"`
	Price           int64             `json:"price"`
	AvatarImage     string            `json:"avatar_image"`
	BannerImage     string            `json:"banner_image"`
	EditorColorType string            `json:"editor_color_type"`
	EditorColors    map[string]string `json:"editor_colors"`
}
