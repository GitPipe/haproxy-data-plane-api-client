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

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// ClearRuntimeMapReader is a Reader for the ClearRuntimeMap structure.
type ClearRuntimeMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClearRuntimeMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewClearRuntimeMapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewClearRuntimeMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewClearRuntimeMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClearRuntimeMapNoContent creates a ClearRuntimeMapNoContent with default headers values
func NewClearRuntimeMapNoContent() *ClearRuntimeMapNoContent {
	return &ClearRuntimeMapNoContent{}
}

/*ClearRuntimeMapNoContent handles this case with default header values.

All map entries deleted
*/
type ClearRuntimeMapNoContent struct {
}

func (o *ClearRuntimeMapNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/maps/{name}][%d] clearRuntimeMapNoContent ", 204)
}

func (o *ClearRuntimeMapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClearRuntimeMapNotFound creates a ClearRuntimeMapNotFound with default headers values
func NewClearRuntimeMapNotFound() *ClearRuntimeMapNotFound {
	return &ClearRuntimeMapNotFound{}
}

/*ClearRuntimeMapNotFound handles this case with default header values.

The specified resource was not found
*/
type ClearRuntimeMapNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ClearRuntimeMapNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/maps/{name}][%d] clearRuntimeMapNotFound  %+v", 404, o.Payload)
}

func (o *ClearRuntimeMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClearRuntimeMapDefault creates a ClearRuntimeMapDefault with default headers values
func NewClearRuntimeMapDefault(code int) *ClearRuntimeMapDefault {
	return &ClearRuntimeMapDefault{
		_statusCode: code,
	}
}

/*ClearRuntimeMapDefault handles this case with default header values.

General Error
*/
type ClearRuntimeMapDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the clear runtime map default response
func (o *ClearRuntimeMapDefault) Code() int {
	return o._statusCode
}

func (o *ClearRuntimeMapDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/runtime/maps/{name}][%d] clearRuntimeMap default  %+v", o._statusCode, o.Payload)
}

func (o *ClearRuntimeMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
