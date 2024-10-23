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

// DirectorBackend struct for DirectorBackend
type DirectorBackend struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ServiceID *string      `json:"service_id,omitempty"`
	Version   *int32       `json:"version,omitempty"`
	// The name of the backend.
	BackendName *string `json:"backend_name,omitempty"`
	// Name for the Director.
	Director             *string `json:"director,omitempty"`
	AdditionalProperties map[string]any
}

type _DirectorBackend DirectorBackend

// NewDirectorBackend instantiates a new DirectorBackend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectorBackend() *DirectorBackend {
	this := DirectorBackend{}
	return &this
}

// NewDirectorBackendWithDefaults instantiates a new DirectorBackend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectorBackendWithDefaults() *DirectorBackend {
	this := DirectorBackend{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectorBackend) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectorBackend) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DirectorBackend) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *DirectorBackend) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *DirectorBackend) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *DirectorBackend) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectorBackend) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectorBackend) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DirectorBackend) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *DirectorBackend) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *DirectorBackend) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *DirectorBackend) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectorBackend) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectorBackend) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DirectorBackend) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *DirectorBackend) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *DirectorBackend) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *DirectorBackend) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *DirectorBackend) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorBackend) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *DirectorBackend) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *DirectorBackend) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DirectorBackend) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorBackend) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DirectorBackend) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DirectorBackend) SetVersion(v int32) {
	o.Version = &v
}

// GetBackendName returns the BackendName field value if set, zero value otherwise.
func (o *DirectorBackend) GetBackendName() string {
	if o == nil || o.BackendName == nil {
		var ret string
		return ret
	}
	return *o.BackendName
}

// GetBackendNameOk returns a tuple with the BackendName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorBackend) GetBackendNameOk() (*string, bool) {
	if o == nil || o.BackendName == nil {
		return nil, false
	}
	return o.BackendName, true
}

// HasBackendName returns a boolean if a field has been set.
func (o *DirectorBackend) HasBackendName() bool {
	if o != nil && o.BackendName != nil {
		return true
	}

	return false
}

// SetBackendName gets a reference to the given string and assigns it to the BackendName field.
func (o *DirectorBackend) SetBackendName(v string) {
	o.BackendName = &v
}

// GetDirector returns the Director field value if set, zero value otherwise.
func (o *DirectorBackend) GetDirector() string {
	if o == nil || o.Director == nil {
		var ret string
		return ret
	}
	return *o.Director
}

// GetDirectorOk returns a tuple with the Director field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorBackend) GetDirectorOk() (*string, bool) {
	if o == nil || o.Director == nil {
		return nil, false
	}
	return o.Director, true
}

// HasDirector returns a boolean if a field has been set.
func (o *DirectorBackend) HasDirector() bool {
	if o != nil && o.Director != nil {
		return true
	}

	return false
}

// SetDirector gets a reference to the given string and assigns it to the Director field.
func (o *DirectorBackend) SetDirector(v string) {
	o.Director = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DirectorBackend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.BackendName != nil {
		toSerialize["backend_name"] = o.BackendName
	}
	if o.Director != nil {
		toSerialize["director"] = o.Director
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DirectorBackend) UnmarshalJSON(bytes []byte) (err error) {
	varDirectorBackend := _DirectorBackend{}

	if err = json.Unmarshal(bytes, &varDirectorBackend); err == nil {
		*o = DirectorBackend(varDirectorBackend)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "backend_name")
		delete(additionalProperties, "director")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDirectorBackend is a helper abstraction for handling nullable directorbackend types.
type NullableDirectorBackend struct {
	value *DirectorBackend
	isSet bool
}

// Get returns the value.
func (v NullableDirectorBackend) Get() *DirectorBackend {
	return v.value
}

// Set modifies the value.
func (v *NullableDirectorBackend) Set(val *DirectorBackend) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDirectorBackend) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDirectorBackend) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDirectorBackend returns a pointer to a new instance of NullableDirectorBackend.
func NewNullableDirectorBackend(val *DirectorBackend) *NullableDirectorBackend {
	return &NullableDirectorBackend{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDirectorBackend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDirectorBackend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
