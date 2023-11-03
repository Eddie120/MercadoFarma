// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BusinessOpeningHour business opening hour
//
// swagger:model BusinessOpeningHour
type BusinessOpeningHour struct {

	// day
	// Enum: [monday tuesday wednesday thursday friday saturday sunday]
	Day string `json:"day,omitempty"`

	// ending time
	EndingTime string `json:"endingTime,omitempty"`

	// start time
	StartTime string `json:"startTime,omitempty"`
}

// Validate validates this business opening hour
func (m *BusinessOpeningHour) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var businessOpeningHourTypeDayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["monday","tuesday","wednesday","thursday","friday","saturday","sunday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		businessOpeningHourTypeDayPropEnum = append(businessOpeningHourTypeDayPropEnum, v)
	}
}

const (

	// BusinessOpeningHourDayMonday captures enum value "monday"
	BusinessOpeningHourDayMonday string = "monday"

	// BusinessOpeningHourDayTuesday captures enum value "tuesday"
	BusinessOpeningHourDayTuesday string = "tuesday"

	// BusinessOpeningHourDayWednesday captures enum value "wednesday"
	BusinessOpeningHourDayWednesday string = "wednesday"

	// BusinessOpeningHourDayThursday captures enum value "thursday"
	BusinessOpeningHourDayThursday string = "thursday"

	// BusinessOpeningHourDayFriday captures enum value "friday"
	BusinessOpeningHourDayFriday string = "friday"

	// BusinessOpeningHourDaySaturday captures enum value "saturday"
	BusinessOpeningHourDaySaturday string = "saturday"

	// BusinessOpeningHourDaySunday captures enum value "sunday"
	BusinessOpeningHourDaySunday string = "sunday"
)

// prop value enum
func (m *BusinessOpeningHour) validateDayEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, businessOpeningHourTypeDayPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BusinessOpeningHour) validateDay(formats strfmt.Registry) error {
	if swag.IsZero(m.Day) { // not required
		return nil
	}

	// value enum
	if err := m.validateDayEnum("day", "body", m.Day); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this business opening hour based on context it is used
func (m *BusinessOpeningHour) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BusinessOpeningHour) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BusinessOpeningHour) UnmarshalBinary(b []byte) error {
	var res BusinessOpeningHour
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}