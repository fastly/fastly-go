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

// TimeseriesMeta struct for TimeseriesMeta
type TimeseriesMeta struct {
	// Start time for the query as supplied in the request.
	From *string `json:"from,omitempty"`
	// End time for the query as supplied in the request.
	To *string `json:"to,omitempty"`
	// The granularity of the time buckets in the response.
	Granularity *string `json:"granularity,omitempty"`
	// Maximum number of results returned in the request.
	Limit                *string `json:"limit,omitempty"`
	AdditionalProperties map[string]any
}

type _TimeseriesMeta TimeseriesMeta

// NewTimeseriesMeta instantiates a new TimeseriesMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesMeta() *TimeseriesMeta {
	this := TimeseriesMeta{}
	return &this
}

// NewTimeseriesMetaWithDefaults instantiates a new TimeseriesMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesMetaWithDefaults() *TimeseriesMeta {
	this := TimeseriesMeta{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *TimeseriesMeta) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesMeta) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *TimeseriesMeta) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *TimeseriesMeta) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *TimeseriesMeta) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesMeta) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *TimeseriesMeta) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *TimeseriesMeta) SetTo(v string) {
	o.To = &v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *TimeseriesMeta) GetGranularity() string {
	if o == nil || o.Granularity == nil {
		var ret string
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesMeta) GetGranularityOk() (*string, bool) {
	if o == nil || o.Granularity == nil {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *TimeseriesMeta) HasGranularity() bool {
	if o != nil && o.Granularity != nil {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given string and assigns it to the Granularity field.
func (o *TimeseriesMeta) SetGranularity(v string) {
	o.Granularity = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TimeseriesMeta) GetLimit() string {
	if o == nil || o.Limit == nil {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesMeta) GetLimitOk() (*string, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TimeseriesMeta) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *TimeseriesMeta) SetLimit(v string) {
	o.Limit = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TimeseriesMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.Granularity != nil {
		toSerialize["granularity"] = o.Granularity
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TimeseriesMeta) UnmarshalJSON(bytes []byte) (err error) {
	varTimeseriesMeta := _TimeseriesMeta{}

	if err = json.Unmarshal(bytes, &varTimeseriesMeta); err == nil {
		*o = TimeseriesMeta(varTimeseriesMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "granularity")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTimeseriesMeta is a helper abstraction for handling nullable timeseriesmeta types.
type NullableTimeseriesMeta struct {
	value *TimeseriesMeta
	isSet bool
}

// Get returns the value.
func (v NullableTimeseriesMeta) Get() *TimeseriesMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableTimeseriesMeta) Set(val *TimeseriesMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTimeseriesMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTimeseriesMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTimeseriesMeta returns a pointer to a new instance of NullableTimeseriesMeta.
func NewNullableTimeseriesMeta(val *TimeseriesMeta) *NullableTimeseriesMeta {
	return &NullableTimeseriesMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTimeseriesMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTimeseriesMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
