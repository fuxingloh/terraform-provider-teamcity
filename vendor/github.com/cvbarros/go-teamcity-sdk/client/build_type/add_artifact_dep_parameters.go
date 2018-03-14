// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// NewAddArtifactDepParams creates a new AddArtifactDepParams object
// with the default values initialized.
func NewAddArtifactDepParams() *AddArtifactDepParams {
	var ()
	return &AddArtifactDepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddArtifactDepParamsWithTimeout creates a new AddArtifactDepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddArtifactDepParamsWithTimeout(timeout time.Duration) *AddArtifactDepParams {
	var ()
	return &AddArtifactDepParams{

		timeout: timeout,
	}
}

// NewAddArtifactDepParamsWithContext creates a new AddArtifactDepParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddArtifactDepParamsWithContext(ctx context.Context) *AddArtifactDepParams {
	var ()
	return &AddArtifactDepParams{

		Context: ctx,
	}
}

// NewAddArtifactDepParamsWithHTTPClient creates a new AddArtifactDepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddArtifactDepParamsWithHTTPClient(client *http.Client) *AddArtifactDepParams {
	var ()
	return &AddArtifactDepParams{
		HTTPClient: client,
	}
}

/*AddArtifactDepParams contains all the parameters to send to the API endpoint
for the add artifact dep operation typically these are written to a http.Request
*/
type AddArtifactDepParams struct {

	/*Body*/
	Body *models.ArtifactDependency
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add artifact dep params
func (o *AddArtifactDepParams) WithTimeout(timeout time.Duration) *AddArtifactDepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add artifact dep params
func (o *AddArtifactDepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add artifact dep params
func (o *AddArtifactDepParams) WithContext(ctx context.Context) *AddArtifactDepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add artifact dep params
func (o *AddArtifactDepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add artifact dep params
func (o *AddArtifactDepParams) WithHTTPClient(client *http.Client) *AddArtifactDepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add artifact dep params
func (o *AddArtifactDepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add artifact dep params
func (o *AddArtifactDepParams) WithBody(body *models.ArtifactDependency) *AddArtifactDepParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add artifact dep params
func (o *AddArtifactDepParams) SetBody(body *models.ArtifactDependency) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the add artifact dep params
func (o *AddArtifactDepParams) WithBtLocator(btLocator string) *AddArtifactDepParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the add artifact dep params
func (o *AddArtifactDepParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the add artifact dep params
func (o *AddArtifactDepParams) WithFields(fields *string) *AddArtifactDepParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add artifact dep params
func (o *AddArtifactDepParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddArtifactDepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}