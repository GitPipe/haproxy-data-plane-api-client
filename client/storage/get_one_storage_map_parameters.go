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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetOneStorageMapParams creates a new GetOneStorageMapParams object
// with the default values initialized.
func NewGetOneStorageMapParams() *GetOneStorageMapParams {
	var ()
	return &GetOneStorageMapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOneStorageMapParamsWithTimeout creates a new GetOneStorageMapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOneStorageMapParamsWithTimeout(timeout time.Duration) *GetOneStorageMapParams {
	var ()
	return &GetOneStorageMapParams{

		timeout: timeout,
	}
}

// NewGetOneStorageMapParamsWithContext creates a new GetOneStorageMapParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOneStorageMapParamsWithContext(ctx context.Context) *GetOneStorageMapParams {
	var ()
	return &GetOneStorageMapParams{

		Context: ctx,
	}
}

// NewGetOneStorageMapParamsWithHTTPClient creates a new GetOneStorageMapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOneStorageMapParamsWithHTTPClient(client *http.Client) *GetOneStorageMapParams {
	var ()
	return &GetOneStorageMapParams{
		HTTPClient: client,
	}
}

/*GetOneStorageMapParams contains all the parameters to send to the API endpoint
for the get one storage map operation typically these are written to a http.Request
*/
type GetOneStorageMapParams struct {

	/*Name
	  Map file storage_name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get one storage map params
func (o *GetOneStorageMapParams) WithTimeout(timeout time.Duration) *GetOneStorageMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get one storage map params
func (o *GetOneStorageMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get one storage map params
func (o *GetOneStorageMapParams) WithContext(ctx context.Context) *GetOneStorageMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get one storage map params
func (o *GetOneStorageMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get one storage map params
func (o *GetOneStorageMapParams) WithHTTPClient(client *http.Client) *GetOneStorageMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get one storage map params
func (o *GetOneStorageMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get one storage map params
func (o *GetOneStorageMapParams) WithName(name string) *GetOneStorageMapParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get one storage map params
func (o *GetOneStorageMapParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetOneStorageMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
