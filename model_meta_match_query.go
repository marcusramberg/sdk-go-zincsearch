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

// MetaMatchQuery struct for MetaMatchQuery
type MetaMatchQuery struct {
	Analyzer *string `json:"analyzer,omitempty"`
	Boost *float32 `json:"boost,omitempty"`
	// auto, 1,2,3,n
	Fuzziness map[string]interface{} `json:"fuzziness,omitempty"`
	// or(default), and
	Operator *string `json:"operator,omitempty"`
	PrefixLength *float32 `json:"prefix_length,omitempty"`
	Query *string `json:"query,omitempty"`
}

// NewMetaMatchQuery instantiates a new MetaMatchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaMatchQuery() *MetaMatchQuery {
	this := MetaMatchQuery{}
	return &this
}

// NewMetaMatchQueryWithDefaults instantiates a new MetaMatchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaMatchQueryWithDefaults() *MetaMatchQuery {
	this := MetaMatchQuery{}
	return &this
}

// GetAnalyzer returns the Analyzer field value if set, zero value otherwise.
func (o *MetaMatchQuery) GetAnalyzer() string {
	if o == nil || o.Analyzer == nil {
		var ret string
		return ret
	}
	return *o.Analyzer
}

// GetAnalyzerOk returns a tuple with the Analyzer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMatchQuery) GetAnalyzerOk() (*string, bool) {
	if o == nil || o.Analyzer == nil {
		return nil, false
	}
	return o.Analyzer, true
}

// HasAnalyzer returns a boolean if a field has been set.
func (o *MetaMatchQuery) HasAnalyzer() bool {
	if o != nil && o.Analyzer != nil {
		return true
	}

	return false
}

// SetAnalyzer gets a reference to the given string and assigns it to the Analyzer field.
func (o *MetaMatchQuery) SetAnalyzer(v string) {
	o.Analyzer = &v
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *MetaMatchQuery) GetBoost() float32 {
	if o == nil || o.Boost == nil {
		var ret float32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMatchQuery) GetBoostOk() (*float32, bool) {
	if o == nil || o.Boost == nil {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *MetaMatchQuery) HasBoost() bool {
	if o != nil && o.Boost != nil {
		return true
	}

	return false
}

// SetBoost gets a reference to the given float32 and assigns it to the Boost field.
func (o *MetaMatchQuery) SetBoost(v float32) {
	o.Boost = &v
}

// GetFuzziness returns the Fuzziness field value if set, zero value otherwise.
func (o *MetaMatchQuery) GetFuzziness() map[string]interface{} {
	if o == nil || o.Fuzziness == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Fuzziness
}

// GetFuzzinessOk returns a tuple with the Fuzziness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMatchQuery) GetFuzzinessOk() (map[string]interface{}, bool) {
	if o == nil || o.Fuzziness == nil {
		return nil, false
	}
	return o.Fuzziness, true
}

// HasFuzziness returns a boolean if a field has been set.
func (o *MetaMatchQuery) HasFuzziness() bool {
	if o != nil && o.Fuzziness != nil {
		return true
	}

	return false
}

// SetFuzziness gets a reference to the given map[string]interface{} and assigns it to the Fuzziness field.
func (o *MetaMatchQuery) SetFuzziness(v map[string]interface{}) {
	o.Fuzziness = v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *MetaMatchQuery) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMatchQuery) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *MetaMatchQuery) HasOperator() bool {
	if o != nil && o.Operator != nil {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *MetaMatchQuery) SetOperator(v string) {
	o.Operator = &v
}

// GetPrefixLength returns the PrefixLength field value if set, zero value otherwise.
func (o *MetaMatchQuery) GetPrefixLength() float32 {
	if o == nil || o.PrefixLength == nil {
		var ret float32
		return ret
	}
	return *o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMatchQuery) GetPrefixLengthOk() (*float32, bool) {
	if o == nil || o.PrefixLength == nil {
		return nil, false
	}
	return o.PrefixLength, true
}

// HasPrefixLength returns a boolean if a field has been set.
func (o *MetaMatchQuery) HasPrefixLength() bool {
	if o != nil && o.PrefixLength != nil {
		return true
	}

	return false
}

// SetPrefixLength gets a reference to the given float32 and assigns it to the PrefixLength field.
func (o *MetaMatchQuery) SetPrefixLength(v float32) {
	o.PrefixLength = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MetaMatchQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMatchQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MetaMatchQuery) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *MetaMatchQuery) SetQuery(v string) {
	o.Query = &v
}

func (o MetaMatchQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Analyzer != nil {
		toSerialize["analyzer"] = o.Analyzer
	}
	if o.Boost != nil {
		toSerialize["boost"] = o.Boost
	}
	if o.Fuzziness != nil {
		toSerialize["fuzziness"] = o.Fuzziness
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.PrefixLength != nil {
		toSerialize["prefix_length"] = o.PrefixLength
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableMetaMatchQuery struct {
	value *MetaMatchQuery
	isSet bool
}

func (v NullableMetaMatchQuery) Get() *MetaMatchQuery {
	return v.value
}

func (v *NullableMetaMatchQuery) Set(val *MetaMatchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaMatchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaMatchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaMatchQuery(val *MetaMatchQuery) *NullableMetaMatchQuery {
	return &NullableMetaMatchQuery{value: val, isSet: true}
}

func (v NullableMetaMatchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaMatchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


