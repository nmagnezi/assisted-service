// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// ListComponentVersionsReader is a Reader for the ListComponentVersions structure.
type ListComponentVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListComponentVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListComponentVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListComponentVersionsOK creates a ListComponentVersionsOK with default headers values
func NewListComponentVersionsOK() *ListComponentVersionsOK {
	return &ListComponentVersionsOK{}
}

/*ListComponentVersionsOK handles this case with default header values.

Success.
*/
type ListComponentVersionsOK struct {
	Payload *models.ListVersions
}

func (o *ListComponentVersionsOK) Error() string {
	return fmt.Sprintf("[GET /component_versions][%d] listComponentVersionsOK  %+v", 200, o.Payload)
}

func (o *ListComponentVersionsOK) GetPayload() *models.ListVersions {
	return o.Payload
}

func (o *ListComponentVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListVersions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
