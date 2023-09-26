// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"AcmeTerraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetUserv1Request struct {
	// Numeric ID of the user to get
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetUserv1Response struct {
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
