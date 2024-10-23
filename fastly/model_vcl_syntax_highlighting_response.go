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

// VclSyntaxHighlightingResponse struct for VclSyntaxHighlightingResponse
type VclSyntaxHighlightingResponse struct {
	// VCL with HTML syntax highlighting.
	Content              *string `json:"content,omitempty"`
	AdditionalProperties map[string]any
}

type _VclSyntaxHighlightingResponse VclSyntaxHighlightingResponse

// NewVclSyntaxHighlightingResponse instantiates a new VclSyntaxHighlightingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVclSyntaxHighlightingResponse() *VclSyntaxHighlightingResponse {
	this := VclSyntaxHighlightingResponse{}
	return &this
}

// NewVclSyntaxHighlightingResponseWithDefaults instantiates a new VclSyntaxHighlightingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVclSyntaxHighlightingResponseWithDefaults() *VclSyntaxHighlightingResponse {
	this := VclSyntaxHighlightingResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *VclSyntaxHighlightingResponse) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VclSyntaxHighlightingResponse) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *VclSyntaxHighlightingResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *VclSyntaxHighlightingResponse) SetContent(v string) {
	o.Content = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o VclSyntaxHighlightingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *VclSyntaxHighlightingResponse) UnmarshalJSON(bytes []byte) (err error) {
	varVclSyntaxHighlightingResponse := _VclSyntaxHighlightingResponse{}

	if err = json.Unmarshal(bytes, &varVclSyntaxHighlightingResponse); err == nil {
		*o = VclSyntaxHighlightingResponse(varVclSyntaxHighlightingResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableVclSyntaxHighlightingResponse is a helper abstraction for handling nullable vclsyntaxhighlightingresponse types.
type NullableVclSyntaxHighlightingResponse struct {
	value *VclSyntaxHighlightingResponse
	isSet bool
}

// Get returns the value.
func (v NullableVclSyntaxHighlightingResponse) Get() *VclSyntaxHighlightingResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableVclSyntaxHighlightingResponse) Set(val *VclSyntaxHighlightingResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableVclSyntaxHighlightingResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableVclSyntaxHighlightingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableVclSyntaxHighlightingResponse returns a pointer to a new instance of NullableVclSyntaxHighlightingResponse.
func NewNullableVclSyntaxHighlightingResponse(val *VclSyntaxHighlightingResponse) *NullableVclSyntaxHighlightingResponse {
	return &NullableVclSyntaxHighlightingResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableVclSyntaxHighlightingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableVclSyntaxHighlightingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
