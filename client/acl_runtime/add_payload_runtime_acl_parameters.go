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

package acl_runtime

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

	models "github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// NewAddPayloadRuntimeACLParams creates a new AddPayloadRuntimeACLParams object
// with the default values initialized.
func NewAddPayloadRuntimeACLParams() *AddPayloadRuntimeACLParams {
	var ()
	return &AddPayloadRuntimeACLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddPayloadRuntimeACLParamsWithTimeout creates a new AddPayloadRuntimeACLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddPayloadRuntimeACLParamsWithTimeout(timeout time.Duration) *AddPayloadRuntimeACLParams {
	var ()
	return &AddPayloadRuntimeACLParams{

		timeout: timeout,
	}
}

// NewAddPayloadRuntimeACLParamsWithContext creates a new AddPayloadRuntimeACLParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddPayloadRuntimeACLParamsWithContext(ctx context.Context) *AddPayloadRuntimeACLParams {
	var ()
	return &AddPayloadRuntimeACLParams{

		Context: ctx,
	}
}

// NewAddPayloadRuntimeACLParamsWithHTTPClient creates a new AddPayloadRuntimeACLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddPayloadRuntimeACLParamsWithHTTPClient(client *http.Client) *AddPayloadRuntimeACLParams {
	var ()
	return &AddPayloadRuntimeACLParams{
		HTTPClient: client,
	}
}

/*AddPayloadRuntimeACLParams contains all the parameters to send to the API endpoint
for the add payload runtime ACL operation typically these are written to a http.Request
*/
type AddPayloadRuntimeACLParams struct {

	/*ACLID
	  ACL ID

	*/
	ACLID string
	/*Data*/
	Data models.ACLFilesEntries

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithTimeout(timeout time.Duration) *AddPayloadRuntimeACLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithContext(ctx context.Context) *AddPayloadRuntimeACLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithHTTPClient(client *http.Client) *AddPayloadRuntimeACLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLID adds the aCLID to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithACLID(aCLID string) *AddPayloadRuntimeACLParams {
	o.SetACLID(aCLID)
	return o
}

// SetACLID adds the aclId to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetACLID(aCLID string) {
	o.ACLID = aCLID
}

// WithData adds the data to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) WithData(data models.ACLFilesEntries) *AddPayloadRuntimeACLParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the add payload runtime ACL params
func (o *AddPayloadRuntimeACLParams) SetData(data models.ACLFilesEntries) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *AddPayloadRuntimeACLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param acl_id
	qrACLID := o.ACLID
	qACLID := qrACLID
	if qACLID != "" {
		if err := r.SetQueryParam("acl_id", qACLID); err != nil {
			return err
		}
	}

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
