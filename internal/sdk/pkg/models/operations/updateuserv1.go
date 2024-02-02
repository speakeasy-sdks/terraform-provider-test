// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/vitor-test/terraform-provider-AcmeTerraform/v3/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateUserv1Request struct {
	User shared.UserInput `request:"mediaType=application/json"`
	// UserID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateUserv1Request) GetUser() shared.UserInput {
	if o == nil {
		return shared.UserInput{}
	}
	return o.User
}

func (o *UpdateUserv1Request) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateUserv1Response struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	User *shared.User
}

func (o *UpdateUserv1Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUserv1Response) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *UpdateUserv1Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUserv1Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateUserv1Response) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}
