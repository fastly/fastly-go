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

// ArcSessionsMeta Metadata about the request.
type ArcSessionsMeta struct {
	// The `limit` value used when making the request.
	Limit *int32 `json:"limit,omitempty"`
	// The `sort` value used when making the request.
	Sort *string `json:"sort,omitempty"`
	// The count of results matching the request.
	Total *int32 `json:"total,omitempty"`
	// The cursor to use to request the next page of results. `null` when there are no more results.
	NextCursor NullableString `json:"next_cursor,omitempty"`
	// The cursor to use to request the previous page of results. `null` when there is no previous page.
	PrevCursor           NullableString `json:"prev_cursor,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcSessionsMeta ArcSessionsMeta

// NewArcSessionsMeta instantiates a new ArcSessionsMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcSessionsMeta() *ArcSessionsMeta {
	this := ArcSessionsMeta{}
	return &this
}

// NewArcSessionsMetaWithDefaults instantiates a new ArcSessionsMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcSessionsMetaWithDefaults() *ArcSessionsMeta {
	this := ArcSessionsMeta{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ArcSessionsMeta) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionsMeta) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ArcSessionsMeta) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ArcSessionsMeta) SetLimit(v int32) {
	o.Limit = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ArcSessionsMeta) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionsMeta) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ArcSessionsMeta) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *ArcSessionsMeta) SetSort(v string) {
	o.Sort = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ArcSessionsMeta) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcSessionsMeta) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ArcSessionsMeta) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ArcSessionsMeta) SetTotal(v int32) {
	o.Total = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcSessionsMeta) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcSessionsMeta) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *ArcSessionsMeta) HasNextCursor() bool {
	if o != nil && o.NextCursor.IsSet() {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given NullableString and assigns it to the NextCursor field.
func (o *ArcSessionsMeta) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// SetNextCursorNil sets the value for NextCursor to be an explicit nil
func (o *ArcSessionsMeta) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
func (o *ArcSessionsMeta) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetPrevCursor returns the PrevCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcSessionsMeta) GetPrevCursor() string {
	if o == nil || o.PrevCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrevCursor.Get()
}

// GetPrevCursorOk returns a tuple with the PrevCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcSessionsMeta) GetPrevCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevCursor.Get(), o.PrevCursor.IsSet()
}

// HasPrevCursor returns a boolean if a field has been set.
func (o *ArcSessionsMeta) HasPrevCursor() bool {
	if o != nil && o.PrevCursor.IsSet() {
		return true
	}

	return false
}

// SetPrevCursor gets a reference to the given NullableString and assigns it to the PrevCursor field.
func (o *ArcSessionsMeta) SetPrevCursor(v string) {
	o.PrevCursor.Set(&v)
}

// SetPrevCursorNil sets the value for PrevCursor to be an explicit nil
func (o *ArcSessionsMeta) SetPrevCursorNil() {
	o.PrevCursor.Set(nil)
}

// UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil
func (o *ArcSessionsMeta) UnsetPrevCursor() {
	o.PrevCursor.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcSessionsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.NextCursor.IsSet() {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if o.PrevCursor.IsSet() {
		toSerialize["prev_cursor"] = o.PrevCursor.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcSessionsMeta) UnmarshalJSON(bytes []byte) (err error) {
	varArcSessionsMeta := _ArcSessionsMeta{}

	if err = json.Unmarshal(bytes, &varArcSessionsMeta); err == nil {
		*o = ArcSessionsMeta(varArcSessionsMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "limit")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "total")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "prev_cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcSessionsMeta is a helper abstraction for handling nullable arcsessionsmeta types.
type NullableArcSessionsMeta struct {
	value *ArcSessionsMeta
	isSet bool
}

// Get returns the value.
func (v NullableArcSessionsMeta) Get() *ArcSessionsMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableArcSessionsMeta) Set(val *ArcSessionsMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcSessionsMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcSessionsMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcSessionsMeta returns a pointer to a new instance of NullableArcSessionsMeta.
func NewNullableArcSessionsMeta(val *ArcSessionsMeta) *NullableArcSessionsMeta {
	return &NullableArcSessionsMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcSessionsMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcSessionsMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
