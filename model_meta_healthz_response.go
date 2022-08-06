/*
Zinc Search engine API

Zinc Search engine API documents https://docs.zincsearch.com

API version: 0.2.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// MetaHealthzResponse struct for MetaHealthzResponse
type MetaHealthzResponse struct {
	Status *string `json:"status,omitempty"`
}

// NewMetaHealthzResponse instantiates a new MetaHealthzResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaHealthzResponse() *MetaHealthzResponse {
	this := MetaHealthzResponse{}
	return &this
}

// NewMetaHealthzResponseWithDefaults instantiates a new MetaHealthzResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaHealthzResponseWithDefaults() *MetaHealthzResponse {
	this := MetaHealthzResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MetaHealthzResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHealthzResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MetaHealthzResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MetaHealthzResponse) SetStatus(v string) {
	o.Status = &v
}

func (o MetaHealthzResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMetaHealthzResponse struct {
	value *MetaHealthzResponse
	isSet bool
}

func (v NullableMetaHealthzResponse) Get() *MetaHealthzResponse {
	return v.value
}

func (v *NullableMetaHealthzResponse) Set(val *MetaHealthzResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaHealthzResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaHealthzResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaHealthzResponse(val *MetaHealthzResponse) *NullableMetaHealthzResponse {
	return &NullableMetaHealthzResponse{value: val, isSet: true}
}

func (v NullableMetaHealthzResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaHealthzResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


