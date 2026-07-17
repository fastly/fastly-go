// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/)

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.

import (
	"encoding/json"
)

// PathResponseAllOf struct for PathResponseAllOf
type PathResponseAllOf struct {
	// Alphanumeric string identifying the path. Stable across versions of the routing config.
	Id *string `json:"id,omitempty"`
	// The URL path pattern, beginning with `/`. Maximum 2048 characters.
	Path *string `json:"path,omitempty"`
	// HATEOAS links to related resources.
	Links                *map[string]string `json:"links,omitempty"`
	AdditionalProperties map[string]any
}

type _PathResponseAllOf PathResponseAllOf

// NewPathResponseAllOf instantiates a new PathResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathResponseAllOf() *PathResponseAllOf {
	this := PathResponseAllOf{}
	return &this
}

// NewPathResponseAllOfWithDefaults instantiates a new PathResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathResponseAllOfWithDefaults() *PathResponseAllOf {
	this := PathResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PathResponseAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PathResponseAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PathResponseAllOf) SetId(v string) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PathResponseAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathResponseAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PathResponseAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PathResponseAllOf) SetPath(v string) {
	o.Path = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PathResponseAllOf) GetLinks() map[string]string {
	if o == nil || o.Links == nil {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathResponseAllOf) GetLinksOk() (*map[string]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PathResponseAllOf) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *PathResponseAllOf) SetLinks(v map[string]string) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PathResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PathResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPathResponseAllOf := _PathResponseAllOf{}

	if err = json.Unmarshal(bytes, &varPathResponseAllOf); err == nil {
		*o = PathResponseAllOf(varPathResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "path")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePathResponseAllOf is a helper abstraction for handling nullable pathresponseallof types.
type NullablePathResponseAllOf struct {
	value *PathResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullablePathResponseAllOf) Get() *PathResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullablePathResponseAllOf) Set(val *PathResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePathResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePathResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePathResponseAllOf returns a pointer to a new instance of NullablePathResponseAllOf.
func NewNullablePathResponseAllOf(val *PathResponseAllOf) *NullablePathResponseAllOf {
	return &NullablePathResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePathResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePathResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
