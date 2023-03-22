// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// Pop struct for Pop
type Pop struct {
	// the three-letter code for the [POP](https://developer.fastly.com/learning/concepts/pop/)
	Code string `json:"code"`
	// the name of the POP
	Name string `json:"name"`
	Group string `json:"group"`
	Region string `json:"region"`
	// the region used for stats reporting
	StatsRegion string `json:"stats_region"`
	// the region used for billing
	BillingRegion string `json:"billing_region"`
	Coordinates *PopCoordinates `json:"coordinates,omitempty"`
	// the name of the [shield code](https://developer.fastly.com/learning/concepts/shielding/#choosing-a-shield-location) if this POP is suitable for shielding
	Shield *string `json:"shield,omitempty"`
	AdditionalProperties map[string]any
}

type _Pop Pop

// NewPop instantiates a new Pop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPop(code string, name string, group string, region string, statsRegion string, billingRegion string) *Pop {
	this := Pop{}
	this.Code = code
	this.Name = name
	this.Group = group
	this.Region = region
	this.StatsRegion = statsRegion
	this.BillingRegion = billingRegion
	return &this
}

// NewPopWithDefaults instantiates a new Pop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopWithDefaults() *Pop {
	this := Pop{}
	return &this
}

// GetCode returns the Code field value
func (o *Pop) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Pop) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Pop) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *Pop) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Pop) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Pop) SetName(v string) {
	o.Name = v
}

// GetGroup returns the Group field value
func (o *Pop) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *Pop) GetGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *Pop) SetGroup(v string) {
	o.Group = v
}

// GetRegion returns the Region field value
func (o *Pop) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Pop) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *Pop) SetRegion(v string) {
	o.Region = v
}

// GetStatsRegion returns the StatsRegion field value
func (o *Pop) GetStatsRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatsRegion
}

// GetStatsRegionOk returns a tuple with the StatsRegion field value
// and a boolean to check if the value has been set.
func (o *Pop) GetStatsRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StatsRegion, true
}

// SetStatsRegion sets field value
func (o *Pop) SetStatsRegion(v string) {
	o.StatsRegion = v
}

// GetBillingRegion returns the BillingRegion field value
func (o *Pop) GetBillingRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingRegion
}

// GetBillingRegionOk returns a tuple with the BillingRegion field value
// and a boolean to check if the value has been set.
func (o *Pop) GetBillingRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BillingRegion, true
}

// SetBillingRegion sets field value
func (o *Pop) SetBillingRegion(v string) {
	o.BillingRegion = v
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *Pop) GetCoordinates() PopCoordinates {
	if o == nil || o.Coordinates == nil {
		var ret PopCoordinates
		return ret
	}
	return *o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pop) GetCoordinatesOk() (*PopCoordinates, bool) {
	if o == nil || o.Coordinates == nil {
		return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *Pop) HasCoordinates() bool {
	if o != nil && o.Coordinates != nil {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given PopCoordinates and assigns it to the Coordinates field.
func (o *Pop) SetCoordinates(v PopCoordinates) {
	o.Coordinates = &v
}

// GetShield returns the Shield field value if set, zero value otherwise.
func (o *Pop) GetShield() string {
	if o == nil || o.Shield == nil {
		var ret string
		return ret
	}
	return *o.Shield
}

// GetShieldOk returns a tuple with the Shield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pop) GetShieldOk() (*string, bool) {
	if o == nil || o.Shield == nil {
		return nil, false
	}
	return o.Shield, true
}

// HasShield returns a boolean if a field has been set.
func (o *Pop) HasShield() bool {
	if o != nil && o.Shield != nil {
		return true
	}

	return false
}

// SetShield gets a reference to the given string and assigns it to the Shield field.
func (o *Pop) SetShield(v string) {
	o.Shield = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Pop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["stats_region"] = o.StatsRegion
	}
	if true {
		toSerialize["billing_region"] = o.BillingRegion
	}
	if o.Coordinates != nil {
		toSerialize["coordinates"] = o.Coordinates
	}
	if o.Shield != nil {
		toSerialize["shield"] = o.Shield
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Pop) UnmarshalJSON(bytes []byte) (err error) {
	varPop := _Pop{}

	if err = json.Unmarshal(bytes, &varPop); err == nil {
		*o = Pop(varPop)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "name")
		delete(additionalProperties, "group")
		delete(additionalProperties, "region")
		delete(additionalProperties, "stats_region")
		delete(additionalProperties, "billing_region")
		delete(additionalProperties, "coordinates")
		delete(additionalProperties, "shield")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePop is a helper abstraction for handling nullable pop types. 
type NullablePop struct {
	value *Pop
	isSet bool
}

// Get returns the value.
func (v NullablePop) Get() *Pop {
	return v.value
}

// Set modifies the value.
func (v *NullablePop) Set(val *Pop) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePop) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePop) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePop returns a pointer to a new instance of NullablePop.
func NewNullablePop(val *Pop) *NullablePop {
	return &NullablePop{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
