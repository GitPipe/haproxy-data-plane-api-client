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

// DeleteSpoeScopeReader is a Reader for the DeleteSpoeScope structure.
type DeleteSpoeScopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeScopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSpoeScopeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteSpoeScopeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSpoeScopeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeScopeNoContent creates a DeleteSpoeScopeNoContent with default headers values
func NewDeleteSpoeScopeNoContent() *DeleteSpoeScopeNoContent {
	return &DeleteSpoeScopeNoContent{}
}

/*DeleteSpoeScopeNoContent handles this case with default header values.

Spoe scope deleted
*/
type DeleteSpoeScopeNoContent struct {
}

func (o *DeleteSpoeScopeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_scopes/{name}][%d] deleteSpoeScopeNoContent ", 204)
}

func (o *DeleteSpoeScopeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeScopeNotFound creates a DeleteSpoeScopeNotFound with default headers values
func NewDeleteSpoeScopeNotFound() *DeleteSpoeScopeNotFound {
	return &DeleteSpoeScopeNotFound{}
}

/*DeleteSpoeScopeNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteSpoeScopeNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteSpoeScopeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_scopes/{name}][%d] deleteSpoeScopeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeScopeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSpoeScopeDefault creates a DeleteSpoeScopeDefault with default headers values
func NewDeleteSpoeScopeDefault(code int) *DeleteSpoeScopeDefault {
	return &DeleteSpoeScopeDefault{
		_statusCode: code,
	}
}

/*DeleteSpoeScopeDefault handles this case with default header values.

General Error
*/
type DeleteSpoeScopeDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete spoe scope default response
func (o *DeleteSpoeScopeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSpoeScopeDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_scopes/{name}][%d] deleteSpoeScope default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeScopeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
