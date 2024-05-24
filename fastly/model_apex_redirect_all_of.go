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

// ApexRedirectAllOf struct for ApexRedirectAllOf
type ApexRedirectAllOf struct {
	// HTTP status code used to redirect the client.
	StatusCode *int32 `json:"status_code,omitempty"`
	// Array of apex domains that should redirect to their WWW subdomain.
	Domains []string `json:"domains,omitempty"`
	// Revision number of the apex redirect feature implementation. Defaults to the most recent revision.
	FeatureRevision *int32 `json:"feature_revision,omitempty"`
	AdditionalProperties map[string]any
}

type _ApexRedirectAllOf ApexRedirectAllOf

// NewApexRedirectAllOf instantiates a new ApexRedirectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApexRedirectAllOf() *ApexRedirectAllOf {
	this := ApexRedirectAllOf{}
	return &this
}

// NewApexRedirectAllOfWithDefaults instantiates a new ApexRedirectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApexRedirectAllOfWithDefaults() *ApexRedirectAllOf {
	this := ApexRedirectAllOf{}
	return &this
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ApexRedirectAllOf) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApexRedirectAllOf) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ApexRedirectAllOf) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ApexRedirectAllOf) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *ApexRedirectAllOf) GetDomains() []string {
	if o == nil || o.Domains == nil {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApexRedirectAllOf) GetDomainsOk() ([]string, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *ApexRedirectAllOf) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *ApexRedirectAllOf) SetDomains(v []string) {
	o.Domains = v
}

// GetFeatureRevision returns the FeatureRevision field value if set, zero value otherwise.
func (o *ApexRedirectAllOf) GetFeatureRevision() int32 {
	if o == nil || o.FeatureRevision == nil {
		var ret int32
		return ret
	}
	return *o.FeatureRevision
}

// GetFeatureRevisionOk returns a tuple with the FeatureRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApexRedirectAllOf) GetFeatureRevisionOk() (*int32, bool) {
	if o == nil || o.FeatureRevision == nil {
		return nil, false
	}
	return o.FeatureRevision, true
}

// HasFeatureRevision returns a boolean if a field has been set.
func (o *ApexRedirectAllOf) HasFeatureRevision() bool {
	if o != nil && o.FeatureRevision != nil {
		return true
	}

	return false
}

// SetFeatureRevision gets a reference to the given int32 and assigns it to the FeatureRevision field.
func (o *ApexRedirectAllOf) SetFeatureRevision(v int32) {
	o.FeatureRevision = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ApexRedirectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.FeatureRevision != nil {
		toSerialize["feature_revision"] = o.FeatureRevision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ApexRedirectAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApexRedirectAllOf := _ApexRedirectAllOf{}

	if err = json.Unmarshal(bytes, &varApexRedirectAllOf); err == nil {
		*o = ApexRedirectAllOf(varApexRedirectAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status_code")
		delete(additionalProperties, "domains")
		delete(additionalProperties, "feature_revision")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableApexRedirectAllOf is a helper abstraction for handling nullable apexredirectallof types. 
type NullableApexRedirectAllOf struct {
	value *ApexRedirectAllOf
	isSet bool
}

// Get returns the value.
func (v NullableApexRedirectAllOf) Get() *ApexRedirectAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableApexRedirectAllOf) Set(val *ApexRedirectAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableApexRedirectAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableApexRedirectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableApexRedirectAllOf returns a pointer to a new instance of NullableApexRedirectAllOf.
func NewNullableApexRedirectAllOf(val *ApexRedirectAllOf) *NullableApexRedirectAllOf {
	return &NullableApexRedirectAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableApexRedirectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableApexRedirectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
