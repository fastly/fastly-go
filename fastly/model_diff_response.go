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

// DiffResponse struct for DiffResponse
type DiffResponse struct {
	// The version number being diffed from.
	From *int32 `json:"from,omitempty"`
	// The version number being diffed to.
	To *int32 `json:"to,omitempty"`
	// The format the diff is being returned in (`text`, `html` or `html_simple`).
	Format *string `json:"format,omitempty"`
	// The differences between two specified service versions. Returns the full config if the version configurations are identical.
	Diff *string `json:"diff,omitempty"`
	AdditionalProperties map[string]any
}

type _DiffResponse DiffResponse

// NewDiffResponse instantiates a new DiffResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiffResponse() *DiffResponse {
	this := DiffResponse{}
	return &this
}

// NewDiffResponseWithDefaults instantiates a new DiffResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiffResponseWithDefaults() *DiffResponse {
	this := DiffResponse{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *DiffResponse) GetFrom() int32 {
	if o == nil || o.From == nil {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffResponse) GetFromOk() (*int32, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *DiffResponse) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *DiffResponse) SetFrom(v int32) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *DiffResponse) GetTo() int32 {
	if o == nil || o.To == nil {
		var ret int32
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffResponse) GetToOk() (*int32, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *DiffResponse) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given int32 and assigns it to the To field.
func (o *DiffResponse) SetTo(v int32) {
	o.To = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DiffResponse) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffResponse) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DiffResponse) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *DiffResponse) SetFormat(v string) {
	o.Format = &v
}

// GetDiff returns the Diff field value if set, zero value otherwise.
func (o *DiffResponse) GetDiff() string {
	if o == nil || o.Diff == nil {
		var ret string
		return ret
	}
	return *o.Diff
}

// GetDiffOk returns a tuple with the Diff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffResponse) GetDiffOk() (*string, bool) {
	if o == nil || o.Diff == nil {
		return nil, false
	}
	return o.Diff, true
}

// HasDiff returns a boolean if a field has been set.
func (o *DiffResponse) HasDiff() bool {
	if o != nil && o.Diff != nil {
		return true
	}

	return false
}

// SetDiff gets a reference to the given string and assigns it to the Diff field.
func (o *DiffResponse) SetDiff(v string) {
	o.Diff = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DiffResponse) MarshalJSON() ([]byte, error) {
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
func (o *DiffResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDiffResponse := _DiffResponse{}

	if err = json.Unmarshal(bytes, &varDiffResponse); err == nil {
		*o = DiffResponse(varDiffResponse)
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

// NullableDiffResponse is a helper abstraction for handling nullable diffresponse types. 
type NullableDiffResponse struct {
	value *DiffResponse
	isSet bool
}

// Get returns the value.
func (v NullableDiffResponse) Get() *DiffResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableDiffResponse) Set(val *DiffResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDiffResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDiffResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDiffResponse returns a pointer to a new instance of NullableDiffResponse.
func NewNullableDiffResponse(val *DiffResponse) *NullableDiffResponse {
	return &NullableDiffResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDiffResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableDiffResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
