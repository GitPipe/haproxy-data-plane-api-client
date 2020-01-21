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

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/haproxytech/models"
)

// DeleteSiteReader is a Reader for the DeleteSite structure.
type DeleteSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteSiteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteSiteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSiteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSiteAccepted creates a DeleteSiteAccepted with default headers values
func NewDeleteSiteAccepted() *DeleteSiteAccepted {
	return &DeleteSiteAccepted{}
}

/*DeleteSiteAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteSiteAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteSiteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/sites/{name}][%d] deleteSiteAccepted ", 202)
}

func (o *DeleteSiteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteSiteNoContent creates a DeleteSiteNoContent with default headers values
func NewDeleteSiteNoContent() *DeleteSiteNoContent {
	return &DeleteSiteNoContent{}
}

/*DeleteSiteNoContent handles this case with default header values.

Site deleted
*/
type DeleteSiteNoContent struct {
}

func (o *DeleteSiteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/sites/{name}][%d] deleteSiteNoContent ", 204)
}

func (o *DeleteSiteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSiteNotFound creates a DeleteSiteNotFound with default headers values
func NewDeleteSiteNotFound() *DeleteSiteNotFound {
	return &DeleteSiteNotFound{}
}

/*DeleteSiteNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteSiteNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeleteSiteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/sites/{name}][%d] deleteSiteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSiteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSiteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSiteDefault creates a DeleteSiteDefault with default headers values
func NewDeleteSiteDefault(code int) *DeleteSiteDefault {
	return &DeleteSiteDefault{
		_statusCode: code,
	}
}

/*DeleteSiteDefault handles this case with default header values.

General Error
*/
type DeleteSiteDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete site default response
func (o *DeleteSiteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSiteDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/sites/{name}][%d] deleteSite default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSiteDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
