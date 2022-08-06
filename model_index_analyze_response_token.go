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

// IndexAnalyzeResponseToken struct for IndexAnalyzeResponseToken
type IndexAnalyzeResponseToken struct {
	EndOffset *int32 `json:"end_offset,omitempty"`
	Keyword *bool `json:"keyword,omitempty"`
	Position *int32 `json:"position,omitempty"`
	StartOffset *int32 `json:"start_offset,omitempty"`
	Token *string `json:"token,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewIndexAnalyzeResponseToken instantiates a new IndexAnalyzeResponseToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexAnalyzeResponseToken() *IndexAnalyzeResponseToken {
	this := IndexAnalyzeResponseToken{}
	return &this
}

// NewIndexAnalyzeResponseTokenWithDefaults instantiates a new IndexAnalyzeResponseToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexAnalyzeResponseTokenWithDefaults() *IndexAnalyzeResponseToken {
	this := IndexAnalyzeResponseToken{}
	return &this
}

// GetEndOffset returns the EndOffset field value if set, zero value otherwise.
func (o *IndexAnalyzeResponseToken) GetEndOffset() int32 {
	if o == nil || o.EndOffset == nil {
		var ret int32
		return ret
	}
	return *o.EndOffset
}

// GetEndOffsetOk returns a tuple with the EndOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexAnalyzeResponseToken) GetEndOffsetOk() (*int32, bool) {
	if o == nil || o.EndOffset == nil {
		return nil, false
	}
	return o.EndOffset, true
}

// HasEndOffset returns a boolean if a field has been set.
func (o *IndexAnalyzeResponseToken) HasEndOffset() bool {
	if o != nil && o.EndOffset != nil {
		return true
	}

	return false
}

// SetEndOffset gets a reference to the given int32 and assigns it to the EndOffset field.
func (o *IndexAnalyzeResponseToken) SetEndOffset(v int32) {
	o.EndOffset = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *IndexAnalyzeResponseToken) GetKeyword() bool {
	if o == nil || o.Keyword == nil {
		var ret bool
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexAnalyzeResponseToken) GetKeywordOk() (*bool, bool) {
	if o == nil || o.Keyword == nil {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *IndexAnalyzeResponseToken) HasKeyword() bool {
	if o != nil && o.Keyword != nil {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given bool and assigns it to the Keyword field.
func (o *IndexAnalyzeResponseToken) SetKeyword(v bool) {
	o.Keyword = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *IndexAnalyzeResponseToken) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexAnalyzeResponseToken) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *IndexAnalyzeResponseToken) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *IndexAnalyzeResponseToken) SetPosition(v int32) {
	o.Position = &v
}

// GetStartOffset returns the StartOffset field value if set, zero value otherwise.
func (o *IndexAnalyzeResponseToken) GetStartOffset() int32 {
	if o == nil || o.StartOffset == nil {
		var ret int32
		return ret
	}
	return *o.StartOffset
}

// GetStartOffsetOk returns a tuple with the StartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexAnalyzeResponseToken) GetStartOffsetOk() (*int32, bool) {
	if o == nil || o.StartOffset == nil {
		return nil, false
	}
	return o.StartOffset, true
}

// HasStartOffset returns a boolean if a field has been set.
func (o *IndexAnalyzeResponseToken) HasStartOffset() bool {
	if o != nil && o.StartOffset != nil {
		return true
	}

	return false
}

// SetStartOffset gets a reference to the given int32 and assigns it to the StartOffset field.
func (o *IndexAnalyzeResponseToken) SetStartOffset(v int32) {
	o.StartOffset = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *IndexAnalyzeResponseToken) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexAnalyzeResponseToken) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *IndexAnalyzeResponseToken) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *IndexAnalyzeResponseToken) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IndexAnalyzeResponseToken) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexAnalyzeResponseToken) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IndexAnalyzeResponseToken) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IndexAnalyzeResponseToken) SetType(v string) {
	o.Type = &v
}

func (o IndexAnalyzeResponseToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndOffset != nil {
		toSerialize["end_offset"] = o.EndOffset
	}
	if o.Keyword != nil {
		toSerialize["keyword"] = o.Keyword
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.StartOffset != nil {
		toSerialize["start_offset"] = o.StartOffset
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIndexAnalyzeResponseToken struct {
	value *IndexAnalyzeResponseToken
	isSet bool
}

func (v NullableIndexAnalyzeResponseToken) Get() *IndexAnalyzeResponseToken {
	return v.value
}

func (v *NullableIndexAnalyzeResponseToken) Set(val *IndexAnalyzeResponseToken) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexAnalyzeResponseToken) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexAnalyzeResponseToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexAnalyzeResponseToken(val *IndexAnalyzeResponseToken) *NullableIndexAnalyzeResponseToken {
	return &NullableIndexAnalyzeResponseToken{value: val, isSet: true}
}

func (v NullableIndexAnalyzeResponseToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexAnalyzeResponseToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


