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

// AiAcceleratorResponseLinksLinks struct for AiAcceleratorResponseLinksLinks
type AiAcceleratorResponseLinksLinks struct {
	// Location of resource
	Self                 *string `json:"self,omitempty"`
	AdditionalProperties map[string]any
}

type _AiAcceleratorResponseLinksLinks AiAcceleratorResponseLinksLinks

// NewAiAcceleratorResponseLinksLinks instantiates a new AiAcceleratorResponseLinksLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAcceleratorResponseLinksLinks() *AiAcceleratorResponseLinksLinks {
	this := AiAcceleratorResponseLinksLinks{}
	return &this
}

// NewAiAcceleratorResponseLinksLinksWithDefaults instantiates a new AiAcceleratorResponseLinksLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAcceleratorResponseLinksLinksWithDefaults() *AiAcceleratorResponseLinksLinks {
	this := AiAcceleratorResponseLinksLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AiAcceleratorResponseLinksLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAcceleratorResponseLinksLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AiAcceleratorResponseLinksLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *AiAcceleratorResponseLinksLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AiAcceleratorResponseLinksLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AiAcceleratorResponseLinksLinks) UnmarshalJSON(bytes []byte) (err error) {
	varAiAcceleratorResponseLinksLinks := _AiAcceleratorResponseLinksLinks{}

	if err = json.Unmarshal(bytes, &varAiAcceleratorResponseLinksLinks); err == nil {
		*o = AiAcceleratorResponseLinksLinks(varAiAcceleratorResponseLinksLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAiAcceleratorResponseLinksLinks is a helper abstraction for handling nullable aiacceleratorresponselinkslinks types.
type NullableAiAcceleratorResponseLinksLinks struct {
	value *AiAcceleratorResponseLinksLinks
	isSet bool
}

// Get returns the value.
func (v NullableAiAcceleratorResponseLinksLinks) Get() *AiAcceleratorResponseLinksLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableAiAcceleratorResponseLinksLinks) Set(val *AiAcceleratorResponseLinksLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAiAcceleratorResponseLinksLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAiAcceleratorResponseLinksLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAiAcceleratorResponseLinksLinks returns a pointer to a new instance of NullableAiAcceleratorResponseLinksLinks.
func NewNullableAiAcceleratorResponseLinksLinks(val *AiAcceleratorResponseLinksLinks) *NullableAiAcceleratorResponseLinksLinks {
	return &NullableAiAcceleratorResponseLinksLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAiAcceleratorResponseLinksLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAiAcceleratorResponseLinksLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
