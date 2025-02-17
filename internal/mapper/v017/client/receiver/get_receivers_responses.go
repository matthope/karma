// Code generated by go-swagger; DO NOT EDIT.

package receiver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prymitive/karma/internal/mapper/v017/models"
)

// GetReceiversReader is a Reader for the GetReceivers structure.
type GetReceiversReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReceiversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReceiversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReceiversOK creates a GetReceiversOK with default headers values
func NewGetReceiversOK() *GetReceiversOK {
	return &GetReceiversOK{}
}

/*
GetReceiversOK describes a response with status code 200, with default header values.

Get receivers response
*/
type GetReceiversOK struct {
	Payload []*models.Receiver
}

// IsSuccess returns true when this get receivers o k response has a 2xx status code
func (o *GetReceiversOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get receivers o k response has a 3xx status code
func (o *GetReceiversOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get receivers o k response has a 4xx status code
func (o *GetReceiversOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get receivers o k response has a 5xx status code
func (o *GetReceiversOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get receivers o k response a status code equal to that given
func (o *GetReceiversOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get receivers o k response
func (o *GetReceiversOK) Code() int {
	return 200
}

func (o *GetReceiversOK) Error() string {
	return fmt.Sprintf("[GET /receivers][%d] getReceiversOK  %+v", 200, o.Payload)
}

func (o *GetReceiversOK) String() string {
	return fmt.Sprintf("[GET /receivers][%d] getReceiversOK  %+v", 200, o.Payload)
}

func (o *GetReceiversOK) GetPayload() []*models.Receiver {
	return o.Payload
}

func (o *GetReceiversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
