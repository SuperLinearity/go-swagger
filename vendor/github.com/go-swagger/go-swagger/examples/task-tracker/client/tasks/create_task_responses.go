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

// CreateTaskReader is a Reader for the CreateTask structure.
type CreateTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateTaskCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTaskCreated creates a CreateTaskCreated with default headers values
func NewCreateTaskCreated() *CreateTaskCreated {
	return &CreateTaskCreated{}
}

/*CreateTaskCreated handles this case with default header values.

Task created
*/
type CreateTaskCreated struct {
}

func (o *CreateTaskCreated) Error() string {
	return fmt.Sprintf("[POST /tasks][%d] createTaskCreated ", 201)
}

func (o *CreateTaskCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTaskDefault creates a CreateTaskDefault with default headers values
func NewCreateTaskDefault(code int) *CreateTaskDefault {
	return &CreateTaskDefault{
		_statusCode: code,
	}
}

/*CreateTaskDefault handles this case with default header values.

Error response
*/
type CreateTaskDefault struct {
	_statusCode int

	XErrorCode string

	Payload *models.Error
}

// Code gets the status code for the create task default response
func (o *CreateTaskDefault) Code() int {
	return o._statusCode
}

func (o *CreateTaskDefault) Error() string {
	return fmt.Sprintf("[POST /tasks][%d] createTask default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Error-Code
	o.XErrorCode = response.GetHeader("X-Error-Code")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
