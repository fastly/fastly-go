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

// WafFirewallsResponse struct for WafFirewallsResponse
type WafFirewallsResponse struct {
	Links *PaginationLinks `json:"links,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Data []WafFirewallResponseData `json:"data,omitempty"`
	Included []SchemasWafFirewallVersion `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallsResponse WafFirewallsResponse

// NewWafFirewallsResponse instantiates a new WafFirewallsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallsResponse() *WafFirewallsResponse {
	this := WafFirewallsResponse{}
	return &this
}

// NewWafFirewallsResponseWithDefaults instantiates a new WafFirewallsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallsResponseWithDefaults() *WafFirewallsResponse {
	this := WafFirewallsResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WafFirewallsResponse) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallsResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WafFirewallsResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *WafFirewallsResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *WafFirewallsResponse) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallsResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *WafFirewallsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *WafFirewallsResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafFirewallsResponse) GetData() []WafFirewallResponseData {
	if o == nil || o.Data == nil {
		var ret []WafFirewallResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallsResponse) GetDataOk() ([]WafFirewallResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafFirewallsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []WafFirewallResponseData and assigns it to the Data field.
func (o *WafFirewallsResponse) SetData(v []WafFirewallResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafFirewallsResponse) GetIncluded() []SchemasWafFirewallVersion {
	if o == nil || o.Included == nil {
		var ret []SchemasWafFirewallVersion
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallsResponse) GetIncludedOk() ([]SchemasWafFirewallVersion, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafFirewallsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SchemasWafFirewallVersion and assigns it to the Included field.
func (o *WafFirewallsResponse) SetIncluded(v []SchemasWafFirewallVersion) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallsResponse) MarshalJSON() ([]byte, error) {
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
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafFirewallsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallsResponse := _WafFirewallsResponse{}

	if err = json.Unmarshal(bytes, &varWafFirewallsResponse); err == nil {
		*o = WafFirewallsResponse(varWafFirewallsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallsResponse is a helper abstraction for handling nullable waffirewallsresponse types. 
type NullableWafFirewallsResponse struct {
	value *WafFirewallsResponse
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallsResponse) Get() *WafFirewallsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallsResponse) Set(val *WafFirewallsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallsResponse returns a pointer to a new instance of NullableWafFirewallsResponse.
func NewNullableWafFirewallsResponse(val *WafFirewallsResponse) *NullableWafFirewallsResponse {
	return &NullableWafFirewallsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
