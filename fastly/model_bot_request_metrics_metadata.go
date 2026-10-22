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
	"time"
)

// BotRequestMetricsMetadata Meta information about the scope of the current response and optional additional information for subsequent requests.
type BotRequestMetricsMetadata struct {
	// An RFC-8339-formatted date and time string indicating the inclusive start of the query time range.
	From *time.Time `json:"from,omitempty"`
	// An RFC-8339-formatted date and time indicating the exclusive end of the query time range.
	To *time.Time `json:"to,omitempty"`
	// List of dimensions (comma-separated) being summed over in the results.
	GroupBy *string `json:"group_by,omitempty"`
	// Optional string that can be used to request the next page of results, if any.
	NextCursor *string `json:"next_cursor,omitempty"`
	// The maximum number of results to return. Defaults to 1000.
	Limit                *int32 `json:"limit,omitempty"`
	AdditionalProperties map[string]any
}

type _BotRequestMetricsMetadata BotRequestMetricsMetadata

// NewBotRequestMetricsMetadata instantiates a new BotRequestMetricsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotRequestMetricsMetadata() *BotRequestMetricsMetadata {
	this := BotRequestMetricsMetadata{}
	return &this
}

// NewBotRequestMetricsMetadataWithDefaults instantiates a new BotRequestMetricsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotRequestMetricsMetadataWithDefaults() *BotRequestMetricsMetadata {
	this := BotRequestMetricsMetadata{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *BotRequestMetricsMetadata) GetFrom() time.Time {
	if o == nil || o.From == nil {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsMetadata) GetFromOk() (*time.Time, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *BotRequestMetricsMetadata) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *BotRequestMetricsMetadata) SetFrom(v time.Time) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *BotRequestMetricsMetadata) GetTo() time.Time {
	if o == nil || o.To == nil {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsMetadata) GetToOk() (*time.Time, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *BotRequestMetricsMetadata) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *BotRequestMetricsMetadata) SetTo(v time.Time) {
	o.To = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *BotRequestMetricsMetadata) GetGroupBy() string {
	if o == nil || o.GroupBy == nil {
		var ret string
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsMetadata) GetGroupByOk() (*string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *BotRequestMetricsMetadata) HasGroupBy() bool {
	if o != nil && o.GroupBy != nil {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given string and assigns it to the GroupBy field.
func (o *BotRequestMetricsMetadata) SetGroupBy(v string) {
	o.GroupBy = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *BotRequestMetricsMetadata) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsMetadata) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *BotRequestMetricsMetadata) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *BotRequestMetricsMetadata) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *BotRequestMetricsMetadata) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsMetadata) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *BotRequestMetricsMetadata) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *BotRequestMetricsMetadata) SetLimit(v int32) {
	o.Limit = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotRequestMetricsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
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
func (o *BotRequestMetricsMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varBotRequestMetricsMetadata := _BotRequestMetricsMetadata{}

	if err = json.Unmarshal(bytes, &varBotRequestMetricsMetadata); err == nil {
		*o = BotRequestMetricsMetadata(varBotRequestMetricsMetadata)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "group_by")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotRequestMetricsMetadata is a helper abstraction for handling nullable botrequestmetricsmetadata types.
type NullableBotRequestMetricsMetadata struct {
	value *BotRequestMetricsMetadata
	isSet bool
}

// Get returns the value.
func (v NullableBotRequestMetricsMetadata) Get() *BotRequestMetricsMetadata {
	return v.value
}

// Set modifies the value.
func (v *NullableBotRequestMetricsMetadata) Set(val *BotRequestMetricsMetadata) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotRequestMetricsMetadata) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotRequestMetricsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotRequestMetricsMetadata returns a pointer to a new instance of NullableBotRequestMetricsMetadata.
func NewNullableBotRequestMetricsMetadata(val *BotRequestMetricsMetadata) *NullableBotRequestMetricsMetadata {
	return &NullableBotRequestMetricsMetadata{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotRequestMetricsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotRequestMetricsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
