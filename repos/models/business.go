package models

import "time"

type BusinessType string

const (
	PharmacyBusinessType        BusinessType = "FARMACIA"
	HealthFoodStoreBusinessType BusinessType = "TIENDA_NATURISTA"
	CosmeticStoreBusinessType   BusinessType = "TIENDA_COSMETICA"
	OtherBusinessType           BusinessType = "OTRO"
)

var IsValidBusinessTypes = map[BusinessType]bool{
	PharmacyBusinessType:        true,
	HealthFoodStoreBusinessType: true,
	CosmeticStoreBusinessType:   true,
	OtherBusinessType:           true,
}

type BusinessOpeningHour struct {
	Day        string     `json:"day"` // Monday,Tuesday,Wednesday...
	StartTime  *time.Time `json:"start_time"`
	EndingTime *time.Time `json:"ending_time"`
}

type Business struct {
	BusinessId           string                `json:"business_id"`
	TaxId                string                `json:"tax_id"`
	Name                 string                `json:"name"`
	BusinessType         BusinessType          `json:"business_type"`
	BusinessOpeningHours []BusinessOpeningHour `json:"business_opening_hours"`
	Address              string                `json:"address"`
	PhoneNumber          string                `json:"phone_number"`
	UserId               string                `json:"user_id"`
	SectorId             string                `json:"sector_id"`
	Active               bool                  `json:"active"`
	CreationDate         *time.Time            `json:"creation_date"`
	UpdateDate           *time.Time            `json:"update_date"`
}
