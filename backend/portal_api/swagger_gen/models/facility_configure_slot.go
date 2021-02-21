// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FacilityConfigureSlot facility configure slot
//
// swagger:model FacilityConfigureSlot
type FacilityConfigureSlot struct {

	// appointment schedule
	AppointmentSchedule []*FacilityAppointmentSchedule `json:"appointmentSchedule"`

	// facility Id
	// Required: true
	FacilityID *string `json:"facilityId"`

	// program Id
	// Required: true
	ProgramID *string `json:"programId"`

	// walk in schedule
	WalkInSchedule []*FacilityWalkInSchedule `json:"walkInSchedule"`
}

// Validate validates this facility configure slot
func (m *FacilityConfigureSlot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppointmentSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgramID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWalkInSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacilityConfigureSlot) validateAppointmentSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.AppointmentSchedule) { // not required
		return nil
	}

	for i := 0; i < len(m.AppointmentSchedule); i++ {
		if swag.IsZero(m.AppointmentSchedule[i]) { // not required
			continue
		}

		if m.AppointmentSchedule[i] != nil {
			if err := m.AppointmentSchedule[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appointmentSchedule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FacilityConfigureSlot) validateFacilityID(formats strfmt.Registry) error {

	if err := validate.Required("facilityId", "body", m.FacilityID); err != nil {
		return err
	}

	return nil
}

func (m *FacilityConfigureSlot) validateProgramID(formats strfmt.Registry) error {

	if err := validate.Required("programId", "body", m.ProgramID); err != nil {
		return err
	}

	return nil
}

func (m *FacilityConfigureSlot) validateWalkInSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.WalkInSchedule) { // not required
		return nil
	}

	for i := 0; i < len(m.WalkInSchedule); i++ {
		if swag.IsZero(m.WalkInSchedule[i]) { // not required
			continue
		}

		if m.WalkInSchedule[i] != nil {
			if err := m.WalkInSchedule[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("walkInSchedule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FacilityConfigureSlot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacilityConfigureSlot) UnmarshalBinary(b []byte) error {
	var res FacilityConfigureSlot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
