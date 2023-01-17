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

// SchemasVersionResponse struct for SchemasVersionResponse
type SchemasVersionResponse struct {
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
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	AdditionalProperties map[string]any
}

type _SchemasVersionResponse SchemasVersionResponse

// NewSchemasVersionResponse instantiates a new SchemasVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasVersionResponse() *SchemasVersionResponse {
	this := SchemasVersionResponse{}
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

// NewSchemasVersionResponseWithDefaults instantiates a new SchemasVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasVersionResponseWithDefaults() *SchemasVersionResponse {
	this := SchemasVersionResponse{}
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
func (o *SchemasVersionResponse) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersionResponse) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SchemasVersionResponse) SetActive(v bool) {
	o.Active = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasVersionResponse) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasVersionResponse) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *SchemasVersionResponse) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *SchemasVersionResponse) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *SchemasVersionResponse) UnsetComment() {
	o.Comment.Unset()
}

// GetDeployed returns the Deployed field value if set, zero value otherwise.
func (o *SchemasVersionResponse) GetDeployed() bool {
	if o == nil || o.Deployed == nil {
		var ret bool
		return ret
	}
	return *o.Deployed
}

// GetDeployedOk returns a tuple with the Deployed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersionResponse) GetDeployedOk() (*bool, bool) {
	if o == nil || o.Deployed == nil {
		return nil, false
	}
	return o.Deployed, true
}

// HasDeployed returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasDeployed() bool {
	if o != nil && o.Deployed != nil {
		return true
	}

	return false
}

// SetDeployed gets a reference to the given bool and assigns it to the Deployed field.
func (o *SchemasVersionResponse) SetDeployed(v bool) {
	o.Deployed = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *SchemasVersionResponse) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersionResponse) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *SchemasVersionResponse) SetLocked(v bool) {
	o.Locked = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *SchemasVersionResponse) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersionResponse) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *SchemasVersionResponse) SetNumber(v int32) {
	o.Number = &v
}

// GetStaging returns the Staging field value if set, zero value otherwise.
func (o *SchemasVersionResponse) GetStaging() bool {
	if o == nil || o.Staging == nil {
		var ret bool
		return ret
	}
	return *o.Staging
}

// GetStagingOk returns a tuple with the Staging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersionResponse) GetStagingOk() (*bool, bool) {
	if o == nil || o.Staging == nil {
		return nil, false
	}
	return o.Staging, true
}

// HasStaging returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasStaging() bool {
	if o != nil && o.Staging != nil {
		return true
	}

	return false
}

// SetStaging gets a reference to the given bool and assigns it to the Staging field.
func (o *SchemasVersionResponse) SetStaging(v bool) {
	o.Staging = &v
}

// GetTesting returns the Testing field value if set, zero value otherwise.
func (o *SchemasVersionResponse) GetTesting() bool {
	if o == nil || o.Testing == nil {
		var ret bool
		return ret
	}
	return *o.Testing
}

// GetTestingOk returns a tuple with the Testing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersionResponse) GetTestingOk() (*bool, bool) {
	if o == nil || o.Testing == nil {
		return nil, false
	}
	return o.Testing, true
}

// HasTesting returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasTesting() bool {
	if o != nil && o.Testing != nil {
		return true
	}

	return false
}

// SetTesting gets a reference to the given bool and assigns it to the Testing field.
func (o *SchemasVersionResponse) SetTesting(v bool) {
	o.Testing = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasVersionResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasVersionResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *SchemasVersionResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *SchemasVersionResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *SchemasVersionResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasVersionResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasVersionResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *SchemasVersionResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *SchemasVersionResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *SchemasVersionResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasVersionResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasVersionResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *SchemasVersionResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *SchemasVersionResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *SchemasVersionResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *SchemasVersionResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasVersionResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *SchemasVersionResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *SchemasVersionResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SchemasVersionResponse) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *SchemasVersionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSchemasVersionResponse := _SchemasVersionResponse{}

	if err = json.Unmarshal(bytes, &varSchemasVersionResponse); err == nil {
		*o = SchemasVersionResponse(varSchemasVersionResponse)
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
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "service_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSchemasVersionResponse is a helper abstraction for handling nullable schemasversionresponse types. 
type NullableSchemasVersionResponse struct {
	value *SchemasVersionResponse
	isSet bool
}

// Get returns the value.
func (v NullableSchemasVersionResponse) Get() *SchemasVersionResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableSchemasVersionResponse) Set(val *SchemasVersionResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSchemasVersionResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSchemasVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSchemasVersionResponse returns a pointer to a new instance of NullableSchemasVersionResponse.
func NewNullableSchemasVersionResponse(val *SchemasVersionResponse) *NullableSchemasVersionResponse {
	return &NullableSchemasVersionResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSchemasVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSchemasVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
