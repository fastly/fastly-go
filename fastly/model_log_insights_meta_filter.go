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

// LogInsightsMetaFilter The filters that were supplied in the request.
type LogInsightsMetaFilter struct {
	// Specifies the ID of the service for which data should be returned.
	ServiceID *string `json:"service_id,omitempty"`
	// Start time for the query as supplied in the request.
	Start *string `json:"start,omitempty"`
	// End time for the query as supplied in the request.
	End *string `json:"end,omitempty"`
	// Value of the `domain_exact_match` filter as supplied in the request.
	DomainExactMatch *bool `json:"domain_exact_match,omitempty"`
	// Number of records per page.
	Limit                *int32 `json:"limit,omitempty"`
	AdditionalProperties map[string]any
}

type _LogInsightsMetaFilter LogInsightsMetaFilter

// NewLogInsightsMetaFilter instantiates a new LogInsightsMetaFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogInsightsMetaFilter() *LogInsightsMetaFilter {
	this := LogInsightsMetaFilter{}
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// NewLogInsightsMetaFilterWithDefaults instantiates a new LogInsightsMetaFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogInsightsMetaFilterWithDefaults() *LogInsightsMetaFilter {
	this := LogInsightsMetaFilter{}
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *LogInsightsMetaFilter) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInsightsMetaFilter) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *LogInsightsMetaFilter) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *LogInsightsMetaFilter) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *LogInsightsMetaFilter) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInsightsMetaFilter) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *LogInsightsMetaFilter) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *LogInsightsMetaFilter) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *LogInsightsMetaFilter) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInsightsMetaFilter) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *LogInsightsMetaFilter) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *LogInsightsMetaFilter) SetEnd(v string) {
	o.End = &v
}

// GetDomainExactMatch returns the DomainExactMatch field value if set, zero value otherwise.
func (o *LogInsightsMetaFilter) GetDomainExactMatch() bool {
	if o == nil || o.DomainExactMatch == nil {
		var ret bool
		return ret
	}
	return *o.DomainExactMatch
}

// GetDomainExactMatchOk returns a tuple with the DomainExactMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInsightsMetaFilter) GetDomainExactMatchOk() (*bool, bool) {
	if o == nil || o.DomainExactMatch == nil {
		return nil, false
	}
	return o.DomainExactMatch, true
}

// HasDomainExactMatch returns a boolean if a field has been set.
func (o *LogInsightsMetaFilter) HasDomainExactMatch() bool {
	if o != nil && o.DomainExactMatch != nil {
		return true
	}

	return false
}

// SetDomainExactMatch gets a reference to the given bool and assigns it to the DomainExactMatch field.
func (o *LogInsightsMetaFilter) SetDomainExactMatch(v bool) {
	o.DomainExactMatch = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *LogInsightsMetaFilter) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInsightsMetaFilter) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *LogInsightsMetaFilter) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *LogInsightsMetaFilter) SetLimit(v int32) {
	o.Limit = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogInsightsMetaFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.DomainExactMatch != nil {
		toSerialize["domain_exact_match"] = o.DomainExactMatch
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
func (o *LogInsightsMetaFilter) UnmarshalJSON(bytes []byte) (err error) {
	varLogInsightsMetaFilter := _LogInsightsMetaFilter{}

	if err = json.Unmarshal(bytes, &varLogInsightsMetaFilter); err == nil {
		*o = LogInsightsMetaFilter(varLogInsightsMetaFilter)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "start")
		delete(additionalProperties, "end")
		delete(additionalProperties, "domain_exact_match")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogInsightsMetaFilter is a helper abstraction for handling nullable loginsightsmetafilter types.
type NullableLogInsightsMetaFilter struct {
	value *LogInsightsMetaFilter
	isSet bool
}

// Get returns the value.
func (v NullableLogInsightsMetaFilter) Get() *LogInsightsMetaFilter {
	return v.value
}

// Set modifies the value.
func (v *NullableLogInsightsMetaFilter) Set(val *LogInsightsMetaFilter) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogInsightsMetaFilter) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogInsightsMetaFilter) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogInsightsMetaFilter returns a pointer to a new instance of NullableLogInsightsMetaFilter.
func NewNullableLogInsightsMetaFilter(val *LogInsightsMetaFilter) *NullableLogInsightsMetaFilter {
	return &NullableLogInsightsMetaFilter{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogInsightsMetaFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogInsightsMetaFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
