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

// IamV1RoleListResponse Paginated list of IAM roles.
type IamV1RoleListResponse struct {
	// Maximum number of results returned.
	Limit *int32 `json:"limit,omitempty"`
	// Cursor for the next page.
	NextCursor *string `json:"next_cursor,omitempty"`
	// Page of IAM roles (length ≤ limit).
	Data                 []IamV1RoleResponse `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _IamV1RoleListResponse IamV1RoleListResponse

// NewIamV1RoleListResponse instantiates a new IamV1RoleListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamV1RoleListResponse() *IamV1RoleListResponse {
	this := IamV1RoleListResponse{}
	return &this
}

// NewIamV1RoleListResponseWithDefaults instantiates a new IamV1RoleListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamV1RoleListResponseWithDefaults() *IamV1RoleListResponse {
	this := IamV1RoleListResponse{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *IamV1RoleListResponse) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV1RoleListResponse) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *IamV1RoleListResponse) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *IamV1RoleListResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *IamV1RoleListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV1RoleListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *IamV1RoleListResponse) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *IamV1RoleListResponse) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IamV1RoleListResponse) GetData() []IamV1RoleResponse {
	if o == nil || o.Data == nil {
		var ret []IamV1RoleResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV1RoleListResponse) GetDataOk() ([]IamV1RoleResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IamV1RoleListResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []IamV1RoleResponse and assigns it to the Data field.
func (o *IamV1RoleListResponse) SetData(v []IamV1RoleResponse) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IamV1RoleListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *IamV1RoleListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIamV1RoleListResponse := _IamV1RoleListResponse{}

	if err = json.Unmarshal(bytes, &varIamV1RoleListResponse); err == nil {
		*o = IamV1RoleListResponse(varIamV1RoleListResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableIamV1RoleListResponse is a helper abstraction for handling nullable iamv1rolelistresponse types.
type NullableIamV1RoleListResponse struct {
	value *IamV1RoleListResponse
	isSet bool
}

// Get returns the value.
func (v NullableIamV1RoleListResponse) Get() *IamV1RoleListResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableIamV1RoleListResponse) Set(val *IamV1RoleListResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIamV1RoleListResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIamV1RoleListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIamV1RoleListResponse returns a pointer to a new instance of NullableIamV1RoleListResponse.
func NewNullableIamV1RoleListResponse(val *IamV1RoleListResponse) *NullableIamV1RoleListResponse {
	return &NullableIamV1RoleListResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIamV1RoleListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableIamV1RoleListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
