// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateOAuthUser Create a user from the OAuth 'social_type'
*/
func (a *Client) CreateOAuthUser(params *CreateOAuthUserParams) (*CreateOAuthUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOAuthUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOAuthUser",
		Method:             "POST",
		PathPattern:        "/user/{social_type}/oauth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOAuthUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateOAuthUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateOAuthUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteUserAccount The authenticated user delete itself from the system.
*/
func (a *Client) DeleteUserAccount(params *DeleteUserAccountParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserAccountNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUserAccount",
		Method:             "DELETE",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUserAccountNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteUserAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EmailVerification Verify that the email address is own by the user.
*/
func (a *Client) EmailVerification(params *EmailVerificationParams) (*EmailVerificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmailVerificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "emailVerification",
		Method:             "PUT",
		PathPattern:        "/user/email/verification/{verification_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmailVerificationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmailVerificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EmailVerificationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EmailVerificationResend Re-send the verification user's email to the indicated address.
*/
func (a *Client) EmailVerificationResend(params *EmailVerificationResendParams) (*EmailVerificationResendNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEmailVerificationResendParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "emailVerificationResend",
		Method:             "POST",
		PathPattern:        "/user/email/verification",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EmailVerificationResendReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EmailVerificationResendNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EmailVerificationResendDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetOAuthUser Used to know if a user from the platform exists on that 'social_type'. If it exists we'll return the JWT 'token', if it does not we'll return the data of that user on the 'user' so it can be confirmed and created
*/
func (a *Client) GetOAuthUser(params *GetOAuthUserParams) (*GetOAuthUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOAuthUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOAuthUser",
		Method:             "GET",
		PathPattern:        "/user/{social_type}/oauth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOAuthUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOAuthUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOAuthUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetUserAccount Get the information of the account of the authenticated user.
*/
func (a *Client) GetUserAccount(params *GetUserAccountParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserAccount",
		Method:             "GET",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetUserAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Login Authenticate a user and return a new JWT token.
*/
func (a *Client) Login(params *LoginParams) (*LoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "login",
		Method:             "POST",
		PathPattern:        "/user/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LoginDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PasswordResetReq Request to reset the password. Due to security reasons, this endpoint doesn't return Not Found (404) when the email doesn't exist or belongs to a user primary email.
*/
func (a *Client) PasswordResetReq(params *PasswordResetReqParams) (*PasswordResetReqNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPasswordResetReqParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "passwordResetReq",
		Method:             "POST",
		PathPattern:        "/user/reset_password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PasswordResetReqReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PasswordResetReqNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PasswordResetReqDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PasswordResetUpdate Reset the user password when it has been forgotten. Due to security reasons, the endpoint doesn't return a Unprocessable Entity (422) when the token is invalid. 404 Status code is returned if the user has been deleted of the system between the user password request and this request.
*/
func (a *Client) PasswordResetUpdate(params *PasswordResetUpdateParams) (*PasswordResetUpdateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPasswordResetUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "passwordResetUpdate",
		Method:             "PUT",
		PathPattern:        "/user/reset_password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PasswordResetUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PasswordResetUpdateNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PasswordResetUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RefreshToken Refresh the user JWT and returns a new one if the previous is valid. The 'organization_canonical_query' has to be of an organization in which the user belongs to, and the 'child_canonical_query' of a child of the 'organization_canonical_query' in any level (could be of a grand child).
*/
func (a *Client) RefreshToken(params *RefreshTokenParams, authInfo runtime.ClientAuthInfoWriter) (*RefreshTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "refreshToken",
		Method:             "GET",
		PathPattern:        "/user/refresh_token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RefreshTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RefreshTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RefreshTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SignUp Create a new User (sign-up).
*/
func (a *Client) SignUp(params *SignUpParams) (*SignUpNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSignUpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "signUp",
		Method:             "POST",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SignUpReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SignUpNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SignUpDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateUserAccount Update the information of the account of the authenticated user.
*/
func (a *Client) UpdateUserAccount(params *UpdateUserAccountParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserAccount",
		Method:             "PUT",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateUserAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateUserGuide Update user's guide progress.
*/
func (a *Client) UpdateUserGuide(params *UpdateUserGuideParams) (*UpdateUserGuideNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserGuideParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserGuide",
		Method:             "PUT",
		PathPattern:        "/user/guide",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/vnd.cycloid.io.v1+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserGuideReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserGuideNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateUserGuideDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
