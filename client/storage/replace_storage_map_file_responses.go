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

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/haproxy-data-plane-api-client/models"
)

// ReplaceStorageMapFileReader is a Reader for the ReplaceStorageMapFile structure.
type ReplaceStorageMapFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceStorageMapFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewReplaceStorageMapFileAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewReplaceStorageMapFileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReplaceStorageMapFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReplaceStorageMapFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewReplaceStorageMapFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceStorageMapFileAccepted creates a ReplaceStorageMapFileAccepted with default headers values
func NewReplaceStorageMapFileAccepted() *ReplaceStorageMapFileAccepted {
	return &ReplaceStorageMapFileAccepted{}
}

/*ReplaceStorageMapFileAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceStorageMapFileAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *ReplaceStorageMapFileAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/maps/{name}][%d] replaceStorageMapFileAccepted ", 202)
}

func (o *ReplaceStorageMapFileAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewReplaceStorageMapFileNoContent creates a ReplaceStorageMapFileNoContent with default headers values
func NewReplaceStorageMapFileNoContent() *ReplaceStorageMapFileNoContent {
	return &ReplaceStorageMapFileNoContent{}
}

/*ReplaceStorageMapFileNoContent handles this case with default header values.

Map file replaced
*/
type ReplaceStorageMapFileNoContent struct {
}

func (o *ReplaceStorageMapFileNoContent) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/maps/{name}][%d] replaceStorageMapFileNoContent ", 204)
}

func (o *ReplaceStorageMapFileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceStorageMapFileBadRequest creates a ReplaceStorageMapFileBadRequest with default headers values
func NewReplaceStorageMapFileBadRequest() *ReplaceStorageMapFileBadRequest {
	return &ReplaceStorageMapFileBadRequest{}
}

/*ReplaceStorageMapFileBadRequest handles this case with default header values.

Bad request
*/
type ReplaceStorageMapFileBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceStorageMapFileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/maps/{name}][%d] replaceStorageMapFileBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceStorageMapFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStorageMapFileNotFound creates a ReplaceStorageMapFileNotFound with default headers values
func NewReplaceStorageMapFileNotFound() *ReplaceStorageMapFileNotFound {
	return &ReplaceStorageMapFileNotFound{}
}

/*ReplaceStorageMapFileNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceStorageMapFileNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceStorageMapFileNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/maps/{name}][%d] replaceStorageMapFileNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceStorageMapFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStorageMapFileDefault creates a ReplaceStorageMapFileDefault with default headers values
func NewReplaceStorageMapFileDefault(code int) *ReplaceStorageMapFileDefault {
	return &ReplaceStorageMapFileDefault{
		_statusCode: code,
	}
}

/*ReplaceStorageMapFileDefault handles this case with default header values.

General Error
*/
type ReplaceStorageMapFileDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace storage map file default response
func (o *ReplaceStorageMapFileDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceStorageMapFileDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/storage/maps/{name}][%d] replaceStorageMapFile default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceStorageMapFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	o.ConfigurationVersion = response.GetHeader("Configuration-Version")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
