// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MppAggregateResult mpp aggregate result
//
// swagger:model MppAggregateResult
type MppAggregateResult struct {

	// antibiotics
	Antibiotics []*Antibiotics `json:"Antibiotics"`

	// compare drug costs
	CompareDrugCosts []*CompareDrugCosts `json:"CompareDrugCosts"`

	// demographics
	Demographics []*Demographics `json:"Demographics"`

	// gender
	Gender []*Gender `json:"Gender"`

	// lis nonlis
	LisNonlis []*LisNonlis `json:"LisNonlis"`

	// opioids
	Opioids []*Opioids `json:"Opioids"`

	// physician drug list
	PhysicianDrugList []*PhysicianDrugList `json:"PhysicianDrugList"`
}

// Validate validates this mpp aggregate result
func (m *MppAggregateResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAntibiotics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompareDrugCosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDemographics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLisNonlis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpioids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicianDrugList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MppAggregateResult) validateAntibiotics(formats strfmt.Registry) error {
	if swag.IsZero(m.Antibiotics) { // not required
		return nil
	}

	for i := 0; i < len(m.Antibiotics); i++ {
		if swag.IsZero(m.Antibiotics[i]) { // not required
			continue
		}

		if m.Antibiotics[i] != nil {
			if err := m.Antibiotics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Antibiotics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) validateCompareDrugCosts(formats strfmt.Registry) error {
	if swag.IsZero(m.CompareDrugCosts) { // not required
		return nil
	}

	for i := 0; i < len(m.CompareDrugCosts); i++ {
		if swag.IsZero(m.CompareDrugCosts[i]) { // not required
			continue
		}

		if m.CompareDrugCosts[i] != nil {
			if err := m.CompareDrugCosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CompareDrugCosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) validateDemographics(formats strfmt.Registry) error {
	if swag.IsZero(m.Demographics) { // not required
		return nil
	}

	for i := 0; i < len(m.Demographics); i++ {
		if swag.IsZero(m.Demographics[i]) { // not required
			continue
		}

		if m.Demographics[i] != nil {
			if err := m.Demographics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Demographics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) validateGender(formats strfmt.Registry) error {
	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	for i := 0; i < len(m.Gender); i++ {
		if swag.IsZero(m.Gender[i]) { // not required
			continue
		}

		if m.Gender[i] != nil {
			if err := m.Gender[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Gender" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) validateLisNonlis(formats strfmt.Registry) error {
	if swag.IsZero(m.LisNonlis) { // not required
		return nil
	}

	for i := 0; i < len(m.LisNonlis); i++ {
		if swag.IsZero(m.LisNonlis[i]) { // not required
			continue
		}

		if m.LisNonlis[i] != nil {
			if err := m.LisNonlis[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LisNonlis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) validateOpioids(formats strfmt.Registry) error {
	if swag.IsZero(m.Opioids) { // not required
		return nil
	}

	for i := 0; i < len(m.Opioids); i++ {
		if swag.IsZero(m.Opioids[i]) { // not required
			continue
		}

		if m.Opioids[i] != nil {
			if err := m.Opioids[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Opioids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) validatePhysicianDrugList(formats strfmt.Registry) error {
	if swag.IsZero(m.PhysicianDrugList) { // not required
		return nil
	}

	for i := 0; i < len(m.PhysicianDrugList); i++ {
		if swag.IsZero(m.PhysicianDrugList[i]) { // not required
			continue
		}

		if m.PhysicianDrugList[i] != nil {
			if err := m.PhysicianDrugList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PhysicianDrugList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this mpp aggregate result based on the context it is used
func (m *MppAggregateResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAntibiotics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCompareDrugCosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDemographics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGender(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLisNonlis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpioids(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhysicianDrugList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MppAggregateResult) contextValidateAntibiotics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Antibiotics); i++ {

		if m.Antibiotics[i] != nil {
			if err := m.Antibiotics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Antibiotics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) contextValidateCompareDrugCosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CompareDrugCosts); i++ {

		if m.CompareDrugCosts[i] != nil {
			if err := m.CompareDrugCosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CompareDrugCosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) contextValidateDemographics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Demographics); i++ {

		if m.Demographics[i] != nil {
			if err := m.Demographics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Demographics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) contextValidateGender(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Gender); i++ {

		if m.Gender[i] != nil {
			if err := m.Gender[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Gender" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) contextValidateLisNonlis(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LisNonlis); i++ {

		if m.LisNonlis[i] != nil {
			if err := m.LisNonlis[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LisNonlis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) contextValidateOpioids(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Opioids); i++ {

		if m.Opioids[i] != nil {
			if err := m.Opioids[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Opioids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MppAggregateResult) contextValidatePhysicianDrugList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhysicianDrugList); i++ {

		if m.PhysicianDrugList[i] != nil {
			if err := m.PhysicianDrugList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PhysicianDrugList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MppAggregateResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MppAggregateResult) UnmarshalBinary(b []byte) error {
	var res MppAggregateResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
