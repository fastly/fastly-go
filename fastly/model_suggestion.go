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

// Suggestion All attributes for a suggested domain.
type Suggestion struct {
	// The suggested domain, consisting of a subdomain and zone.
	Domain *string `json:"domain,omitempty"`
	// The subdomain of the suggested domain.
	Subdomain *string `json:"subdomain,omitempty"`
	// The zone of the suggested domain.
	Zone *string `json:"zone,omitempty"`
	// If present, the path is to be appended to the domain to complete the suggestion.
	Path                 *string `json:"path,omitempty"`
	AdditionalProperties map[string]any
}

type _Suggestion Suggestion

// NewSuggestion instantiates a new Suggestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestion() *Suggestion {
	this := Suggestion{}
	return &this
}

// NewSuggestionWithDefaults instantiates a new Suggestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestionWithDefaults() *Suggestion {
	this := Suggestion{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Suggestion) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Suggestion) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Suggestion) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Suggestion) SetDomain(v string) {
	o.Domain = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *Suggestion) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Suggestion) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *Suggestion) HasSubdomain() bool {
	if o != nil && o.Subdomain != nil {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *Suggestion) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *Suggestion) GetZone() string {
	if o == nil || o.Zone == nil {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Suggestion) GetZoneOk() (*string, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *Suggestion) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *Suggestion) SetZone(v string) {
	o.Zone = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Suggestion) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Suggestion) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Suggestion) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Suggestion) SetPath(v string) {
	o.Path = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Suggestion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Suggestion) UnmarshalJSON(bytes []byte) (err error) {
	varSuggestion := _Suggestion{}

	if err = json.Unmarshal(bytes, &varSuggestion); err == nil {
		*o = Suggestion(varSuggestion)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "zone")
		delete(additionalProperties, "path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSuggestion is a helper abstraction for handling nullable suggestion types.
type NullableSuggestion struct {
	value *Suggestion
	isSet bool
}

// Get returns the value.
func (v NullableSuggestion) Get() *Suggestion {
	return v.value
}

// Set modifies the value.
func (v *NullableSuggestion) Set(val *Suggestion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSuggestion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSuggestion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSuggestion returns a pointer to a new instance of NullableSuggestion.
func NewNullableSuggestion(val *Suggestion) *NullableSuggestion {
	return &NullableSuggestion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSuggestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSuggestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
