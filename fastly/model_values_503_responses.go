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

// Values503Responses struct for Values503Responses
type Values503Responses struct {
	// The HTTP request path.
	URL *string `json:"url,omitempty"`
	// The rate at which the reason in this dimension occurs among responses to this URL with a 503 status code.
	RatePerURL *float32 `json:"rate_per_url,omitempty"`
	// The rate at which 503 status codes are returned for this URL.
	Var503RatePerURL     *float32 `json:"503_rate_per_url,omitempty"`
	AdditionalProperties map[string]any
}

type _Values503Responses Values503Responses

// NewValues503Responses instantiates a new Values503Responses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValues503Responses() *Values503Responses {
	this := Values503Responses{}
	return &this
}

// NewValues503ResponsesWithDefaults instantiates a new Values503Responses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValues503ResponsesWithDefaults() *Values503Responses {
	this := Values503Responses{}
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *Values503Responses) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values503Responses) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *Values503Responses) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *Values503Responses) SetURL(v string) {
	o.URL = &v
}

// GetRatePerURL returns the RatePerURL field value if set, zero value otherwise.
func (o *Values503Responses) GetRatePerURL() float32 {
	if o == nil || o.RatePerURL == nil {
		var ret float32
		return ret
	}
	return *o.RatePerURL
}

// GetRatePerURLOk returns a tuple with the RatePerURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values503Responses) GetRatePerURLOk() (*float32, bool) {
	if o == nil || o.RatePerURL == nil {
		return nil, false
	}
	return o.RatePerURL, true
}

// HasRatePerURL returns a boolean if a field has been set.
func (o *Values503Responses) HasRatePerURL() bool {
	if o != nil && o.RatePerURL != nil {
		return true
	}

	return false
}

// SetRatePerURL gets a reference to the given float32 and assigns it to the RatePerURL field.
func (o *Values503Responses) SetRatePerURL(v float32) {
	o.RatePerURL = &v
}

// GetVar503RatePerURL returns the Var503RatePerURL field value if set, zero value otherwise.
func (o *Values503Responses) GetVar503RatePerURL() float32 {
	if o == nil || o.Var503RatePerURL == nil {
		var ret float32
		return ret
	}
	return *o.Var503RatePerURL
}

// GetVar503RatePerURLOk returns a tuple with the Var503RatePerURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values503Responses) GetVar503RatePerURLOk() (*float32, bool) {
	if o == nil || o.Var503RatePerURL == nil {
		return nil, false
	}
	return o.Var503RatePerURL, true
}

// HasVar503RatePerURL returns a boolean if a field has been set.
func (o *Values503Responses) HasVar503RatePerURL() bool {
	if o != nil && o.Var503RatePerURL != nil {
		return true
	}

	return false
}

// SetVar503RatePerURL gets a reference to the given float32 and assigns it to the Var503RatePerURL field.
func (o *Values503Responses) SetVar503RatePerURL(v float32) {
	o.Var503RatePerURL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Values503Responses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.URL != nil {
		toSerialize["url"] = o.URL
	}
	if o.RatePerURL != nil {
		toSerialize["rate_per_url"] = o.RatePerURL
	}
	if o.Var503RatePerURL != nil {
		toSerialize["503_rate_per_url"] = o.Var503RatePerURL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Values503Responses) UnmarshalJSON(bytes []byte) (err error) {
	varValues503Responses := _Values503Responses{}

	if err = json.Unmarshal(bytes, &varValues503Responses); err == nil {
		*o = Values503Responses(varValues503Responses)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "rate_per_url")
		delete(additionalProperties, "503_rate_per_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValues503Responses is a helper abstraction for handling nullable values503responses types.
type NullableValues503Responses struct {
	value *Values503Responses
	isSet bool
}

// Get returns the value.
func (v NullableValues503Responses) Get() *Values503Responses {
	return v.value
}

// Set modifies the value.
func (v *NullableValues503Responses) Set(val *Values503Responses) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValues503Responses) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValues503Responses) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValues503Responses returns a pointer to a new instance of NullableValues503Responses.
func NewNullableValues503Responses(val *Values503Responses) *NullableValues503Responses {
	return &NullableValues503Responses{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValues503Responses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValues503Responses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
