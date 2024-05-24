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

// BillingEstimateResponse struct for BillingEstimateResponse
type BillingEstimateResponse struct {
	// Date and time in ISO 8601 format.
	EndTime NullableTime `json:"end_time,omitempty"`
	// Date and time in ISO 8601 format.
	StartTime NullableTime `json:"start_time,omitempty"`
	InvoiceID *string `json:"invoice_id,omitempty"`
	CustomerID *string `json:"customer_id,omitempty"`
	// The current state of our third-party billing vendor. One of `up` or `down`.
	VendorState *string `json:"vendor_state,omitempty"`
	Status *BillingStatus `json:"status,omitempty"`
	Total *BillingTotal `json:"total,omitempty"`
	// Breakdown of regional data for products that are region based.
	Regions *map[string]map[string]map[string]any `json:"regions,omitempty"`
	LineItems []BillingEstimateLinesLineItems `json:"line_items,omitempty"`
	AdditionalProperties map[string]any
}

type _BillingEstimateResponse BillingEstimateResponse

// NewBillingEstimateResponse instantiates a new BillingEstimateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingEstimateResponse() *BillingEstimateResponse {
	this := BillingEstimateResponse{}
	return &this
}

// NewBillingEstimateResponseWithDefaults instantiates a new BillingEstimateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingEstimateResponseWithDefaults() *BillingEstimateResponse {
	this := BillingEstimateResponse{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingEstimateResponse) GetEndTime() time.Time {
	if o == nil || o.EndTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingEstimateResponse) GetEndTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *BillingEstimateResponse) HasEndTime() bool {
	if o != nil && o.EndTime.IsSet() {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given NullableTime and assigns it to the EndTime field.
func (o *BillingEstimateResponse) SetEndTime(v time.Time) {
	o.EndTime.Set(&v)
}
// SetEndTimeNil sets the value for EndTime to be an explicit nil
func (o *BillingEstimateResponse) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
func (o *BillingEstimateResponse) UnsetEndTime() {
	o.EndTime.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingEstimateResponse) GetStartTime() time.Time {
	if o == nil || o.StartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingEstimateResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *BillingEstimateResponse) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *BillingEstimateResponse) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *BillingEstimateResponse) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *BillingEstimateResponse) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetInvoiceID returns the InvoiceID field value if set, zero value otherwise.
func (o *BillingEstimateResponse) GetInvoiceID() string {
	if o == nil || o.InvoiceID == nil {
		var ret string
		return ret
	}
	return *o.InvoiceID
}

// GetInvoiceIDOk returns a tuple with the InvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponse) GetInvoiceIDOk() (*string, bool) {
	if o == nil || o.InvoiceID == nil {
		return nil, false
	}
	return o.InvoiceID, true
}

// HasInvoiceID returns a boolean if a field has been set.
func (o *BillingEstimateResponse) HasInvoiceID() bool {
	if o != nil && o.InvoiceID != nil {
		return true
	}

	return false
}

// SetInvoiceID gets a reference to the given string and assigns it to the InvoiceID field.
func (o *BillingEstimateResponse) SetInvoiceID(v string) {
	o.InvoiceID = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *BillingEstimateResponse) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponse) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *BillingEstimateResponse) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *BillingEstimateResponse) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetVendorState returns the VendorState field value if set, zero value otherwise.
func (o *BillingEstimateResponse) GetVendorState() string {
	if o == nil || o.VendorState == nil {
		var ret string
		return ret
	}
	return *o.VendorState
}

// GetVendorStateOk returns a tuple with the VendorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponse) GetVendorStateOk() (*string, bool) {
	if o == nil || o.VendorState == nil {
		return nil, false
	}
	return o.VendorState, true
}

// HasVendorState returns a boolean if a field has been set.
func (o *BillingEstimateResponse) HasVendorState() bool {
	if o != nil && o.VendorState != nil {
		return true
	}

	return false
}

// SetVendorState gets a reference to the given string and assigns it to the VendorState field.
func (o *BillingEstimateResponse) SetVendorState(v string) {
	o.VendorState = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BillingEstimateResponse) GetStatus() BillingStatus {
	if o == nil || o.Status == nil {
		var ret BillingStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponse) GetStatusOk() (*BillingStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BillingEstimateResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BillingStatus and assigns it to the Status field.
func (o *BillingEstimateResponse) SetStatus(v BillingStatus) {
	o.Status = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BillingEstimateResponse) GetTotal() BillingTotal {
	if o == nil || o.Total == nil {
		var ret BillingTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponse) GetTotalOk() (*BillingTotal, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BillingEstimateResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given BillingTotal and assigns it to the Total field.
func (o *BillingEstimateResponse) SetTotal(v BillingTotal) {
	o.Total = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *BillingEstimateResponse) GetRegions() map[string]map[string]map[string]any {
	if o == nil || o.Regions == nil {
		var ret map[string]map[string]map[string]any
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponse) GetRegionsOk() (*map[string]map[string]map[string]any, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *BillingEstimateResponse) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given map[string]map[string]map[string]any and assigns it to the Regions field.
func (o *BillingEstimateResponse) SetRegions(v map[string]map[string]map[string]any) {
	o.Regions = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *BillingEstimateResponse) GetLineItems() []BillingEstimateLinesLineItems {
	if o == nil || o.LineItems == nil {
		var ret []BillingEstimateLinesLineItems
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingEstimateResponse) GetLineItemsOk() ([]BillingEstimateLinesLineItems, bool) {
	if o == nil || o.LineItems == nil {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *BillingEstimateResponse) HasLineItems() bool {
	if o != nil && o.LineItems != nil {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []BillingEstimateLinesLineItems and assigns it to the LineItems field.
func (o *BillingEstimateResponse) SetLineItems(v []BillingEstimateLinesLineItems) {
	o.LineItems = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BillingEstimateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.EndTime.IsSet() {
		toSerialize["end_time"] = o.EndTime.Get()
	}
	if o.StartTime.IsSet() {
		toSerialize["start_time"] = o.StartTime.Get()
	}
	if o.InvoiceID != nil {
		toSerialize["invoice_id"] = o.InvoiceID
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.VendorState != nil {
		toSerialize["vendor_state"] = o.VendorState
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	if o.LineItems != nil {
		toSerialize["line_items"] = o.LineItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *BillingEstimateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBillingEstimateResponse := _BillingEstimateResponse{}

	if err = json.Unmarshal(bytes, &varBillingEstimateResponse); err == nil {
		*o = BillingEstimateResponse(varBillingEstimateResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "end_time")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "vendor_state")
		delete(additionalProperties, "status")
		delete(additionalProperties, "total")
		delete(additionalProperties, "regions")
		delete(additionalProperties, "line_items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBillingEstimateResponse is a helper abstraction for handling nullable billingestimateresponse types. 
type NullableBillingEstimateResponse struct {
	value *BillingEstimateResponse
	isSet bool
}

// Get returns the value.
func (v NullableBillingEstimateResponse) Get() *BillingEstimateResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBillingEstimateResponse) Set(val *BillingEstimateResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBillingEstimateResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBillingEstimateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBillingEstimateResponse returns a pointer to a new instance of NullableBillingEstimateResponse.
func NewNullableBillingEstimateResponse(val *BillingEstimateResponse) *NullableBillingEstimateResponse {
	return &NullableBillingEstimateResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBillingEstimateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableBillingEstimateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
