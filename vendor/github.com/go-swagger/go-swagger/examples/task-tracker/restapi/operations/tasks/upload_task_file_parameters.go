// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUploadTaskFileParams creates a new UploadTaskFileParams object
// with the default values initialized.
func NewUploadTaskFileParams() UploadTaskFileParams {
	var ()
	return UploadTaskFileParams{}
}

// UploadTaskFileParams contains all the bound params for the upload task file operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadTaskFile
type UploadTaskFileParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Extra information describing the file
	  In: formData
	*/
	Description *string
	/*The file to upload
	  In: formData
	*/
	File *runtime.File
	/*The id of the item
	  Required: true
	  In: path
	*/
	ID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UploadTaskFileParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	fdDescription, fdhkDescription, _ := fds.GetOK("description")
	if err := o.bindDescription(fdDescription, fdhkDescription, route.Formats); err != nil {
		res = append(res, err)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "file", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindFile(file, fileHeader); err != nil {
		res = append(res, err)
	} else {
		o.File = &runtime.File{Data: file, Header: fileHeader}
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UploadTaskFileParams) bindDescription(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Description = &raw

	return nil
}

func (o *UploadTaskFileParams) bindFile(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *UploadTaskFileParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}
