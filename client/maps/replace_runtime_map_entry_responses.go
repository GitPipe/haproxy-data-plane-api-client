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

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// ReplaceRuntimeMapEntryReader is a Reader for the ReplaceRuntimeMapEntry structure.
type ReplaceRuntimeMapEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceRuntimeMapEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceRuntimeMapEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReplaceRuntimeMapEntryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReplaceRuntimeMapEntryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewReplaceRuntimeMapEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceRuntimeMapEntryOK creates a ReplaceRuntimeMapEntryOK with default headers values
func NewReplaceRuntimeMapEntryOK() *ReplaceRuntimeMapEntryOK {
	return &ReplaceRuntimeMapEntryOK{}
}

/*ReplaceRuntimeMapEntryOK handles this case with default header values.

Map value replaced
*/
type ReplaceRuntimeMapEntryOK struct {
	Payload *models.MapEntry
}

func (o *ReplaceRuntimeMapEntryOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps_entries/{id}][%d] replaceRuntimeMapEntryOK  %+v", 200, o.Payload)
}

func (o *ReplaceRuntimeMapEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MapEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRuntimeMapEntryBadRequest creates a ReplaceRuntimeMapEntryBadRequest with default headers values
func NewReplaceRuntimeMapEntryBadRequest() *ReplaceRuntimeMapEntryBadRequest {
	return &ReplaceRuntimeMapEntryBadRequest{}
}

/*ReplaceRuntimeMapEntryBadRequest handles this case with default header values.

Bad request
*/
type ReplaceRuntimeMapEntryBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceRuntimeMapEntryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps_entries/{id}][%d] replaceRuntimeMapEntryBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceRuntimeMapEntryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRuntimeMapEntryNotFound creates a ReplaceRuntimeMapEntryNotFound with default headers values
func NewReplaceRuntimeMapEntryNotFound() *ReplaceRuntimeMapEntryNotFound {
	return &ReplaceRuntimeMapEntryNotFound{}
}

/*ReplaceRuntimeMapEntryNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceRuntimeMapEntryNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceRuntimeMapEntryNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps_entries/{id}][%d] replaceRuntimeMapEntryNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceRuntimeMapEntryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRuntimeMapEntryDefault creates a ReplaceRuntimeMapEntryDefault with default headers values
func NewReplaceRuntimeMapEntryDefault(code int) *ReplaceRuntimeMapEntryDefault {
	return &ReplaceRuntimeMapEntryDefault{
		_statusCode: code,
	}
}

/*ReplaceRuntimeMapEntryDefault handles this case with default header values.

General Error
*/
type ReplaceRuntimeMapEntryDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace runtime map entry default response
func (o *ReplaceRuntimeMapEntryDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceRuntimeMapEntryDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/runtime/maps_entries/{id}][%d] replaceRuntimeMapEntry default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceRuntimeMapEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ReplaceRuntimeMapEntryBody replace runtime map entry body
swagger:model ReplaceRuntimeMapEntryBody
*/
type ReplaceRuntimeMapEntryBody struct {

	// Map value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this replace runtime map entry body
func (o *ReplaceRuntimeMapEntryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReplaceRuntimeMapEntryBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ReplaceRuntimeMapEntryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplaceRuntimeMapEntryBody) UnmarshalBinary(b []byte) error {
	var res ReplaceRuntimeMapEntryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
