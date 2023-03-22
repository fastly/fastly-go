// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
	"time"
)

// ServerResponse struct for ServerResponse
type ServerResponse struct {
	// Weight (`1-100`) used to load balance this server against others.
	Weight *int32 `json:"weight,omitempty"`
	// Maximum number of connections. If the value is `0`, it inherits the value from pool's `max_conn_default`.
	MaxConn *int32 `json:"max_conn,omitempty"`
	// Port number. Setting port `443` does not force TLS. Set `use_tls` in pool to force TLS.
	Port *int32 `json:"port,omitempty"`
	// A hostname, IPv4, or IPv6 address for the server. Required.
	Address *string `json:"address,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// Allows servers to be enabled and disabled in a pool.
	Disabled *bool `json:"disabled,omitempty"`
	// The hostname to override the Host header. Defaults to `null` meaning no override of the Host header if not set. This setting can also be added to a Pool definition. However, the server setting will override the Pool setting.
	OverrideHost NullableString `json:"override_host,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	ID *string `json:"id,omitempty"`
	PoolID *string `json:"pool_id,omitempty"`
	AdditionalProperties map[string]any
}

type _ServerResponse ServerResponse

// NewServerResponse instantiates a new ServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerResponse() *ServerResponse {
	this := ServerResponse{}
	var weight int32 = 100
	this.Weight = &weight
	var maxConn int32 = 0
	this.MaxConn = &maxConn
	var port int32 = 80
	this.Port = &port
	var disabled bool = false
	this.Disabled = &disabled
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	return &this
}

// NewServerResponseWithDefaults instantiates a new ServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerResponseWithDefaults() *ServerResponse {
	this := ServerResponse{}
	var weight int32 = 100
	this.Weight = &weight
	var maxConn int32 = 0
	this.MaxConn = &maxConn
	var port int32 = 80
	this.Port = &port
	var disabled bool = false
	this.Disabled = &disabled
	var overrideHost string = "null"
	this.OverrideHost = *NewNullableString(&overrideHost)
	return &this
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ServerResponse) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponse) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ServerResponse) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *ServerResponse) SetWeight(v int32) {
	o.Weight = &v
}

// GetMaxConn returns the MaxConn field value if set, zero value otherwise.
func (o *ServerResponse) GetMaxConn() int32 {
	if o == nil || o.MaxConn == nil {
		var ret int32
		return ret
	}
	return *o.MaxConn
}

// GetMaxConnOk returns a tuple with the MaxConn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponse) GetMaxConnOk() (*int32, bool) {
	if o == nil || o.MaxConn == nil {
		return nil, false
	}
	return o.MaxConn, true
}

// HasMaxConn returns a boolean if a field has been set.
func (o *ServerResponse) HasMaxConn() bool {
	if o != nil && o.MaxConn != nil {
		return true
	}

	return false
}

// SetMaxConn gets a reference to the given int32 and assigns it to the MaxConn field.
func (o *ServerResponse) SetMaxConn(v int32) {
	o.MaxConn = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ServerResponse) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponse) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ServerResponse) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ServerResponse) SetPort(v int32) {
	o.Port = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ServerResponse) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponse) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ServerResponse) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ServerResponse) SetAddress(v string) {
	o.Address = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerResponse) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerResponse) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *ServerResponse) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *ServerResponse) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *ServerResponse) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *ServerResponse) UnsetComment() {
	o.Comment.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ServerResponse) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponse) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ServerResponse) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ServerResponse) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetOverrideHost returns the OverrideHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerResponse) GetOverrideHost() string {
	if o == nil || o.OverrideHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideHost.Get()
}

// GetOverrideHostOk returns a tuple with the OverrideHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerResponse) GetOverrideHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverrideHost.Get(), o.OverrideHost.IsSet()
}

// HasOverrideHost returns a boolean if a field has been set.
func (o *ServerResponse) HasOverrideHost() bool {
	if o != nil && o.OverrideHost.IsSet() {
		return true
	}

	return false
}

// SetOverrideHost gets a reference to the given NullableString and assigns it to the OverrideHost field.
func (o *ServerResponse) SetOverrideHost(v string) {
	o.OverrideHost.Set(&v)
}
// SetOverrideHostNil sets the value for OverrideHost to be an explicit nil
func (o *ServerResponse) SetOverrideHostNil() {
	o.OverrideHost.Set(nil)
}

// UnsetOverrideHost ensures that no value is present for OverrideHost, not even an explicit nil
func (o *ServerResponse) UnsetOverrideHost() {
	o.OverrideHost.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServerResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *ServerResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ServerResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ServerResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ServerResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *ServerResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *ServerResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *ServerResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServerResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *ServerResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ServerResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ServerResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *ServerResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *ServerResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *ServerResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ServerResponse) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponse) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ServerResponse) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ServerResponse) SetID(v string) {
	o.ID = &v
}

// GetPoolID returns the PoolID field value if set, zero value otherwise.
func (o *ServerResponse) GetPoolID() string {
	if o == nil || o.PoolID == nil {
		var ret string
		return ret
	}
	return *o.PoolID
}

// GetPoolIDOk returns a tuple with the PoolID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerResponse) GetPoolIDOk() (*string, bool) {
	if o == nil || o.PoolID == nil {
		return nil, false
	}
	return o.PoolID, true
}

// HasPoolID returns a boolean if a field has been set.
func (o *ServerResponse) HasPoolID() bool {
	if o != nil && o.PoolID != nil {
		return true
	}

	return false
}

// SetPoolID gets a reference to the given string and assigns it to the PoolID field.
func (o *ServerResponse) SetPoolID(v string) {
	o.PoolID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.MaxConn != nil {
		toSerialize["max_conn"] = o.MaxConn
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.OverrideHost.IsSet() {
		toSerialize["override_host"] = o.OverrideHost.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.PoolID != nil {
		toSerialize["pool_id"] = o.PoolID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServerResponse) UnmarshalJSON(bytes []byte) (err error) {
	varServerResponse := _ServerResponse{}

	if err = json.Unmarshal(bytes, &varServerResponse); err == nil {
		*o = ServerResponse(varServerResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "weight")
		delete(additionalProperties, "max_conn")
		delete(additionalProperties, "port")
		delete(additionalProperties, "address")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "override_host")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "pool_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServerResponse is a helper abstraction for handling nullable serverresponse types. 
type NullableServerResponse struct {
	value *ServerResponse
	isSet bool
}

// Get returns the value.
func (v NullableServerResponse) Get() *ServerResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableServerResponse) Set(val *ServerResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServerResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServerResponse returns a pointer to a new instance of NullableServerResponse.
func NewNullableServerResponse(val *ServerResponse) *NullableServerResponse {
	return &NullableServerResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
