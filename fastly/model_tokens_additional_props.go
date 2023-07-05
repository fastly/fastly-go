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
	"fmt"
)

// TokensAdditionalProps struct for TokensAdditionalProps
type TokensAdditionalProps struct {
	float32 *float32
	string *string
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TokensAdditionalProps) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into float32
	err = json.Unmarshal(data, &o.float32);
	if err == nil {
		jsonfloat32, _ := json.Marshal(o.float32)
		if string(jsonfloat32) != "{}" { // empty struct
			return nil // data stored in o.float32, return on the first match
		}
    o.float32 = nil
	} else {
		o.float32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &o.string);
	if err == nil {
		jsonstring, _ := json.Marshal(o.string)
		if string(jsonstring) != "{}" { // empty struct
			return nil // data stored in o.string, return on the first match
		}
    o.string = nil
	} else {
		o.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TokensAdditionalProps)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *TokensAdditionalProps) MarshalJSON() ([]byte, error) {
	if o.float32 != nil {
		return json.Marshal(&o.float32)
	}

	if o.string != nil {
		return json.Marshal(&o.string)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableTokensAdditionalProps is a helper abstraction for handling nullable tokensadditionalprops types. 
type NullableTokensAdditionalProps struct {
	value *TokensAdditionalProps
	isSet bool
}

// Get returns the value.
func (v NullableTokensAdditionalProps) Get() *TokensAdditionalProps {
	return v.value
}

// Set modifies the value.
func (v *NullableTokensAdditionalProps) Set(val *TokensAdditionalProps) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTokensAdditionalProps) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTokensAdditionalProps) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTokensAdditionalProps returns a pointer to a new instance of NullableTokensAdditionalProps.
func NewNullableTokensAdditionalProps(val *TokensAdditionalProps) *NullableTokensAdditionalProps {
	return &NullableTokensAdditionalProps{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTokensAdditionalProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTokensAdditionalProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
