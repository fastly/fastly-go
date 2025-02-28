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

// LogExplorerInsightsResponseLinksLinks struct for LogExplorerInsightsResponseLinksLinks
type LogExplorerInsightsResponseLinksLinks struct {
	// Location of resource
	Self *string `json:"self,omitempty"`
	// Location of the service resource
	Service              *string `json:"service,omitempty"`
	AdditionalProperties map[string]any
}

type _LogExplorerInsightsResponseLinksLinks LogExplorerInsightsResponseLinksLinks

// NewLogExplorerInsightsResponseLinksLinks instantiates a new LogExplorerInsightsResponseLinksLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogExplorerInsightsResponseLinksLinks() *LogExplorerInsightsResponseLinksLinks {
	this := LogExplorerInsightsResponseLinksLinks{}
	return &this
}

// NewLogExplorerInsightsResponseLinksLinksWithDefaults instantiates a new LogExplorerInsightsResponseLinksLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogExplorerInsightsResponseLinksLinksWithDefaults() *LogExplorerInsightsResponseLinksLinks {
	this := LogExplorerInsightsResponseLinksLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseLinksLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseLinksLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseLinksLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *LogExplorerInsightsResponseLinksLinks) SetSelf(v string) {
	o.Self = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseLinksLinks) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseLinksLinks) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseLinksLinks) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *LogExplorerInsightsResponseLinksLinks) SetService(v string) {
	o.Service = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogExplorerInsightsResponseLinksLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogExplorerInsightsResponseLinksLinks) UnmarshalJSON(bytes []byte) (err error) {
	varLogExplorerInsightsResponseLinksLinks := _LogExplorerInsightsResponseLinksLinks{}

	if err = json.Unmarshal(bytes, &varLogExplorerInsightsResponseLinksLinks); err == nil {
		*o = LogExplorerInsightsResponseLinksLinks(varLogExplorerInsightsResponseLinksLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogExplorerInsightsResponseLinksLinks is a helper abstraction for handling nullable logexplorerinsightsresponselinkslinks types.
type NullableLogExplorerInsightsResponseLinksLinks struct {
	value *LogExplorerInsightsResponseLinksLinks
	isSet bool
}

// Get returns the value.
func (v NullableLogExplorerInsightsResponseLinksLinks) Get() *LogExplorerInsightsResponseLinksLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableLogExplorerInsightsResponseLinksLinks) Set(val *LogExplorerInsightsResponseLinksLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogExplorerInsightsResponseLinksLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogExplorerInsightsResponseLinksLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogExplorerInsightsResponseLinksLinks returns a pointer to a new instance of NullableLogExplorerInsightsResponseLinksLinks.
func NewNullableLogExplorerInsightsResponseLinksLinks(val *LogExplorerInsightsResponseLinksLinks) *NullableLogExplorerInsightsResponseLinksLinks {
	return &NullableLogExplorerInsightsResponseLinksLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogExplorerInsightsResponseLinksLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogExplorerInsightsResponseLinksLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
