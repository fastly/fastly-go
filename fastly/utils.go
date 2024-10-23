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
	"time"
)

// PtrBool is a helper routine that returns a pointer to given boolean value.
func PtrBool(v bool) *bool {
	return &v
}

// PtrInt is a helper routine that returns a pointer to given integer value.
func PtrInt(v int) *int {
	return &v
}

// PtrInt32 is a helper routine that returns a pointer to given integer value.
func PtrInt32(v int32) *int32 {
	return &v
}

// PtrInt64 is a helper routine that returns a pointer to given integer value.
func PtrInt64(v int64) *int64 {
	return &v
}

// PtrFloat32 is a helper routine that returns a pointer to given float value.
func PtrFloat32(v float32) *float32 {
	return &v
}

// PtrFloat64 is a helper routine that returns a pointer to given float value.
func PtrFloat64(v float64) *float64 {
	return &v
}

// PtrString is a helper routine that returns a pointer to given string value.
func PtrString(v string) *string {
	return &v
}

// PtrTime is helper routine that returns a pointer to given Time value.
func PtrTime(v time.Time) *time.Time {
	return &v
}

// NullableBool is a helper abstraction for handling nullable boolean types.
type NullableBool struct {
	value *bool
	isSet bool
}

// Get returns the value.
func (v NullableBool) Get() *bool {
	return v.value
}

// Set modifies the value.
func (v *NullableBool) Set(val *bool) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBool) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBool) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBool returns a pointer to a new instance of NullableBool.
func NewNullableBool(val *bool) *NullableBool {
	return &NullableBool{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// NullableInt is a helper abstraction for handling nullable integer types.
type NullableInt struct {
	value *int
	isSet bool
}

// Get returns the value.
func (v NullableInt) Get() *int {
	return v.value
}

// Set modifies the value.
func (v *NullableInt) Set(val *int) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInt) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInt) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInt returns a pointer to a new instance of NullableInt.
func NewNullableInt(val *int) *NullableInt {
	return &NullableInt{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// NullableInt32 is a helper abstraction for handling nullable int32 types.
type NullableInt32 struct {
	value *int32
	isSet bool
}

// Get returns the value.
func (v NullableInt32) Get() *int32 {
	return v.value
}

// Set modifies the value.
func (v *NullableInt32) Set(val *int32) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInt32) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInt32) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInt32 returns a pointer to a new instance of NullableInt32.
func NewNullableInt32(val *int32) *NullableInt32 {
	return &NullableInt32{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInt32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// NullableInt64 is a helper abstraction for handling nullable int64 types.
type NullableInt64 struct {
	value *int64
	isSet bool
}

// Get returns the value.
func (v NullableInt64) Get() *int64 {
	return v.value
}

// Set modifies the value.
func (v *NullableInt64) Set(val *int64) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInt64) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInt64) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInt64 returns a pointer to a new instance of NullableInt64.
func NewNullableInt64(val *int64) *NullableInt64 {
	return &NullableInt64{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInt64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// NullableFloat32 is a helper abstraction for handling nullable float32 types.
type NullableFloat32 struct {
	value *float32
	isSet bool
}

// Get returns the value.
func (v NullableFloat32) Get() *float32 {
	return v.value
}

// Set modifies the value.
func (v *NullableFloat32) Set(val *float32) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableFloat32) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableFloat32) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFloat32 returns a pointer to a new instance of NullableFloat32.
func NewNullableFloat32(val *float32) *NullableFloat32 {
	return &NullableFloat32{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableFloat32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// NullableFloat64 is a helper abstraction for handling nullable float64 types.
type NullableFloat64 struct {
	value *float64
	isSet bool
}

// Get returns the value.
func (v NullableFloat64) Get() *float64 {
	return v.value
}

// Set modifies the value.
func (v *NullableFloat64) Set(val *float64) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableFloat64) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableFloat64) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFloat64 returns a pointer to a new instance of NullableFloat64.
func NewNullableFloat64(val *float64) *NullableFloat64 {
	return &NullableFloat64{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableFloat64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// NullableString is a helper abstraction for handling nullable string types.
type NullableString struct {
	value *string
	isSet bool
}

// Get returns the value.
func (v NullableString) Get() *string {
	return v.value
}

// Set modifies the value.
func (v *NullableString) Set(val *string) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableString) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableString) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableString returns a pointer to a new instance of NullableString.
func NewNullableString(val *string) *NullableString {
	return &NullableString{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// NullableTime is a helper abstraction for handling nullable time types.
type NullableTime struct {
	value *time.Time
	isSet bool
}

// Get returns the value.
func (v NullableTime) Get() *time.Time {
	return v.value
}

// Set modifies the value.
func (v *NullableTime) Set(val *time.Time) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTime) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTime) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTime returns a pointer to a new instance of NullableTime.
func NewNullableTime(val *time.Time) *NullableTime {
	return &NullableTime{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTime) MarshalJSON() ([]byte, error) {
	return v.value.MarshalJSON()
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
