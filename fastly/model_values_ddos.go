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

// ValuesDdos The results of the query, optionally filtered and grouped over the requested timespan.
type ValuesDdos struct {
	// For HTTP/2, the number of connections the limit-streams action was applied to. The limit-streams action caps the allowed number of concurrent streams in a connection.
	DdosActionLimitStreamsConnections *int32 `json:"ddos_action_limit_streams_connections,omitempty"`
	// For HTTP/2, the number of requests made on a connection for which the limit-streams action was taken. The limit-streams action caps the allowed number of concurrent streams in a connection.
	DdosActionLimitStreamsRequests *int32 `json:"ddos_action_limit_streams_requests,omitempty"`
	// The number of times the tarpit-accept action was taken. The tarpit-accept action adds a delay when accepting future connections.
	DdosActionTarpitAccept *int32 `json:"ddos_action_tarpit_accept,omitempty"`
	// The number of times the tarpit action was taken. The tarpit action delays writing the response to the client.
	DdosActionTarpit *int32 `json:"ddos_action_tarpit,omitempty"`
	// The number of times the close action was taken. The close action aborts the connection as soon as possible. The close action takes effect either right after accept, right after the client hello, or right after the response was sent.
	DdosActionClose *int32 `json:"ddos_action_close,omitempty"`
	// The number of times the blackhole action was taken. The blackhole action quietly closes a TCP connection without sending a reset. The blackhole action quietly closes a TCP connection without notifying its peer (all TCP state is dropped).
	DdosActionBlackhole *int32 `json:"ddos_action_blackhole,omitempty"`
	AdditionalProperties map[string]any
}

type _ValuesDdos ValuesDdos

// NewValuesDdos instantiates a new ValuesDdos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesDdos() *ValuesDdos {
	this := ValuesDdos{}
	return &this
}

// NewValuesDdosWithDefaults instantiates a new ValuesDdos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesDdosWithDefaults() *ValuesDdos {
	this := ValuesDdos{}
	return &this
}

// GetDdosActionLimitStreamsConnections returns the DdosActionLimitStreamsConnections field value if set, zero value otherwise.
func (o *ValuesDdos) GetDdosActionLimitStreamsConnections() int32 {
	if o == nil || o.DdosActionLimitStreamsConnections == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionLimitStreamsConnections
}

// GetDdosActionLimitStreamsConnectionsOk returns a tuple with the DdosActionLimitStreamsConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesDdos) GetDdosActionLimitStreamsConnectionsOk() (*int32, bool) {
	if o == nil || o.DdosActionLimitStreamsConnections == nil {
		return nil, false
	}
	return o.DdosActionLimitStreamsConnections, true
}

// HasDdosActionLimitStreamsConnections returns a boolean if a field has been set.
func (o *ValuesDdos) HasDdosActionLimitStreamsConnections() bool {
	if o != nil && o.DdosActionLimitStreamsConnections != nil {
		return true
	}

	return false
}

// SetDdosActionLimitStreamsConnections gets a reference to the given int32 and assigns it to the DdosActionLimitStreamsConnections field.
func (o *ValuesDdos) SetDdosActionLimitStreamsConnections(v int32) {
	o.DdosActionLimitStreamsConnections = &v
}

// GetDdosActionLimitStreamsRequests returns the DdosActionLimitStreamsRequests field value if set, zero value otherwise.
func (o *ValuesDdos) GetDdosActionLimitStreamsRequests() int32 {
	if o == nil || o.DdosActionLimitStreamsRequests == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionLimitStreamsRequests
}

// GetDdosActionLimitStreamsRequestsOk returns a tuple with the DdosActionLimitStreamsRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesDdos) GetDdosActionLimitStreamsRequestsOk() (*int32, bool) {
	if o == nil || o.DdosActionLimitStreamsRequests == nil {
		return nil, false
	}
	return o.DdosActionLimitStreamsRequests, true
}

// HasDdosActionLimitStreamsRequests returns a boolean if a field has been set.
func (o *ValuesDdos) HasDdosActionLimitStreamsRequests() bool {
	if o != nil && o.DdosActionLimitStreamsRequests != nil {
		return true
	}

	return false
}

// SetDdosActionLimitStreamsRequests gets a reference to the given int32 and assigns it to the DdosActionLimitStreamsRequests field.
func (o *ValuesDdos) SetDdosActionLimitStreamsRequests(v int32) {
	o.DdosActionLimitStreamsRequests = &v
}

// GetDdosActionTarpitAccept returns the DdosActionTarpitAccept field value if set, zero value otherwise.
func (o *ValuesDdos) GetDdosActionTarpitAccept() int32 {
	if o == nil || o.DdosActionTarpitAccept == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionTarpitAccept
}

// GetDdosActionTarpitAcceptOk returns a tuple with the DdosActionTarpitAccept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesDdos) GetDdosActionTarpitAcceptOk() (*int32, bool) {
	if o == nil || o.DdosActionTarpitAccept == nil {
		return nil, false
	}
	return o.DdosActionTarpitAccept, true
}

// HasDdosActionTarpitAccept returns a boolean if a field has been set.
func (o *ValuesDdos) HasDdosActionTarpitAccept() bool {
	if o != nil && o.DdosActionTarpitAccept != nil {
		return true
	}

	return false
}

// SetDdosActionTarpitAccept gets a reference to the given int32 and assigns it to the DdosActionTarpitAccept field.
func (o *ValuesDdos) SetDdosActionTarpitAccept(v int32) {
	o.DdosActionTarpitAccept = &v
}

// GetDdosActionTarpit returns the DdosActionTarpit field value if set, zero value otherwise.
func (o *ValuesDdos) GetDdosActionTarpit() int32 {
	if o == nil || o.DdosActionTarpit == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionTarpit
}

// GetDdosActionTarpitOk returns a tuple with the DdosActionTarpit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesDdos) GetDdosActionTarpitOk() (*int32, bool) {
	if o == nil || o.DdosActionTarpit == nil {
		return nil, false
	}
	return o.DdosActionTarpit, true
}

// HasDdosActionTarpit returns a boolean if a field has been set.
func (o *ValuesDdos) HasDdosActionTarpit() bool {
	if o != nil && o.DdosActionTarpit != nil {
		return true
	}

	return false
}

// SetDdosActionTarpit gets a reference to the given int32 and assigns it to the DdosActionTarpit field.
func (o *ValuesDdos) SetDdosActionTarpit(v int32) {
	o.DdosActionTarpit = &v
}

// GetDdosActionClose returns the DdosActionClose field value if set, zero value otherwise.
func (o *ValuesDdos) GetDdosActionClose() int32 {
	if o == nil || o.DdosActionClose == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionClose
}

// GetDdosActionCloseOk returns a tuple with the DdosActionClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesDdos) GetDdosActionCloseOk() (*int32, bool) {
	if o == nil || o.DdosActionClose == nil {
		return nil, false
	}
	return o.DdosActionClose, true
}

// HasDdosActionClose returns a boolean if a field has been set.
func (o *ValuesDdos) HasDdosActionClose() bool {
	if o != nil && o.DdosActionClose != nil {
		return true
	}

	return false
}

// SetDdosActionClose gets a reference to the given int32 and assigns it to the DdosActionClose field.
func (o *ValuesDdos) SetDdosActionClose(v int32) {
	o.DdosActionClose = &v
}

// GetDdosActionBlackhole returns the DdosActionBlackhole field value if set, zero value otherwise.
func (o *ValuesDdos) GetDdosActionBlackhole() int32 {
	if o == nil || o.DdosActionBlackhole == nil {
		var ret int32
		return ret
	}
	return *o.DdosActionBlackhole
}

// GetDdosActionBlackholeOk returns a tuple with the DdosActionBlackhole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesDdos) GetDdosActionBlackholeOk() (*int32, bool) {
	if o == nil || o.DdosActionBlackhole == nil {
		return nil, false
	}
	return o.DdosActionBlackhole, true
}

// HasDdosActionBlackhole returns a boolean if a field has been set.
func (o *ValuesDdos) HasDdosActionBlackhole() bool {
	if o != nil && o.DdosActionBlackhole != nil {
		return true
	}

	return false
}

// SetDdosActionBlackhole gets a reference to the given int32 and assigns it to the DdosActionBlackhole field.
func (o *ValuesDdos) SetDdosActionBlackhole(v int32) {
	o.DdosActionBlackhole = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesDdos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.DdosActionLimitStreamsConnections != nil {
		toSerialize["ddos_action_limit_streams_connections"] = o.DdosActionLimitStreamsConnections
	}
	if o.DdosActionLimitStreamsRequests != nil {
		toSerialize["ddos_action_limit_streams_requests"] = o.DdosActionLimitStreamsRequests
	}
	if o.DdosActionTarpitAccept != nil {
		toSerialize["ddos_action_tarpit_accept"] = o.DdosActionTarpitAccept
	}
	if o.DdosActionTarpit != nil {
		toSerialize["ddos_action_tarpit"] = o.DdosActionTarpit
	}
	if o.DdosActionClose != nil {
		toSerialize["ddos_action_close"] = o.DdosActionClose
	}
	if o.DdosActionBlackhole != nil {
		toSerialize["ddos_action_blackhole"] = o.DdosActionBlackhole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ValuesDdos) UnmarshalJSON(bytes []byte) (err error) {
	varValuesDdos := _ValuesDdos{}

	if err = json.Unmarshal(bytes, &varValuesDdos); err == nil {
		*o = ValuesDdos(varValuesDdos)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ddos_action_limit_streams_connections")
		delete(additionalProperties, "ddos_action_limit_streams_requests")
		delete(additionalProperties, "ddos_action_tarpit_accept")
		delete(additionalProperties, "ddos_action_tarpit")
		delete(additionalProperties, "ddos_action_close")
		delete(additionalProperties, "ddos_action_blackhole")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesDdos is a helper abstraction for handling nullable valuesddos types. 
type NullableValuesDdos struct {
	value *ValuesDdos
	isSet bool
}

// Get returns the value.
func (v NullableValuesDdos) Get() *ValuesDdos {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesDdos) Set(val *ValuesDdos) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesDdos) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesDdos) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesDdos returns a pointer to a new instance of NullableValuesDdos.
func NewNullableValuesDdos(val *ValuesDdos) *NullableValuesDdos {
	return &NullableValuesDdos{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesDdos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableValuesDdos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
