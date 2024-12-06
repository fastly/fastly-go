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

// LogInsightsValues - struct for LogInsightsValues
type LogInsightsValues struct {
	Values503Responses  *Values503Responses
	ValuesBandwidth     *ValuesBandwidth
	ValuesBrowser       *ValuesBrowser
	ValuesCacheHitRatio *ValuesCacheHitRatio
	ValuesCountryStats  *ValuesCountryStats
	ValuesDuration      *ValuesDuration
	ValuesMisses        *ValuesMisses
	ValuesRate          *ValuesRate
	ValuesRequests      *ValuesRequests
	ValuesStatusCodes   *ValuesStatusCodes
}

// Values503ResponsesAsLogInsightsValues is a convenience function that returns Values503Responses wrapped in LogInsightsValues
func Values503ResponsesAsLogInsightsValues(v *Values503Responses) LogInsightsValues {
	return LogInsightsValues{
		Values503Responses: v,
	}
}

// ValuesBandwidthAsLogInsightsValues is a convenience function that returns ValuesBandwidth wrapped in LogInsightsValues
func ValuesBandwidthAsLogInsightsValues(v *ValuesBandwidth) LogInsightsValues {
	return LogInsightsValues{
		ValuesBandwidth: v,
	}
}

// ValuesBrowserAsLogInsightsValues is a convenience function that returns ValuesBrowser wrapped in LogInsightsValues
func ValuesBrowserAsLogInsightsValues(v *ValuesBrowser) LogInsightsValues {
	return LogInsightsValues{
		ValuesBrowser: v,
	}
}

// ValuesCacheHitRatioAsLogInsightsValues is a convenience function that returns ValuesCacheHitRatio wrapped in LogInsightsValues
func ValuesCacheHitRatioAsLogInsightsValues(v *ValuesCacheHitRatio) LogInsightsValues {
	return LogInsightsValues{
		ValuesCacheHitRatio: v,
	}
}

// ValuesCountryStatsAsLogInsightsValues is a convenience function that returns ValuesCountryStats wrapped in LogInsightsValues
func ValuesCountryStatsAsLogInsightsValues(v *ValuesCountryStats) LogInsightsValues {
	return LogInsightsValues{
		ValuesCountryStats: v,
	}
}

// ValuesDurationAsLogInsightsValues is a convenience function that returns ValuesDuration wrapped in LogInsightsValues
func ValuesDurationAsLogInsightsValues(v *ValuesDuration) LogInsightsValues {
	return LogInsightsValues{
		ValuesDuration: v,
	}
}

// ValuesMissesAsLogInsightsValues is a convenience function that returns ValuesMisses wrapped in LogInsightsValues
func ValuesMissesAsLogInsightsValues(v *ValuesMisses) LogInsightsValues {
	return LogInsightsValues{
		ValuesMisses: v,
	}
}

// ValuesRateAsLogInsightsValues is a convenience function that returns ValuesRate wrapped in LogInsightsValues
func ValuesRateAsLogInsightsValues(v *ValuesRate) LogInsightsValues {
	return LogInsightsValues{
		ValuesRate: v,
	}
}

// ValuesRequestsAsLogInsightsValues is a convenience function that returns ValuesRequests wrapped in LogInsightsValues
func ValuesRequestsAsLogInsightsValues(v *ValuesRequests) LogInsightsValues {
	return LogInsightsValues{
		ValuesRequests: v,
	}
}

// ValuesStatusCodesAsLogInsightsValues is a convenience function that returns ValuesStatusCodes wrapped in LogInsightsValues
func ValuesStatusCodesAsLogInsightsValues(v *ValuesStatusCodes) LogInsightsValues {
	return LogInsightsValues{
		ValuesStatusCodes: v,
	}
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogInsightsValues) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Values503Responses
	err = newStrictDecoder(data).Decode(&o.Values503Responses)
	if err == nil {
		jsonValues503Responses, _ := json.Marshal(o.Values503Responses)
		if string(jsonValues503Responses) == "{}" { // empty struct
			o.Values503Responses = nil
		} else {
			match++
		}
	} else {
		o.Values503Responses = nil
	}

	// try to unmarshal data into ValuesBandwidth
	err = newStrictDecoder(data).Decode(&o.ValuesBandwidth)
	if err == nil {
		jsonValuesBandwidth, _ := json.Marshal(o.ValuesBandwidth)
		if string(jsonValuesBandwidth) == "{}" { // empty struct
			o.ValuesBandwidth = nil
		} else {
			match++
		}
	} else {
		o.ValuesBandwidth = nil
	}

	// try to unmarshal data into ValuesBrowser
	err = newStrictDecoder(data).Decode(&o.ValuesBrowser)
	if err == nil {
		jsonValuesBrowser, _ := json.Marshal(o.ValuesBrowser)
		if string(jsonValuesBrowser) == "{}" { // empty struct
			o.ValuesBrowser = nil
		} else {
			match++
		}
	} else {
		o.ValuesBrowser = nil
	}

	// try to unmarshal data into ValuesCacheHitRatio
	err = newStrictDecoder(data).Decode(&o.ValuesCacheHitRatio)
	if err == nil {
		jsonValuesCacheHitRatio, _ := json.Marshal(o.ValuesCacheHitRatio)
		if string(jsonValuesCacheHitRatio) == "{}" { // empty struct
			o.ValuesCacheHitRatio = nil
		} else {
			match++
		}
	} else {
		o.ValuesCacheHitRatio = nil
	}

	// try to unmarshal data into ValuesCountryStats
	err = newStrictDecoder(data).Decode(&o.ValuesCountryStats)
	if err == nil {
		jsonValuesCountryStats, _ := json.Marshal(o.ValuesCountryStats)
		if string(jsonValuesCountryStats) == "{}" { // empty struct
			o.ValuesCountryStats = nil
		} else {
			match++
		}
	} else {
		o.ValuesCountryStats = nil
	}

	// try to unmarshal data into ValuesDuration
	err = newStrictDecoder(data).Decode(&o.ValuesDuration)
	if err == nil {
		jsonValuesDuration, _ := json.Marshal(o.ValuesDuration)
		if string(jsonValuesDuration) == "{}" { // empty struct
			o.ValuesDuration = nil
		} else {
			match++
		}
	} else {
		o.ValuesDuration = nil
	}

	// try to unmarshal data into ValuesMisses
	err = newStrictDecoder(data).Decode(&o.ValuesMisses)
	if err == nil {
		jsonValuesMisses, _ := json.Marshal(o.ValuesMisses)
		if string(jsonValuesMisses) == "{}" { // empty struct
			o.ValuesMisses = nil
		} else {
			match++
		}
	} else {
		o.ValuesMisses = nil
	}

	// try to unmarshal data into ValuesRate
	err = newStrictDecoder(data).Decode(&o.ValuesRate)
	if err == nil {
		jsonValuesRate, _ := json.Marshal(o.ValuesRate)
		if string(jsonValuesRate) == "{}" { // empty struct
			o.ValuesRate = nil
		} else {
			match++
		}
	} else {
		o.ValuesRate = nil
	}

	// try to unmarshal data into ValuesRequests
	err = newStrictDecoder(data).Decode(&o.ValuesRequests)
	if err == nil {
		jsonValuesRequests, _ := json.Marshal(o.ValuesRequests)
		if string(jsonValuesRequests) == "{}" { // empty struct
			o.ValuesRequests = nil
		} else {
			match++
		}
	} else {
		o.ValuesRequests = nil
	}

	// try to unmarshal data into ValuesStatusCodes
	err = newStrictDecoder(data).Decode(&o.ValuesStatusCodes)
	if err == nil {
		jsonValuesStatusCodes, _ := json.Marshal(o.ValuesStatusCodes)
		if string(jsonValuesStatusCodes) == "{}" { // empty struct
			o.ValuesStatusCodes = nil
		} else {
			match++
		}
	} else {
		o.ValuesStatusCodes = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		o.Values503Responses = nil
		o.ValuesBandwidth = nil
		o.ValuesBrowser = nil
		o.ValuesCacheHitRatio = nil
		o.ValuesCountryStats = nil
		o.ValuesDuration = nil
		o.ValuesMisses = nil
		o.ValuesRate = nil
		o.ValuesRequests = nil
		o.ValuesStatusCodes = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LogInsightsValues)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LogInsightsValues)")
	}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogInsightsValues) MarshalJSON() ([]byte, error) {
	if o.Values503Responses != nil {
		return json.Marshal(&o.Values503Responses)
	}

	if o.ValuesBandwidth != nil {
		return json.Marshal(&o.ValuesBandwidth)
	}

	if o.ValuesBrowser != nil {
		return json.Marshal(&o.ValuesBrowser)
	}

	if o.ValuesCacheHitRatio != nil {
		return json.Marshal(&o.ValuesCacheHitRatio)
	}

	if o.ValuesCountryStats != nil {
		return json.Marshal(&o.ValuesCountryStats)
	}

	if o.ValuesDuration != nil {
		return json.Marshal(&o.ValuesDuration)
	}

	if o.ValuesMisses != nil {
		return json.Marshal(&o.ValuesMisses)
	}

	if o.ValuesRate != nil {
		return json.Marshal(&o.ValuesRate)
	}

	if o.ValuesRequests != nil {
		return json.Marshal(&o.ValuesRequests)
	}

	if o.ValuesStatusCodes != nil {
		return json.Marshal(&o.ValuesStatusCodes)
	}

	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns a specific instance of the type.
func (o *LogInsightsValues) GetActualInstance() any {
	if o == nil {
		return nil
	}
	if o.Values503Responses != nil {
		return o.Values503Responses
	}

	if o.ValuesBandwidth != nil {
		return o.ValuesBandwidth
	}

	if o.ValuesBrowser != nil {
		return o.ValuesBrowser
	}

	if o.ValuesCacheHitRatio != nil {
		return o.ValuesCacheHitRatio
	}

	if o.ValuesCountryStats != nil {
		return o.ValuesCountryStats
	}

	if o.ValuesDuration != nil {
		return o.ValuesDuration
	}

	if o.ValuesMisses != nil {
		return o.ValuesMisses
	}

	if o.ValuesRate != nil {
		return o.ValuesRate
	}

	if o.ValuesRequests != nil {
		return o.ValuesRequests
	}

	if o.ValuesStatusCodes != nil {
		return o.ValuesStatusCodes
	}

	// all schemas are nil
	return nil
}

// NullableLogInsightsValues is a helper abstraction for handling nullable loginsightsvalues types.
type NullableLogInsightsValues struct {
	value *LogInsightsValues
	isSet bool
}

// Get returns the value.
func (v NullableLogInsightsValues) Get() *LogInsightsValues {
	return v.value
}

// Set modifies the value.
func (v *NullableLogInsightsValues) Set(val *LogInsightsValues) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogInsightsValues) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogInsightsValues) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogInsightsValues returns a pointer to a new instance of NullableLogInsightsValues.
func NewNullableLogInsightsValues(val *LogInsightsValues) *NullableLogInsightsValues {
	return &NullableLogInsightsValues{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogInsightsValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogInsightsValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
