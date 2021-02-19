// Code generated by go-swagger; DO NOT EDIT.

package change_password

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/JCbayfisher/swaggerTesting/models"
)

// NewPostChangePasswordParams creates a new PostChangePasswordParams object
//
// There are no default values defined in the spec.
func NewPostChangePasswordParams() PostChangePasswordParams {

	return PostChangePasswordParams{}
}

// PostChangePasswordParams contains all the bound params for the post change password operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostChangePassword
type PostChangePasswordParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	ChangePassword *models.ChangePassword
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostChangePasswordParams() beforehand.
func (o *PostChangePasswordParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ChangePassword
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("changePassword", "body", ""))
			} else {
				res = append(res, errors.NewParseError("changePassword", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ChangePassword = &body
			}
		}
	} else {
		res = append(res, errors.Required("changePassword", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
