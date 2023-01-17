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

// MutualAuthenticationsResponse struct for MutualAuthenticationsResponse
type MutualAuthenticationsResponse struct {
	Links *PaginationLinks `json:"links,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Data []MutualAuthenticationResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _MutualAuthenticationsResponse MutualAuthenticationsResponse

// NewMutualAuthenticationsResponse instantiates a new MutualAuthenticationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualAuthenticationsResponse() *MutualAuthenticationsResponse {
	this := MutualAuthenticationsResponse{}
	return &this
}

// NewMutualAuthenticationsResponseWithDefaults instantiates a new MutualAuthenticationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualAuthenticationsResponseWithDefaults() *MutualAuthenticationsResponse {
	this := MutualAuthenticationsResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MutualAuthenticationsResponse) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationsResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MutualAuthenticationsResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *MutualAuthenticationsResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MutualAuthenticationsResponse) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationsResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MutualAuthenticationsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *MutualAuthenticationsResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MutualAuthenticationsResponse) GetData() []MutualAuthenticationResponseData {
	if o == nil || o.Data == nil {
		var ret []MutualAuthenticationResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationsResponse) GetDataOk() ([]MutualAuthenticationResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MutualAuthenticationsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []MutualAuthenticationResponseData and assigns it to the Data field.
func (o *MutualAuthenticationsResponse) SetData(v []MutualAuthenticationResponseData) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o MutualAuthenticationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *MutualAuthenticationsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varMutualAuthenticationsResponse := _MutualAuthenticationsResponse{}

	if err = json.Unmarshal(bytes, &varMutualAuthenticationsResponse); err == nil {
		*o = MutualAuthenticationsResponse(varMutualAuthenticationsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableMutualAuthenticationsResponse is a helper abstraction for handling nullable mutualauthenticationsresponse types. 
type NullableMutualAuthenticationsResponse struct {
	value *MutualAuthenticationsResponse
	isSet bool
}

// Get returns the value.
func (v NullableMutualAuthenticationsResponse) Get() *MutualAuthenticationsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableMutualAuthenticationsResponse) Set(val *MutualAuthenticationsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableMutualAuthenticationsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableMutualAuthenticationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMutualAuthenticationsResponse returns a pointer to a new instance of NullableMutualAuthenticationsResponse.
func NewNullableMutualAuthenticationsResponse(val *MutualAuthenticationsResponse) *NullableMutualAuthenticationsResponse {
	return &NullableMutualAuthenticationsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableMutualAuthenticationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableMutualAuthenticationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
