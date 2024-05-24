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

// HeaderResponseAdditional struct for HeaderResponseAdditional
type HeaderResponseAdditional struct {
	// Don't add the header if it is added already. Only applies to 'set' action. Numerical value (\"0\" = false, \"1\" = true)
	IgnoreIfSet *string `json:"ignore_if_set,omitempty"`
	// Priority determines execution order. Lower numbers execute first.
	Priority *string `json:"priority,omitempty"`
	AdditionalProperties map[string]any
}

type _HeaderResponseAdditional HeaderResponseAdditional

// NewHeaderResponseAdditional instantiates a new HeaderResponseAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderResponseAdditional() *HeaderResponseAdditional {
	this := HeaderResponseAdditional{}
	var priority string = "100"
	this.Priority = &priority
	return &this
}

// NewHeaderResponseAdditionalWithDefaults instantiates a new HeaderResponseAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderResponseAdditionalWithDefaults() *HeaderResponseAdditional {
	this := HeaderResponseAdditional{}
	var priority string = "100"
	this.Priority = &priority
	return &this
}

// GetIgnoreIfSet returns the IgnoreIfSet field value if set, zero value otherwise.
func (o *HeaderResponseAdditional) GetIgnoreIfSet() string {
	if o == nil || o.IgnoreIfSet == nil {
		var ret string
		return ret
	}
	return *o.IgnoreIfSet
}

// GetIgnoreIfSetOk returns a tuple with the IgnoreIfSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponseAdditional) GetIgnoreIfSetOk() (*string, bool) {
	if o == nil || o.IgnoreIfSet == nil {
		return nil, false
	}
	return o.IgnoreIfSet, true
}

// HasIgnoreIfSet returns a boolean if a field has been set.
func (o *HeaderResponseAdditional) HasIgnoreIfSet() bool {
	if o != nil && o.IgnoreIfSet != nil {
		return true
	}

	return false
}

// SetIgnoreIfSet gets a reference to the given string and assigns it to the IgnoreIfSet field.
func (o *HeaderResponseAdditional) SetIgnoreIfSet(v string) {
	o.IgnoreIfSet = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *HeaderResponseAdditional) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderResponseAdditional) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *HeaderResponseAdditional) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *HeaderResponseAdditional) SetPriority(v string) {
	o.Priority = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HeaderResponseAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.IgnoreIfSet != nil {
		toSerialize["ignore_if_set"] = o.IgnoreIfSet
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *HeaderResponseAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varHeaderResponseAdditional := _HeaderResponseAdditional{}

	if err = json.Unmarshal(bytes, &varHeaderResponseAdditional); err == nil {
		*o = HeaderResponseAdditional(varHeaderResponseAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ignore_if_set")
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHeaderResponseAdditional is a helper abstraction for handling nullable headerresponseadditional types. 
type NullableHeaderResponseAdditional struct {
	value *HeaderResponseAdditional
	isSet bool
}

// Get returns the value.
func (v NullableHeaderResponseAdditional) Get() *HeaderResponseAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableHeaderResponseAdditional) Set(val *HeaderResponseAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHeaderResponseAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHeaderResponseAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHeaderResponseAdditional returns a pointer to a new instance of NullableHeaderResponseAdditional.
func NewNullableHeaderResponseAdditional(val *HeaderResponseAdditional) *NullableHeaderResponseAdditional {
	return &NullableHeaderResponseAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHeaderResponseAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHeaderResponseAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
