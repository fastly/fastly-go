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

// PurgeKeys Purge multiple surrogate key tags using a JSON POST body. Not required if the `Surrogate-Key` header is specified.
type PurgeKeys struct {
	SurrogateKeys []string `json:"surrogate_keys,omitempty"`
	AdditionalProperties map[string]any
}

type _PurgeKeys PurgeKeys

// NewPurgeKeys instantiates a new PurgeKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurgeKeys() *PurgeKeys {
	this := PurgeKeys{}
	return &this
}

// NewPurgeKeysWithDefaults instantiates a new PurgeKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurgeKeysWithDefaults() *PurgeKeys {
	this := PurgeKeys{}
	return &this
}

// GetSurrogateKeys returns the SurrogateKeys field value if set, zero value otherwise.
func (o *PurgeKeys) GetSurrogateKeys() []string {
	if o == nil || o.SurrogateKeys == nil {
		var ret []string
		return ret
	}
	return o.SurrogateKeys
}

// GetSurrogateKeysOk returns a tuple with the SurrogateKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurgeKeys) GetSurrogateKeysOk() ([]string, bool) {
	if o == nil || o.SurrogateKeys == nil {
		return nil, false
	}
	return o.SurrogateKeys, true
}

// HasSurrogateKeys returns a boolean if a field has been set.
func (o *PurgeKeys) HasSurrogateKeys() bool {
	if o != nil && o.SurrogateKeys != nil {
		return true
	}

	return false
}

// SetSurrogateKeys gets a reference to the given []string and assigns it to the SurrogateKeys field.
func (o *PurgeKeys) SetSurrogateKeys(v []string) {
	o.SurrogateKeys = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PurgeKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.SurrogateKeys != nil {
		toSerialize["surrogate_keys"] = o.SurrogateKeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *PurgeKeys) UnmarshalJSON(bytes []byte) (err error) {
	varPurgeKeys := _PurgeKeys{}

	if err = json.Unmarshal(bytes, &varPurgeKeys); err == nil {
		*o = PurgeKeys(varPurgeKeys)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "surrogate_keys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePurgeKeys is a helper abstraction for handling nullable purgekeys types. 
type NullablePurgeKeys struct {
	value *PurgeKeys
	isSet bool
}

// Get returns the value.
func (v NullablePurgeKeys) Get() *PurgeKeys {
	return v.value
}

// Set modifies the value.
func (v *NullablePurgeKeys) Set(val *PurgeKeys) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePurgeKeys) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePurgeKeys) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePurgeKeys returns a pointer to a new instance of NullablePurgeKeys.
func NewNullablePurgeKeys(val *PurgeKeys) *NullablePurgeKeys {
	return &NullablePurgeKeys{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePurgeKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePurgeKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
