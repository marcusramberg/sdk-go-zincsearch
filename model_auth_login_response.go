/*
Zinc Search engine API

Zinc Search engine API documents https://docs.zincsearch.com

API version: 0.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AuthLoginResponse struct for AuthLoginResponse
type AuthLoginResponse struct {
	User *AuthLoginUser `json:"user,omitempty"`
	Validated *bool `json:"validated,omitempty"`
}

// NewAuthLoginResponse instantiates a new AuthLoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthLoginResponse() *AuthLoginResponse {
	this := AuthLoginResponse{}
	return &this
}

// NewAuthLoginResponseWithDefaults instantiates a new AuthLoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthLoginResponseWithDefaults() *AuthLoginResponse {
	this := AuthLoginResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AuthLoginResponse) GetUser() AuthLoginUser {
	if o == nil || o.User == nil {
		var ret AuthLoginUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLoginResponse) GetUserOk() (*AuthLoginUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AuthLoginResponse) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given AuthLoginUser and assigns it to the User field.
func (o *AuthLoginResponse) SetUser(v AuthLoginUser) {
	o.User = &v
}

// GetValidated returns the Validated field value if set, zero value otherwise.
func (o *AuthLoginResponse) GetValidated() bool {
	if o == nil || o.Validated == nil {
		var ret bool
		return ret
	}
	return *o.Validated
}

// GetValidatedOk returns a tuple with the Validated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLoginResponse) GetValidatedOk() (*bool, bool) {
	if o == nil || o.Validated == nil {
		return nil, false
	}
	return o.Validated, true
}

// HasValidated returns a boolean if a field has been set.
func (o *AuthLoginResponse) HasValidated() bool {
	if o != nil && o.Validated != nil {
		return true
	}

	return false
}

// SetValidated gets a reference to the given bool and assigns it to the Validated field.
func (o *AuthLoginResponse) SetValidated(v bool) {
	o.Validated = &v
}

func (o AuthLoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Validated != nil {
		toSerialize["validated"] = o.Validated
	}
	return json.Marshal(toSerialize)
}

type NullableAuthLoginResponse struct {
	value *AuthLoginResponse
	isSet bool
}

func (v NullableAuthLoginResponse) Get() *AuthLoginResponse {
	return v.value
}

func (v *NullableAuthLoginResponse) Set(val *AuthLoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthLoginResponse(val *AuthLoginResponse) *NullableAuthLoginResponse {
	return &NullableAuthLoginResponse{value: val, isSet: true}
}

func (v NullableAuthLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


