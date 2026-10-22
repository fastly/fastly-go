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

// ArcSessionLogEntry A single request and response pair recorded within a session.
type ArcSessionLogEntry struct {
	// The raw request data.
	Request *string `json:"request,omitempty"`
	// The raw response data.
	Response *string `json:"response,omitempty"`
	// Additional attributes recorded for this request/response pair.
	Attrs                map[string]map[string]interface{} `json:"attrs,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcSessionLogEntry ArcSessionLogEntry

// NewArcSessionLogEntry instantiates a new ArcSessionLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcSessionLogEntry() *ArcSessionLogEntry {
	this := ArcSessionLogEntry{}
	return &this
}

// NewArcSessionLogEntryWithDefaults instantiates a new ArcSessionLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcSessionLogEntryWithDefaults() *ArcSessionLogEntry {
	this := ArcSessionLogEntry{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *ArcSessionLogEntry) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionLogEntry) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ArcSessionLogEntry) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *ArcSessionLogEntry) SetRequest(v string) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ArcSessionLogEntry) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionLogEntry) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ArcSessionLogEntry) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *ArcSessionLogEntry) SetResponse(v string) {
	o.Response = &v
}

// GetAttrs returns the Attrs field value if set, zero value otherwise.
func (o *ArcSessionLogEntry) GetAttrs() map[string]map[string]interface{} {
	if o == nil || o.Attrs == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Attrs
}

// GetAttrsOk returns a tuple with the Attrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionLogEntry) GetAttrsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Attrs == nil {
		return nil, false
	}
	return o.Attrs, true
}

// HasAttrs returns a boolean if a field has been set.
func (o *ArcSessionLogEntry) HasAttrs() bool {
	if o != nil && o.Attrs != nil {
		return true
	}

	return false
}

// SetAttrs gets a reference to the given map[string]map[string]interface{} and assigns it to the Attrs field.
func (o *ArcSessionLogEntry) SetAttrs(v map[string]map[string]interface{}) {
	o.Attrs = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcSessionLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.Attrs != nil {
		toSerialize["attrs"] = o.Attrs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcSessionLogEntry) UnmarshalJSON(bytes []byte) (err error) {
	varArcSessionLogEntry := _ArcSessionLogEntry{}

	if err = json.Unmarshal(bytes, &varArcSessionLogEntry); err == nil {
		*o = ArcSessionLogEntry(varArcSessionLogEntry)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "response")
		delete(additionalProperties, "attrs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcSessionLogEntry is a helper abstraction for handling nullable arcsessionlogentry types.
type NullableArcSessionLogEntry struct {
	value *ArcSessionLogEntry
	isSet bool
}

// Get returns the value.
func (v NullableArcSessionLogEntry) Get() *ArcSessionLogEntry {
	return v.value
}

// Set modifies the value.
func (v *NullableArcSessionLogEntry) Set(val *ArcSessionLogEntry) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcSessionLogEntry) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcSessionLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcSessionLogEntry returns a pointer to a new instance of NullableArcSessionLogEntry.
func NewNullableArcSessionLogEntry(val *ArcSessionLogEntry) *NullableArcSessionLogEntry {
	return &NullableArcSessionLogEntry{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcSessionLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcSessionLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
