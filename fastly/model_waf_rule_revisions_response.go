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

// WafRuleRevisionsResponse struct for WafRuleRevisionsResponse
type WafRuleRevisionsResponse struct {
	Links *PaginationLinks `json:"links,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Data []WafRuleRevisionResponseData `json:"data,omitempty"`
	Included []WafRule `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRuleRevisionsResponse WafRuleRevisionsResponse

// NewWafRuleRevisionsResponse instantiates a new WafRuleRevisionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleRevisionsResponse() *WafRuleRevisionsResponse {
	this := WafRuleRevisionsResponse{}
	return &this
}

// NewWafRuleRevisionsResponseWithDefaults instantiates a new WafRuleRevisionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleRevisionsResponseWithDefaults() *WafRuleRevisionsResponse {
	this := WafRuleRevisionsResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WafRuleRevisionsResponse) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionsResponse) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WafRuleRevisionsResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *WafRuleRevisionsResponse) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *WafRuleRevisionsResponse) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionsResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *WafRuleRevisionsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *WafRuleRevisionsResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafRuleRevisionsResponse) GetData() []WafRuleRevisionResponseData {
	if o == nil || o.Data == nil {
		var ret []WafRuleRevisionResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionsResponse) GetDataOk() ([]WafRuleRevisionResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafRuleRevisionsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []WafRuleRevisionResponseData and assigns it to the Data field.
func (o *WafRuleRevisionsResponse) SetData(v []WafRuleRevisionResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafRuleRevisionsResponse) GetIncluded() []WafRule {
	if o == nil || o.Included == nil {
		var ret []WafRule
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionsResponse) GetIncludedOk() ([]WafRule, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafRuleRevisionsResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []WafRule and assigns it to the Included field.
func (o *WafRuleRevisionsResponse) SetIncluded(v []WafRule) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleRevisionsResponse) MarshalJSON() ([]byte, error) {
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
func (o *WafRuleRevisionsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWafRuleRevisionsResponse := _WafRuleRevisionsResponse{}

	if err = json.Unmarshal(bytes, &varWafRuleRevisionsResponse); err == nil {
		*o = WafRuleRevisionsResponse(varWafRuleRevisionsResponse)
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

// NullableWafRuleRevisionsResponse is a helper abstraction for handling nullable wafrulerevisionsresponse types. 
type NullableWafRuleRevisionsResponse struct {
	value *WafRuleRevisionsResponse
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleRevisionsResponse) Get() *WafRuleRevisionsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleRevisionsResponse) Set(val *WafRuleRevisionsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleRevisionsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleRevisionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleRevisionsResponse returns a pointer to a new instance of NullableWafRuleRevisionsResponse.
func NewNullableWafRuleRevisionsResponse(val *WafRuleRevisionsResponse) *NullableWafRuleRevisionsResponse {
	return &NullableWafRuleRevisionsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleRevisionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafRuleRevisionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
