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

// KvStoreUpsertBatch struct for KvStoreUpsertBatch
type KvStoreUpsertBatch struct {
	// The key of the item to be added, updated, retrieved, or deleted.
	Key string `json:"key"`
	// Indicates that the item should be deleted after the specified number of seconds have elapsed. Deletion is not immediate but will occur within 24 hours of the requested expiration.
	TimeToLiveSec *int32 `json:"time_to_live_sec,omitempty"`
	// An arbitrary data field which can contain up to 2000 bytes of data.
	Metadata *string `json:"metadata,omitempty"`
	// If set to true, the new value for the item will not be immediately visible to other users of the KV store; they will receive the existing (stale) value while the platform updates cached copies. Setting this to true ensures that other users of the KV store will receive responses to 'get' operations for this item quickly, although they will be slightly out of date.
	BackgroundFetch *bool `json:"background_fetch,omitempty"`
	// Value for the item (in Base64 encoding).
	Value                string `json:"value"`
	AdditionalProperties map[string]any
}

type _KvStoreUpsertBatch KvStoreUpsertBatch

// NewKvStoreUpsertBatch instantiates a new KvStoreUpsertBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvStoreUpsertBatch(key string, value string) *KvStoreUpsertBatch {
	this := KvStoreUpsertBatch{}
	this.Key = key
	var backgroundFetch bool = false
	this.BackgroundFetch = &backgroundFetch
	this.Value = value
	return &this
}

// NewKvStoreUpsertBatchWithDefaults instantiates a new KvStoreUpsertBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvStoreUpsertBatchWithDefaults() *KvStoreUpsertBatch {
	this := KvStoreUpsertBatch{}
	var backgroundFetch bool = false
	this.BackgroundFetch = &backgroundFetch
	return &this
}

// GetKey returns the Key field value
func (o *KvStoreUpsertBatch) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *KvStoreUpsertBatch) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *KvStoreUpsertBatch) SetKey(v string) {
	o.Key = v
}

// GetTimeToLiveSec returns the TimeToLiveSec field value if set, zero value otherwise.
func (o *KvStoreUpsertBatch) GetTimeToLiveSec() int32 {
	if o == nil || o.TimeToLiveSec == nil {
		var ret int32
		return ret
	}
	return *o.TimeToLiveSec
}

// GetTimeToLiveSecOk returns a tuple with the TimeToLiveSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreUpsertBatch) GetTimeToLiveSecOk() (*int32, bool) {
	if o == nil || o.TimeToLiveSec == nil {
		return nil, false
	}
	return o.TimeToLiveSec, true
}

// HasTimeToLiveSec returns a boolean if a field has been set.
func (o *KvStoreUpsertBatch) HasTimeToLiveSec() bool {
	if o != nil && o.TimeToLiveSec != nil {
		return true
	}

	return false
}

// SetTimeToLiveSec gets a reference to the given int32 and assigns it to the TimeToLiveSec field.
func (o *KvStoreUpsertBatch) SetTimeToLiveSec(v int32) {
	o.TimeToLiveSec = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *KvStoreUpsertBatch) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreUpsertBatch) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *KvStoreUpsertBatch) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *KvStoreUpsertBatch) SetMetadata(v string) {
	o.Metadata = &v
}

// GetBackgroundFetch returns the BackgroundFetch field value if set, zero value otherwise.
func (o *KvStoreUpsertBatch) GetBackgroundFetch() bool {
	if o == nil || o.BackgroundFetch == nil {
		var ret bool
		return ret
	}
	return *o.BackgroundFetch
}

// GetBackgroundFetchOk returns a tuple with the BackgroundFetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvStoreUpsertBatch) GetBackgroundFetchOk() (*bool, bool) {
	if o == nil || o.BackgroundFetch == nil {
		return nil, false
	}
	return o.BackgroundFetch, true
}

// HasBackgroundFetch returns a boolean if a field has been set.
func (o *KvStoreUpsertBatch) HasBackgroundFetch() bool {
	if o != nil && o.BackgroundFetch != nil {
		return true
	}

	return false
}

// SetBackgroundFetch gets a reference to the given bool and assigns it to the BackgroundFetch field.
func (o *KvStoreUpsertBatch) SetBackgroundFetch(v bool) {
	o.BackgroundFetch = &v
}

// GetValue returns the Value field value
func (o *KvStoreUpsertBatch) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *KvStoreUpsertBatch) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *KvStoreUpsertBatch) SetValue(v string) {
	o.Value = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o KvStoreUpsertBatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.TimeToLiveSec != nil {
		toSerialize["time_to_live_sec"] = o.TimeToLiveSec
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.BackgroundFetch != nil {
		toSerialize["background_fetch"] = o.BackgroundFetch
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *KvStoreUpsertBatch) UnmarshalJSON(bytes []byte) (err error) {
	varKvStoreUpsertBatch := _KvStoreUpsertBatch{}

	if err = json.Unmarshal(bytes, &varKvStoreUpsertBatch); err == nil {
		*o = KvStoreUpsertBatch(varKvStoreUpsertBatch)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "time_to_live_sec")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "background_fetch")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableKvStoreUpsertBatch is a helper abstraction for handling nullable kvstoreupsertbatch types.
type NullableKvStoreUpsertBatch struct {
	value *KvStoreUpsertBatch
	isSet bool
}

// Get returns the value.
func (v NullableKvStoreUpsertBatch) Get() *KvStoreUpsertBatch {
	return v.value
}

// Set modifies the value.
func (v *NullableKvStoreUpsertBatch) Set(val *KvStoreUpsertBatch) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableKvStoreUpsertBatch) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableKvStoreUpsertBatch) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableKvStoreUpsertBatch returns a pointer to a new instance of NullableKvStoreUpsertBatch.
func NewNullableKvStoreUpsertBatch(val *KvStoreUpsertBatch) *NullableKvStoreUpsertBatch {
	return &NullableKvStoreUpsertBatch{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableKvStoreUpsertBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableKvStoreUpsertBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
