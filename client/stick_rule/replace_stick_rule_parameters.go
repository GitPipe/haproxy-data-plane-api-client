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

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models"
)

// NewReplaceStickRuleParams creates a new ReplaceStickRuleParams object
// with the default values initialized.
func NewReplaceStickRuleParams() *ReplaceStickRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceStickRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceStickRuleParamsWithTimeout creates a new ReplaceStickRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceStickRuleParamsWithTimeout(timeout time.Duration) *ReplaceStickRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceStickRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewReplaceStickRuleParamsWithContext creates a new ReplaceStickRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceStickRuleParamsWithContext(ctx context.Context) *ReplaceStickRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceStickRuleParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewReplaceStickRuleParamsWithHTTPClient creates a new ReplaceStickRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceStickRuleParamsWithHTTPClient(client *http.Client) *ReplaceStickRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceStickRuleParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*ReplaceStickRuleParams contains all the parameters to send to the API endpoint
for the replace stick rule operation typically these are written to a http.Request
*/
type ReplaceStickRuleParams struct {

	/*Backend
	  Backend name

	*/
	Backend string
	/*Data*/
	Data *models.StickRule
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*ID
	  Stick Rule ID

	*/
	ID int64
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string
	/*Version
	  Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace stick rule params
func (o *ReplaceStickRuleParams) WithTimeout(timeout time.Duration) *ReplaceStickRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace stick rule params
func (o *ReplaceStickRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace stick rule params
func (o *ReplaceStickRuleParams) WithContext(ctx context.Context) *ReplaceStickRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace stick rule params
func (o *ReplaceStickRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace stick rule params
func (o *ReplaceStickRuleParams) WithHTTPClient(client *http.Client) *ReplaceStickRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace stick rule params
func (o *ReplaceStickRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the replace stick rule params
func (o *ReplaceStickRuleParams) WithBackend(backend string) *ReplaceStickRuleParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the replace stick rule params
func (o *ReplaceStickRuleParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithData adds the data to the replace stick rule params
func (o *ReplaceStickRuleParams) WithData(data *models.StickRule) *ReplaceStickRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace stick rule params
func (o *ReplaceStickRuleParams) SetData(data *models.StickRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace stick rule params
func (o *ReplaceStickRuleParams) WithForceReload(forceReload *bool) *ReplaceStickRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace stick rule params
func (o *ReplaceStickRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithID adds the id to the replace stick rule params
func (o *ReplaceStickRuleParams) WithID(id int64) *ReplaceStickRuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the replace stick rule params
func (o *ReplaceStickRuleParams) SetID(id int64) {
	o.ID = id
}

// WithTransactionID adds the transactionID to the replace stick rule params
func (o *ReplaceStickRuleParams) WithTransactionID(transactionID *string) *ReplaceStickRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace stick rule params
func (o *ReplaceStickRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace stick rule params
func (o *ReplaceStickRuleParams) WithVersion(version *int64) *ReplaceStickRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace stick rule params
func (o *ReplaceStickRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceStickRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param backend
	qrBackend := o.Backend
	qBackend := qrBackend
	if qBackend != "" {
		if err := r.SetQueryParam("backend", qBackend); err != nil {
			return err
		}
	}

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if o.ForceReload != nil {

		// query param force_reload
		var qrForceReload bool
		if o.ForceReload != nil {
			qrForceReload = *o.ForceReload
		}
		qForceReload := swag.FormatBool(qrForceReload)
		if qForceReload != "" {
			if err := r.SetQueryParam("force_reload", qForceReload); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
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

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}