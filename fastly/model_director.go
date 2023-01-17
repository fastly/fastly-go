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
)

// Director struct for Director
type Director struct {
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
	// Selected POP to serve as a shield for the backends. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](/reference/api/utils/pops/) to get a list of available POPs used for shielding.
	Shield NullableString `json:"shield,omitempty"`
	// What type of load balance group to use.
	Type *int32 `json:"type,omitempty"`
	// How many backends to search if it fails.
	Retries *int32 `json:"retries,omitempty"`
	AdditionalProperties map[string]any
}

type _Director Director

// NewDirector instantiates a new Director object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirector() *Director {
	this := Director{}
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

// NewDirectorWithDefaults instantiates a new Director object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectorWithDefaults() *Director {
	this := Director{}
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
func (o *Director) GetBackends() []Backend {
	if o == nil || o.Backends == nil {
		var ret []Backend
		return ret
	}
	return o.Backends
}

// GetBackendsOk returns a tuple with the Backends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Director) GetBackendsOk() ([]Backend, bool) {
	if o == nil || o.Backends == nil {
		return nil, false
	}
	return o.Backends, true
}

// HasBackends returns a boolean if a field has been set.
func (o *Director) HasBackends() bool {
	if o != nil && o.Backends != nil {
		return true
	}

	return false
}

// SetBackends gets a reference to the given []Backend and assigns it to the Backends field.
func (o *Director) SetBackends(v []Backend) {
	o.Backends = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *Director) GetCapacity() int32 {
	if o == nil || o.Capacity == nil {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Director) GetCapacityOk() (*int32, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *Director) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *Director) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Director) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Director) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *Director) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *Director) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *Director) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *Director) UnsetComment() {
	o.Comment.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Director) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Director) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Director) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Director) SetName(v string) {
	o.Name = &v
}

// GetQuorum returns the Quorum field value if set, zero value otherwise.
func (o *Director) GetQuorum() int32 {
	if o == nil || o.Quorum == nil {
		var ret int32
		return ret
	}
	return *o.Quorum
}

// GetQuorumOk returns a tuple with the Quorum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Director) GetQuorumOk() (*int32, bool) {
	if o == nil || o.Quorum == nil {
		return nil, false
	}
	return o.Quorum, true
}

// HasQuorum returns a boolean if a field has been set.
func (o *Director) HasQuorum() bool {
	if o != nil && o.Quorum != nil {
		return true
	}

	return false
}

// SetQuorum gets a reference to the given int32 and assigns it to the Quorum field.
func (o *Director) SetQuorum(v int32) {
	o.Quorum = &v
}

// GetShield returns the Shield field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Director) GetShield() string {
	if o == nil || o.Shield.Get() == nil {
		var ret string
		return ret
	}
	return *o.Shield.Get()
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Director) GetShieldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Shield.Get(), o.Shield.IsSet()
}

// HasShield returns a boolean if a field has been set.
func (o *Director) HasShield() bool {
	if o != nil && o.Shield.IsSet() {
		return true
	}

	return false
}

// SetShield gets a reference to the given NullableString and assigns it to the Shield field.
func (o *Director) SetShield(v string) {
	o.Shield.Set(&v)
}
// SetShieldNil sets the value for Shield to be an explicit nil
func (o *Director) SetShieldNil() {
	o.Shield.Set(nil)
}

// UnsetShield ensures that no value is present for Shield, not even an explicit nil
func (o *Director) UnsetShield() {
	o.Shield.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Director) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Director) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Director) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *Director) SetType(v int32) {
	o.Type = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *Director) GetRetries() int32 {
	if o == nil || o.Retries == nil {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Director) GetRetriesOk() (*int32, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *Director) HasRetries() bool {
	if o != nil && o.Retries != nil {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *Director) SetRetries(v int32) {
	o.Retries = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Director) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Director) UnmarshalJSON(bytes []byte) (err error) {
	varDirector := _Director{}

	if err = json.Unmarshal(bytes, &varDirector); err == nil {
		*o = Director(varDirector)
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDirector is a helper abstraction for handling nullable director types. 
type NullableDirector struct {
	value *Director
	isSet bool
}

// Get returns the value.
func (v NullableDirector) Get() *Director {
	return v.value
}

// Set modifies the value.
func (v *NullableDirector) Set(val *Director) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDirector) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDirector) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDirector returns a pointer to a new instance of NullableDirector.
func NewNullableDirector(val *Director) *NullableDirector {
	return &NullableDirector{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDirector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableDirector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
