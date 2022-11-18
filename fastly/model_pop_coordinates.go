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

// PopCoordinates struct for PopCoordinates
type PopCoordinates struct {
	X *float32 `json:"x,omitempty"`
	Y *float32 `json:"y,omitempty"`
	Latitude *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
	AdditionalProperties map[string]any
}

type _PopCoordinates PopCoordinates

// NewPopCoordinates instantiates a new PopCoordinates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopCoordinates() *PopCoordinates {
	this := PopCoordinates{}
	return &this
}

// NewPopCoordinatesWithDefaults instantiates a new PopCoordinates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopCoordinatesWithDefaults() *PopCoordinates {
	this := PopCoordinates{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise.
func (o *PopCoordinates) GetX() float32 {
	if o == nil || o.X == nil {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopCoordinates) GetXOk() (*float32, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *PopCoordinates) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *PopCoordinates) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *PopCoordinates) GetY() float32 {
	if o == nil || o.Y == nil {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopCoordinates) GetYOk() (*float32, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *PopCoordinates) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *PopCoordinates) SetY(v float32) {
	o.Y = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *PopCoordinates) GetLatitude() float32 {
	if o == nil || o.Latitude == nil {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopCoordinates) GetLatitudeOk() (*float32, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *PopCoordinates) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *PopCoordinates) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *PopCoordinates) GetLongitude() float32 {
	if o == nil || o.Longitude == nil {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PopCoordinates) GetLongitudeOk() (*float32, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *PopCoordinates) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *PopCoordinates) SetLongitude(v float32) {
	o.Longitude = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PopCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *PopCoordinates) UnmarshalJSON(bytes []byte) (err error) {
	varPopCoordinates := _PopCoordinates{}

	if err = json.Unmarshal(bytes, &varPopCoordinates); err == nil {
		*o = PopCoordinates(varPopCoordinates)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePopCoordinates is a helper abstraction for handling nullable popcoordinates types. 
type NullablePopCoordinates struct {
	value *PopCoordinates
	isSet bool
}

// Get returns the value.
func (v NullablePopCoordinates) Get() *PopCoordinates {
	return v.value
}

// Set modifies the value.
func (v *NullablePopCoordinates) Set(val *PopCoordinates) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePopCoordinates) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePopCoordinates) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePopCoordinates returns a pointer to a new instance of NullablePopCoordinates.
func NewNullablePopCoordinates(val *PopCoordinates) *NullablePopCoordinates {
	return &NullablePopCoordinates{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePopCoordinates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePopCoordinates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
