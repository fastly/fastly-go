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

// WafFirewallResponse struct for WafFirewallResponse
type WafFirewallResponse struct {
	Data *WafFirewallResponseData `json:"data,omitempty"`
	Included []SchemasWafFirewallVersion `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallResponse WafFirewallResponse

// NewWafFirewallResponse instantiates a new WafFirewallResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallResponse() *WafFirewallResponse {
	this := WafFirewallResponse{}
	return &this
}

// NewWafFirewallResponseWithDefaults instantiates a new WafFirewallResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallResponseWithDefaults() *WafFirewallResponse {
	this := WafFirewallResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafFirewallResponse) GetData() WafFirewallResponseData {
	if o == nil || o.Data == nil {
		var ret WafFirewallResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponse) GetDataOk() (*WafFirewallResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafFirewallResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given WafFirewallResponseData and assigns it to the Data field.
func (o *WafFirewallResponse) SetData(v WafFirewallResponseData) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafFirewallResponse) GetIncluded() []SchemasWafFirewallVersion {
	if o == nil || o.Included == nil {
		var ret []SchemasWafFirewallVersion
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallResponse) GetIncludedOk() ([]SchemasWafFirewallVersion, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafFirewallResponse) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SchemasWafFirewallVersion and assigns it to the Included field.
func (o *WafFirewallResponse) SetIncluded(v []SchemasWafFirewallVersion) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *WafFirewallResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallResponse := _WafFirewallResponse{}

	if err = json.Unmarshal(bytes, &varWafFirewallResponse); err == nil {
		*o = WafFirewallResponse(varWafFirewallResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallResponse is a helper abstraction for handling nullable waffirewallresponse types. 
type NullableWafFirewallResponse struct {
	value *WafFirewallResponse
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallResponse) Get() *WafFirewallResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallResponse) Set(val *WafFirewallResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallResponse returns a pointer to a new instance of NullableWafFirewallResponse.
func NewNullableWafFirewallResponse(val *WafFirewallResponse) *NullableWafFirewallResponse {
	return &NullableWafFirewallResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
