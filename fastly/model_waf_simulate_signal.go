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

// WafSimulateSignal A signal detected during WAF simulation. The `type`, `detector`, `detector_scope`, and `redaction` fields are always present. The `location`, `name`, and `value` fields are present only when applicable to the signal category.
type WafSimulateSignal struct {
	// The type of signal detected (e.g., `SQLI`, `XSS`, `CMDEXE`, `TRAVERSAL`, `BACKDOOR`, `LOG4J-JNDI`, `BLOCKED`).
	Type string `json:"type"`
	// The detector engine that identified the signal (e.g., `SQLI`, `LIBINJECTIONV5`, `LIBINJECTIONJS`, or a rule ID).
	Detector string `json:"detector"`
	// The scope of the detector that identified the signal. Derived from the signal type and detection type at simulation time. `system` — built-in WAF rule (e.g., `SQLI`, `XSS`). `workspace` — workspace-level custom rule or signal (e.g., `site.*` prefix). `account` — account-level custom signal (e.g., `corp.*` prefix). `unknown` — scope could not be determined (e.g., tags fetch failed or unrecognized type).
	DetectorScope string `json:"detector_scope"`
	// The redaction level applied to the detected value. Clients should handle unexpected string values gracefully, as new redaction types may be added.
	Redaction string `json:"redaction"`
	// Where in the request the signal was detected (e.g., `QUERYSTRING`, `POSTBODY`, `HEADER`, `HEADEROUT`, `POSTARG`). Present for detection signals; absent for custom and action signals.
	Location *string `json:"location,omitempty"`
	// The parameter or header name that triggered detection. Present when the WAF engine identifies a specific parameter or header.
	Name *string `json:"name,omitempty"`
	// The matched payload value that triggered signal detection. For detection signals, contains the matched content. For `BLOCKED` signals, carries the WAF response code as a string. Absent for custom signals.
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]any
}

type _WafSimulateSignal WafSimulateSignal

// NewWafSimulateSignal instantiates a new WafSimulateSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafSimulateSignal(type_ string, detector string, detectorScope string, redaction string) *WafSimulateSignal {
	this := WafSimulateSignal{}
	this.Type = type_
	this.Detector = detector
	this.DetectorScope = detectorScope
	this.Redaction = redaction
	return &this
}

// NewWafSimulateSignalWithDefaults instantiates a new WafSimulateSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafSimulateSignalWithDefaults() *WafSimulateSignal {
	this := WafSimulateSignal{}
	return &this
}

// GetType returns the Type field value
func (o *WafSimulateSignal) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WafSimulateSignal) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WafSimulateSignal) SetType(v string) {
	o.Type = v
}

// GetDetector returns the Detector field value
func (o *WafSimulateSignal) GetDetector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detector
}

// GetDetectorOk returns a tuple with the Detector field value
// and a boolean to check if the value has been set.
func (o *WafSimulateSignal) GetDetectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detector, true
}

// SetDetector sets field value
func (o *WafSimulateSignal) SetDetector(v string) {
	o.Detector = v
}

// GetDetectorScope returns the DetectorScope field value
func (o *WafSimulateSignal) GetDetectorScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DetectorScope
}

// GetDetectorScopeOk returns a tuple with the DetectorScope field value
// and a boolean to check if the value has been set.
func (o *WafSimulateSignal) GetDetectorScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectorScope, true
}

// SetDetectorScope sets field value
func (o *WafSimulateSignal) SetDetectorScope(v string) {
	o.DetectorScope = v
}

// GetRedaction returns the Redaction field value
func (o *WafSimulateSignal) GetRedaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Redaction
}

// GetRedactionOk returns a tuple with the Redaction field value
// and a boolean to check if the value has been set.
func (o *WafSimulateSignal) GetRedactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redaction, true
}

// SetRedaction sets field value
func (o *WafSimulateSignal) SetRedaction(v string) {
	o.Redaction = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *WafSimulateSignal) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSimulateSignal) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *WafSimulateSignal) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *WafSimulateSignal) SetLocation(v string) {
	o.Location = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafSimulateSignal) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSimulateSignal) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafSimulateSignal) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafSimulateSignal) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WafSimulateSignal) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSimulateSignal) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WafSimulateSignal) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WafSimulateSignal) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafSimulateSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["detector"] = o.Detector
	}
	if true {
		toSerialize["detector_scope"] = o.DetectorScope
	}
	if true {
		toSerialize["redaction"] = o.Redaction
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafSimulateSignal) UnmarshalJSON(bytes []byte) (err error) {
	varWafSimulateSignal := _WafSimulateSignal{}

	if err = json.Unmarshal(bytes, &varWafSimulateSignal); err == nil {
		*o = WafSimulateSignal(varWafSimulateSignal)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "detector")
		delete(additionalProperties, "detector_scope")
		delete(additionalProperties, "redaction")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafSimulateSignal is a helper abstraction for handling nullable wafsimulatesignal types.
type NullableWafSimulateSignal struct {
	value *WafSimulateSignal
	isSet bool
}

// Get returns the value.
func (v NullableWafSimulateSignal) Get() *WafSimulateSignal {
	return v.value
}

// Set modifies the value.
func (v *NullableWafSimulateSignal) Set(val *WafSimulateSignal) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafSimulateSignal) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafSimulateSignal) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafSimulateSignal returns a pointer to a new instance of NullableWafSimulateSignal.
func NewNullableWafSimulateSignal(val *WafSimulateSignal) *NullableWafSimulateSignal {
	return &NullableWafSimulateSignal{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafSimulateSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafSimulateSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
