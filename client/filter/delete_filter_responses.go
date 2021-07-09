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

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// DeleteFilterReader is a Reader for the DeleteFilter structure.
type DeleteFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteFilterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeleteFilterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteFilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFilterAccepted creates a DeleteFilterAccepted with default headers values
func NewDeleteFilterAccepted() *DeleteFilterAccepted {
	return &DeleteFilterAccepted{}
}

/*DeleteFilterAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteFilterAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteFilterAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterAccepted ", 202)
}

func (o *DeleteFilterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteFilterNoContent creates a DeleteFilterNoContent with default headers values
func NewDeleteFilterNoContent() *DeleteFilterNoContent {
	return &DeleteFilterNoContent{}
}

/*DeleteFilterNoContent handles this case with default header values.

Filter deleted
*/
type DeleteFilterNoContent struct {
}

func (o *DeleteFilterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterNoContent ", 204)
}

func (o *DeleteFilterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFilterNotFound creates a DeleteFilterNotFound with default headers values
func NewDeleteFilterNotFound() *DeleteFilterNotFound {
	return &DeleteFilterNotFound{}
}

/*DeleteFilterNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteFilterNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteFilterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilterNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFilterDefault creates a DeleteFilterDefault with default headers values
func NewDeleteFilterDefault(code int) *DeleteFilterDefault {
	return &DeleteFilterDefault{
		_statusCode: code,
	}
}

/*DeleteFilterDefault handles this case with default header values.

General Error
*/
type DeleteFilterDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete filter default response
func (o *DeleteFilterDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFilterDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/filters/{index}][%d] deleteFilter default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
