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

// Report struct for Report
type Report struct {
	ID                   *string    `json:"id,omitempty"`
	PolicyID             *string    `json:"policy_id,omitempty"`
	BlockedURI           *string    `json:"blocked_uri,omitempty"`
	DocumentURI          *string    `json:"document_uri,omitempty"`
	ViolatedDirective    *string    `json:"violated_directive,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	AdditionalProperties map[string]any
}

type _Report Report

// NewReport instantiates a new Report object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReport() *Report {
	this := Report{}
	return &this
}

// NewReportWithDefaults instantiates a new Report object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportWithDefaults() *Report {
	this := Report{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *Report) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *Report) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *Report) SetID(v string) {
	o.ID = &v
}

// GetPolicyID returns the PolicyID field value if set, zero value otherwise.
func (o *Report) GetPolicyID() string {
	if o == nil || o.PolicyID == nil {
		var ret string
		return ret
	}
	return *o.PolicyID
}

// GetPolicyIDOk returns a tuple with the PolicyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetPolicyIDOk() (*string, bool) {
	if o == nil || o.PolicyID == nil {
		return nil, false
	}
	return o.PolicyID, true
}

// HasPolicyID returns a boolean if a field has been set.
func (o *Report) HasPolicyID() bool {
	if o != nil && o.PolicyID != nil {
		return true
	}

	return false
}

// SetPolicyID gets a reference to the given string and assigns it to the PolicyID field.
func (o *Report) SetPolicyID(v string) {
	o.PolicyID = &v
}

// GetBlockedURI returns the BlockedURI field value if set, zero value otherwise.
func (o *Report) GetBlockedURI() string {
	if o == nil || o.BlockedURI == nil {
		var ret string
		return ret
	}
	return *o.BlockedURI
}

// GetBlockedURIOk returns a tuple with the BlockedURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetBlockedURIOk() (*string, bool) {
	if o == nil || o.BlockedURI == nil {
		return nil, false
	}
	return o.BlockedURI, true
}

// HasBlockedURI returns a boolean if a field has been set.
func (o *Report) HasBlockedURI() bool {
	if o != nil && o.BlockedURI != nil {
		return true
	}

	return false
}

// SetBlockedURI gets a reference to the given string and assigns it to the BlockedURI field.
func (o *Report) SetBlockedURI(v string) {
	o.BlockedURI = &v
}

// GetDocumentURI returns the DocumentURI field value if set, zero value otherwise.
func (o *Report) GetDocumentURI() string {
	if o == nil || o.DocumentURI == nil {
		var ret string
		return ret
	}
	return *o.DocumentURI
}

// GetDocumentURIOk returns a tuple with the DocumentURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetDocumentURIOk() (*string, bool) {
	if o == nil || o.DocumentURI == nil {
		return nil, false
	}
	return o.DocumentURI, true
}

// HasDocumentURI returns a boolean if a field has been set.
func (o *Report) HasDocumentURI() bool {
	if o != nil && o.DocumentURI != nil {
		return true
	}

	return false
}

// SetDocumentURI gets a reference to the given string and assigns it to the DocumentURI field.
func (o *Report) SetDocumentURI(v string) {
	o.DocumentURI = &v
}

// GetViolatedDirective returns the ViolatedDirective field value if set, zero value otherwise.
func (o *Report) GetViolatedDirective() string {
	if o == nil || o.ViolatedDirective == nil {
		var ret string
		return ret
	}
	return *o.ViolatedDirective
}

// GetViolatedDirectiveOk returns a tuple with the ViolatedDirective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetViolatedDirectiveOk() (*string, bool) {
	if o == nil || o.ViolatedDirective == nil {
		return nil, false
	}
	return o.ViolatedDirective, true
}

// HasViolatedDirective returns a boolean if a field has been set.
func (o *Report) HasViolatedDirective() bool {
	if o != nil && o.ViolatedDirective != nil {
		return true
	}

	return false
}

// SetViolatedDirective gets a reference to the given string and assigns it to the ViolatedDirective field.
func (o *Report) SetViolatedDirective(v string) {
	o.ViolatedDirective = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Report) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Report) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Report) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Report) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.PolicyID != nil {
		toSerialize["policy_id"] = o.PolicyID
	}
	if o.BlockedURI != nil {
		toSerialize["blocked_uri"] = o.BlockedURI
	}
	if o.DocumentURI != nil {
		toSerialize["document_uri"] = o.DocumentURI
	}
	if o.ViolatedDirective != nil {
		toSerialize["violated_directive"] = o.ViolatedDirective
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Report) UnmarshalJSON(bytes []byte) (err error) {
	varReport := _Report{}

	if err = json.Unmarshal(bytes, &varReport); err == nil {
		*o = Report(varReport)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "policy_id")
		delete(additionalProperties, "blocked_uri")
		delete(additionalProperties, "document_uri")
		delete(additionalProperties, "violated_directive")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableReport is a helper abstraction for handling nullable report types.
type NullableReport struct {
	value *Report
	isSet bool
}

// Get returns the value.
func (v NullableReport) Get() *Report {
	return v.value
}

// Set modifies the value.
func (v *NullableReport) Set(val *Report) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableReport) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableReport) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableReport returns a pointer to a new instance of NullableReport.
func NewNullableReport(val *Report) *NullableReport {
	return &NullableReport{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
