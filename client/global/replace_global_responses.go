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

package global

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// ReplaceGlobalReader is a Reader for the ReplaceGlobal structure.
type ReplaceGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceGlobalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewReplaceGlobalAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReplaceGlobalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewReplaceGlobalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceGlobalOK creates a ReplaceGlobalOK with default headers values
func NewReplaceGlobalOK() *ReplaceGlobalOK {
	return &ReplaceGlobalOK{}
}

/*ReplaceGlobalOK handles this case with default header values.

Global replaced
*/
type ReplaceGlobalOK struct {
	Payload *models.Global
}

func (o *ReplaceGlobalOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalOK  %+v", 200, o.Payload)
}

func (o *ReplaceGlobalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Global)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceGlobalAccepted creates a ReplaceGlobalAccepted with default headers values
func NewReplaceGlobalAccepted() *ReplaceGlobalAccepted {
	return &ReplaceGlobalAccepted{}
}

/*ReplaceGlobalAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceGlobalAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.Global
}

func (o *ReplaceGlobalAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceGlobalAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.Global)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceGlobalBadRequest creates a ReplaceGlobalBadRequest with default headers values
func NewReplaceGlobalBadRequest() *ReplaceGlobalBadRequest {
	return &ReplaceGlobalBadRequest{}
}

/*ReplaceGlobalBadRequest handles this case with default header values.

Bad request
*/
type ReplaceGlobalBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceGlobalBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobalBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceGlobalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceGlobalDefault creates a ReplaceGlobalDefault with default headers values
func NewReplaceGlobalDefault(code int) *ReplaceGlobalDefault {
	return &ReplaceGlobalDefault{
		_statusCode: code,
	}
}

/*ReplaceGlobalDefault handles this case with default header values.

General Error
*/
type ReplaceGlobalDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace global default response
func (o *ReplaceGlobalDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceGlobalDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/global][%d] replaceGlobal default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceGlobalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
