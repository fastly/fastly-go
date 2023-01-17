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
)

// RateLimiterResponse1 Custom response to be sent when the rate limit is exceeded. Required if `action` is `response`.
type RateLimiterResponse1 struct {
	// HTTP status code for custom limit enforcement response.
	Status *int32 `json:"status,omitempty"`
	// MIME type for custom limit enforcement response.
	ContentType *string `json:"content_type,omitempty"`
	// Response body for custom limit enforcement response.
	Content *string `json:"content,omitempty"`
	AdditionalProperties map[string]any
}

type _RateLimiterResponse1 RateLimiterResponse1

// NewRateLimiterResponse1 instantiates a new RateLimiterResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimiterResponse1() *RateLimiterResponse1 {
	this := RateLimiterResponse1{}
	return &this
}

// NewRateLimiterResponse1WithDefaults instantiates a new RateLimiterResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimiterResponse1WithDefaults() *RateLimiterResponse1 {
	this := RateLimiterResponse1{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RateLimiterResponse1) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiterResponse1) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RateLimiterResponse1) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *RateLimiterResponse1) SetStatus(v int32) {
	o.Status = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *RateLimiterResponse1) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiterResponse1) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *RateLimiterResponse1) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *RateLimiterResponse1) SetContentType(v string) {
	o.ContentType = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *RateLimiterResponse1) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiterResponse1) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *RateLimiterResponse1) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *RateLimiterResponse1) SetContent(v string) {
	o.Content = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RateLimiterResponse1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RateLimiterResponse1) UnmarshalJSON(bytes []byte) (err error) {
	varRateLimiterResponse1 := _RateLimiterResponse1{}

	if err = json.Unmarshal(bytes, &varRateLimiterResponse1); err == nil {
		*o = RateLimiterResponse1(varRateLimiterResponse1)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRateLimiterResponse1 is a helper abstraction for handling nullable ratelimiterresponse1 types. 
type NullableRateLimiterResponse1 struct {
	value *RateLimiterResponse1
	isSet bool
}

// Get returns the value.
func (v NullableRateLimiterResponse1) Get() *RateLimiterResponse1 {
	return v.value
}

// Set modifies the value.
func (v *NullableRateLimiterResponse1) Set(val *RateLimiterResponse1) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRateLimiterResponse1) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRateLimiterResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRateLimiterResponse1 returns a pointer to a new instance of NullableRateLimiterResponse1.
func NewNullableRateLimiterResponse1(val *RateLimiterResponse1) *NullableRateLimiterResponse1 {
	return &NullableRateLimiterResponse1{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRateLimiterResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRateLimiterResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
