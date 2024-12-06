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

// LogInsightsDimensions - struct for LogInsightsDimensions
type LogInsightsDimensions struct {
	DimensionBrowser     *DimensionBrowser
	DimensionContentType *DimensionContentType
	DimensionCountry     *DimensionCountry
	DimensionDevice      *DimensionDevice
	DimensionOs          *DimensionOs
	DimensionResponse    *DimensionResponse
	DimensionStatusCode  *DimensionStatusCode
	DimensionURL         *DimensionURL
}

// DimensionBrowserAsLogInsightsDimensions is a convenience function that returns DimensionBrowser wrapped in LogInsightsDimensions
func DimensionBrowserAsLogInsightsDimensions(v *DimensionBrowser) LogInsightsDimensions {
	return LogInsightsDimensions{
		DimensionBrowser: v,
	}
}

// DimensionContentTypeAsLogInsightsDimensions is a convenience function that returns DimensionContentType wrapped in LogInsightsDimensions
func DimensionContentTypeAsLogInsightsDimensions(v *DimensionContentType) LogInsightsDimensions {
	return LogInsightsDimensions{
		DimensionContentType: v,
	}
}

// DimensionCountryAsLogInsightsDimensions is a convenience function that returns DimensionCountry wrapped in LogInsightsDimensions
func DimensionCountryAsLogInsightsDimensions(v *DimensionCountry) LogInsightsDimensions {
	return LogInsightsDimensions{
		DimensionCountry: v,
	}
}

// DimensionDeviceAsLogInsightsDimensions is a convenience function that returns DimensionDevice wrapped in LogInsightsDimensions
func DimensionDeviceAsLogInsightsDimensions(v *DimensionDevice) LogInsightsDimensions {
	return LogInsightsDimensions{
		DimensionDevice: v,
	}
}

// DimensionOsAsLogInsightsDimensions is a convenience function that returns DimensionOs wrapped in LogInsightsDimensions
func DimensionOsAsLogInsightsDimensions(v *DimensionOs) LogInsightsDimensions {
	return LogInsightsDimensions{
		DimensionOs: v,
	}
}

// DimensionResponseAsLogInsightsDimensions is a convenience function that returns DimensionResponse wrapped in LogInsightsDimensions
func DimensionResponseAsLogInsightsDimensions(v *DimensionResponse) LogInsightsDimensions {
	return LogInsightsDimensions{
		DimensionResponse: v,
	}
}

// DimensionStatusCodeAsLogInsightsDimensions is a convenience function that returns DimensionStatusCode wrapped in LogInsightsDimensions
func DimensionStatusCodeAsLogInsightsDimensions(v *DimensionStatusCode) LogInsightsDimensions {
	return LogInsightsDimensions{
		DimensionStatusCode: v,
	}
}

// DimensionURLAsLogInsightsDimensions is a convenience function that returns DimensionURL wrapped in LogInsightsDimensions
func DimensionURLAsLogInsightsDimensions(v *DimensionURL) LogInsightsDimensions {
	return LogInsightsDimensions{
		DimensionURL: v,
	}
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogInsightsDimensions) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DimensionBrowser
	err = newStrictDecoder(data).Decode(&o.DimensionBrowser)
	if err == nil {
		jsonDimensionBrowser, _ := json.Marshal(o.DimensionBrowser)
		if string(jsonDimensionBrowser) == "{}" { // empty struct
			o.DimensionBrowser = nil
		} else {
			match++
		}
	} else {
		o.DimensionBrowser = nil
	}

	// try to unmarshal data into DimensionContentType
	err = newStrictDecoder(data).Decode(&o.DimensionContentType)
	if err == nil {
		jsonDimensionContentType, _ := json.Marshal(o.DimensionContentType)
		if string(jsonDimensionContentType) == "{}" { // empty struct
			o.DimensionContentType = nil
		} else {
			match++
		}
	} else {
		o.DimensionContentType = nil
	}

	// try to unmarshal data into DimensionCountry
	err = newStrictDecoder(data).Decode(&o.DimensionCountry)
	if err == nil {
		jsonDimensionCountry, _ := json.Marshal(o.DimensionCountry)
		if string(jsonDimensionCountry) == "{}" { // empty struct
			o.DimensionCountry = nil
		} else {
			match++
		}
	} else {
		o.DimensionCountry = nil
	}

	// try to unmarshal data into DimensionDevice
	err = newStrictDecoder(data).Decode(&o.DimensionDevice)
	if err == nil {
		jsonDimensionDevice, _ := json.Marshal(o.DimensionDevice)
		if string(jsonDimensionDevice) == "{}" { // empty struct
			o.DimensionDevice = nil
		} else {
			match++
		}
	} else {
		o.DimensionDevice = nil
	}

	// try to unmarshal data into DimensionOs
	err = newStrictDecoder(data).Decode(&o.DimensionOs)
	if err == nil {
		jsonDimensionOs, _ := json.Marshal(o.DimensionOs)
		if string(jsonDimensionOs) == "{}" { // empty struct
			o.DimensionOs = nil
		} else {
			match++
		}
	} else {
		o.DimensionOs = nil
	}

	// try to unmarshal data into DimensionResponse
	err = newStrictDecoder(data).Decode(&o.DimensionResponse)
	if err == nil {
		jsonDimensionResponse, _ := json.Marshal(o.DimensionResponse)
		if string(jsonDimensionResponse) == "{}" { // empty struct
			o.DimensionResponse = nil
		} else {
			match++
		}
	} else {
		o.DimensionResponse = nil
	}

	// try to unmarshal data into DimensionStatusCode
	err = newStrictDecoder(data).Decode(&o.DimensionStatusCode)
	if err == nil {
		jsonDimensionStatusCode, _ := json.Marshal(o.DimensionStatusCode)
		if string(jsonDimensionStatusCode) == "{}" { // empty struct
			o.DimensionStatusCode = nil
		} else {
			match++
		}
	} else {
		o.DimensionStatusCode = nil
	}

	// try to unmarshal data into DimensionURL
	err = newStrictDecoder(data).Decode(&o.DimensionURL)
	if err == nil {
		jsonDimensionURL, _ := json.Marshal(o.DimensionURL)
		if string(jsonDimensionURL) == "{}" { // empty struct
			o.DimensionURL = nil
		} else {
			match++
		}
	} else {
		o.DimensionURL = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		o.DimensionBrowser = nil
		o.DimensionContentType = nil
		o.DimensionCountry = nil
		o.DimensionDevice = nil
		o.DimensionOs = nil
		o.DimensionResponse = nil
		o.DimensionStatusCode = nil
		o.DimensionURL = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LogInsightsDimensions)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LogInsightsDimensions)")
	}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogInsightsDimensions) MarshalJSON() ([]byte, error) {
	if o.DimensionBrowser != nil {
		return json.Marshal(&o.DimensionBrowser)
	}

	if o.DimensionContentType != nil {
		return json.Marshal(&o.DimensionContentType)
	}

	if o.DimensionCountry != nil {
		return json.Marshal(&o.DimensionCountry)
	}

	if o.DimensionDevice != nil {
		return json.Marshal(&o.DimensionDevice)
	}

	if o.DimensionOs != nil {
		return json.Marshal(&o.DimensionOs)
	}

	if o.DimensionResponse != nil {
		return json.Marshal(&o.DimensionResponse)
	}

	if o.DimensionStatusCode != nil {
		return json.Marshal(&o.DimensionStatusCode)
	}

	if o.DimensionURL != nil {
		return json.Marshal(&o.DimensionURL)
	}

	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns a specific instance of the type.
func (o *LogInsightsDimensions) GetActualInstance() any {
	if o == nil {
		return nil
	}
	if o.DimensionBrowser != nil {
		return o.DimensionBrowser
	}

	if o.DimensionContentType != nil {
		return o.DimensionContentType
	}

	if o.DimensionCountry != nil {
		return o.DimensionCountry
	}

	if o.DimensionDevice != nil {
		return o.DimensionDevice
	}

	if o.DimensionOs != nil {
		return o.DimensionOs
	}

	if o.DimensionResponse != nil {
		return o.DimensionResponse
	}

	if o.DimensionStatusCode != nil {
		return o.DimensionStatusCode
	}

	if o.DimensionURL != nil {
		return o.DimensionURL
	}

	// all schemas are nil
	return nil
}

// NullableLogInsightsDimensions is a helper abstraction for handling nullable loginsightsdimensions types.
type NullableLogInsightsDimensions struct {
	value *LogInsightsDimensions
	isSet bool
}

// Get returns the value.
func (v NullableLogInsightsDimensions) Get() *LogInsightsDimensions {
	return v.value
}

// Set modifies the value.
func (v *NullableLogInsightsDimensions) Set(val *LogInsightsDimensions) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogInsightsDimensions) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogInsightsDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogInsightsDimensions returns a pointer to a new instance of NullableLogInsightsDimensions.
func NewNullableLogInsightsDimensions(val *LogInsightsDimensions) *NullableLogInsightsDimensions {
	return &NullableLogInsightsDimensions{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogInsightsDimensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogInsightsDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
