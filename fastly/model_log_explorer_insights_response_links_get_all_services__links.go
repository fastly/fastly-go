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

// LogExplorerInsightsResponseLinksGetAllServicesLinks struct for LogExplorerInsightsResponseLinksGetAllServicesLinks
type LogExplorerInsightsResponseLinksGetAllServicesLinks struct {
	// Location of resource
	Self                 *string `json:"self,omitempty"`
	AdditionalProperties map[string]any
}

type _LogExplorerInsightsResponseLinksGetAllServicesLinks LogExplorerInsightsResponseLinksGetAllServicesLinks

// NewLogExplorerInsightsResponseLinksGetAllServicesLinks instantiates a new LogExplorerInsightsResponseLinksGetAllServicesLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogExplorerInsightsResponseLinksGetAllServicesLinks() *LogExplorerInsightsResponseLinksGetAllServicesLinks {
	this := LogExplorerInsightsResponseLinksGetAllServicesLinks{}
	return &this
}

// NewLogExplorerInsightsResponseLinksGetAllServicesLinksWithDefaults instantiates a new LogExplorerInsightsResponseLinksGetAllServicesLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogExplorerInsightsResponseLinksGetAllServicesLinksWithDefaults() *LogExplorerInsightsResponseLinksGetAllServicesLinks {
	this := LogExplorerInsightsResponseLinksGetAllServicesLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseLinksGetAllServicesLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseLinksGetAllServicesLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseLinksGetAllServicesLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *LogExplorerInsightsResponseLinksGetAllServicesLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogExplorerInsightsResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogExplorerInsightsResponseLinksGetAllServicesLinks) UnmarshalJSON(bytes []byte) (err error) {
	varLogExplorerInsightsResponseLinksGetAllServicesLinks := _LogExplorerInsightsResponseLinksGetAllServicesLinks{}

	if err = json.Unmarshal(bytes, &varLogExplorerInsightsResponseLinksGetAllServicesLinks); err == nil {
		*o = LogExplorerInsightsResponseLinksGetAllServicesLinks(varLogExplorerInsightsResponseLinksGetAllServicesLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogExplorerInsightsResponseLinksGetAllServicesLinks is a helper abstraction for handling nullable logexplorerinsightsresponselinksgetallserviceslinks types.
type NullableLogExplorerInsightsResponseLinksGetAllServicesLinks struct {
	value *LogExplorerInsightsResponseLinksGetAllServicesLinks
	isSet bool
}

// Get returns the value.
func (v NullableLogExplorerInsightsResponseLinksGetAllServicesLinks) Get() *LogExplorerInsightsResponseLinksGetAllServicesLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableLogExplorerInsightsResponseLinksGetAllServicesLinks) Set(val *LogExplorerInsightsResponseLinksGetAllServicesLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogExplorerInsightsResponseLinksGetAllServicesLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogExplorerInsightsResponseLinksGetAllServicesLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogExplorerInsightsResponseLinksGetAllServicesLinks returns a pointer to a new instance of NullableLogExplorerInsightsResponseLinksGetAllServicesLinks.
func NewNullableLogExplorerInsightsResponseLinksGetAllServicesLinks(val *LogExplorerInsightsResponseLinksGetAllServicesLinks) *NullableLogExplorerInsightsResponseLinksGetAllServicesLinks {
	return &NullableLogExplorerInsightsResponseLinksGetAllServicesLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogExplorerInsightsResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogExplorerInsightsResponseLinksGetAllServicesLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
