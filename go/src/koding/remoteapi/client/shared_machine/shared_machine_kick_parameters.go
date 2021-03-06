package shared_machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewSharedMachineKickParams creates a new SharedMachineKickParams object
// with the default values initialized.
func NewSharedMachineKickParams() *SharedMachineKickParams {
	var ()
	return &SharedMachineKickParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSharedMachineKickParamsWithTimeout creates a new SharedMachineKickParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSharedMachineKickParamsWithTimeout(timeout time.Duration) *SharedMachineKickParams {
	var ()
	return &SharedMachineKickParams{

		timeout: timeout,
	}
}

// NewSharedMachineKickParamsWithContext creates a new SharedMachineKickParams object
// with the default values initialized, and the ability to set a context for a request
func NewSharedMachineKickParamsWithContext(ctx context.Context) *SharedMachineKickParams {
	var ()
	return &SharedMachineKickParams{

		Context: ctx,
	}
}

/*SharedMachineKickParams contains all the parameters to send to the API endpoint
for the shared machine kick operation typically these are written to a http.Request
*/
type SharedMachineKickParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the shared machine kick params
func (o *SharedMachineKickParams) WithTimeout(timeout time.Duration) *SharedMachineKickParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shared machine kick params
func (o *SharedMachineKickParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shared machine kick params
func (o *SharedMachineKickParams) WithContext(ctx context.Context) *SharedMachineKickParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shared machine kick params
func (o *SharedMachineKickParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the shared machine kick params
func (o *SharedMachineKickParams) WithBody(body models.DefaultSelector) *SharedMachineKickParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the shared machine kick params
func (o *SharedMachineKickParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SharedMachineKickParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
