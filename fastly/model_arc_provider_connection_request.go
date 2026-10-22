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
	"fmt"
)

// ArcProviderConnectionRequest - struct for ArcProviderConnectionRequest
type ArcProviderConnectionRequest struct {
	ArcProviderConnectionApiKey *ArcProviderConnectionApiKey
	ArcProviderConnectionAwsIam *ArcProviderConnectionAwsIam
}

// ArcProviderConnectionApiKeyAsArcProviderConnectionRequest is a convenience function that returns ArcProviderConnectionApiKey wrapped in ArcProviderConnectionRequest
func ArcProviderConnectionApiKeyAsArcProviderConnectionRequest(v *ArcProviderConnectionApiKey) ArcProviderConnectionRequest {
	return ArcProviderConnectionRequest{
		ArcProviderConnectionApiKey: v,
	}
}

// ArcProviderConnectionAwsIamAsArcProviderConnectionRequest is a convenience function that returns ArcProviderConnectionAwsIam wrapped in ArcProviderConnectionRequest
func ArcProviderConnectionAwsIamAsArcProviderConnectionRequest(v *ArcProviderConnectionAwsIam) ArcProviderConnectionRequest {
	return ArcProviderConnectionRequest{
		ArcProviderConnectionAwsIam: v,
	}
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcProviderConnectionRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArcProviderConnectionApiKey
	err = newStrictDecoder(data).Decode(&o.ArcProviderConnectionApiKey)
	if err == nil {
		jsonArcProviderConnectionApiKey, _ := json.Marshal(o.ArcProviderConnectionApiKey)
		if string(jsonArcProviderConnectionApiKey) == "{}" { // empty struct
			o.ArcProviderConnectionApiKey = nil
		} else {
			match++
		}
	} else {
		o.ArcProviderConnectionApiKey = nil
	}

	// try to unmarshal data into ArcProviderConnectionAwsIam
	err = newStrictDecoder(data).Decode(&o.ArcProviderConnectionAwsIam)
	if err == nil {
		jsonArcProviderConnectionAwsIam, _ := json.Marshal(o.ArcProviderConnectionAwsIam)
		if string(jsonArcProviderConnectionAwsIam) == "{}" { // empty struct
			o.ArcProviderConnectionAwsIam = nil
		} else {
			match++
		}
	} else {
		o.ArcProviderConnectionAwsIam = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		o.ArcProviderConnectionApiKey = nil
		o.ArcProviderConnectionAwsIam = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ArcProviderConnectionRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ArcProviderConnectionRequest)")
	}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcProviderConnectionRequest) MarshalJSON() ([]byte, error) {
	if o.ArcProviderConnectionApiKey != nil {
		return json.Marshal(&o.ArcProviderConnectionApiKey)
	}

	if o.ArcProviderConnectionAwsIam != nil {
		return json.Marshal(&o.ArcProviderConnectionAwsIam)
	}

	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns a specific instance of the type.
func (o *ArcProviderConnectionRequest) GetActualInstance() any {
	if o == nil {
		return nil
	}
	if o.ArcProviderConnectionApiKey != nil {
		return o.ArcProviderConnectionApiKey
	}

	if o.ArcProviderConnectionAwsIam != nil {
		return o.ArcProviderConnectionAwsIam
	}

	// all schemas are nil
	return nil
}

// NullableArcProviderConnectionRequest is a helper abstraction for handling nullable arcproviderconnectionrequest types.
type NullableArcProviderConnectionRequest struct {
	value *ArcProviderConnectionRequest
	isSet bool
}

// Get returns the value.
func (v NullableArcProviderConnectionRequest) Get() *ArcProviderConnectionRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableArcProviderConnectionRequest) Set(val *ArcProviderConnectionRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcProviderConnectionRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcProviderConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcProviderConnectionRequest returns a pointer to a new instance of NullableArcProviderConnectionRequest.
func NewNullableArcProviderConnectionRequest(val *ArcProviderConnectionRequest) *NullableArcProviderConnectionRequest {
	return &NullableArcProviderConnectionRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcProviderConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcProviderConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
