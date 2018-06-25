package storage

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

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// NewImageJoinParams creates a new ImageJoinParams object
// with the default values initialized.
func NewImageJoinParams() *ImageJoinParams {
	var ()
	return &ImageJoinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImageJoinParamsWithTimeout creates a new ImageJoinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageJoinParamsWithTimeout(timeout time.Duration) *ImageJoinParams {
	var ()
	return &ImageJoinParams{

		timeout: timeout,
	}
}

// NewImageJoinParamsWithContext creates a new ImageJoinParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageJoinParamsWithContext(ctx context.Context) *ImageJoinParams {
	var ()
	return &ImageJoinParams{

		Context: ctx,
	}
}

// NewImageJoinParamsWithHTTPClient creates a new ImageJoinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageJoinParamsWithHTTPClient(client *http.Client) *ImageJoinParams {
	var ()
	return &ImageJoinParams{
		HTTPClient: client,
	}
}

/*ImageJoinParams contains all the parameters to send to the API endpoint
for the image join operation typically these are written to a http.Request
*/
type ImageJoinParams struct {

	/*Config*/
	Config *models.ImageJoinConfig
	/*ID*/
	ID string
	/*StoreName*/
	StoreName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image join params
func (o *ImageJoinParams) WithTimeout(timeout time.Duration) *ImageJoinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image join params
func (o *ImageJoinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image join params
func (o *ImageJoinParams) WithContext(ctx context.Context) *ImageJoinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image join params
func (o *ImageJoinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image join params
func (o *ImageJoinParams) WithHTTPClient(client *http.Client) *ImageJoinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image join params
func (o *ImageJoinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the image join params
func (o *ImageJoinParams) WithConfig(config *models.ImageJoinConfig) *ImageJoinParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the image join params
func (o *ImageJoinParams) SetConfig(config *models.ImageJoinConfig) {
	o.Config = config
}

// WithID adds the id to the image join params
func (o *ImageJoinParams) WithID(id string) *ImageJoinParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the image join params
func (o *ImageJoinParams) SetID(id string) {
	o.ID = id
}

// WithStoreName adds the storeName to the image join params
func (o *ImageJoinParams) WithStoreName(storeName string) *ImageJoinParams {
	o.SetStoreName(storeName)
	return o
}

// SetStoreName adds the storeName to the image join params
func (o *ImageJoinParams) SetStoreName(storeName string) {
	o.StoreName = storeName
}

// WriteToRequest writes these params to a swagger request
func (o *ImageJoinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Config == nil {
		o.Config = new(models.ImageJoinConfig)
	}

	if err := r.SetBodyParam(o.Config); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param store_name
	if err := r.SetPathParam("store_name", o.StoreName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}