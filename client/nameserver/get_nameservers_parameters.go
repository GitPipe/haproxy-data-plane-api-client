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

package nameserver

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

// NewGetNameserversParams creates a new GetNameserversParams object
// with the default values initialized.
func NewGetNameserversParams() *GetNameserversParams {
	var ()
	return &GetNameserversParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNameserversParamsWithTimeout creates a new GetNameserversParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNameserversParamsWithTimeout(timeout time.Duration) *GetNameserversParams {
	var ()
	return &GetNameserversParams{

		timeout: timeout,
	}
}

// NewGetNameserversParamsWithContext creates a new GetNameserversParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNameserversParamsWithContext(ctx context.Context) *GetNameserversParams {
	var ()
	return &GetNameserversParams{

		Context: ctx,
	}
}

// NewGetNameserversParamsWithHTTPClient creates a new GetNameserversParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNameserversParamsWithHTTPClient(client *http.Client) *GetNameserversParams {
	var ()
	return &GetNameserversParams{
		HTTPClient: client,
	}
}

/*GetNameserversParams contains all the parameters to send to the API endpoint
for the get nameservers operation typically these are written to a http.Request
*/
type GetNameserversParams struct {

	/*Resolver
	  Parent resolver name

	*/
	Resolver string
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nameservers params
func (o *GetNameserversParams) WithTimeout(timeout time.Duration) *GetNameserversParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nameservers params
func (o *GetNameserversParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nameservers params
func (o *GetNameserversParams) WithContext(ctx context.Context) *GetNameserversParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nameservers params
func (o *GetNameserversParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nameservers params
func (o *GetNameserversParams) WithHTTPClient(client *http.Client) *GetNameserversParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nameservers params
func (o *GetNameserversParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResolver adds the resolver to the get nameservers params
func (o *GetNameserversParams) WithResolver(resolver string) *GetNameserversParams {
	o.SetResolver(resolver)
	return o
}

// SetResolver adds the resolver to the get nameservers params
func (o *GetNameserversParams) SetResolver(resolver string) {
	o.Resolver = resolver
}

// WithTransactionID adds the transactionID to the get nameservers params
func (o *GetNameserversParams) WithTransactionID(transactionID *string) *GetNameserversParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get nameservers params
func (o *GetNameserversParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNameserversParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param resolver
	qrResolver := o.Resolver
	qResolver := qrResolver
	if qResolver != "" {
		if err := r.SetQueryParam("resolver", qResolver); err != nil {
			return err
		}
	}

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string
		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {
			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
