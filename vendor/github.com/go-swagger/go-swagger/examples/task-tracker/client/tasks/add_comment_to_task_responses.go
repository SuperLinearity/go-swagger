// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-swagger/go-swagger/examples/task-tracker/models"
)

// AddCommentToTaskReader is a Reader for the AddCommentToTask structure.
type AddCommentToTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCommentToTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddCommentToTaskCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddCommentToTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddCommentToTaskCreated creates a AddCommentToTaskCreated with default headers values
func NewAddCommentToTaskCreated() *AddCommentToTaskCreated {
	return &AddCommentToTaskCreated{}
}

/*AddCommentToTaskCreated handles this case with default header values.

Comment added
*/
type AddCommentToTaskCreated struct {
}

func (o *AddCommentToTaskCreated) Error() string {
	return fmt.Sprintf("[POST /tasks/{id}/comments][%d] addCommentToTaskCreated ", 201)
}

func (o *AddCommentToTaskCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddCommentToTaskDefault creates a AddCommentToTaskDefault with default headers values
func NewAddCommentToTaskDefault(code int) *AddCommentToTaskDefault {
	return &AddCommentToTaskDefault{
		_statusCode: code,
	}
}

/*AddCommentToTaskDefault handles this case with default header values.

Error response
*/
type AddCommentToTaskDefault struct {
	_statusCode int

	XErrorCode string

	Payload *models.Error
}

// Code gets the status code for the add comment to task default response
func (o *AddCommentToTaskDefault) Code() int {
	return o._statusCode
}

func (o *AddCommentToTaskDefault) Error() string {
	return fmt.Sprintf("[POST /tasks/{id}/comments][%d] addCommentToTask default  %+v", o._statusCode, o.Payload)
}

func (o *AddCommentToTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Error-Code
	o.XErrorCode = response.GetHeader("X-Error-Code")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
