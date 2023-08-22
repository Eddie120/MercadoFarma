// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SearchInputResponse Search Input response
// Example: {"searchInputId":"be16910858a41fd19ea5c1b4e9decca9a784d1024cb00b2158defe2f29dc86dd"}
//
// swagger:model SearchInputResponse
type SearchInputResponse struct {

	// search input Id
	SearchInputID string `json:"searchInputId,omitempty"`
}

// Validate validates this search input response
func (m *SearchInputResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search input response based on context it is used
func (m *SearchInputResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchInputResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchInputResponse) UnmarshalBinary(b []byte) error {
	var res SearchInputResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}