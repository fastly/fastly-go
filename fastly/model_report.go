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
	Id                   *string    `json:"id,omitempty"`
	PolicyId             *string    `json:"policy_id,omitempty"`
	BlockedUri           *string    `json:"blocked_uri,omitempty"`
	DocumentUri          *string    `json:"document_uri,omitempty"`
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *Report) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Report) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Report) SetId(v string) {
	o.Id = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *Report) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *Report) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *Report) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetBlockedUri returns the BlockedUri field value if set, zero value otherwise.
func (o *Report) GetBlockedUri() string {
	if o == nil || o.BlockedUri == nil {
		var ret string
		return ret
	}
	return *o.BlockedUri
}

// GetBlockedUriOk returns a tuple with the BlockedUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetBlockedUriOk() (*string, bool) {
	if o == nil || o.BlockedUri == nil {
		return nil, false
	}
	return o.BlockedUri, true
}

// HasBlockedUri returns a boolean if a field has been set.
func (o *Report) HasBlockedUri() bool {
	if o != nil && o.BlockedUri != nil {
		return true
	}

	return false
}

// SetBlockedUri gets a reference to the given string and assigns it to the BlockedUri field.
func (o *Report) SetBlockedUri(v string) {
	o.BlockedUri = &v
}

// GetDocumentUri returns the DocumentUri field value if set, zero value otherwise.
func (o *Report) GetDocumentUri() string {
	if o == nil || o.DocumentUri == nil {
		var ret string
		return ret
	}
	return *o.DocumentUri
}

// GetDocumentUriOk returns a tuple with the DocumentUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetDocumentUriOk() (*string, bool) {
	if o == nil || o.DocumentUri == nil {
		return nil, false
	}
	return o.DocumentUri, true
}

// HasDocumentUri returns a boolean if a field has been set.
func (o *Report) HasDocumentUri() bool {
	if o != nil && o.DocumentUri != nil {
		return true
	}

	return false
}

// SetDocumentUri gets a reference to the given string and assigns it to the DocumentUri field.
func (o *Report) SetDocumentUri(v string) {
	o.DocumentUri = &v
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PolicyId != nil {
		toSerialize["policy_id"] = o.PolicyId
	}
	if o.BlockedUri != nil {
		toSerialize["blocked_uri"] = o.BlockedUri
	}
	if o.DocumentUri != nil {
		toSerialize["document_uri"] = o.DocumentUri
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
