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

// SchemasVersion struct for SchemasVersion
type SchemasVersion struct {
	// Whether this is the active version or not.
	Active *bool `json:"active,omitempty"`
	// A freeform descriptive note.
	Comment NullableString `json:"comment,omitempty"`
	// Unused at this time.
	Deployed *bool `json:"deployed,omitempty"`
	// Whether this version is locked or not. Objects can not be added or edited on locked versions.
	Locked *bool `json:"locked,omitempty"`
	// The number of this version.
	Number *int32 `json:"number,omitempty"`
	// Unused at this time.
	Staging *bool `json:"staging,omitempty"`
	// Unused at this time.
	Testing *bool `json:"testing,omitempty"`
	AdditionalProperties map[string]any
}

type _SchemasVersion SchemasVersion

// NewSchemasVersion instantiates a new SchemasVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasVersion() *SchemasVersion {
	this := SchemasVersion{}
	var active bool = false
	this.Active = &active
	var locked bool = false
	this.Locked = &locked
	var staging bool = false
	this.Staging = &staging
	var testing bool = false
	this.Testing = &testing
	return &this
}

// NewSchemasVersionWithDefaults instantiates a new SchemasVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasVersionWithDefaults() *SchemasVersion {
	this := SchemasVersion{}
	var active bool = false
	this.Active = &active
	var locked bool = false
	this.Locked = &locked
	var staging bool = false
	this.Staging = &staging
	var testing bool = false
	this.Testing = &testing
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SchemasVersion) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersion) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SchemasVersion) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SchemasVersion) SetActive(v bool) {
	o.Active = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasVersion) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasVersion) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *SchemasVersion) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *SchemasVersion) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *SchemasVersion) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *SchemasVersion) UnsetComment() {
	o.Comment.Unset()
}

// GetDeployed returns the Deployed field value if set, zero value otherwise.
func (o *SchemasVersion) GetDeployed() bool {
	if o == nil || o.Deployed == nil {
		var ret bool
		return ret
	}
	return *o.Deployed
}

// GetDeployedOk returns a tuple with the Deployed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersion) GetDeployedOk() (*bool, bool) {
	if o == nil || o.Deployed == nil {
		return nil, false
	}
	return o.Deployed, true
}

// HasDeployed returns a boolean if a field has been set.
func (o *SchemasVersion) HasDeployed() bool {
	if o != nil && o.Deployed != nil {
		return true
	}

	return false
}

// SetDeployed gets a reference to the given bool and assigns it to the Deployed field.
func (o *SchemasVersion) SetDeployed(v bool) {
	o.Deployed = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *SchemasVersion) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersion) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *SchemasVersion) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *SchemasVersion) SetLocked(v bool) {
	o.Locked = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SchemasVersion) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersion) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SchemasVersion) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *SchemasVersion) SetNumber(v int32) {
	o.Number = &v
}

// GetStaging returns the Staging field value if set, zero value otherwise.
func (o *SchemasVersion) GetStaging() bool {
	if o == nil || o.Staging == nil {
		var ret bool
		return ret
	}
	return *o.Staging
}

// GetStagingOk returns a tuple with the Staging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersion) GetStagingOk() (*bool, bool) {
	if o == nil || o.Staging == nil {
		return nil, false
	}
	return o.Staging, true
}

// HasStaging returns a boolean if a field has been set.
func (o *SchemasVersion) HasStaging() bool {
	if o != nil && o.Staging != nil {
		return true
	}

	return false
}

// SetStaging gets a reference to the given bool and assigns it to the Staging field.
func (o *SchemasVersion) SetStaging(v bool) {
	o.Staging = &v
}

// GetTesting returns the Testing field value if set, zero value otherwise.
func (o *SchemasVersion) GetTesting() bool {
	if o == nil || o.Testing == nil {
		var ret bool
		return ret
	}
	return *o.Testing
}

// GetTestingOk returns a tuple with the Testing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersion) GetTestingOk() (*bool, bool) {
	if o == nil || o.Testing == nil {
		return nil, false
	}
	return o.Testing, true
}

// HasTesting returns a boolean if a field has been set.
func (o *SchemasVersion) HasTesting() bool {
	if o != nil && o.Testing != nil {
		return true
	}

	return false
}

// SetTesting gets a reference to the given bool and assigns it to the Testing field.
func (o *SchemasVersion) SetTesting(v bool) {
	o.Testing = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SchemasVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Deployed != nil {
		toSerialize["deployed"] = o.Deployed
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Staging != nil {
		toSerialize["staging"] = o.Staging
	}
	if o.Testing != nil {
		toSerialize["testing"] = o.Testing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *SchemasVersion) UnmarshalJSON(bytes []byte) (err error) {
	varSchemasVersion := _SchemasVersion{}

	if err = json.Unmarshal(bytes, &varSchemasVersion); err == nil {
		*o = SchemasVersion(varSchemasVersion)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "deployed")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "number")
		delete(additionalProperties, "staging")
		delete(additionalProperties, "testing")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSchemasVersion is a helper abstraction for handling nullable schemasversion types. 
type NullableSchemasVersion struct {
	value *SchemasVersion
	isSet bool
}

// Get returns the value.
func (v NullableSchemasVersion) Get() *SchemasVersion {
	return v.value
}

// Set modifies the value.
func (v *NullableSchemasVersion) Set(val *SchemasVersion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSchemasVersion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSchemasVersion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSchemasVersion returns a pointer to a new instance of NullableSchemasVersion.
func NewNullableSchemasVersion(val *SchemasVersion) *NullableSchemasVersion {
	return &NullableSchemasVersion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSchemasVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSchemasVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
