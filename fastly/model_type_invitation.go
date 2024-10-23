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
	"fmt"
)

// TypeInvitation Resource type
type TypeInvitation string

// List of resourceTypeinvitation
const (
	TYPEINVITATION_INVITATION TypeInvitation = "invitation"
)

// AllowedTypeInvitationEnumValues All allowed values of TypeInvitation enum
var AllowedTypeInvitationEnumValues = []TypeInvitation{
	"invitation",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeInvitation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeInvitation(value)
	for _, existing := range AllowedTypeInvitationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeInvitation", value)
}

// NewTypeInvitationFromValue returns a pointer to a valid TypeInvitation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeInvitationFromValue(v string) (*TypeInvitation, error) {
	ev := TypeInvitation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeInvitation: valid values are %v", v, AllowedTypeInvitationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeInvitation) IsValid() bool {
	for _, existing := range AllowedTypeInvitationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypeinvitation value
func (v TypeInvitation) Ptr() *TypeInvitation {
	return &v
}

// NullableTypeInvitation is a helper abstraction for handling nullable typeinvitation types.
type NullableTypeInvitation struct {
	value *TypeInvitation
	isSet bool
}

// Get returns the value.
func (v NullableTypeInvitation) Get() *TypeInvitation {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeInvitation) Set(val *TypeInvitation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeInvitation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeInvitation returns a pointer to a new instance of NullableTypeInvitation.
func NewNullableTypeInvitation(val *TypeInvitation) *NullableTypeInvitation {
	return &NullableTypeInvitation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
