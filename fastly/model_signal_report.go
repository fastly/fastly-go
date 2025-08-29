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

// SignalReport struct for SignalReport
type SignalReport struct {
	// Name of the attack type.
	Name *string `json:"name,omitempty"`
	// Display name of the attack type.
	DisplayName *string `json:"display_name,omitempty"`
	// Total count of attacks of this type.
	Count *int32 `json:"count,omitempty"`
	// Top workspaces affected by this attack type.
	TopWorkspaces        []TopWorkspace `json:"top_workspaces,omitempty"`
	AdditionalProperties map[string]any
}

type _SignalReport SignalReport

// NewSignalReport instantiates a new SignalReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalReport() *SignalReport {
	this := SignalReport{}
	return &this
}

// NewSignalReportWithDefaults instantiates a new SignalReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalReportWithDefaults() *SignalReport {
	this := SignalReport{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SignalReport) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalReport) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SignalReport) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SignalReport) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SignalReport) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalReport) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SignalReport) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SignalReport) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SignalReport) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalReport) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SignalReport) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SignalReport) SetCount(v int32) {
	o.Count = &v
}

// GetTopWorkspaces returns the TopWorkspaces field value if set, zero value otherwise.
func (o *SignalReport) GetTopWorkspaces() []TopWorkspace {
	if o == nil || o.TopWorkspaces == nil {
		var ret []TopWorkspace
		return ret
	}
	return o.TopWorkspaces
}

// GetTopWorkspacesOk returns a tuple with the TopWorkspaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalReport) GetTopWorkspacesOk() ([]TopWorkspace, bool) {
	if o == nil || o.TopWorkspaces == nil {
		return nil, false
	}
	return o.TopWorkspaces, true
}

// HasTopWorkspaces returns a boolean if a field has been set.
func (o *SignalReport) HasTopWorkspaces() bool {
	if o != nil && o.TopWorkspaces != nil {
		return true
	}

	return false
}

// SetTopWorkspaces gets a reference to the given []TopWorkspace and assigns it to the TopWorkspaces field.
func (o *SignalReport) SetTopWorkspaces(v []TopWorkspace) {
	o.TopWorkspaces = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SignalReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.TopWorkspaces != nil {
		toSerialize["top_workspaces"] = o.TopWorkspaces
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *SignalReport) UnmarshalJSON(bytes []byte) (err error) {
	varSignalReport := _SignalReport{}

	if err = json.Unmarshal(bytes, &varSignalReport); err == nil {
		*o = SignalReport(varSignalReport)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "count")
		delete(additionalProperties, "top_workspaces")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSignalReport is a helper abstraction for handling nullable signalreport types.
type NullableSignalReport struct {
	value *SignalReport
	isSet bool
}

// Get returns the value.
func (v NullableSignalReport) Get() *SignalReport {
	return v.value
}

// Set modifies the value.
func (v *NullableSignalReport) Set(val *SignalReport) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSignalReport) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSignalReport) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSignalReport returns a pointer to a new instance of NullableSignalReport.
func NewNullableSignalReport(val *SignalReport) *NullableSignalReport {
	return &NullableSignalReport{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSignalReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSignalReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
