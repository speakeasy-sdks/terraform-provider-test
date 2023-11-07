// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"AcmeTerraform/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type SearchUsersv1Response struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Users *shared.Users
}

func (o *SearchUsersv1Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SearchUsersv1Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SearchUsersv1Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SearchUsersv1Response) GetUsers() *shared.Users {
	if o == nil {
		return nil
	}
	return o.Users
}
