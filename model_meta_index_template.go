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

// MetaIndexTemplate struct for MetaIndexTemplate
type MetaIndexTemplate struct {
	CreatedAt *string `json:"created_at,omitempty"`
	IndexPatterns []string `json:"index_patterns,omitempty"`
	// highest priority is chosen
	Priority *int32 `json:"priority,omitempty"`
	Template *MetaTemplateTemplate `json:"template,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewMetaIndexTemplate instantiates a new MetaIndexTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaIndexTemplate() *MetaIndexTemplate {
	this := MetaIndexTemplate{}
	return &this
}

// NewMetaIndexTemplateWithDefaults instantiates a new MetaIndexTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaIndexTemplateWithDefaults() *MetaIndexTemplate {
	this := MetaIndexTemplate{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MetaIndexTemplate) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexTemplate) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MetaIndexTemplate) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *MetaIndexTemplate) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetIndexPatterns returns the IndexPatterns field value if set, zero value otherwise.
func (o *MetaIndexTemplate) GetIndexPatterns() []string {
	if o == nil || o.IndexPatterns == nil {
		var ret []string
		return ret
	}
	return o.IndexPatterns
}

// GetIndexPatternsOk returns a tuple with the IndexPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexTemplate) GetIndexPatternsOk() ([]string, bool) {
	if o == nil || o.IndexPatterns == nil {
		return nil, false
	}
	return o.IndexPatterns, true
}

// HasIndexPatterns returns a boolean if a field has been set.
func (o *MetaIndexTemplate) HasIndexPatterns() bool {
	if o != nil && o.IndexPatterns != nil {
		return true
	}

	return false
}

// SetIndexPatterns gets a reference to the given []string and assigns it to the IndexPatterns field.
func (o *MetaIndexTemplate) SetIndexPatterns(v []string) {
	o.IndexPatterns = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *MetaIndexTemplate) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexTemplate) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *MetaIndexTemplate) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *MetaIndexTemplate) SetPriority(v int32) {
	o.Priority = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *MetaIndexTemplate) GetTemplate() MetaTemplateTemplate {
	if o == nil || o.Template == nil {
		var ret MetaTemplateTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexTemplate) GetTemplateOk() (*MetaTemplateTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *MetaIndexTemplate) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given MetaTemplateTemplate and assigns it to the Template field.
func (o *MetaIndexTemplate) SetTemplate(v MetaTemplateTemplate) {
	o.Template = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MetaIndexTemplate) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaIndexTemplate) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MetaIndexTemplate) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MetaIndexTemplate) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o MetaIndexTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.IndexPatterns != nil {
		toSerialize["index_patterns"] = o.IndexPatterns
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableMetaIndexTemplate struct {
	value *MetaIndexTemplate
	isSet bool
}

func (v NullableMetaIndexTemplate) Get() *MetaIndexTemplate {
	return v.value
}

func (v *NullableMetaIndexTemplate) Set(val *MetaIndexTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaIndexTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaIndexTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaIndexTemplate(val *MetaIndexTemplate) *NullableMetaIndexTemplate {
	return &NullableMetaIndexTemplate{value: val, isSet: true}
}

func (v NullableMetaIndexTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaIndexTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


