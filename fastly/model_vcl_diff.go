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

// VclDiff struct for VclDiff
type VclDiff struct {
	// The version number of the service to which changes in the generated VCL are being compared.
	From *int32 `json:"from,omitempty"`
	// The version number of the service from which changes in the generated VCL are being compared.
	To *int32 `json:"to,omitempty"`
	// The format in which compared VCL changes are being returned in.
	Format *string `json:"format,omitempty"`
	// The differences between two specified versions.
	Diff                 *string `json:"diff,omitempty"`
	AdditionalProperties map[string]any
}

type _VclDiff VclDiff

// NewVclDiff instantiates a new VclDiff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVclDiff() *VclDiff {
	this := VclDiff{}
	return &this
}

// NewVclDiffWithDefaults instantiates a new VclDiff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVclDiffWithDefaults() *VclDiff {
	this := VclDiff{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *VclDiff) GetFrom() int32 {
	if o == nil || o.From == nil {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VclDiff) GetFromOk() (*int32, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *VclDiff) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *VclDiff) SetFrom(v int32) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *VclDiff) GetTo() int32 {
	if o == nil || o.To == nil {
		var ret int32
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VclDiff) GetToOk() (*int32, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *VclDiff) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given int32 and assigns it to the To field.
func (o *VclDiff) SetTo(v int32) {
	o.To = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *VclDiff) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VclDiff) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *VclDiff) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *VclDiff) SetFormat(v string) {
	o.Format = &v
}

// GetDiff returns the Diff field value if set, zero value otherwise.
func (o *VclDiff) GetDiff() string {
	if o == nil || o.Diff == nil {
		var ret string
		return ret
	}
	return *o.Diff
}

// GetDiffOk returns a tuple with the Diff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VclDiff) GetDiffOk() (*string, bool) {
	if o == nil || o.Diff == nil {
		return nil, false
	}
	return o.Diff, true
}

// HasDiff returns a boolean if a field has been set.
func (o *VclDiff) HasDiff() bool {
	if o != nil && o.Diff != nil {
		return true
	}

	return false
}

// SetDiff gets a reference to the given string and assigns it to the Diff field.
func (o *VclDiff) SetDiff(v string) {
	o.Diff = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o VclDiff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Diff != nil {
		toSerialize["diff"] = o.Diff
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *VclDiff) UnmarshalJSON(bytes []byte) (err error) {
	varVclDiff := _VclDiff{}

	if err = json.Unmarshal(bytes, &varVclDiff); err == nil {
		*o = VclDiff(varVclDiff)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		delete(additionalProperties, "format")
		delete(additionalProperties, "diff")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableVclDiff is a helper abstraction for handling nullable vcldiff types.
type NullableVclDiff struct {
	value *VclDiff
	isSet bool
}

// Get returns the value.
func (v NullableVclDiff) Get() *VclDiff {
	return v.value
}

// Set modifies the value.
func (v *NullableVclDiff) Set(val *VclDiff) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableVclDiff) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableVclDiff) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableVclDiff returns a pointer to a new instance of NullableVclDiff.
func NewNullableVclDiff(val *VclDiff) *NullableVclDiff {
	return &NullableVclDiff{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableVclDiff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableVclDiff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
