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

// Gzip struct for Gzip
type Gzip struct {
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition NullableString `json:"cache_condition,omitempty"`
	// Space-separated list of content types to compress. If you omit this field a default list will be used.
	ContentTypes NullableString `json:"content_types,omitempty"`
	// Space-separated list of file extensions to compress. If you omit this field a default list will be used.
	Extensions NullableString `json:"extensions,omitempty"`
	// Name of the gzip configuration.
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _Gzip Gzip

// NewGzip instantiates a new Gzip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGzip() *Gzip {
	this := Gzip{}
	return &this
}

// NewGzipWithDefaults instantiates a new Gzip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGzipWithDefaults() *Gzip {
	this := Gzip{}
	return &this
}

// GetCacheCondition returns the CacheCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gzip) GetCacheCondition() string {
	if o == nil || o.CacheCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.CacheCondition.Get()
}

// GetCacheConditionOk returns a tuple with the CacheCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gzip) GetCacheConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CacheCondition.Get(), o.CacheCondition.IsSet()
}

// HasCacheCondition returns a boolean if a field has been set.
func (o *Gzip) HasCacheCondition() bool {
	if o != nil && o.CacheCondition.IsSet() {
		return true
	}

	return false
}

// SetCacheCondition gets a reference to the given NullableString and assigns it to the CacheCondition field.
func (o *Gzip) SetCacheCondition(v string) {
	o.CacheCondition.Set(&v)
}

// SetCacheConditionNil sets the value for CacheCondition to be an explicit nil
func (o *Gzip) SetCacheConditionNil() {
	o.CacheCondition.Set(nil)
}

// UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
func (o *Gzip) UnsetCacheCondition() {
	o.CacheCondition.Unset()
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gzip) GetContentTypes() string {
	if o == nil || o.ContentTypes.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentTypes.Get()
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gzip) GetContentTypesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes.Get(), o.ContentTypes.IsSet()
}

// HasContentTypes returns a boolean if a field has been set.
func (o *Gzip) HasContentTypes() bool {
	if o != nil && o.ContentTypes.IsSet() {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given NullableString and assigns it to the ContentTypes field.
func (o *Gzip) SetContentTypes(v string) {
	o.ContentTypes.Set(&v)
}

// SetContentTypesNil sets the value for ContentTypes to be an explicit nil
func (o *Gzip) SetContentTypesNil() {
	o.ContentTypes.Set(nil)
}

// UnsetContentTypes ensures that no value is present for ContentTypes, not even an explicit nil
func (o *Gzip) UnsetContentTypes() {
	o.ContentTypes.Unset()
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gzip) GetExtensions() string {
	if o == nil || o.Extensions.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extensions.Get()
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gzip) GetExtensionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extensions.Get(), o.Extensions.IsSet()
}

// HasExtensions returns a boolean if a field has been set.
func (o *Gzip) HasExtensions() bool {
	if o != nil && o.Extensions.IsSet() {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given NullableString and assigns it to the Extensions field.
func (o *Gzip) SetExtensions(v string) {
	o.Extensions.Set(&v)
}

// SetExtensionsNil sets the value for Extensions to be an explicit nil
func (o *Gzip) SetExtensionsNil() {
	o.Extensions.Set(nil)
}

// UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
func (o *Gzip) UnsetExtensions() {
	o.Extensions.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Gzip) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gzip) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Gzip) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Gzip) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Gzip) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Gzip) UnmarshalJSON(bytes []byte) (err error) {
	varGzip := _Gzip{}

	if err = json.Unmarshal(bytes, &varGzip); err == nil {
		*o = Gzip(varGzip)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cache_condition")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableGzip is a helper abstraction for handling nullable gzip types.
type NullableGzip struct {
	value *Gzip
	isSet bool
}

// Get returns the value.
func (v NullableGzip) Get() *Gzip {
	return v.value
}

// Set modifies the value.
func (v *NullableGzip) Set(val *Gzip) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableGzip) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableGzip) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGzip returns a pointer to a new instance of NullableGzip.
func NewNullableGzip(val *Gzip) *NullableGzip {
	return &NullableGzip{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableGzip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableGzip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
