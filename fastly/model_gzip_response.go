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

// GzipResponse struct for GzipResponse
type GzipResponse struct {
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition NullableString `json:"cache_condition,omitempty"`
	// Space-separated list of content types to compress. If you omit this field a default list will be used.
	ContentTypes NullableString `json:"content_types,omitempty"`
	// Space-separated list of file extensions to compress. If you omit this field a default list will be used.
	Extensions NullableString `json:"extensions,omitempty"`
	// Name of the gzip configuration.
	Name      *string `json:"name,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	Version   *string `json:"version,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt            NullableTime `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _GzipResponse GzipResponse

// NewGzipResponse instantiates a new GzipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGzipResponse() *GzipResponse {
	this := GzipResponse{}
	return &this
}

// NewGzipResponseWithDefaults instantiates a new GzipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGzipResponseWithDefaults() *GzipResponse {
	this := GzipResponse{}
	return &this
}

// GetCacheCondition returns the CacheCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GzipResponse) GetCacheCondition() string {
	if o == nil || o.CacheCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.CacheCondition.Get()
}

// GetCacheConditionOk returns a tuple with the CacheCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GzipResponse) GetCacheConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CacheCondition.Get(), o.CacheCondition.IsSet()
}

// HasCacheCondition returns a boolean if a field has been set.
func (o *GzipResponse) HasCacheCondition() bool {
	if o != nil && o.CacheCondition.IsSet() {
		return true
	}

	return false
}

// SetCacheCondition gets a reference to the given NullableString and assigns it to the CacheCondition field.
func (o *GzipResponse) SetCacheCondition(v string) {
	o.CacheCondition.Set(&v)
}

// SetCacheConditionNil sets the value for CacheCondition to be an explicit nil
func (o *GzipResponse) SetCacheConditionNil() {
	o.CacheCondition.Set(nil)
}

// UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
func (o *GzipResponse) UnsetCacheCondition() {
	o.CacheCondition.Unset()
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GzipResponse) GetContentTypes() string {
	if o == nil || o.ContentTypes.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentTypes.Get()
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GzipResponse) GetContentTypesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes.Get(), o.ContentTypes.IsSet()
}

// HasContentTypes returns a boolean if a field has been set.
func (o *GzipResponse) HasContentTypes() bool {
	if o != nil && o.ContentTypes.IsSet() {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given NullableString and assigns it to the ContentTypes field.
func (o *GzipResponse) SetContentTypes(v string) {
	o.ContentTypes.Set(&v)
}

// SetContentTypesNil sets the value for ContentTypes to be an explicit nil
func (o *GzipResponse) SetContentTypesNil() {
	o.ContentTypes.Set(nil)
}

// UnsetContentTypes ensures that no value is present for ContentTypes, not even an explicit nil
func (o *GzipResponse) UnsetContentTypes() {
	o.ContentTypes.Unset()
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GzipResponse) GetExtensions() string {
	if o == nil || o.Extensions.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extensions.Get()
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GzipResponse) GetExtensionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extensions.Get(), o.Extensions.IsSet()
}

// HasExtensions returns a boolean if a field has been set.
func (o *GzipResponse) HasExtensions() bool {
	if o != nil && o.Extensions.IsSet() {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given NullableString and assigns it to the Extensions field.
func (o *GzipResponse) SetExtensions(v string) {
	o.Extensions.Set(&v)
}

// SetExtensionsNil sets the value for Extensions to be an explicit nil
func (o *GzipResponse) SetExtensionsNil() {
	o.Extensions.Set(nil)
}

// UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
func (o *GzipResponse) UnsetExtensions() {
	o.Extensions.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GzipResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GzipResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GzipResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GzipResponse) SetName(v string) {
	o.Name = &v
}

// GetServiceID returns the ServiceID field value if set, zero value otherwise.
func (o *GzipResponse) GetServiceID() string {
	if o == nil || o.ServiceID == nil {
		var ret string
		return ret
	}
	return *o.ServiceID
}

// GetServiceIDOk returns a tuple with the ServiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GzipResponse) GetServiceIDOk() (*string, bool) {
	if o == nil || o.ServiceID == nil {
		return nil, false
	}
	return o.ServiceID, true
}

// HasServiceID returns a boolean if a field has been set.
func (o *GzipResponse) HasServiceID() bool {
	if o != nil && o.ServiceID != nil {
		return true
	}

	return false
}

// SetServiceID gets a reference to the given string and assigns it to the ServiceID field.
func (o *GzipResponse) SetServiceID(v string) {
	o.ServiceID = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GzipResponse) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GzipResponse) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GzipResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GzipResponse) SetVersion(v string) {
	o.Version = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GzipResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GzipResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GzipResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *GzipResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GzipResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GzipResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GzipResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GzipResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *GzipResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *GzipResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *GzipResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *GzipResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GzipResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GzipResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GzipResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *GzipResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *GzipResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *GzipResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o GzipResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CacheCondition.IsSet() {
		toSerialize["cache_condition"] = o.CacheCondition.Get()
	}
	if o.ContentTypes.IsSet() {
		toSerialize["content_types"] = o.ContentTypes.Get()
	}
	if o.Extensions.IsSet() {
		toSerialize["extensions"] = o.Extensions.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
func (o *GzipResponse) UnmarshalJSON(bytes []byte) (err error) {
	varGzipResponse := _GzipResponse{}

	if err = json.Unmarshal(bytes, &varGzipResponse); err == nil {
		*o = GzipResponse(varGzipResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cache_condition")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "name")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableGzipResponse is a helper abstraction for handling nullable gzipresponse types.
type NullableGzipResponse struct {
	value *GzipResponse
	isSet bool
}

// Get returns the value.
func (v NullableGzipResponse) Get() *GzipResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableGzipResponse) Set(val *GzipResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableGzipResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableGzipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGzipResponse returns a pointer to a new instance of NullableGzipResponse.
func NewNullableGzipResponse(val *GzipResponse) *NullableGzipResponse {
	return &NullableGzipResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableGzipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableGzipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
