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

// DomainCheckItem struct for DomainCheckItem
type DomainCheckItem struct {
	Domain *Domain
	bool *bool
	string *string
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *DomainCheckItem) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Domain
	err = json.Unmarshal(data, &o.Domain);
	if err == nil {
		jsonDomain, _ := json.Marshal(o.Domain)
		if string(jsonDomain) != "{}" { // empty struct
			return nil // data stored in o.Domain, return on the first match
		}
    o.Domain = nil
	} else {
		o.Domain = nil
	}

	// try to unmarshal JSON data into bool
	err = json.Unmarshal(data, &o.bool);
	if err == nil {
		jsonbool, _ := json.Marshal(o.bool)
		if string(jsonbool) != "{}" { // empty struct
			return nil // data stored in o.bool, return on the first match
		}
    o.bool = nil
	} else {
		o.bool = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DomainCheckItem)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *DomainCheckItem) MarshalJSON() ([]byte, error) {
	if o.Domain != nil {
		return json.Marshal(&o.Domain)
	}

	if o.bool != nil {
		return json.Marshal(&o.bool)
	}

	if o.string != nil {
		return json.Marshal(&o.string)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableDomainCheckItem is a helper abstraction for handling nullable domaincheckitem types. 
type NullableDomainCheckItem struct {
	value *DomainCheckItem
	isSet bool
}

// Get returns the value.
func (v NullableDomainCheckItem) Get() *DomainCheckItem {
	return v.value
}

// Set modifies the value.
func (v *NullableDomainCheckItem) Set(val *DomainCheckItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDomainCheckItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDomainCheckItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDomainCheckItem returns a pointer to a new instance of NullableDomainCheckItem.
func NewNullableDomainCheckItem(val *DomainCheckItem) *NullableDomainCheckItem {
	return &NullableDomainCheckItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDomainCheckItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableDomainCheckItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
