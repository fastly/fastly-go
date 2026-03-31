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

// OperationBulkCreateOperations struct for OperationBulkCreateOperations
type OperationBulkCreateOperations struct {
	// The HTTP method for the operation.
	Method string `json:"method"`
	// The domain for the operation.
	Domain string `json:"domain"`
	// The path for the operation.
	Path string `json:"path"`
	// A description of what the operation does.
	Description *string `json:"description,omitempty"`
	// An array of tag IDs to associate with this operation.
	TagIds []string `json:"tag_ids,omitempty"`
	// The status to assign to the operation. Defaults to SAVED if omitted.
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]any
}

type _OperationBulkCreateOperations OperationBulkCreateOperations

// NewOperationBulkCreateOperations instantiates a new OperationBulkCreateOperations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationBulkCreateOperations(method string, domain string, path string) *OperationBulkCreateOperations {
	this := OperationBulkCreateOperations{}
	this.Method = method
	this.Domain = domain
	this.Path = path
	var status string = "SAVED"
	this.Status = &status
	return &this
}

// NewOperationBulkCreateOperationsWithDefaults instantiates a new OperationBulkCreateOperations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationBulkCreateOperationsWithDefaults() *OperationBulkCreateOperations {
	this := OperationBulkCreateOperations{}
	var status string = "SAVED"
	this.Status = &status
	return &this
}

// GetMethod returns the Method field value
func (o *OperationBulkCreateOperations) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *OperationBulkCreateOperations) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *OperationBulkCreateOperations) SetMethod(v string) {
	o.Method = v
}

// GetDomain returns the Domain field value
func (o *OperationBulkCreateOperations) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *OperationBulkCreateOperations) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *OperationBulkCreateOperations) SetDomain(v string) {
	o.Domain = v
}

// GetPath returns the Path field value
func (o *OperationBulkCreateOperations) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *OperationBulkCreateOperations) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *OperationBulkCreateOperations) SetPath(v string) {
	o.Path = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OperationBulkCreateOperations) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationBulkCreateOperations) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OperationBulkCreateOperations) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OperationBulkCreateOperations) SetDescription(v string) {
	o.Description = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *OperationBulkCreateOperations) GetTagIds() []string {
	if o == nil || o.TagIds == nil {
		var ret []string
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationBulkCreateOperations) GetTagIdsOk() ([]string, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *OperationBulkCreateOperations) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []string and assigns it to the TagIds field.
func (o *OperationBulkCreateOperations) SetTagIds(v []string) {
	o.TagIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OperationBulkCreateOperations) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationBulkCreateOperations) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OperationBulkCreateOperations) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OperationBulkCreateOperations) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OperationBulkCreateOperations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TagIds != nil {
		toSerialize["tag_ids"] = o.TagIds
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OperationBulkCreateOperations) UnmarshalJSON(bytes []byte) (err error) {
	varOperationBulkCreateOperations := _OperationBulkCreateOperations{}

	if err = json.Unmarshal(bytes, &varOperationBulkCreateOperations); err == nil {
		*o = OperationBulkCreateOperations(varOperationBulkCreateOperations)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "path")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tag_ids")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOperationBulkCreateOperations is a helper abstraction for handling nullable operationbulkcreateoperations types.
type NullableOperationBulkCreateOperations struct {
	value *OperationBulkCreateOperations
	isSet bool
}

// Get returns the value.
func (v NullableOperationBulkCreateOperations) Get() *OperationBulkCreateOperations {
	return v.value
}

// Set modifies the value.
func (v *NullableOperationBulkCreateOperations) Set(val *OperationBulkCreateOperations) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOperationBulkCreateOperations) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOperationBulkCreateOperations) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOperationBulkCreateOperations returns a pointer to a new instance of NullableOperationBulkCreateOperations.
func NewNullableOperationBulkCreateOperations(val *OperationBulkCreateOperations) *NullableOperationBulkCreateOperations {
	return &NullableOperationBulkCreateOperations{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOperationBulkCreateOperations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOperationBulkCreateOperations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
