// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

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
)

// NewGetPoolProjectsParams creates a new GetPoolProjectsParams object
// with the default values initialized.
func NewGetPoolProjectsParams() *GetPoolProjectsParams {
	var ()
	return &GetPoolProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPoolProjectsParamsWithTimeout creates a new GetPoolProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPoolProjectsParamsWithTimeout(timeout time.Duration) *GetPoolProjectsParams {
	var ()
	return &GetPoolProjectsParams{

		timeout: timeout,
	}
}

// NewGetPoolProjectsParamsWithContext creates a new GetPoolProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPoolProjectsParamsWithContext(ctx context.Context) *GetPoolProjectsParams {
	var ()
	return &GetPoolProjectsParams{

		Context: ctx,
	}
}

// NewGetPoolProjectsParamsWithHTTPClient creates a new GetPoolProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPoolProjectsParamsWithHTTPClient(client *http.Client) *GetPoolProjectsParams {
	var ()
	return &GetPoolProjectsParams{
		HTTPClient: client,
	}
}

/*GetPoolProjectsParams contains all the parameters to send to the API endpoint
for the get pool projects operation typically these are written to a http.Request
*/
type GetPoolProjectsParams struct {

	/*AgentPoolLocator*/
	AgentPoolLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pool projects params
func (o *GetPoolProjectsParams) WithTimeout(timeout time.Duration) *GetPoolProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pool projects params
func (o *GetPoolProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pool projects params
func (o *GetPoolProjectsParams) WithContext(ctx context.Context) *GetPoolProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pool projects params
func (o *GetPoolProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pool projects params
func (o *GetPoolProjectsParams) WithHTTPClient(client *http.Client) *GetPoolProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pool projects params
func (o *GetPoolProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentPoolLocator adds the agentPoolLocator to the get pool projects params
func (o *GetPoolProjectsParams) WithAgentPoolLocator(agentPoolLocator string) *GetPoolProjectsParams {
	o.SetAgentPoolLocator(agentPoolLocator)
	return o
}

// SetAgentPoolLocator adds the agentPoolLocator to the get pool projects params
func (o *GetPoolProjectsParams) SetAgentPoolLocator(agentPoolLocator string) {
	o.AgentPoolLocator = agentPoolLocator
}

// WithFields adds the fields to the get pool projects params
func (o *GetPoolProjectsParams) WithFields(fields *string) *GetPoolProjectsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get pool projects params
func (o *GetPoolProjectsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetPoolProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentPoolLocator
	if err := r.SetPathParam("agentPoolLocator", o.AgentPoolLocator); err != nil {
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