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
	"github.com/go-openapi/validate"
)

// Plan Plan
//
// The plan is what represent a concourse build.
//
// swagger:model Plan
type Plan struct {

	// aggregate
	Aggregate string `json:"aggregate,omitempty"`

	// attempts
	Attempts []uint32 `json:"attempts"`

	// do
	Do []*Plan `json:"do"`

	// ensure
	Ensure *EnsurePlan `json:"ensure,omitempty"`

	// get
	Get *GetPlan `json:"get,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// on failure
	OnFailure *OnFailurePlan `json:"on_failure,omitempty"`

	// on success
	OnSuccess *OnSuccessPlan `json:"on_success,omitempty"`

	// put
	Put *PutPlan `json:"put,omitempty"`

	// retry
	Retry []*Plan `json:"retry"`

	// task
	Task *TaskPlan `json:"task,omitempty"`

	// timeout
	Timeout *TimeoutPlan `json:"timeout,omitempty"`

	// try
	Try *TryPlan `json:"try,omitempty"`
}

// Validate validates this plan
func (m *Plan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnsure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnSuccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePut(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plan) validateDo(formats strfmt.Registry) error {
	if swag.IsZero(m.Do) { // not required
		return nil
	}

	for i := 0; i < len(m.Do); i++ {
		if swag.IsZero(m.Do[i]) { // not required
			continue
		}

		if m.Do[i] != nil {
			if err := m.Do[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("do" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("do" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Plan) validateEnsure(formats strfmt.Registry) error {
	if swag.IsZero(m.Ensure) { // not required
		return nil
	}

	if m.Ensure != nil {
		if err := m.Ensure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ensure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ensure")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) validateGet(formats strfmt.Registry) error {
	if swag.IsZero(m.Get) { // not required
		return nil
	}

	if m.Get != nil {
		if err := m.Get.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("get")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("get")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validateOnFailure(formats strfmt.Registry) error {
	if swag.IsZero(m.OnFailure) { // not required
		return nil
	}

	if m.OnFailure != nil {
		if err := m.OnFailure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("on_failure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("on_failure")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) validateOnSuccess(formats strfmt.Registry) error {
	if swag.IsZero(m.OnSuccess) { // not required
		return nil
	}

	if m.OnSuccess != nil {
		if err := m.OnSuccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("on_success")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("on_success")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) validatePut(formats strfmt.Registry) error {
	if swag.IsZero(m.Put) { // not required
		return nil
	}

	if m.Put != nil {
		if err := m.Put.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("put")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("put")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) validateRetry(formats strfmt.Registry) error {
	if swag.IsZero(m.Retry) { // not required
		return nil
	}

	for i := 0; i < len(m.Retry); i++ {
		if swag.IsZero(m.Retry[i]) { // not required
			continue
		}

		if m.Retry[i] != nil {
			if err := m.Retry[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("retry" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("retry" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Plan) validateTask(formats strfmt.Registry) error {
	if swag.IsZero(m.Task) { // not required
		return nil
	}

	if m.Task != nil {
		if err := m.Task.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("task")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) validateTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.Timeout) { // not required
		return nil
	}

	if m.Timeout != nil {
		if err := m.Timeout.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeout")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeout")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) validateTry(formats strfmt.Registry) error {
	if swag.IsZero(m.Try) { // not required
		return nil
	}

	if m.Try != nil {
		if err := m.Try.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("try")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("try")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plan based on the context it is used
func (m *Plan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnsure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnFailure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnSuccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePut(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRetry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeout(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plan) contextValidateDo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Do); i++ {

		if m.Do[i] != nil {

			if swag.IsZero(m.Do[i]) { // not required
				return nil
			}

			if err := m.Do[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("do" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("do" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Plan) contextValidateEnsure(ctx context.Context, formats strfmt.Registry) error {

	if m.Ensure != nil {

		if swag.IsZero(m.Ensure) { // not required
			return nil
		}

		if err := m.Ensure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ensure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ensure")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) contextValidateGet(ctx context.Context, formats strfmt.Registry) error {

	if m.Get != nil {

		if swag.IsZero(m.Get) { // not required
			return nil
		}

		if err := m.Get.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("get")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("get")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) contextValidateOnFailure(ctx context.Context, formats strfmt.Registry) error {

	if m.OnFailure != nil {

		if swag.IsZero(m.OnFailure) { // not required
			return nil
		}

		if err := m.OnFailure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("on_failure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("on_failure")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) contextValidateOnSuccess(ctx context.Context, formats strfmt.Registry) error {

	if m.OnSuccess != nil {

		if swag.IsZero(m.OnSuccess) { // not required
			return nil
		}

		if err := m.OnSuccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("on_success")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("on_success")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) contextValidatePut(ctx context.Context, formats strfmt.Registry) error {

	if m.Put != nil {

		if swag.IsZero(m.Put) { // not required
			return nil
		}

		if err := m.Put.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("put")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("put")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) contextValidateRetry(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Retry); i++ {

		if m.Retry[i] != nil {

			if swag.IsZero(m.Retry[i]) { // not required
				return nil
			}

			if err := m.Retry[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("retry" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("retry" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Plan) contextValidateTask(ctx context.Context, formats strfmt.Registry) error {

	if m.Task != nil {

		if swag.IsZero(m.Task) { // not required
			return nil
		}

		if err := m.Task.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("task")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("task")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) contextValidateTimeout(ctx context.Context, formats strfmt.Registry) error {

	if m.Timeout != nil {

		if swag.IsZero(m.Timeout) { // not required
			return nil
		}

		if err := m.Timeout.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeout")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeout")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) contextValidateTry(ctx context.Context, formats strfmt.Registry) error {

	if m.Try != nil {

		if swag.IsZero(m.Try) { // not required
			return nil
		}

		if err := m.Try.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("try")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("try")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plan) UnmarshalBinary(b []byte) error {
	var res Plan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
