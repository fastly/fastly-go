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

// HistoricalDomainsMetaFilters Filters that were applied when calculating the results for this query.
type HistoricalDomainsMetaFilters struct {
	Region               *string `json:"region,omitempty"`
	Datacenter           *string `json:"datacenter,omitempty"`
	Domain               *string `json:"domain,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalDomainsMetaFilters HistoricalDomainsMetaFilters

// NewHistoricalDomainsMetaFilters instantiates a new HistoricalDomainsMetaFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalDomainsMetaFilters() *HistoricalDomainsMetaFilters {
	this := HistoricalDomainsMetaFilters{}
	return &this
}

// NewHistoricalDomainsMetaFiltersWithDefaults instantiates a new HistoricalDomainsMetaFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalDomainsMetaFiltersWithDefaults() *HistoricalDomainsMetaFilters {
	this := HistoricalDomainsMetaFilters{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *HistoricalDomainsMetaFilters) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDomainsMetaFilters) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *HistoricalDomainsMetaFilters) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *HistoricalDomainsMetaFilters) SetRegion(v string) {
	o.Region = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *HistoricalDomainsMetaFilters) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDomainsMetaFilters) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *HistoricalDomainsMetaFilters) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *HistoricalDomainsMetaFilters) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *HistoricalDomainsMetaFilters) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDomainsMetaFilters) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *HistoricalDomainsMetaFilters) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *HistoricalDomainsMetaFilters) SetDomain(v string) {
	o.Domain = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalDomainsMetaFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Datacenter != nil {
		toSerialize["datacenter"] = o.Datacenter
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *HistoricalDomainsMetaFilters) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalDomainsMetaFilters := _HistoricalDomainsMetaFilters{}

	if err = json.Unmarshal(bytes, &varHistoricalDomainsMetaFilters); err == nil {
		*o = HistoricalDomainsMetaFilters(varHistoricalDomainsMetaFilters)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "region")
		delete(additionalProperties, "datacenter")
		delete(additionalProperties, "domain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalDomainsMetaFilters is a helper abstraction for handling nullable historicaldomainsmetafilters types.
type NullableHistoricalDomainsMetaFilters struct {
	value *HistoricalDomainsMetaFilters
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalDomainsMetaFilters) Get() *HistoricalDomainsMetaFilters {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalDomainsMetaFilters) Set(val *HistoricalDomainsMetaFilters) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalDomainsMetaFilters) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalDomainsMetaFilters) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalDomainsMetaFilters returns a pointer to a new instance of NullableHistoricalDomainsMetaFilters.
func NewNullableHistoricalDomainsMetaFilters(val *HistoricalDomainsMetaFilters) *NullableHistoricalDomainsMetaFilters {
	return &NullableHistoricalDomainsMetaFilters{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalDomainsMetaFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableHistoricalDomainsMetaFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
