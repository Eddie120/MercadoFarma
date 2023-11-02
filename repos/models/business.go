package models

import "time"

type BusinessType string

const (
	PharmacyBusinessType        BusinessType = "FARMACIA"
	HealthFoodStoreBusinessType BusinessType = "TIENDA_NATURISTA"
	CosmeticStoreBusinessType   BusinessType = "TIENDA_COSMETICA"
	OtherBusinessType           BusinessType = "OTRO"
)

type Day string

const (
	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"
	Sunday    Day = "sunday"
)

var (
	IsValidBusinessTypes = map[BusinessType]bool{
		PharmacyBusinessType:        true,
		HealthFoodStoreBusinessType: true,
		CosmeticStoreBusinessType:   true,
		OtherBusinessType:           true,
	}

	IsValidDayOfWeek = map[Day]bool{
		Monday:    true,
		Tuesday:   true,
		Wednesday: true,
		Thursday:  true,
		Friday:    true,
		Saturday:  true,
		Sunday:    true,
	}
)

type BusinessOpeningHour struct {
	BusinessOpeningHourId int        `json:"business_opening_hour_id,omitempty"`
	Day                   string     `json:"day" json:"day,omitempty"` // Monday,Tuesday,Wednesday...
	StartTime             *time.Time `json:"start_time" json:"start_time,omitempty"`
	EndingTime            *time.Time `json:"ending_time" json:"ending_time,omitempty"`
}

type Business struct {
	BusinessId   string       `json:"business_id"`
	TaxId        string       `json:"tax_id"`
	Name         string       `json:"name"`
	BusinessType BusinessType `json:"business_type"`
	Address      string       `json:"address"`
	PhoneNumber  string       `json:"phone_number"`
	UserId       string       `json:"user_id"`
	SectorId     string       `json:"sector_id"`
	Active       bool         `json:"active"`
	CreationDate *time.Time   `json:"creation_date"`
	UpdateDate   *time.Time   `json:"update_date"`
}
