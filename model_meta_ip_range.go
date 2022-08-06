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

// MetaIPRange struct for MetaIPRange
type MetaIPRange struct {
	From *string `json:"from,omitempty"`
	To *string `json:"to,omitempty"`
}

// NewMetaIPRange instantiates a new MetaIPRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaIPRange() *MetaIPRange {
	this := MetaIPRange{}
	return &this
}

// NewMetaIPRangeWithDefaults instantiates a new MetaIPRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaIPRangeWithDefaults() *MetaIPRange {
	this := MetaIPRange{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *MetaIPRange) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIPRange) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *MetaIPRange) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *MetaIPRange) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *MetaIPRange) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIPRange) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *MetaIPRange) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *MetaIPRange) SetTo(v string) {
	o.To = &v
}

func (o MetaIPRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableMetaIPRange struct {
	value *MetaIPRange
	isSet bool
}

func (v NullableMetaIPRange) Get() *MetaIPRange {
	return v.value
}

func (v *NullableMetaIPRange) Set(val *MetaIPRange) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaIPRange) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaIPRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaIPRange(val *MetaIPRange) *NullableMetaIPRange {
	return &NullableMetaIPRange{value: val, isSet: true}
}

func (v NullableMetaIPRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaIPRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


