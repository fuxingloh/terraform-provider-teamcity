// Code generated by go-swagger; DO NOT EDIT.

package build

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

// NewGetArtifactsDirectoryParams creates a new GetArtifactsDirectoryParams object
// with the default values initialized.
func NewGetArtifactsDirectoryParams() *GetArtifactsDirectoryParams {
	var ()
	return &GetArtifactsDirectoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArtifactsDirectoryParamsWithTimeout creates a new GetArtifactsDirectoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArtifactsDirectoryParamsWithTimeout(timeout time.Duration) *GetArtifactsDirectoryParams {
	var ()
	return &GetArtifactsDirectoryParams{

		timeout: timeout,
	}
}

// NewGetArtifactsDirectoryParamsWithContext creates a new GetArtifactsDirectoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArtifactsDirectoryParamsWithContext(ctx context.Context) *GetArtifactsDirectoryParams {
	var ()
	return &GetArtifactsDirectoryParams{

		Context: ctx,
	}
}

// NewGetArtifactsDirectoryParamsWithHTTPClient creates a new GetArtifactsDirectoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArtifactsDirectoryParamsWithHTTPClient(client *http.Client) *GetArtifactsDirectoryParams {
	var ()
	return &GetArtifactsDirectoryParams{
		HTTPClient: client,
	}
}

/*GetArtifactsDirectoryParams contains all the parameters to send to the API endpoint
for the get artifacts directory operation typically these are written to a http.Request
*/
type GetArtifactsDirectoryParams struct {

	/*BuildLocator*/
	BuildLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get artifacts directory params
func (o *GetArtifactsDirectoryParams) WithTimeout(timeout time.Duration) *GetArtifactsDirectoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get artifacts directory params
func (o *GetArtifactsDirectoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get artifacts directory params
func (o *GetArtifactsDirectoryParams) WithContext(ctx context.Context) *GetArtifactsDirectoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get artifacts directory params
func (o *GetArtifactsDirectoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get artifacts directory params
func (o *GetArtifactsDirectoryParams) WithHTTPClient(client *http.Client) *GetArtifactsDirectoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get artifacts directory params
func (o *GetArtifactsDirectoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the get artifacts directory params
func (o *GetArtifactsDirectoryParams) WithBuildLocator(buildLocator string) *GetArtifactsDirectoryParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get artifacts directory params
func (o *GetArtifactsDirectoryParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetArtifactsDirectoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}