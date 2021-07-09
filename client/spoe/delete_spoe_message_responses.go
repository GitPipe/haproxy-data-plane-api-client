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

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// DeleteSpoeMessageReader is a Reader for the DeleteSpoeMessage structure.
type DeleteSpoeMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSpoeMessageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteSpoeMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSpoeMessageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeMessageNoContent creates a DeleteSpoeMessageNoContent with default headers values
func NewDeleteSpoeMessageNoContent() *DeleteSpoeMessageNoContent {
	return &DeleteSpoeMessageNoContent{}
}

/*DeleteSpoeMessageNoContent handles this case with default header values.

Spoe message deleted
*/
type DeleteSpoeMessageNoContent struct {
}

func (o *DeleteSpoeMessageNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_messages/{name}][%d] deleteSpoeMessageNoContent ", 204)
}

func (o *DeleteSpoeMessageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeMessageNotFound creates a DeleteSpoeMessageNotFound with default headers values
func NewDeleteSpoeMessageNotFound() *DeleteSpoeMessageNotFound {
	return &DeleteSpoeMessageNotFound{}
}

/*DeleteSpoeMessageNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteSpoeMessageNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteSpoeMessageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_messages/{name}][%d] deleteSpoeMessageNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSpoeMessageDefault creates a DeleteSpoeMessageDefault with default headers values
func NewDeleteSpoeMessageDefault(code int) *DeleteSpoeMessageDefault {
	return &DeleteSpoeMessageDefault{
		_statusCode: code,
	}
}

/*DeleteSpoeMessageDefault handles this case with default header values.

General Error
*/
type DeleteSpoeMessageDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete spoe message default response
func (o *DeleteSpoeMessageDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSpoeMessageDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_messages/{name}][%d] deleteSpoeMessage default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeMessageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
