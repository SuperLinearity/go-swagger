// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdatePetWithFormParams creates a new UpdatePetWithFormParams object
// with the default values initialized.
func NewUpdatePetWithFormParams() UpdatePetWithFormParams {
	var ()
	return UpdatePetWithFormParams{}
}

// UpdatePetWithFormParams contains all the bound params for the update pet with form operation
// typically these are obtained from a http.Request
//
// swagger:parameters updatePetWithForm
type UpdatePetWithFormParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Updated name of the pet
	  Required: true
	  In: formData
	*/
	Name string
	/*ID of pet that needs to be updated
	  Required: true
	  In: path
	*/
	PetID string
	/*Updated status of the pet
	  Required: true
	  In: formData
	*/
	Status string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdatePetWithFormParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	fdName, fdhkName, _ := fds.GetOK("name")
	if err := o.bindName(fdName, fdhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rPetID, rhkPetID, _ := route.Params.GetOK("petId")
	if err := o.bindPetID(rPetID, rhkPetID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdStatus, fdhkStatus, _ := fds.GetOK("status")
	if err := o.bindStatus(fdStatus, fdhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePetWithFormParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("name", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("name", "formData", raw); err != nil {
		return err
	}

	o.Name = raw

	return nil
}

func (o *UpdatePetWithFormParams) bindPetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.PetID = raw

	return nil
}

func (o *UpdatePetWithFormParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("status", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("status", "formData", raw); err != nil {
		return err
	}

	o.Status = raw

	return nil
}