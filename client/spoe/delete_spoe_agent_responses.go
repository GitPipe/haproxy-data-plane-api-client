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

// DeleteSpoeAgentReader is a Reader for the DeleteSpoeAgent structure.
type DeleteSpoeAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSpoeAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteSpoeAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSpoeAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeAgentNoContent creates a DeleteSpoeAgentNoContent with default headers values
func NewDeleteSpoeAgentNoContent() *DeleteSpoeAgentNoContent {
	return &DeleteSpoeAgentNoContent{}
}

/*DeleteSpoeAgentNoContent handles this case with default header values.

Spoe agent deleted
*/
type DeleteSpoeAgentNoContent struct {
}

func (o *DeleteSpoeAgentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_agents/{name}][%d] deleteSpoeAgentNoContent ", 204)
}

func (o *DeleteSpoeAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeAgentNotFound creates a DeleteSpoeAgentNotFound with default headers values
func NewDeleteSpoeAgentNotFound() *DeleteSpoeAgentNotFound {
	return &DeleteSpoeAgentNotFound{}
}

/*DeleteSpoeAgentNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteSpoeAgentNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteSpoeAgentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_agents/{name}][%d] deleteSpoeAgentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSpoeAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSpoeAgentDefault creates a DeleteSpoeAgentDefault with default headers values
func NewDeleteSpoeAgentDefault(code int) *DeleteSpoeAgentDefault {
	return &DeleteSpoeAgentDefault{
		_statusCode: code,
	}
}

/*DeleteSpoeAgentDefault handles this case with default header values.

General Error
*/
type DeleteSpoeAgentDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete spoe agent default response
func (o *DeleteSpoeAgentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSpoeAgentDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_agents/{name}][%d] deleteSpoeAgent default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSpoeAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
