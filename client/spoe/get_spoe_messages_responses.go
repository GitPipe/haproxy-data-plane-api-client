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
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/models"
)

// GetSpoeMessagesReader is a Reader for the GetSpoeMessages structure.
type GetSpoeMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpoeMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpoeMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSpoeMessagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSpoeMessagesOK creates a GetSpoeMessagesOK with default headers values
func NewGetSpoeMessagesOK() *GetSpoeMessagesOK {
	return &GetSpoeMessagesOK{}
}

/*GetSpoeMessagesOK handles this case with default header values.

Successful operation
*/
type GetSpoeMessagesOK struct {
	/*Spoe configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetSpoeMessagesOKBody
}

func (o *GetSpoeMessagesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_messages][%d] getSpoeMessagesOK  %+v", 200, o.Payload)
}

func (o *GetSpoeMessagesOK) GetPayload() *GetSpoeMessagesOKBody {
	return o.Payload
}

func (o *GetSpoeMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetSpoeMessagesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpoeMessagesDefault creates a GetSpoeMessagesDefault with default headers values
func NewGetSpoeMessagesDefault(code int) *GetSpoeMessagesDefault {
	return &GetSpoeMessagesDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetSpoeMessagesDefault handles this case with default header values.

General Error
*/
type GetSpoeMessagesDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get spoe messages default response
func (o *GetSpoeMessagesDefault) Code() int {
	return o._statusCode
}

func (o *GetSpoeMessagesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/spoe/spoe_messages][%d] getSpoeMessages default  %+v", o._statusCode, o.Payload)
}

func (o *GetSpoeMessagesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSpoeMessagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSpoeMessagesOKBody get spoe messages o k body
swagger:model GetSpoeMessagesOKBody
*/
type GetSpoeMessagesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.SpoeMessages `json:"data"`
}

// Validate validates this get spoe messages o k body
func (o *GetSpoeMessagesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSpoeMessagesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getSpoeMessagesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getSpoeMessagesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSpoeMessagesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSpoeMessagesOKBody) UnmarshalBinary(b []byte) error {
	var res GetSpoeMessagesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
