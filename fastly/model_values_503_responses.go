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
	Url *string `json:"url,omitempty"`
	// The rate at which the reason in this dimension occurs among responses to this URL with a 503 status code.
	RatePerUrl *float32 `json:"rate_per_url,omitempty"`
	// The rate at which 503 status codes are returned for this URL.
	Var503RatePerUrl     *float32 `json:"503_rate_per_url,omitempty"`
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

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Values503Responses) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values503Responses) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Values503Responses) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Values503Responses) SetUrl(v string) {
	o.Url = &v
}

// GetRatePerUrl returns the RatePerUrl field value if set, zero value otherwise.
func (o *Values503Responses) GetRatePerUrl() float32 {
	if o == nil || o.RatePerUrl == nil {
		var ret float32
		return ret
	}
	return *o.RatePerUrl
}

// GetRatePerUrlOk returns a tuple with the RatePerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values503Responses) GetRatePerUrlOk() (*float32, bool) {
	if o == nil || o.RatePerUrl == nil {
		return nil, false
	}
	return o.RatePerUrl, true
}

// HasRatePerUrl returns a boolean if a field has been set.
func (o *Values503Responses) HasRatePerUrl() bool {
	if o != nil && o.RatePerUrl != nil {
		return true
	}

	return false
}

// SetRatePerUrl gets a reference to the given float32 and assigns it to the RatePerUrl field.
func (o *Values503Responses) SetRatePerUrl(v float32) {
	o.RatePerUrl = &v
}

// GetVar503RatePerUrl returns the Var503RatePerUrl field value if set, zero value otherwise.
func (o *Values503Responses) GetVar503RatePerUrl() float32 {
	if o == nil || o.Var503RatePerUrl == nil {
		var ret float32
		return ret
	}
	return *o.Var503RatePerUrl
}

// GetVar503RatePerUrlOk returns a tuple with the Var503RatePerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values503Responses) GetVar503RatePerUrlOk() (*float32, bool) {
	if o == nil || o.Var503RatePerUrl == nil {
		return nil, false
	}
	return o.Var503RatePerUrl, true
}

// HasVar503RatePerUrl returns a boolean if a field has been set.
func (o *Values503Responses) HasVar503RatePerUrl() bool {
	if o != nil && o.Var503RatePerUrl != nil {
		return true
	}

	return false
}

// SetVar503RatePerUrl gets a reference to the given float32 and assigns it to the Var503RatePerUrl field.
func (o *Values503Responses) SetVar503RatePerUrl(v float32) {
	o.Var503RatePerUrl = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Values503Responses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.RatePerUrl != nil {
		toSerialize["rate_per_url"] = o.RatePerUrl
	}
	if o.Var503RatePerUrl != nil {
		toSerialize["503_rate_per_url"] = o.Var503RatePerUrl
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
