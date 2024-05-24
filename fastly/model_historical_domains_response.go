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

// HistoricalDomainsResponse struct for HistoricalDomainsResponse
type HistoricalDomainsResponse struct {
	// Whether or not we were able to successfully execute the query.
	Status *string `json:"status,omitempty"`
	Meta *HistoricalDomainsMeta `json:"meta,omitempty"`
	// If the query was not successful, this will provide a string that explains why.
	Msg NullableString `json:"msg,omitempty"`
	// A list of timeseries. Each individual timeseries represents a unique combination of dimensions, such as domain, region or POP.
	Data []DomainInspectorEntry `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalDomainsResponse HistoricalDomainsResponse

// NewHistoricalDomainsResponse instantiates a new HistoricalDomainsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalDomainsResponse() *HistoricalDomainsResponse {
	this := HistoricalDomainsResponse{}
	return &this
}

// NewHistoricalDomainsResponseWithDefaults instantiates a new HistoricalDomainsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalDomainsResponseWithDefaults() *HistoricalDomainsResponse {
	this := HistoricalDomainsResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HistoricalDomainsResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDomainsResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HistoricalDomainsResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HistoricalDomainsResponse) SetStatus(v string) {
	o.Status = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *HistoricalDomainsResponse) GetMeta() HistoricalDomainsMeta {
	if o == nil || o.Meta == nil {
		var ret HistoricalDomainsMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDomainsResponse) GetMetaOk() (*HistoricalDomainsMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *HistoricalDomainsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given HistoricalDomainsMeta and assigns it to the Meta field.
func (o *HistoricalDomainsResponse) SetMeta(v HistoricalDomainsMeta) {
	o.Meta = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricalDomainsResponse) GetMsg() string {
	if o == nil || o.Msg.Get() == nil {
		var ret string
		return ret
	}
	return *o.Msg.Get()
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricalDomainsResponse) GetMsgOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Msg.Get(), o.Msg.IsSet()
}

// HasMsg returns a boolean if a field has been set.
func (o *HistoricalDomainsResponse) HasMsg() bool {
	if o != nil && o.Msg.IsSet() {
		return true
	}

	return false
}

// SetMsg gets a reference to the given NullableString and assigns it to the Msg field.
func (o *HistoricalDomainsResponse) SetMsg(v string) {
	o.Msg.Set(&v)
}
// SetMsgNil sets the value for Msg to be an explicit nil
func (o *HistoricalDomainsResponse) SetMsgNil() {
	o.Msg.Set(nil)
}

// UnsetMsg ensures that no value is present for Msg, not even an explicit nil
func (o *HistoricalDomainsResponse) UnsetMsg() {
	o.Msg.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HistoricalDomainsResponse) GetData() []DomainInspectorEntry {
	if o == nil || o.Data == nil {
		var ret []DomainInspectorEntry
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDomainsResponse) GetDataOk() ([]DomainInspectorEntry, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoricalDomainsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []DomainInspectorEntry and assigns it to the Data field.
func (o *HistoricalDomainsResponse) SetData(v []DomainInspectorEntry) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalDomainsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Msg.IsSet() {
		toSerialize["msg"] = o.Msg.Get()
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
func (o *HistoricalDomainsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalDomainsResponse := _HistoricalDomainsResponse{}

	if err = json.Unmarshal(bytes, &varHistoricalDomainsResponse); err == nil {
		*o = HistoricalDomainsResponse(varHistoricalDomainsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalDomainsResponse is a helper abstraction for handling nullable historicaldomainsresponse types. 
type NullableHistoricalDomainsResponse struct {
	value *HistoricalDomainsResponse
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalDomainsResponse) Get() *HistoricalDomainsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalDomainsResponse) Set(val *HistoricalDomainsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalDomainsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalDomainsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalDomainsResponse returns a pointer to a new instance of NullableHistoricalDomainsResponse.
func NewNullableHistoricalDomainsResponse(val *HistoricalDomainsResponse) *NullableHistoricalDomainsResponse {
	return &NullableHistoricalDomainsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalDomainsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalDomainsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
