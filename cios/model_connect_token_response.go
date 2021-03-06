/*
 * Cios Openapi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cios

import (
	"encoding/json"
)

// ConnectTokenResponse struct for ConnectTokenResponse
type ConnectTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn int32 `json:"expires_in"`
	IdToken *string `json:"id_token,omitempty"`
	Scope string `json:"scope"`
}

// NewConnectTokenResponse instantiates a new ConnectTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectTokenResponse(accessToken string, tokenType string, refreshToken string, expiresIn int32, scope string, ) *ConnectTokenResponse {
	this := ConnectTokenResponse{}
	this.AccessToken = accessToken
	this.TokenType = tokenType
	this.RefreshToken = refreshToken
	this.ExpiresIn = expiresIn
	this.Scope = scope
	return &this
}

// NewConnectTokenResponseWithDefaults instantiates a new ConnectTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectTokenResponseWithDefaults() *ConnectTokenResponse {
	this := ConnectTokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *ConnectTokenResponse) GetAccessToken() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ConnectTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ConnectTokenResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetTokenType returns the TokenType field value
func (o *ConnectTokenResponse) GetTokenType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *ConnectTokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *ConnectTokenResponse) SetTokenType(v string) {
	o.TokenType = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *ConnectTokenResponse) GetRefreshToken() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *ConnectTokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *ConnectTokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *ConnectTokenResponse) GetExpiresIn() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *ConnectTokenResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *ConnectTokenResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *ConnectTokenResponse) GetIdToken() string {
	if o == nil || o.IdToken == nil {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectTokenResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *ConnectTokenResponse) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *ConnectTokenResponse) SetIdToken(v string) {
	o.IdToken = &v
}

// GetScope returns the Scope field value
func (o *ConnectTokenResponse) GetScope() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ConnectTokenResponse) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *ConnectTokenResponse) SetScope(v string) {
	o.Scope = v
}

func (o ConnectTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	if true {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if true {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.IdToken != nil {
		toSerialize["id_token"] = o.IdToken
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableConnectTokenResponse struct {
	value *ConnectTokenResponse
	isSet bool
}

func (v NullableConnectTokenResponse) Get() *ConnectTokenResponse {
	return v.value
}

func (v *NullableConnectTokenResponse) Set(val *ConnectTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectTokenResponse(val *ConnectTokenResponse) *NullableConnectTokenResponse {
	return &NullableConnectTokenResponse{value: val, isSet: true}
}

func (v NullableConnectTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


