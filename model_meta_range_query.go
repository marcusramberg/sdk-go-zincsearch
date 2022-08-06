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

// MetaRangeQuery struct for MetaRangeQuery
type MetaRangeQuery struct {
	Boost *float32 `json:"boost,omitempty"`
	// Date format used to convert date values in the query.
	Format *string `json:"format,omitempty"`
	// string, float64
	Gt *string `json:"gt,omitempty"`
	// string, float64
	Gte *string `json:"gte,omitempty"`
	// string, float64
	Lt *string `json:"lt,omitempty"`
	// string, float64
	Lte *string `json:"lte,omitempty"`
	// used to convert date values in the query to UTC.
	TimeZone *string `json:"time_zone,omitempty"`
}

// NewMetaRangeQuery instantiates a new MetaRangeQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaRangeQuery() *MetaRangeQuery {
	this := MetaRangeQuery{}
	return &this
}

// NewMetaRangeQueryWithDefaults instantiates a new MetaRangeQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaRangeQueryWithDefaults() *MetaRangeQuery {
	this := MetaRangeQuery{}
	return &this
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *MetaRangeQuery) GetBoost() float32 {
	if o == nil || o.Boost == nil {
		var ret float32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRangeQuery) GetBoostOk() (*float32, bool) {
	if o == nil || o.Boost == nil {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *MetaRangeQuery) HasBoost() bool {
	if o != nil && o.Boost != nil {
		return true
	}

	return false
}

// SetBoost gets a reference to the given float32 and assigns it to the Boost field.
func (o *MetaRangeQuery) SetBoost(v float32) {
	o.Boost = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *MetaRangeQuery) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRangeQuery) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MetaRangeQuery) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *MetaRangeQuery) SetFormat(v string) {
	o.Format = &v
}

// GetGt returns the Gt field value if set, zero value otherwise.
func (o *MetaRangeQuery) GetGt() string {
	if o == nil || o.Gt == nil {
		var ret string
		return ret
	}
	return *o.Gt
}

// GetGtOk returns a tuple with the Gt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRangeQuery) GetGtOk() (*string, bool) {
	if o == nil || o.Gt == nil {
		return nil, false
	}
	return o.Gt, true
}

// HasGt returns a boolean if a field has been set.
func (o *MetaRangeQuery) HasGt() bool {
	if o != nil && o.Gt != nil {
		return true
	}

	return false
}

// SetGt gets a reference to the given string and assigns it to the Gt field.
func (o *MetaRangeQuery) SetGt(v string) {
	o.Gt = &v
}

// GetGte returns the Gte field value if set, zero value otherwise.
func (o *MetaRangeQuery) GetGte() string {
	if o == nil || o.Gte == nil {
		var ret string
		return ret
	}
	return *o.Gte
}

// GetGteOk returns a tuple with the Gte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRangeQuery) GetGteOk() (*string, bool) {
	if o == nil || o.Gte == nil {
		return nil, false
	}
	return o.Gte, true
}

// HasGte returns a boolean if a field has been set.
func (o *MetaRangeQuery) HasGte() bool {
	if o != nil && o.Gte != nil {
		return true
	}

	return false
}

// SetGte gets a reference to the given string and assigns it to the Gte field.
func (o *MetaRangeQuery) SetGte(v string) {
	o.Gte = &v
}

// GetLt returns the Lt field value if set, zero value otherwise.
func (o *MetaRangeQuery) GetLt() string {
	if o == nil || o.Lt == nil {
		var ret string
		return ret
	}
	return *o.Lt
}

// GetLtOk returns a tuple with the Lt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRangeQuery) GetLtOk() (*string, bool) {
	if o == nil || o.Lt == nil {
		return nil, false
	}
	return o.Lt, true
}

// HasLt returns a boolean if a field has been set.
func (o *MetaRangeQuery) HasLt() bool {
	if o != nil && o.Lt != nil {
		return true
	}

	return false
}

// SetLt gets a reference to the given string and assigns it to the Lt field.
func (o *MetaRangeQuery) SetLt(v string) {
	o.Lt = &v
}

// GetLte returns the Lte field value if set, zero value otherwise.
func (o *MetaRangeQuery) GetLte() string {
	if o == nil || o.Lte == nil {
		var ret string
		return ret
	}
	return *o.Lte
}

// GetLteOk returns a tuple with the Lte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRangeQuery) GetLteOk() (*string, bool) {
	if o == nil || o.Lte == nil {
		return nil, false
	}
	return o.Lte, true
}

// HasLte returns a boolean if a field has been set.
func (o *MetaRangeQuery) HasLte() bool {
	if o != nil && o.Lte != nil {
		return true
	}

	return false
}

// SetLte gets a reference to the given string and assigns it to the Lte field.
func (o *MetaRangeQuery) SetLte(v string) {
	o.Lte = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *MetaRangeQuery) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRangeQuery) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *MetaRangeQuery) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *MetaRangeQuery) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o MetaRangeQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Boost != nil {
		toSerialize["boost"] = o.Boost
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Gt != nil {
		toSerialize["gt"] = o.Gt
	}
	if o.Gte != nil {
		toSerialize["gte"] = o.Gte
	}
	if o.Lt != nil {
		toSerialize["lt"] = o.Lt
	}
	if o.Lte != nil {
		toSerialize["lte"] = o.Lte
	}
	if o.TimeZone != nil {
		toSerialize["time_zone"] = o.TimeZone
	}
	return json.Marshal(toSerialize)
}

type NullableMetaRangeQuery struct {
	value *MetaRangeQuery
	isSet bool
}

func (v NullableMetaRangeQuery) Get() *MetaRangeQuery {
	return v.value
}

func (v *NullableMetaRangeQuery) Set(val *MetaRangeQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaRangeQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaRangeQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaRangeQuery(val *MetaRangeQuery) *NullableMetaRangeQuery {
	return &NullableMetaRangeQuery{value: val, isSet: true}
}

func (v NullableMetaRangeQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaRangeQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


