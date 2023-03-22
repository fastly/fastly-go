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

// HistoricalUsageMonthResponseAllOfData struct for HistoricalUsageMonthResponseAllOfData
type HistoricalUsageMonthResponseAllOfData struct {
	CustomerID *string `json:"customer_id,omitempty"`
	Services *map[string]map[string]HistoricalUsageResults `json:"services,omitempty"`
	Total *HistoricalUsageResults `json:"total,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalUsageMonthResponseAllOfData HistoricalUsageMonthResponseAllOfData

// NewHistoricalUsageMonthResponseAllOfData instantiates a new HistoricalUsageMonthResponseAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalUsageMonthResponseAllOfData() *HistoricalUsageMonthResponseAllOfData {
	this := HistoricalUsageMonthResponseAllOfData{}
	return &this
}

// NewHistoricalUsageMonthResponseAllOfDataWithDefaults instantiates a new HistoricalUsageMonthResponseAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalUsageMonthResponseAllOfDataWithDefaults() *HistoricalUsageMonthResponseAllOfData {
	this := HistoricalUsageMonthResponseAllOfData{}
	return &this
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *HistoricalUsageMonthResponseAllOfData) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageMonthResponseAllOfData) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *HistoricalUsageMonthResponseAllOfData) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *HistoricalUsageMonthResponseAllOfData) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *HistoricalUsageMonthResponseAllOfData) GetServices() map[string]map[string]HistoricalUsageResults {
	if o == nil || o.Services == nil {
		var ret map[string]map[string]HistoricalUsageResults
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageMonthResponseAllOfData) GetServicesOk() (*map[string]map[string]HistoricalUsageResults, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *HistoricalUsageMonthResponseAllOfData) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given map[string]map[string]HistoricalUsageResults and assigns it to the Services field.
func (o *HistoricalUsageMonthResponseAllOfData) SetServices(v map[string]map[string]HistoricalUsageResults) {
	o.Services = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *HistoricalUsageMonthResponseAllOfData) GetTotal() HistoricalUsageResults {
	if o == nil || o.Total == nil {
		var ret HistoricalUsageResults
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageMonthResponseAllOfData) GetTotalOk() (*HistoricalUsageResults, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *HistoricalUsageMonthResponseAllOfData) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given HistoricalUsageResults and assigns it to the Total field.
func (o *HistoricalUsageMonthResponseAllOfData) SetTotal(v HistoricalUsageResults) {
	o.Total = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalUsageMonthResponseAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
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
func (o *HistoricalUsageMonthResponseAllOfData) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalUsageMonthResponseAllOfData := _HistoricalUsageMonthResponseAllOfData{}

	if err = json.Unmarshal(bytes, &varHistoricalUsageMonthResponseAllOfData); err == nil {
		*o = HistoricalUsageMonthResponseAllOfData(varHistoricalUsageMonthResponseAllOfData)
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

// NullableHistoricalUsageMonthResponseAllOfData is a helper abstraction for handling nullable historicalusagemonthresponseallofdata types. 
type NullableHistoricalUsageMonthResponseAllOfData struct {
	value *HistoricalUsageMonthResponseAllOfData
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalUsageMonthResponseAllOfData) Get() *HistoricalUsageMonthResponseAllOfData {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalUsageMonthResponseAllOfData) Set(val *HistoricalUsageMonthResponseAllOfData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalUsageMonthResponseAllOfData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalUsageMonthResponseAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalUsageMonthResponseAllOfData returns a pointer to a new instance of NullableHistoricalUsageMonthResponseAllOfData.
func NewNullableHistoricalUsageMonthResponseAllOfData(val *HistoricalUsageMonthResponseAllOfData) *NullableHistoricalUsageMonthResponseAllOfData {
	return &NullableHistoricalUsageMonthResponseAllOfData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalUsageMonthResponseAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalUsageMonthResponseAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
