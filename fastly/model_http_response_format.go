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

// HttpResponseFormat Payload format for delivering to subscribers of whole HTTP responses (`response` hold mode). One of `body` or `body-bin` must be specified.
type HttpResponseFormat struct {
	// The HTTP status code.
	Code *int32 `json:"code,omitempty"`
	// The HTTP status string. Defaults to a string appropriate for `code`.
	Reason *string `json:"reason,omitempty"`
	// A map of arbitrary HTTP response headers to include in the response.
	Headers *map[string]string `json:"headers,omitempty"`
	// The response body as a string.
	Body *string `json:"body,omitempty"`
	// The response body as a base64-encoded binary blob.
	BodyBin              *string `json:"body-bin,omitempty"`
	AdditionalProperties map[string]any
}

type _HttpResponseFormat HttpResponseFormat

// NewHttpResponseFormat instantiates a new HttpResponseFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpResponseFormat() *HttpResponseFormat {
	this := HttpResponseFormat{}
	var code int32 = 200
	this.Code = &code
	return &this
}

// NewHttpResponseFormatWithDefaults instantiates a new HttpResponseFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpResponseFormatWithDefaults() *HttpResponseFormat {
	this := HttpResponseFormat{}
	var code int32 = 200
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *HttpResponseFormat) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpResponseFormat) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *HttpResponseFormat) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *HttpResponseFormat) SetCode(v int32) {
	o.Code = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *HttpResponseFormat) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpResponseFormat) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *HttpResponseFormat) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *HttpResponseFormat) SetReason(v string) {
	o.Reason = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *HttpResponseFormat) GetHeaders() map[string]string {
	if o == nil || o.Headers == nil {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpResponseFormat) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *HttpResponseFormat) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *HttpResponseFormat) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *HttpResponseFormat) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpResponseFormat) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *HttpResponseFormat) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *HttpResponseFormat) SetBody(v string) {
	o.Body = &v
}

// GetBodyBin returns the BodyBin field value if set, zero value otherwise.
func (o *HttpResponseFormat) GetBodyBin() string {
	if o == nil || o.BodyBin == nil {
		var ret string
		return ret
	}
	return *o.BodyBin
}

// GetBodyBinOk returns a tuple with the BodyBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpResponseFormat) GetBodyBinOk() (*string, bool) {
	if o == nil || o.BodyBin == nil {
		return nil, false
	}
	return o.BodyBin, true
}

// HasBodyBin returns a boolean if a field has been set.
func (o *HttpResponseFormat) HasBodyBin() bool {
	if o != nil && o.BodyBin != nil {
		return true
	}

	return false
}

// SetBodyBin gets a reference to the given string and assigns it to the BodyBin field.
func (o *HttpResponseFormat) SetBodyBin(v string) {
	o.BodyBin = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HttpResponseFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.BodyBin != nil {
		toSerialize["body-bin"] = o.BodyBin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *HttpResponseFormat) UnmarshalJSON(bytes []byte) (err error) {
	varHttpResponseFormat := _HttpResponseFormat{}

	if err = json.Unmarshal(bytes, &varHttpResponseFormat); err == nil {
		*o = HttpResponseFormat(varHttpResponseFormat)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "body")
		delete(additionalProperties, "body-bin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHttpResponseFormat is a helper abstraction for handling nullable httpresponseformat types.
type NullableHttpResponseFormat struct {
	value *HttpResponseFormat
	isSet bool
}

// Get returns the value.
func (v NullableHttpResponseFormat) Get() *HttpResponseFormat {
	return v.value
}

// Set modifies the value.
func (v *NullableHttpResponseFormat) Set(val *HttpResponseFormat) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHttpResponseFormat) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHttpResponseFormat) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHttpResponseFormat returns a pointer to a new instance of NullableHttpResponseFormat.
func NewNullableHttpResponseFormat(val *HttpResponseFormat) *NullableHttpResponseFormat {
	return &NullableHttpResponseFormat{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHttpResponseFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableHttpResponseFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
