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

// OriginInspectorHistoricalMeta Meta information about the scope of the query in a human readable format.
type OriginInspectorHistoricalMeta struct {
	// Start time that was used to perform the query as an ISO-8601-formatted date and time.
	Start *string `json:"start,omitempty"`
	// End time that was used to perform the query as an ISO-8601-formatted date and time.
	End *string `json:"end,omitempty"`
	// Downsample that was used to perform the query. One of `minute`, `hour` or `day`.
	Downsample *string `json:"downsample,omitempty"`
	// A comma-separated list of the metrics that were requested.
	Metrics *string `json:"metrics,omitempty"`
	// The maximum number of results shown per page.
	Limit *float32 `json:"limit,omitempty"`
	// A string that can be used to request the next page of results, if any.
	NextCursor *string `json:"next_cursor,omitempty"`
	// A comma-separated list of keys the results are sorted by.
	Sort *string `json:"sort,omitempty"`
	// A comma-separated list of dimensions being summed over in the query.
	GroupBy              *string                               `json:"group_by,omitempty"`
	Filters              *OriginInspectorHistoricalMetaFilters `json:"filters,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorHistoricalMeta OriginInspectorHistoricalMeta

// NewOriginInspectorHistoricalMeta instantiates a new OriginInspectorHistoricalMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorHistoricalMeta() *OriginInspectorHistoricalMeta {
	this := OriginInspectorHistoricalMeta{}
	return &this
}

// NewOriginInspectorHistoricalMetaWithDefaults instantiates a new OriginInspectorHistoricalMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorHistoricalMetaWithDefaults() *OriginInspectorHistoricalMeta {
	this := OriginInspectorHistoricalMeta{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMeta) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMeta) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMeta) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *OriginInspectorHistoricalMeta) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMeta) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMeta) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMeta) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *OriginInspectorHistoricalMeta) SetEnd(v string) {
	o.End = &v
}

// GetDownsample returns the Downsample field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMeta) GetDownsample() string {
	if o == nil || o.Downsample == nil {
		var ret string
		return ret
	}
	return *o.Downsample
}

// GetDownsampleOk returns a tuple with the Downsample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMeta) GetDownsampleOk() (*string, bool) {
	if o == nil || o.Downsample == nil {
		return nil, false
	}
	return o.Downsample, true
}

// HasDownsample returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMeta) HasDownsample() bool {
	if o != nil && o.Downsample != nil {
		return true
	}

	return false
}

// SetDownsample gets a reference to the given string and assigns it to the Downsample field.
func (o *OriginInspectorHistoricalMeta) SetDownsample(v string) {
	o.Downsample = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMeta) GetMetrics() string {
	if o == nil || o.Metrics == nil {
		var ret string
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMeta) GetMetricsOk() (*string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMeta) HasMetrics() bool {
	if o != nil && o.Metrics != nil {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given string and assigns it to the Metrics field.
func (o *OriginInspectorHistoricalMeta) SetMetrics(v string) {
	o.Metrics = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMeta) GetLimit() float32 {
	if o == nil || o.Limit == nil {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMeta) GetLimitOk() (*float32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMeta) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *OriginInspectorHistoricalMeta) SetLimit(v float32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMeta) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMeta) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMeta) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *OriginInspectorHistoricalMeta) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMeta) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMeta) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMeta) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *OriginInspectorHistoricalMeta) SetSort(v string) {
	o.Sort = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMeta) GetGroupBy() string {
	if o == nil || o.GroupBy == nil {
		var ret string
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMeta) GetGroupByOk() (*string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMeta) HasGroupBy() bool {
	if o != nil && o.GroupBy != nil {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given string and assigns it to the GroupBy field.
func (o *OriginInspectorHistoricalMeta) SetGroupBy(v string) {
	o.GroupBy = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalMeta) GetFilters() OriginInspectorHistoricalMetaFilters {
	if o == nil || o.Filters == nil {
		var ret OriginInspectorHistoricalMetaFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalMeta) GetFiltersOk() (*OriginInspectorHistoricalMetaFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalMeta) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given OriginInspectorHistoricalMetaFilters and assigns it to the Filters field.
func (o *OriginInspectorHistoricalMeta) SetFilters(v OriginInspectorHistoricalMetaFilters) {
	o.Filters = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorHistoricalMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Downsample != nil {
		toSerialize["downsample"] = o.Downsample
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OriginInspectorHistoricalMeta) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorHistoricalMeta := _OriginInspectorHistoricalMeta{}

	if err = json.Unmarshal(bytes, &varOriginInspectorHistoricalMeta); err == nil {
		*o = OriginInspectorHistoricalMeta(varOriginInspectorHistoricalMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "start")
		delete(additionalProperties, "end")
		delete(additionalProperties, "downsample")
		delete(additionalProperties, "metrics")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "group_by")
		delete(additionalProperties, "filters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorHistoricalMeta is a helper abstraction for handling nullable origininspectorhistoricalmeta types.
type NullableOriginInspectorHistoricalMeta struct {
	value *OriginInspectorHistoricalMeta
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorHistoricalMeta) Get() *OriginInspectorHistoricalMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorHistoricalMeta) Set(val *OriginInspectorHistoricalMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorHistoricalMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorHistoricalMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorHistoricalMeta returns a pointer to a new instance of NullableOriginInspectorHistoricalMeta.
func NewNullableOriginInspectorHistoricalMeta(val *OriginInspectorHistoricalMeta) *NullableOriginInspectorHistoricalMeta {
	return &NullableOriginInspectorHistoricalMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorHistoricalMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOriginInspectorHistoricalMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
