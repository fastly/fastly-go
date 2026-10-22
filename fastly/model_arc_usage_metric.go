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

// ArcUsageMetric A usage metric measured for a single date (or hour, when `granularity` is `hourly`).
type ArcUsageMetric struct {
	// The date (or hour) the usage was measured. Formatted as `YYYY-MM-DD` when `granularity` is `daily`, or `YYYY-MM-DD HH:00:00` when `granularity` is `hourly`.
	Date string `json:"date"`
	// The type of usage being measured.
	UsageType string `json:"usage_type"`
	// The quantity of the usage type for this date.
	Quantity int32 `json:"quantity"`
	// The ID of the virtual key this usage is attributed to.
	VirtualKeyId *string `json:"virtual_key_id,omitempty"`
	// The display name of the virtual key this usage is attributed to.
	VirtualKeyName *string `json:"virtual_key_name,omitempty"`
	// The ID of the provider connection originally requested, before any failover occurred.
	PrimaryProvider *string `json:"primary_provider,omitempty"`
	// The name of the provider connection that ultimately served the request.
	Provider *string `json:"provider,omitempty"`
	// The ID of the provider connection that ultimately served the request.
	ProviderConnectionId *string `json:"provider_connection_id,omitempty"`
	// The AI model identifier used.
	Model                *string `json:"model,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcUsageMetric ArcUsageMetric

// NewArcUsageMetric instantiates a new ArcUsageMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcUsageMetric(date string, usageType string, quantity int32) *ArcUsageMetric {
	this := ArcUsageMetric{}
	this.Date = date
	this.UsageType = usageType
	this.Quantity = quantity
	return &this
}

// NewArcUsageMetricWithDefaults instantiates a new ArcUsageMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcUsageMetricWithDefaults() *ArcUsageMetric {
	this := ArcUsageMetric{}
	return &this
}

// GetDate returns the Date field value
func (o *ArcUsageMetric) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ArcUsageMetric) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ArcUsageMetric) SetDate(v string) {
	o.Date = v
}

// GetUsageType returns the UsageType field value
func (o *ArcUsageMetric) GetUsageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value
// and a boolean to check if the value has been set.
func (o *ArcUsageMetric) GetUsageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageType, true
}

// SetUsageType sets field value
func (o *ArcUsageMetric) SetUsageType(v string) {
	o.UsageType = v
}

// GetQuantity returns the Quantity field value
func (o *ArcUsageMetric) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ArcUsageMetric) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ArcUsageMetric) SetQuantity(v int32) {
	o.Quantity = v
}

// GetVirtualKeyId returns the VirtualKeyId field value if set, zero value otherwise.
func (o *ArcUsageMetric) GetVirtualKeyId() string {
	if o == nil || o.VirtualKeyId == nil {
		var ret string
		return ret
	}
	return *o.VirtualKeyId
}

// GetVirtualKeyIdOk returns a tuple with the VirtualKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUsageMetric) GetVirtualKeyIdOk() (*string, bool) {
	if o == nil || o.VirtualKeyId == nil {
		return nil, false
	}
	return o.VirtualKeyId, true
}

// HasVirtualKeyId returns a boolean if a field has been set.
func (o *ArcUsageMetric) HasVirtualKeyId() bool {
	if o != nil && o.VirtualKeyId != nil {
		return true
	}

	return false
}

// SetVirtualKeyId gets a reference to the given string and assigns it to the VirtualKeyId field.
func (o *ArcUsageMetric) SetVirtualKeyId(v string) {
	o.VirtualKeyId = &v
}

// GetVirtualKeyName returns the VirtualKeyName field value if set, zero value otherwise.
func (o *ArcUsageMetric) GetVirtualKeyName() string {
	if o == nil || o.VirtualKeyName == nil {
		var ret string
		return ret
	}
	return *o.VirtualKeyName
}

// GetVirtualKeyNameOk returns a tuple with the VirtualKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUsageMetric) GetVirtualKeyNameOk() (*string, bool) {
	if o == nil || o.VirtualKeyName == nil {
		return nil, false
	}
	return o.VirtualKeyName, true
}

// HasVirtualKeyName returns a boolean if a field has been set.
func (o *ArcUsageMetric) HasVirtualKeyName() bool {
	if o != nil && o.VirtualKeyName != nil {
		return true
	}

	return false
}

// SetVirtualKeyName gets a reference to the given string and assigns it to the VirtualKeyName field.
func (o *ArcUsageMetric) SetVirtualKeyName(v string) {
	o.VirtualKeyName = &v
}

// GetPrimaryProvider returns the PrimaryProvider field value if set, zero value otherwise.
func (o *ArcUsageMetric) GetPrimaryProvider() string {
	if o == nil || o.PrimaryProvider == nil {
		var ret string
		return ret
	}
	return *o.PrimaryProvider
}

// GetPrimaryProviderOk returns a tuple with the PrimaryProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUsageMetric) GetPrimaryProviderOk() (*string, bool) {
	if o == nil || o.PrimaryProvider == nil {
		return nil, false
	}
	return o.PrimaryProvider, true
}

// HasPrimaryProvider returns a boolean if a field has been set.
func (o *ArcUsageMetric) HasPrimaryProvider() bool {
	if o != nil && o.PrimaryProvider != nil {
		return true
	}

	return false
}

// SetPrimaryProvider gets a reference to the given string and assigns it to the PrimaryProvider field.
func (o *ArcUsageMetric) SetPrimaryProvider(v string) {
	o.PrimaryProvider = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ArcUsageMetric) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUsageMetric) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ArcUsageMetric) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ArcUsageMetric) SetProvider(v string) {
	o.Provider = &v
}

// GetProviderConnectionId returns the ProviderConnectionId field value if set, zero value otherwise.
func (o *ArcUsageMetric) GetProviderConnectionId() string {
	if o == nil || o.ProviderConnectionId == nil {
		var ret string
		return ret
	}
	return *o.ProviderConnectionId
}

// GetProviderConnectionIdOk returns a tuple with the ProviderConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUsageMetric) GetProviderConnectionIdOk() (*string, bool) {
	if o == nil || o.ProviderConnectionId == nil {
		return nil, false
	}
	return o.ProviderConnectionId, true
}

// HasProviderConnectionId returns a boolean if a field has been set.
func (o *ArcUsageMetric) HasProviderConnectionId() bool {
	if o != nil && o.ProviderConnectionId != nil {
		return true
	}

	return false
}

// SetProviderConnectionId gets a reference to the given string and assigns it to the ProviderConnectionId field.
func (o *ArcUsageMetric) SetProviderConnectionId(v string) {
	o.ProviderConnectionId = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ArcUsageMetric) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUsageMetric) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ArcUsageMetric) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ArcUsageMetric) SetModel(v string) {
	o.Model = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcUsageMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["usage_type"] = o.UsageType
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if o.VirtualKeyId != nil {
		toSerialize["virtual_key_id"] = o.VirtualKeyId
	}
	if o.VirtualKeyName != nil {
		toSerialize["virtual_key_name"] = o.VirtualKeyName
	}
	if o.PrimaryProvider != nil {
		toSerialize["primary_provider"] = o.PrimaryProvider
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.ProviderConnectionId != nil {
		toSerialize["provider_connection_id"] = o.ProviderConnectionId
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcUsageMetric) UnmarshalJSON(bytes []byte) (err error) {
	varArcUsageMetric := _ArcUsageMetric{}

	if err = json.Unmarshal(bytes, &varArcUsageMetric); err == nil {
		*o = ArcUsageMetric(varArcUsageMetric)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "usage_type")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "virtual_key_id")
		delete(additionalProperties, "virtual_key_name")
		delete(additionalProperties, "primary_provider")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "provider_connection_id")
		delete(additionalProperties, "model")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcUsageMetric is a helper abstraction for handling nullable arcusagemetric types.
type NullableArcUsageMetric struct {
	value *ArcUsageMetric
	isSet bool
}

// Get returns the value.
func (v NullableArcUsageMetric) Get() *ArcUsageMetric {
	return v.value
}

// Set modifies the value.
func (v *NullableArcUsageMetric) Set(val *ArcUsageMetric) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcUsageMetric) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcUsageMetric) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcUsageMetric returns a pointer to a new instance of NullableArcUsageMetric.
func NewNullableArcUsageMetric(val *ArcUsageMetric) *NullableArcUsageMetric {
	return &NullableArcUsageMetric{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcUsageMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcUsageMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
