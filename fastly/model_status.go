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

// Status All attributes for a domain status.
type Status struct {
	// The scope provided in the status request.
	Scope *string `json:"scope,omitempty"`
	// The domain provided in the status request.
	Domain *string `json:"domain,omitempty"`
	// The zone of the domain provided of the status request.
	Zone *string `json:"zone,omitempty"`
	// A space-delimited string of the varying statuses associated with the domain provided.
	Status *string `json:"status,omitempty"`
	// A space-delimited string of the varying tags associated with the domain provided.
	Tags                 *string `json:"tags,omitempty"`
	Offers               []Offer `json:"offers,omitempty"`
	AdditionalProperties map[string]any
}

type _Status Status

// NewStatus instantiates a new Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatus() *Status {
	this := Status{}
	return &this
}

// NewStatusWithDefaults instantiates a new Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusWithDefaults() *Status {
	this := Status{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *Status) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *Status) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *Status) SetScope(v string) {
	o.Scope = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Status) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Status) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Status) SetDomain(v string) {
	o.Domain = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *Status) GetZone() string {
	if o == nil || o.Zone == nil {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetZoneOk() (*string, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *Status) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *Status) SetZone(v string) {
	o.Zone = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Status) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Status) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Status) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Status) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Status) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *Status) SetTags(v string) {
	o.Tags = &v
}

// GetOffers returns the Offers field value if set, zero value otherwise.
func (o *Status) GetOffers() []Offer {
	if o == nil || o.Offers == nil {
		var ret []Offer
		return ret
	}
	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Status) GetOffersOk() ([]Offer, bool) {
	if o == nil || o.Offers == nil {
		return nil, false
	}
	return o.Offers, true
}

// HasOffers returns a boolean if a field has been set.
func (o *Status) HasOffers() bool {
	if o != nil && o.Offers != nil {
		return true
	}

	return false
}

// SetOffers gets a reference to the given []Offer and assigns it to the Offers field.
func (o *Status) SetOffers(v []Offer) {
	o.Offers = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Offers != nil {
		toSerialize["offers"] = o.Offers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Status) UnmarshalJSON(bytes []byte) (err error) {
	varStatus := _Status{}

	if err = json.Unmarshal(bytes, &varStatus); err == nil {
		*o = Status(varStatus)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "scope")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "zone")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "offers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableStatus is a helper abstraction for handling nullable status types.
type NullableStatus struct {
	value *Status
	isSet bool
}

// Get returns the value.
func (v NullableStatus) Get() *Status {
	return v.value
}

// Set modifies the value.
func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableStatus) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableStatus returns a pointer to a new instance of NullableStatus.
func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
