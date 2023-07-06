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

// LegacyWafUpdateStatus struct for LegacyWafUpdateStatus
type LegacyWafUpdateStatus struct {
	// Date and time that job was completed.
	CompletedAt *string `json:"completed_at,omitempty"`
	// Date and time that job was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// This field can contain data passed to the background worker as well as output from the background job.
	Data *string `json:"data,omitempty"`
	// Message with information about the status of the update.
	Message *string `json:"message,omitempty"`
	// Current status of the update.
	Status *string `json:"status,omitempty"`
	// Date and time that job was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _LegacyWafUpdateStatus LegacyWafUpdateStatus

// NewLegacyWafUpdateStatus instantiates a new LegacyWafUpdateStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyWafUpdateStatus() *LegacyWafUpdateStatus {
	this := LegacyWafUpdateStatus{}
	return &this
}

// NewLegacyWafUpdateStatusWithDefaults instantiates a new LegacyWafUpdateStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyWafUpdateStatusWithDefaults() *LegacyWafUpdateStatus {
	this := LegacyWafUpdateStatus{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *LegacyWafUpdateStatus) GetCompletedAt() string {
	if o == nil || o.CompletedAt == nil {
		var ret string
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafUpdateStatus) GetCompletedAtOk() (*string, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *LegacyWafUpdateStatus) HasCompletedAt() bool {
	if o != nil && o.CompletedAt != nil {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given string and assigns it to the CompletedAt field.
func (o *LegacyWafUpdateStatus) SetCompletedAt(v string) {
	o.CompletedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LegacyWafUpdateStatus) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafUpdateStatus) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LegacyWafUpdateStatus) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LegacyWafUpdateStatus) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LegacyWafUpdateStatus) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafUpdateStatus) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LegacyWafUpdateStatus) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *LegacyWafUpdateStatus) SetData(v string) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LegacyWafUpdateStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafUpdateStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LegacyWafUpdateStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LegacyWafUpdateStatus) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LegacyWafUpdateStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafUpdateStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LegacyWafUpdateStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LegacyWafUpdateStatus) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LegacyWafUpdateStatus) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyWafUpdateStatus) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LegacyWafUpdateStatus) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *LegacyWafUpdateStatus) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LegacyWafUpdateStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CompletedAt != nil {
		toSerialize["completed_at"] = o.CompletedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LegacyWafUpdateStatus) UnmarshalJSON(bytes []byte) (err error) {
	varLegacyWafUpdateStatus := _LegacyWafUpdateStatus{}

	if err = json.Unmarshal(bytes, &varLegacyWafUpdateStatus); err == nil {
		*o = LegacyWafUpdateStatus(varLegacyWafUpdateStatus)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "data")
		delete(additionalProperties, "message")
		delete(additionalProperties, "status")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLegacyWafUpdateStatus is a helper abstraction for handling nullable legacywafupdatestatus types. 
type NullableLegacyWafUpdateStatus struct {
	value *LegacyWafUpdateStatus
	isSet bool
}

// Get returns the value.
func (v NullableLegacyWafUpdateStatus) Get() *LegacyWafUpdateStatus {
	return v.value
}

// Set modifies the value.
func (v *NullableLegacyWafUpdateStatus) Set(val *LegacyWafUpdateStatus) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLegacyWafUpdateStatus) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLegacyWafUpdateStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLegacyWafUpdateStatus returns a pointer to a new instance of NullableLegacyWafUpdateStatus.
func NewNullableLegacyWafUpdateStatus(val *LegacyWafUpdateStatus) *NullableLegacyWafUpdateStatus {
	return &NullableLegacyWafUpdateStatus{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLegacyWafUpdateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLegacyWafUpdateStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
