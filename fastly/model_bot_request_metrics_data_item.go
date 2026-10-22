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

// BotRequestMetricsDataItem A time series data item containing bot request metric counts. The `bot` field is present when grouped by bot. The `datacenter` field is present when grouped by data center. This dataset is sparse, meaning only keys with non-zero values are included.
type BotRequestMetricsDataItem struct {
	// Timestamp for the time series being reported.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Bot identifier for which data was recorded. Present only when grouped by bot.
	Bot *string `json:"bot,omitempty"`
	// Datacenter (POP) for which data was recorded. Present only when grouped by datacenter.
	Datacenter *string `json:"datacenter,omitempty"`
	// Count of edge requests where a bot was detected and the request was allowed by the configured policy.
	BotEdgeRequestsAllowedCount *int64 `json:"bot_edge_requests_allowed_count,omitempty"`
	// Count of edge requests where a bot was detected and a client challenge was issued for the request by the configured policy.
	BotEdgeRequestsChallengedCount *int64 `json:"bot_edge_requests_challenged_count,omitempty"`
	// Count of edge requests where a bot was detected and the request was blocked by the configured policy.
	BotEdgeRequestsBlockedCount *int64 `json:"bot_edge_requests_blocked_count,omitempty"`
	// Count of edge requests where a bot was detected and the client making the request has already solved a challenge issued by the configured policy.
	BotEdgeRequestsChallengeTokenValidCount *int64 `json:"bot_edge_requests_challenge_token_valid_count,omitempty"`
	AdditionalProperties                    map[string]any
}

type _BotRequestMetricsDataItem BotRequestMetricsDataItem

// NewBotRequestMetricsDataItem instantiates a new BotRequestMetricsDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotRequestMetricsDataItem() *BotRequestMetricsDataItem {
	this := BotRequestMetricsDataItem{}
	return &this
}

// NewBotRequestMetricsDataItemWithDefaults instantiates a new BotRequestMetricsDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotRequestMetricsDataItemWithDefaults() *BotRequestMetricsDataItem {
	this := BotRequestMetricsDataItem{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BotRequestMetricsDataItem) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsDataItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BotRequestMetricsDataItem) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *BotRequestMetricsDataItem) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetBot returns the Bot field value if set, zero value otherwise.
func (o *BotRequestMetricsDataItem) GetBot() string {
	if o == nil || o.Bot == nil {
		var ret string
		return ret
	}
	return *o.Bot
}

// GetBotOk returns a tuple with the Bot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsDataItem) GetBotOk() (*string, bool) {
	if o == nil || o.Bot == nil {
		return nil, false
	}
	return o.Bot, true
}

// HasBot returns a boolean if a field has been set.
func (o *BotRequestMetricsDataItem) HasBot() bool {
	if o != nil && o.Bot != nil {
		return true
	}

	return false
}

// SetBot gets a reference to the given string and assigns it to the Bot field.
func (o *BotRequestMetricsDataItem) SetBot(v string) {
	o.Bot = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *BotRequestMetricsDataItem) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsDataItem) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *BotRequestMetricsDataItem) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *BotRequestMetricsDataItem) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetBotEdgeRequestsAllowedCount returns the BotEdgeRequestsAllowedCount field value if set, zero value otherwise.
func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsAllowedCount() int64 {
	if o == nil || o.BotEdgeRequestsAllowedCount == nil {
		var ret int64
		return ret
	}
	return *o.BotEdgeRequestsAllowedCount
}

// GetBotEdgeRequestsAllowedCountOk returns a tuple with the BotEdgeRequestsAllowedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsAllowedCountOk() (*int64, bool) {
	if o == nil || o.BotEdgeRequestsAllowedCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsAllowedCount, true
}

// HasBotEdgeRequestsAllowedCount returns a boolean if a field has been set.
func (o *BotRequestMetricsDataItem) HasBotEdgeRequestsAllowedCount() bool {
	if o != nil && o.BotEdgeRequestsAllowedCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsAllowedCount gets a reference to the given int64 and assigns it to the BotEdgeRequestsAllowedCount field.
func (o *BotRequestMetricsDataItem) SetBotEdgeRequestsAllowedCount(v int64) {
	o.BotEdgeRequestsAllowedCount = &v
}

// GetBotEdgeRequestsChallengedCount returns the BotEdgeRequestsChallengedCount field value if set, zero value otherwise.
func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsChallengedCount() int64 {
	if o == nil || o.BotEdgeRequestsChallengedCount == nil {
		var ret int64
		return ret
	}
	return *o.BotEdgeRequestsChallengedCount
}

// GetBotEdgeRequestsChallengedCountOk returns a tuple with the BotEdgeRequestsChallengedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsChallengedCountOk() (*int64, bool) {
	if o == nil || o.BotEdgeRequestsChallengedCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsChallengedCount, true
}

// HasBotEdgeRequestsChallengedCount returns a boolean if a field has been set.
func (o *BotRequestMetricsDataItem) HasBotEdgeRequestsChallengedCount() bool {
	if o != nil && o.BotEdgeRequestsChallengedCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsChallengedCount gets a reference to the given int64 and assigns it to the BotEdgeRequestsChallengedCount field.
func (o *BotRequestMetricsDataItem) SetBotEdgeRequestsChallengedCount(v int64) {
	o.BotEdgeRequestsChallengedCount = &v
}

// GetBotEdgeRequestsBlockedCount returns the BotEdgeRequestsBlockedCount field value if set, zero value otherwise.
func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsBlockedCount() int64 {
	if o == nil || o.BotEdgeRequestsBlockedCount == nil {
		var ret int64
		return ret
	}
	return *o.BotEdgeRequestsBlockedCount
}

// GetBotEdgeRequestsBlockedCountOk returns a tuple with the BotEdgeRequestsBlockedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsBlockedCountOk() (*int64, bool) {
	if o == nil || o.BotEdgeRequestsBlockedCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsBlockedCount, true
}

// HasBotEdgeRequestsBlockedCount returns a boolean if a field has been set.
func (o *BotRequestMetricsDataItem) HasBotEdgeRequestsBlockedCount() bool {
	if o != nil && o.BotEdgeRequestsBlockedCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsBlockedCount gets a reference to the given int64 and assigns it to the BotEdgeRequestsBlockedCount field.
func (o *BotRequestMetricsDataItem) SetBotEdgeRequestsBlockedCount(v int64) {
	o.BotEdgeRequestsBlockedCount = &v
}

// GetBotEdgeRequestsChallengeTokenValidCount returns the BotEdgeRequestsChallengeTokenValidCount field value if set, zero value otherwise.
func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsChallengeTokenValidCount() int64 {
	if o == nil || o.BotEdgeRequestsChallengeTokenValidCount == nil {
		var ret int64
		return ret
	}
	return *o.BotEdgeRequestsChallengeTokenValidCount
}

// GetBotEdgeRequestsChallengeTokenValidCountOk returns a tuple with the BotEdgeRequestsChallengeTokenValidCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsChallengeTokenValidCountOk() (*int64, bool) {
	if o == nil || o.BotEdgeRequestsChallengeTokenValidCount == nil {
		return nil, false
	}
	return o.BotEdgeRequestsChallengeTokenValidCount, true
}

// HasBotEdgeRequestsChallengeTokenValidCount returns a boolean if a field has been set.
func (o *BotRequestMetricsDataItem) HasBotEdgeRequestsChallengeTokenValidCount() bool {
	if o != nil && o.BotEdgeRequestsChallengeTokenValidCount != nil {
		return true
	}

	return false
}

// SetBotEdgeRequestsChallengeTokenValidCount gets a reference to the given int64 and assigns it to the BotEdgeRequestsChallengeTokenValidCount field.
func (o *BotRequestMetricsDataItem) SetBotEdgeRequestsChallengeTokenValidCount(v int64) {
	o.BotEdgeRequestsChallengeTokenValidCount = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotRequestMetricsDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Bot != nil {
		toSerialize["bot"] = o.Bot
	}
	if o.Datacenter != nil {
		toSerialize["datacenter"] = o.Datacenter
	}
	if o.BotEdgeRequestsAllowedCount != nil {
		toSerialize["bot_edge_requests_allowed_count"] = o.BotEdgeRequestsAllowedCount
	}
	if o.BotEdgeRequestsChallengedCount != nil {
		toSerialize["bot_edge_requests_challenged_count"] = o.BotEdgeRequestsChallengedCount
	}
	if o.BotEdgeRequestsBlockedCount != nil {
		toSerialize["bot_edge_requests_blocked_count"] = o.BotEdgeRequestsBlockedCount
	}
	if o.BotEdgeRequestsChallengeTokenValidCount != nil {
		toSerialize["bot_edge_requests_challenge_token_valid_count"] = o.BotEdgeRequestsChallengeTokenValidCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotRequestMetricsDataItem) UnmarshalJSON(bytes []byte) (err error) {
	varBotRequestMetricsDataItem := _BotRequestMetricsDataItem{}

	if err = json.Unmarshal(bytes, &varBotRequestMetricsDataItem); err == nil {
		*o = BotRequestMetricsDataItem(varBotRequestMetricsDataItem)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "bot")
		delete(additionalProperties, "datacenter")
		delete(additionalProperties, "bot_edge_requests_allowed_count")
		delete(additionalProperties, "bot_edge_requests_challenged_count")
		delete(additionalProperties, "bot_edge_requests_blocked_count")
		delete(additionalProperties, "bot_edge_requests_challenge_token_valid_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotRequestMetricsDataItem is a helper abstraction for handling nullable botrequestmetricsdataitem types.
type NullableBotRequestMetricsDataItem struct {
	value *BotRequestMetricsDataItem
	isSet bool
}

// Get returns the value.
func (v NullableBotRequestMetricsDataItem) Get() *BotRequestMetricsDataItem {
	return v.value
}

// Set modifies the value.
func (v *NullableBotRequestMetricsDataItem) Set(val *BotRequestMetricsDataItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotRequestMetricsDataItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotRequestMetricsDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotRequestMetricsDataItem returns a pointer to a new instance of NullableBotRequestMetricsDataItem.
func NewNullableBotRequestMetricsDataItem(val *BotRequestMetricsDataItem) *NullableBotRequestMetricsDataItem {
	return &NullableBotRequestMetricsDataItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotRequestMetricsDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotRequestMetricsDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
