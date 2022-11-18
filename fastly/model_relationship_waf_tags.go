// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// RelationshipWafTags struct for RelationshipWafTags
type RelationshipWafTags struct {
	WafTags *RelationshipWafTagsWafTags `json:"waf_tags,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipWafTags RelationshipWafTags

// NewRelationshipWafTags instantiates a new RelationshipWafTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWafTags() *RelationshipWafTags {
	this := RelationshipWafTags{}
	return &this
}

// NewRelationshipWafTagsWithDefaults instantiates a new RelationshipWafTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWafTagsWithDefaults() *RelationshipWafTags {
	this := RelationshipWafTags{}
	return &this
}

// GetWafTags returns the WafTags field value if set, zero value otherwise.
func (o *RelationshipWafTags) GetWafTags() RelationshipWafTagsWafTags {
	if o == nil || o.WafTags == nil {
		var ret RelationshipWafTagsWafTags
		return ret
	}
	return *o.WafTags
}

// GetWafTagsOk returns a tuple with the WafTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWafTags) GetWafTagsOk() (*RelationshipWafTagsWafTags, bool) {
	if o == nil || o.WafTags == nil {
		return nil, false
	}
	return o.WafTags, true
}

// HasWafTags returns a boolean if a field has been set.
func (o *RelationshipWafTags) HasWafTags() bool {
	if o != nil && o.WafTags != nil {
		return true
	}

	return false
}

// SetWafTags gets a reference to the given RelationshipWafTagsWafTags and assigns it to the WafTags field.
func (o *RelationshipWafTags) SetWafTags(v RelationshipWafTagsWafTags) {
	o.WafTags = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipWafTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WafTags != nil {
		toSerialize["waf_tags"] = o.WafTags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipWafTags) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipWafTags := _RelationshipWafTags{}

	if err = json.Unmarshal(bytes, &varRelationshipWafTags); err == nil {
		*o = RelationshipWafTags(varRelationshipWafTags)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipWafTags is a helper abstraction for handling nullable relationshipwaftags types. 
type NullableRelationshipWafTags struct {
	value *RelationshipWafTags
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipWafTags) Get() *RelationshipWafTags {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipWafTags) Set(val *RelationshipWafTags) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipWafTags) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipWafTags) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipWafTags returns a pointer to a new instance of NullableRelationshipWafTags.
func NewNullableRelationshipWafTags(val *RelationshipWafTags) *NullableRelationshipWafTags {
	return &NullableRelationshipWafTags{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipWafTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipWafTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
