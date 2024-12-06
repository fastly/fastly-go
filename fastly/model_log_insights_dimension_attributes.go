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

// LogInsightsDimensionAttributes - struct for LogInsightsDimensionAttributes
type LogInsightsDimensionAttributes struct {
	DimensionAttributesCountryStats *DimensionAttributesCountryStats
	DimensionAttributesRate         *DimensionAttributesRate
}

// DimensionAttributesCountryStatsAsLogInsightsDimensionAttributes is a convenience function that returns DimensionAttributesCountryStats wrapped in LogInsightsDimensionAttributes
func DimensionAttributesCountryStatsAsLogInsightsDimensionAttributes(v *DimensionAttributesCountryStats) LogInsightsDimensionAttributes {
	return LogInsightsDimensionAttributes{
		DimensionAttributesCountryStats: v,
	}
}

// DimensionAttributesRateAsLogInsightsDimensionAttributes is a convenience function that returns DimensionAttributesRate wrapped in LogInsightsDimensionAttributes
func DimensionAttributesRateAsLogInsightsDimensionAttributes(v *DimensionAttributesRate) LogInsightsDimensionAttributes {
	return LogInsightsDimensionAttributes{
		DimensionAttributesRate: v,
	}
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogInsightsDimensionAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DimensionAttributesCountryStats
	err = newStrictDecoder(data).Decode(&o.DimensionAttributesCountryStats)
	if err == nil {
		jsonDimensionAttributesCountryStats, _ := json.Marshal(o.DimensionAttributesCountryStats)
		if string(jsonDimensionAttributesCountryStats) == "{}" { // empty struct
			o.DimensionAttributesCountryStats = nil
		} else {
			match++
		}
	} else {
		o.DimensionAttributesCountryStats = nil
	}

	// try to unmarshal data into DimensionAttributesRate
	err = newStrictDecoder(data).Decode(&o.DimensionAttributesRate)
	if err == nil {
		jsonDimensionAttributesRate, _ := json.Marshal(o.DimensionAttributesRate)
		if string(jsonDimensionAttributesRate) == "{}" { // empty struct
			o.DimensionAttributesRate = nil
		} else {
			match++
		}
	} else {
		o.DimensionAttributesRate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		o.DimensionAttributesCountryStats = nil
		o.DimensionAttributesRate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LogInsightsDimensionAttributes)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LogInsightsDimensionAttributes)")
	}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogInsightsDimensionAttributes) MarshalJSON() ([]byte, error) {
	if o.DimensionAttributesCountryStats != nil {
		return json.Marshal(&o.DimensionAttributesCountryStats)
	}

	if o.DimensionAttributesRate != nil {
		return json.Marshal(&o.DimensionAttributesRate)
	}

	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns a specific instance of the type.
func (o *LogInsightsDimensionAttributes) GetActualInstance() any {
	if o == nil {
		return nil
	}
	if o.DimensionAttributesCountryStats != nil {
		return o.DimensionAttributesCountryStats
	}

	if o.DimensionAttributesRate != nil {
		return o.DimensionAttributesRate
	}

	// all schemas are nil
	return nil
}

// NullableLogInsightsDimensionAttributes is a helper abstraction for handling nullable loginsightsdimensionattributes types.
type NullableLogInsightsDimensionAttributes struct {
	value *LogInsightsDimensionAttributes
	isSet bool
}

// Get returns the value.
func (v NullableLogInsightsDimensionAttributes) Get() *LogInsightsDimensionAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableLogInsightsDimensionAttributes) Set(val *LogInsightsDimensionAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogInsightsDimensionAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogInsightsDimensionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogInsightsDimensionAttributes returns a pointer to a new instance of NullableLogInsightsDimensionAttributes.
func NewNullableLogInsightsDimensionAttributes(val *LogInsightsDimensionAttributes) *NullableLogInsightsDimensionAttributes {
	return &NullableLogInsightsDimensionAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogInsightsDimensionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogInsightsDimensionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
