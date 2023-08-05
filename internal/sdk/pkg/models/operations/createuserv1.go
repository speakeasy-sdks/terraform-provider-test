// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"AcmeTerraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateUserv1Response struct {
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// OK
	User *shared.User
}
