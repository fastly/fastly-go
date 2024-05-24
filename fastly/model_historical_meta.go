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

// HistoricalMeta Meta information about the scope of the query in a human readable format.
type HistoricalMeta struct {
	To *string `json:"to,omitempty"`
	From *string `json:"from,omitempty"`
	By *string `json:"by,omitempty"`
	Region *string `json:"region,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalMeta HistoricalMeta

// NewHistoricalMeta instantiates a new HistoricalMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalMeta() *HistoricalMeta {
	this := HistoricalMeta{}
	return &this
}

// NewHistoricalMetaWithDefaults instantiates a new HistoricalMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalMetaWithDefaults() *HistoricalMeta {
	this := HistoricalMeta{}
	return &this
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *HistoricalMeta) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalMeta) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *HistoricalMeta) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *HistoricalMeta) SetTo(v string) {
	o.To = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *HistoricalMeta) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalMeta) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *HistoricalMeta) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *HistoricalMeta) SetFrom(v string) {
	o.From = &v
}

// GetBy returns the By field value if set, zero value otherwise.
func (o *HistoricalMeta) GetBy() string {
	if o == nil || o.By == nil {
		var ret string
		return ret
	}
	return *o.By
}

// GetByOk returns a tuple with the By field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalMeta) GetByOk() (*string, bool) {
	if o == nil || o.By == nil {
		return nil, false
	}
	return o.By, true
}

// HasBy returns a boolean if a field has been set.
func (o *HistoricalMeta) HasBy() bool {
	if o != nil && o.By != nil {
		return true
	}

	return false
}

// SetBy gets a reference to the given string and assigns it to the By field.
func (o *HistoricalMeta) SetBy(v string) {
	o.By = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *HistoricalMeta) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalMeta) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *HistoricalMeta) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *HistoricalMeta) SetRegion(v string) {
	o.Region = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.By != nil {
		toSerialize["by"] = o.By
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *HistoricalMeta) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalMeta := _HistoricalMeta{}

	if err = json.Unmarshal(bytes, &varHistoricalMeta); err == nil {
		*o = HistoricalMeta(varHistoricalMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "to")
		delete(additionalProperties, "from")
		delete(additionalProperties, "by")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalMeta is a helper abstraction for handling nullable historicalmeta types. 
type NullableHistoricalMeta struct {
	value *HistoricalMeta
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalMeta) Get() *HistoricalMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalMeta) Set(val *HistoricalMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalMeta returns a pointer to a new instance of NullableHistoricalMeta.
func NewNullableHistoricalMeta(val *HistoricalMeta) *NullableHistoricalMeta {
	return &NullableHistoricalMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
