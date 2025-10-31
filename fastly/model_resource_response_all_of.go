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

// ResourceResponseAllOf struct for ResourceResponseAllOf
type ResourceResponseAllOf struct {
	// An alphanumeric string identifying the resource link.
	Id *string `json:"id,omitempty"`
	// The path to the resource.
	Href *string `json:"href,omitempty"`
	// Alphanumeric string identifying the service.
	ServiceId *string `json:"service_id,omitempty"`
	// Integer identifying a service version.
	Version              *int32        `json:"version,omitempty"`
	ResourceType         *TypeResource `json:"resource_type,omitempty"`
	AdditionalProperties map[string]any
}

type _ResourceResponseAllOf ResourceResponseAllOf

// NewResourceResponseAllOf instantiates a new ResourceResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceResponseAllOf() *ResourceResponseAllOf {
	this := ResourceResponseAllOf{}
	return &this
}

// NewResourceResponseAllOfWithDefaults instantiates a new ResourceResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceResponseAllOfWithDefaults() *ResourceResponseAllOf {
	this := ResourceResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceResponseAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceResponseAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceResponseAllOf) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ResourceResponseAllOf) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceResponseAllOf) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ResourceResponseAllOf) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ResourceResponseAllOf) SetHref(v string) {
	o.Href = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ResourceResponseAllOf) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceResponseAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ResourceResponseAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ResourceResponseAllOf) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ResourceResponseAllOf) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceResponseAllOf) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ResourceResponseAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ResourceResponseAllOf) SetVersion(v int32) {
	o.Version = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourceResponseAllOf) GetResourceType() TypeResource {
	if o == nil || o.ResourceType == nil {
		var ret TypeResource
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceResponseAllOf) GetResourceTypeOk() (*TypeResource, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourceResponseAllOf) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given TypeResource and assigns it to the ResourceType field.
func (o *ResourceResponseAllOf) SetResourceType(v TypeResource) {
	o.ResourceType = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ResourceResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ResourceResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourceResponseAllOf := _ResourceResponseAllOf{}

	if err = json.Unmarshal(bytes, &varResourceResponseAllOf); err == nil {
		*o = ResourceResponseAllOf(varResourceResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "href")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "resource_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableResourceResponseAllOf is a helper abstraction for handling nullable resourceresponseallof types.
type NullableResourceResponseAllOf struct {
	value *ResourceResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableResourceResponseAllOf) Get() *ResourceResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableResourceResponseAllOf) Set(val *ResourceResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableResourceResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableResourceResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableResourceResponseAllOf returns a pointer to a new instance of NullableResourceResponseAllOf.
func NewNullableResourceResponseAllOf(val *ResourceResponseAllOf) *NullableResourceResponseAllOf {
	return &NullableResourceResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableResourceResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableResourceResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
