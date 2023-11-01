// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"AcmeTerraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateUserv1Response struct {
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

func (o *CreateUserv1Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUserv1Response) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *CreateUserv1Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUserv1Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUserv1Response) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}
