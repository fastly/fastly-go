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

// WebsiteUpdate struct for WebsiteUpdate
type WebsiteUpdate struct {
	// Website domain
	Domain               *string `json:"domain,omitempty"`
	AdditionalProperties map[string]any
}

type _WebsiteUpdate WebsiteUpdate

// NewWebsiteUpdate instantiates a new WebsiteUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsiteUpdate() *WebsiteUpdate {
	this := WebsiteUpdate{}
	return &this
}

// NewWebsiteUpdateWithDefaults instantiates a new WebsiteUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsiteUpdateWithDefaults() *WebsiteUpdate {
	this := WebsiteUpdate{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *WebsiteUpdate) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsiteUpdate) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *WebsiteUpdate) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *WebsiteUpdate) SetDomain(v string) {
	o.Domain = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WebsiteUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WebsiteUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varWebsiteUpdate := _WebsiteUpdate{}

	if err = json.Unmarshal(bytes, &varWebsiteUpdate); err == nil {
		*o = WebsiteUpdate(varWebsiteUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "domain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWebsiteUpdate is a helper abstraction for handling nullable websiteupdate types.
type NullableWebsiteUpdate struct {
	value *WebsiteUpdate
	isSet bool
}

// Get returns the value.
func (v NullableWebsiteUpdate) Get() *WebsiteUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableWebsiteUpdate) Set(val *WebsiteUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWebsiteUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWebsiteUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWebsiteUpdate returns a pointer to a new instance of NullableWebsiteUpdate.
func NewNullableWebsiteUpdate(val *WebsiteUpdate) *NullableWebsiteUpdate {
	return &NullableWebsiteUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWebsiteUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWebsiteUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
