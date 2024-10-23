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

// DirectorResponse struct for DirectorResponse
type DirectorResponse struct {
	// List of backends associated to a director.
	Backends []Backend `json:"backends,omitempty"`
	// Unused.
	Capacity *int32 `json:"capacity,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// Name for the Director.
	Name *string `json:"name,omitempty"`
	// The percentage of capacity that needs to be up for a director to be considered up. `0` to `100`.
	Quorum *int32 `json:"quorum,omitempty"`
	// Selected POP to serve as a shield for the backends. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding.
	Shield NullableString `json:"shield,omitempty"`
	// What type of load balance group to use.
	Type *int32 `json:"type,omitempty"`
	// How many backends to search if it fails.
	Retries   *int32  `json:"retries,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	Version   *int32  `json:"version,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt            NullableTime `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _DirectorResponse DirectorResponse

// NewDirectorResponse instantiates a new DirectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectorResponse() *DirectorResponse {
	this := DirectorResponse{}
	var quorum int32 = 75
	this.Quorum = &quorum
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var resourceType int32 = 1
	this.Type = &resourceType
	var retries int32 = 5
	this.Retries = &retries
	return &this
}

// NewDirectorResponseWithDefaults instantiates a new DirectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectorResponseWithDefaults() *DirectorResponse {
	this := DirectorResponse{}
	var quorum int32 = 75
	this.Quorum = &quorum
	var shield string = "null"
	this.Shield = *NewNullableString(&shield)
	var resourceType int32 = 1
	this.Type = &resourceType
	var retries int32 = 5
	this.Retries = &retries
	return &this
}

// GetBackends returns the Backends field value if set, zero value otherwise.
func (o *DirectorResponse) GetBackends() []Backend {
	if o == nil || o.Backends == nil {
		var ret []Backend
		return ret
	}
	return o.Backends
}

// GetBackendsOk returns a tuple with the Backends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorResponse) GetBackendsOk() ([]Backend, bool) {
	if o == nil || o.Backends == nil {
		return nil, false
	}
	return o.Backends, true
}

// HasBackends returns a boolean if a field has been set.
func (o *DirectorResponse) HasBackends() bool {
	if o != nil && o.Backends != nil {
		return true
	}

	return false
}

// SetBackends gets a reference to the given []Backend and assigns it to the Backends field.
func (o *DirectorResponse) SetBackends(v []Backend) {
	o.Backends = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *DirectorResponse) GetCapacity() int32 {
	if o == nil || o.Capacity == nil {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorResponse) GetCapacityOk() (*int32, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *DirectorResponse) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *DirectorResponse) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectorResponse) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectorResponse) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *DirectorResponse) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *DirectorResponse) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *DirectorResponse) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *DirectorResponse) UnsetComment() {
	o.Comment.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DirectorResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DirectorResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DirectorResponse) SetName(v string) {
	o.Name = &v
}

// GetQuorum returns the Quorum field value if set, zero value otherwise.
func (o *DirectorResponse) GetQuorum() int32 {
	if o == nil || o.Quorum == nil {
		var ret int32
		return ret
	}
	return *o.Quorum
}

// GetQuorumOk returns a tuple with the Quorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorResponse) GetQuorumOk() (*int32, bool) {
	if o == nil || o.Quorum == nil {
		return nil, false
	}
	return o.Quorum, true
}

// HasQuorum returns a boolean if a field has been set.
func (o *DirectorResponse) HasQuorum() bool {
	if o != nil && o.Quorum != nil {
		return true
	}

	return false
}

// SetQuorum gets a reference to the given int32 and assigns it to the Quorum field.
func (o *DirectorResponse) SetQuorum(v int32) {
	o.Quorum = &v
}

// GetShield returns the Shield field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectorResponse) GetShield() string {
	if o == nil || o.Shield.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shield.Get()
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectorResponse) GetShieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shield.Get(), o.Shield.IsSet()
}

// HasShield returns a boolean if a field has been set.
func (o *DirectorResponse) HasShield() bool {
	if o != nil && o.Shield.IsSet() {
		return true
	}

	return false
}

// SetShield gets a reference to the given NullableString and assigns it to the Shield field.
func (o *DirectorResponse) SetShield(v string) {
	o.Shield.Set(&v)
}

// SetShieldNil sets the value for Shield to be an explicit nil
func (o *DirectorResponse) SetShieldNil() {
	o.Shield.Set(nil)
}

// UnsetShield ensures that no value is present for Shield, not even an explicit nil
func (o *DirectorResponse) UnsetShield() {
	o.Shield.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DirectorResponse) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorResponse) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DirectorResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *DirectorResponse) SetType(v int32) {
	o.Type = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *DirectorResponse) GetRetries() int32 {
	if o == nil || o.Retries == nil {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorResponse) GetRetriesOk() (*int32, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *DirectorResponse) HasRetries() bool {
	if o != nil && o.Retries != nil {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *DirectorResponse) SetRetries(v int32) {
	o.Retries = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *DirectorResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *DirectorResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *DirectorResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DirectorResponse) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectorResponse) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DirectorResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DirectorResponse) SetVersion(v int32) {
	o.Version = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectorResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectorResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DirectorResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *DirectorResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *DirectorResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *DirectorResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectorResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectorResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DirectorResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *DirectorResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *DirectorResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *DirectorResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectorResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectorResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DirectorResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *DirectorResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *DirectorResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *DirectorResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DirectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Backends != nil {
		toSerialize["backends"] = o.Backends
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Quorum != nil {
		toSerialize["quorum"] = o.Quorum
	}
	if o.Shield.IsSet() {
		toSerialize["shield"] = o.Shield.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}
	if o.ServiceID != nil {
		toSerialize["service_id"] = o.ServiceID
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DirectorResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDirectorResponse := _DirectorResponse{}

	if err = json.Unmarshal(bytes, &varDirectorResponse); err == nil {
		*o = DirectorResponse(varDirectorResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "backends")
		delete(additionalProperties, "capacity")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "name")
		delete(additionalProperties, "quorum")
		delete(additionalProperties, "shield")
		delete(additionalProperties, "type")
		delete(additionalProperties, "retries")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDirectorResponse is a helper abstraction for handling nullable directorresponse types.
type NullableDirectorResponse struct {
	value *DirectorResponse
	isSet bool
}

// Get returns the value.
func (v NullableDirectorResponse) Get() *DirectorResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableDirectorResponse) Set(val *DirectorResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDirectorResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDirectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDirectorResponse returns a pointer to a new instance of NullableDirectorResponse.
func NewNullableDirectorResponse(val *DirectorResponse) *NullableDirectorResponse {
	return &NullableDirectorResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDirectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDirectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
