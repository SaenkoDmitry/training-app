package dto

type TelegramUser struct {
	ID           int64  `json:"id"` // chat_id
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	PhotoURL     string `json:"photo_url"`
	AuthDate     int64  `json:"auth_date"`
	Hash         string `json:"hash"`
}

type YandexProfile struct {
	ID           string `json:"id"`
	Login        string `json:"login"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DefaultEmail string `json:"default_email"`
}
