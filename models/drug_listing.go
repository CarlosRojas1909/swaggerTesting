// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DrugListing drug listing
//
// swagger:model DrugListing
type DrugListing struct {

	// cost per claim
	CostPerClaim int64 `json:"cost_per_claim,omitempty"`

	// drug name
	DrugName string `json:"drug_name,omitempty"`

	// generic name
	GenericName string `json:"generic_name,omitempty"`

	// geo
	Geo interface{} `json:"geo,omitempty"`
}

// Validate validates this drug listing
func (m *DrugListing) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this drug listing based on context it is used
func (m *DrugListing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DrugListing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DrugListing) UnmarshalBinary(b []byte) error {
	var res DrugListing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
