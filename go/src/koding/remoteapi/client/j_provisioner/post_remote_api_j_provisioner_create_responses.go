package j_provisioner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJProvisionerCreateReader is a Reader for the PostRemoteAPIJProvisionerCreate structure.
type PostRemoteAPIJProvisionerCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJProvisionerCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJProvisionerCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJProvisionerCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJProvisionerCreateOK creates a PostRemoteAPIJProvisionerCreateOK with default headers values
func NewPostRemoteAPIJProvisionerCreateOK() *PostRemoteAPIJProvisionerCreateOK {
	return &PostRemoteAPIJProvisionerCreateOK{}
}

/*PostRemoteAPIJProvisionerCreateOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJProvisionerCreateOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJProvisionerCreateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProvisioner.create][%d] postRemoteApiJProvisionerCreateOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJProvisionerCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJProvisionerCreateUnauthorized creates a PostRemoteAPIJProvisionerCreateUnauthorized with default headers values
func NewPostRemoteAPIJProvisionerCreateUnauthorized() *PostRemoteAPIJProvisionerCreateUnauthorized {
	return &PostRemoteAPIJProvisionerCreateUnauthorized{}
}

/*PostRemoteAPIJProvisionerCreateUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJProvisionerCreateUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJProvisionerCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProvisioner.create][%d] postRemoteApiJProvisionerCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJProvisionerCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
