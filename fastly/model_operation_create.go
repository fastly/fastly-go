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

// OperationCreate struct for OperationCreate
type OperationCreate struct {
	// The HTTP method for the operation.
	Method *string `json:"method,omitempty"`
	// The domain for the operation.
	Domain *string `json:"domain,omitempty"`
	// The path for the operation, which may include path parameters.
	Path *string `json:"path,omitempty"`
	// A description of what the operation does.
	Description *string `json:"description,omitempty"`
	// An array of operation tag IDs associated with this operation.
	TagIds               []string `json:"tag_ids,omitempty"`
	AdditionalProperties map[string]any
}

type _OperationCreate OperationCreate

// NewOperationCreate instantiates a new OperationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationCreate() *OperationCreate {
	this := OperationCreate{}
	return &this
}

// NewOperationCreateWithDefaults instantiates a new OperationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationCreateWithDefaults() *OperationCreate {
	this := OperationCreate{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *OperationCreate) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationCreate) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *OperationCreate) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *OperationCreate) SetMethod(v string) {
	o.Method = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *OperationCreate) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationCreate) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *OperationCreate) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *OperationCreate) SetDomain(v string) {
	o.Domain = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *OperationCreate) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationCreate) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *OperationCreate) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *OperationCreate) SetPath(v string) {
	o.Path = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OperationCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OperationCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OperationCreate) SetDescription(v string) {
	o.Description = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *OperationCreate) GetTagIds() []string {
	if o == nil || o.TagIds == nil {
		var ret []string
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationCreate) GetTagIdsOk() ([]string, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *OperationCreate) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []string and assigns it to the TagIds field.
func (o *OperationCreate) SetTagIds(v []string) {
	o.TagIds = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OperationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TagIds != nil {
		toSerialize["tag_ids"] = o.TagIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OperationCreate) UnmarshalJSON(bytes []byte) (err error) {
	varOperationCreate := _OperationCreate{}

	if err = json.Unmarshal(bytes, &varOperationCreate); err == nil {
		*o = OperationCreate(varOperationCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "path")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tag_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOperationCreate is a helper abstraction for handling nullable operationcreate types.
type NullableOperationCreate struct {
	value *OperationCreate
	isSet bool
}

// Get returns the value.
func (v NullableOperationCreate) Get() *OperationCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableOperationCreate) Set(val *OperationCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOperationCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOperationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOperationCreate returns a pointer to a new instance of NullableOperationCreate.
func NewNullableOperationCreate(val *OperationCreate) *NullableOperationCreate {
	return &NullableOperationCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOperationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOperationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
