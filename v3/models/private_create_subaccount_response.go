// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivateCreateSubaccountResponse private create subaccount response
// swagger:model private_create_subaccount_response
type PrivateCreateSubaccountResponse struct {
	BaseMessage

	// result
	// Required: true
	Result *PrivateCreateSubaccountResponseAO1Result `json:"result"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PrivateCreateSubaccountResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMessage = aO0

	// AO1
	var dataAO1 struct {
		Result *PrivateCreateSubaccountResponseAO1Result `json:"result"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PrivateCreateSubaccountResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Result *PrivateCreateSubaccountResponseAO1Result `json:"result"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this private create subaccount response
func (m *PrivateCreateSubaccountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMessage
	if err := m.BaseMessage.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateCreateSubaccountResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateCreateSubaccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateCreateSubaccountResponse) UnmarshalBinary(b []byte) error {
	var res PrivateCreateSubaccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrivateCreateSubaccountResponseAO1Result private create subaccount response a o1 result
// swagger:model PrivateCreateSubaccountResponseAO1Result
type PrivateCreateSubaccountResponseAO1Result struct {

	// User email
	// Required: true
	Email *string `json:"email"`

	// Subaccount identifier
	// Required: true
	ID *int64 `json:"id"`

	// `true` when password for the subaccount has been configured
	// Required: true
	IsPassword *bool `json:"is_password"`

	// Informs whether login to the subaccount is enabled
	// Required: true
	LoginEnabled *bool `json:"login_enabled"`

	// portfolio
	Portfolio *Portfolio `json:"portfolio,omitempty"`

	// When `true` - receive all notification emails on the main email
	// Required: true
	ReceiveNotifications *bool `json:"receive_notifications"`

	// System generated user nickname
	// Required: true
	SystemName *string `json:"system_name"`

	// Whether the two factor authentication is enabled
	// Required: true
	TfaEnabled *bool `json:"tfa_enabled"`

	// Account type
	// Required: true
	// Enum: [subaccount]
	Type *string `json:"type"`

	// Account name (given by user)
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this private create subaccount response a o1 result
func (m *PrivateCreateSubaccountResponseAO1Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoginEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortfolio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceiveNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTfaEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validateID(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validateIsPassword(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"is_password", "body", m.IsPassword); err != nil {
		return err
	}

	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validateLoginEnabled(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"login_enabled", "body", m.LoginEnabled); err != nil {
		return err
	}

	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validatePortfolio(formats strfmt.Registry) error {

	if swag.IsZero(m.Portfolio) { // not required
		return nil
	}

	if m.Portfolio != nil {
		if err := m.Portfolio.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + "portfolio")
			}
			return err
		}
	}

	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validateReceiveNotifications(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"receive_notifications", "body", m.ReceiveNotifications); err != nil {
		return err
	}

	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validateSystemName(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"system_name", "body", m.SystemName); err != nil {
		return err
	}

	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validateTfaEnabled(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"tfa_enabled", "body", m.TfaEnabled); err != nil {
		return err
	}

	return nil
}

var privateCreateSubaccountResponseAO1ResultTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["subaccount"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		privateCreateSubaccountResponseAO1ResultTypeTypePropEnum = append(privateCreateSubaccountResponseAO1ResultTypeTypePropEnum, v)
	}
}

const (

	// PrivateCreateSubaccountResponseAO1ResultTypeSubaccount captures enum value "subaccount"
	PrivateCreateSubaccountResponseAO1ResultTypeSubaccount string = "subaccount"
)

// prop value enum
func (m *PrivateCreateSubaccountResponseAO1Result) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, privateCreateSubaccountResponseAO1ResultTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validateType(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("result"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PrivateCreateSubaccountResponseAO1Result) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateCreateSubaccountResponseAO1Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateCreateSubaccountResponseAO1Result) UnmarshalBinary(b []byte) error {
	var res PrivateCreateSubaccountResponseAO1Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
