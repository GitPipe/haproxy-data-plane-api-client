// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreateSpoeParams creates a new CreateSpoeParams object
// with the default values initialized.
func NewCreateSpoeParams() *CreateSpoeParams {
	var ()
	return &CreateSpoeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSpoeParamsWithTimeout creates a new CreateSpoeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSpoeParamsWithTimeout(timeout time.Duration) *CreateSpoeParams {
	var ()
	return &CreateSpoeParams{

		timeout: timeout,
	}
}

// NewCreateSpoeParamsWithContext creates a new CreateSpoeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSpoeParamsWithContext(ctx context.Context) *CreateSpoeParams {
	var ()
	return &CreateSpoeParams{

		Context: ctx,
	}
}

// NewCreateSpoeParamsWithHTTPClient creates a new CreateSpoeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSpoeParamsWithHTTPClient(client *http.Client) *CreateSpoeParams {
	var ()
	return &CreateSpoeParams{
		HTTPClient: client,
	}
}

/*CreateSpoeParams contains all the parameters to send to the API endpoint
for the create spoe operation typically these are written to a http.Request
*/
type CreateSpoeParams struct {

	/*FileUpload
	  The spoe file to upload

	*/
	FileUpload runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create spoe params
func (o *CreateSpoeParams) WithTimeout(timeout time.Duration) *CreateSpoeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create spoe params
func (o *CreateSpoeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create spoe params
func (o *CreateSpoeParams) WithContext(ctx context.Context) *CreateSpoeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create spoe params
func (o *CreateSpoeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create spoe params
func (o *CreateSpoeParams) WithHTTPClient(client *http.Client) *CreateSpoeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create spoe params
func (o *CreateSpoeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileUpload adds the fileUpload to the create spoe params
func (o *CreateSpoeParams) WithFileUpload(fileUpload runtime.NamedReadCloser) *CreateSpoeParams {
	o.SetFileUpload(fileUpload)
	return o
}

// SetFileUpload adds the fileUpload to the create spoe params
func (o *CreateSpoeParams) SetFileUpload(fileUpload runtime.NamedReadCloser) {
	o.FileUpload = fileUpload
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSpoeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FileUpload != nil {

		if o.FileUpload != nil {

			// form file param file_upload
			if err := r.SetFileParam("file_upload", o.FileUpload); err != nil {
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}