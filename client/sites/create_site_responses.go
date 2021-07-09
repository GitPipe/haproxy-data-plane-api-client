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

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// CreateSiteReader is a Reader for the CreateSite structure.
type CreateSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSiteCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewCreateSiteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSiteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateSiteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSiteCreated creates a CreateSiteCreated with default headers values
func NewCreateSiteCreated() *CreateSiteCreated {
	return &CreateSiteCreated{}
}

/*CreateSiteCreated handles this case with default header values.

Site created
*/
type CreateSiteCreated struct {
	Payload *models.Site
}

func (o *CreateSiteCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/sites][%d] createSiteCreated  %+v", 201, o.Payload)
}

func (o *CreateSiteCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteAccepted creates a CreateSiteAccepted with default headers values
func NewCreateSiteAccepted() *CreateSiteAccepted {
	return &CreateSiteAccepted{}
}

/*CreateSiteAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type CreateSiteAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.Site
}

func (o *CreateSiteAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/sites][%d] createSiteAccepted  %+v", 202, o.Payload)
}

func (o *CreateSiteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteBadRequest creates a CreateSiteBadRequest with default headers values
func NewCreateSiteBadRequest() *CreateSiteBadRequest {
	return &CreateSiteBadRequest{}
}

/*CreateSiteBadRequest handles this case with default header values.

Bad request
*/
type CreateSiteBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSiteBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/sites][%d] createSiteBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSiteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteConflict creates a CreateSiteConflict with default headers values
func NewCreateSiteConflict() *CreateSiteConflict {
	return &CreateSiteConflict{}
}

/*CreateSiteConflict handles this case with default header values.

The specified resource already exists
*/
type CreateSiteConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSiteConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/sites][%d] createSiteConflict  %+v", 409, o.Payload)
}

func (o *CreateSiteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteDefault creates a CreateSiteDefault with default headers values
func NewCreateSiteDefault(code int) *CreateSiteDefault {
	return &CreateSiteDefault{
		_statusCode: code,
	}
}

/*CreateSiteDefault handles this case with default header values.

General Error
*/
type CreateSiteDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create site default response
func (o *CreateSiteDefault) Code() int {
	return o._statusCode
}

func (o *CreateSiteDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/sites][%d] createSite default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
