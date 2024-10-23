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

// AwsRegion A named set of [AWS resources](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints) that's in the same geographical area.
type AwsRegion string

// List of aws_region
const (
	AWSREGION_US_EAST_1      AwsRegion = "us-east-1"
	AWSREGION_US_EAST_2      AwsRegion = "us-east-2"
	AWSREGION_US_WEST_1      AwsRegion = "us-west-1"
	AWSREGION_US_WEST_2      AwsRegion = "us-west-2"
	AWSREGION_AF_SOUTH_1     AwsRegion = "af-south-1"
	AWSREGION_AP_EAST_1      AwsRegion = "ap-east-1"
	AWSREGION_AP_SOUTH_1     AwsRegion = "ap-south-1"
	AWSREGION_AP_NORTHEAST_3 AwsRegion = "ap-northeast-3"
	AWSREGION_AP_NORTHEAST_2 AwsRegion = "ap-northeast-2"
	AWSREGION_AP_SOUTHEAST_1 AwsRegion = "ap-southeast-1"
	AWSREGION_AP_SOUTHEAST_2 AwsRegion = "ap-southeast-2"
	AWSREGION_AP_NORTHEAST_1 AwsRegion = "ap-northeast-1"
	AWSREGION_CA_CENTRAL_1   AwsRegion = "ca-central-1"
	AWSREGION_CN_NORTH_1     AwsRegion = "cn-north-1"
	AWSREGION_CN_NORTHWEST_1 AwsRegion = "cn-northwest-1"
	AWSREGION_EU_CENTRAL_1   AwsRegion = "eu-central-1"
	AWSREGION_EU_WEST_1      AwsRegion = "eu-west-1"
	AWSREGION_EU_WEST_2      AwsRegion = "eu-west-2"
	AWSREGION_EU_SOUTH_1     AwsRegion = "eu-south-1"
	AWSREGION_EU_WEST_3      AwsRegion = "eu-west-3"
	AWSREGION_EU_NORTH_1     AwsRegion = "eu-north-1"
	AWSREGION_ME_SOUTH_1     AwsRegion = "me-south-1"
	AWSREGION_SA_EAST_1      AwsRegion = "sa-east-1"
)

// AllowedAwsRegionEnumValues All allowed values of AwsRegion enum
var AllowedAwsRegionEnumValues = []AwsRegion{
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
	"af-south-1",
	"ap-east-1",
	"ap-south-1",
	"ap-northeast-3",
	"ap-northeast-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-northeast-1",
	"ca-central-1",
	"cn-north-1",
	"cn-northwest-1",
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"eu-south-1",
	"eu-west-3",
	"eu-north-1",
	"me-south-1",
	"sa-east-1",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *AwsRegion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsRegion(value)
	for _, existing := range AllowedAwsRegionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsRegion", value)
}

// NewAwsRegionFromValue returns a pointer to a valid AwsRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsRegionFromValue(v string) (*AwsRegion, error) {
	ev := AwsRegion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AwsRegion: valid values are %v", v, AllowedAwsRegionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsRegion) IsValid() bool {
	for _, existing := range AllowedAwsRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to aws_region value
func (v AwsRegion) Ptr() *AwsRegion {
	return &v
}

// NullableAwsRegion is a helper abstraction for handling nullable awsregion types.
type NullableAwsRegion struct {
	value *AwsRegion
	isSet bool
}

// Get returns the value.
func (v NullableAwsRegion) Get() *AwsRegion {
	return v.value
}

// Set modifies the value.
func (v *NullableAwsRegion) Set(val *AwsRegion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAwsRegion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAwsRegion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAwsRegion returns a pointer to a new instance of NullableAwsRegion.
func NewNullableAwsRegion(val *AwsRegion) *NullableAwsRegion {
	return &NullableAwsRegion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAwsRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAwsRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
