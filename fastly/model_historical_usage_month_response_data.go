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

// HistoricalUsageMonthResponseData struct for HistoricalUsageMonthResponseData
type HistoricalUsageMonthResponseData struct {
	CustomerId *string `json:"customer_id,omitempty"`
	// Organized by *service id*.
	Services *map[string]HistoricalUsageService `json:"services,omitempty"`
	// Organized by *region*.
	Total                *map[string]HistoricalUsageData `json:"total,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalUsageMonthResponseData HistoricalUsageMonthResponseData

// NewHistoricalUsageMonthResponseData instantiates a new HistoricalUsageMonthResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalUsageMonthResponseData() *HistoricalUsageMonthResponseData {
	this := HistoricalUsageMonthResponseData{}
	return &this
}

// NewHistoricalUsageMonthResponseDataWithDefaults instantiates a new HistoricalUsageMonthResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalUsageMonthResponseDataWithDefaults() *HistoricalUsageMonthResponseData {
	this := HistoricalUsageMonthResponseData{}
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *HistoricalUsageMonthResponseData) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageMonthResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *HistoricalUsageMonthResponseData) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *HistoricalUsageMonthResponseData) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *HistoricalUsageMonthResponseData) GetServices() map[string]HistoricalUsageService {
	if o == nil || o.Services == nil {
		var ret map[string]HistoricalUsageService
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageMonthResponseData) GetServicesOk() (*map[string]HistoricalUsageService, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *HistoricalUsageMonthResponseData) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given map[string]HistoricalUsageService and assigns it to the Services field.
func (o *HistoricalUsageMonthResponseData) SetServices(v map[string]HistoricalUsageService) {
	o.Services = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *HistoricalUsageMonthResponseData) GetTotal() map[string]HistoricalUsageData {
	if o == nil || o.Total == nil {
		var ret map[string]HistoricalUsageData
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageMonthResponseData) GetTotalOk() (*map[string]HistoricalUsageData, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *HistoricalUsageMonthResponseData) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given map[string]HistoricalUsageData and assigns it to the Total field.
func (o *HistoricalUsageMonthResponseData) SetTotal(v map[string]HistoricalUsageData) {
	o.Total = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalUsageMonthResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *HistoricalUsageMonthResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalUsageMonthResponseData := _HistoricalUsageMonthResponseData{}

	if err = json.Unmarshal(bytes, &varHistoricalUsageMonthResponseData); err == nil {
		*o = HistoricalUsageMonthResponseData(varHistoricalUsageMonthResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "services")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalUsageMonthResponseData is a helper abstraction for handling nullable historicalusagemonthresponsedata types.
type NullableHistoricalUsageMonthResponseData struct {
	value *HistoricalUsageMonthResponseData
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalUsageMonthResponseData) Get() *HistoricalUsageMonthResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalUsageMonthResponseData) Set(val *HistoricalUsageMonthResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalUsageMonthResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalUsageMonthResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalUsageMonthResponseData returns a pointer to a new instance of NullableHistoricalUsageMonthResponseData.
func NewNullableHistoricalUsageMonthResponseData(val *HistoricalUsageMonthResponseData) *NullableHistoricalUsageMonthResponseData {
	return &NullableHistoricalUsageMonthResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalUsageMonthResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableHistoricalUsageMonthResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
