package models

import "github.com/shopspring/decimal"

type Store struct {
	ID          uint64 `json:"id" gorm:"primary_key;not null"`
	StoreName   string `json:"store_name" gorm:"not null"`
	IsActive    bool   `json:"is_active" gorm:"not null"`
	BannerImage string `json:"banner_image"`
	LogoImage   string `json:"logo_image"`
	IconImage   string `json:"icon_image"`
}

type Deal struct {
	ID                 uint64          `json:"id" gorm:"primary_key;not null"`
	StoreID            int             `json:"store_id" gorm:"not null"`
	GameID             int             `json:"game_id" gorm:"not null"`
	SteamAppID         int             `json:"steam_app_id"`
	InternalName       string          `json:"internal_name" gorm:"not null"`
	Title              string          `json:"title" gorm:"not null"`
	MetacriticLink     string          `json:"metacritic_link"`
	SalesPrice         decimal.Decimal `json:"sales_price" gorm:"not null"`
	NormalPrice        decimal.Decimal `json:"normal_price" gorm:"not null"`
	IsOnSale           bool            `json:"is_on_sale" gorm:"not null"`
	Savings            decimal.Decimal `json:"savings" gorm:"not null"`
	MetacriticScore    int             `json:"metacritic_score" gorm:"not null"`
	SteamRatingText    string          `json:"steam_rating_text"`
	SteamRatingPercent int             `json:"steam_rating_percent" gorm:"not null"`
	SteamRatingCount   int             `json:"steam_rating_count" gorm:"not null"`
	ReleaseDate        int             `json:"release_date" gorm:"not null"`
	LastChange         int             `json:"last_change" gorm:"not null"`
	DealRating         decimal.Decimal `json:"deal_rating"`
	Thumb              string          `json:"thumb"`
}
