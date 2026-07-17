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

// KvStoreResponseLinksLinks struct for KvStoreResponseLinksLinks
type KvStoreResponseLinksLinks struct {
	// Location of resource
	Self                 *string `json:"self,omitempty"`
	AdditionalProperties map[string]any
}

type _KvStoreResponseLinksLinks KvStoreResponseLinksLinks

// NewKvStoreResponseLinksLinks instantiates a new KvStoreResponseLinksLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvStoreResponseLinksLinks() *KvStoreResponseLinksLinks {
	this := KvStoreResponseLinksLinks{}
	return &this
}

// NewKvStoreResponseLinksLinksWithDefaults instantiates a new KvStoreResponseLinksLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvStoreResponseLinksLinksWithDefaults() *KvStoreResponseLinksLinks {
	this := KvStoreResponseLinksLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *KvStoreResponseLinksLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreResponseLinksLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *KvStoreResponseLinksLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *KvStoreResponseLinksLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o KvStoreResponseLinksLinks) MarshalJSON() ([]byte, error) {
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
func (o *KvStoreResponseLinksLinks) UnmarshalJSON(bytes []byte) (err error) {
	varKvStoreResponseLinksLinks := _KvStoreResponseLinksLinks{}

	if err = json.Unmarshal(bytes, &varKvStoreResponseLinksLinks); err == nil {
		*o = KvStoreResponseLinksLinks(varKvStoreResponseLinksLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableKvStoreResponseLinksLinks is a helper abstraction for handling nullable kvstoreresponselinkslinks types.
type NullableKvStoreResponseLinksLinks struct {
	value *KvStoreResponseLinksLinks
	isSet bool
}

// Get returns the value.
func (v NullableKvStoreResponseLinksLinks) Get() *KvStoreResponseLinksLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableKvStoreResponseLinksLinks) Set(val *KvStoreResponseLinksLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableKvStoreResponseLinksLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableKvStoreResponseLinksLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableKvStoreResponseLinksLinks returns a pointer to a new instance of NullableKvStoreResponseLinksLinks.
func NewNullableKvStoreResponseLinksLinks(val *KvStoreResponseLinksLinks) *NullableKvStoreResponseLinksLinks {
	return &NullableKvStoreResponseLinksLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableKvStoreResponseLinksLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableKvStoreResponseLinksLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
