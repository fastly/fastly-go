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

// Mtdinvoice struct for Mtdinvoice
type Mtdinvoice struct {
	// The Customer ID associated with the invoice.
	CustomerID *string `json:"customer_id,omitempty"`
	// An alphanumeric string identifying the invoice.
	InvoiceID *string `json:"invoice_id,omitempty"`
	// The date and time (in ISO 8601 format) for the initiation point of a billing cycle, signifying the start of charges for a service or subscription.
	BillingStartDate *time.Time `json:"billing_start_date,omitempty"`
	// The date and time (in ISO 8601 format) for the termination point of a billing cycle, signifying the end of charges for a service or subscription.
	BillingEndDate *time.Time `json:"billing_end_date,omitempty"`
	// The total billable amount for invoiced services charged within a single month.
	MonthlyTransactionAmount *string `json:"monthly_transaction_amount,omitempty"`
	TransactionLineItems []Mtdlineitems `json:"transaction_line_items,omitempty"`
	AdditionalProperties map[string]any
}

type _Mtdinvoice Mtdinvoice

// NewMtdinvoice instantiates a new Mtdinvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtdinvoice() *Mtdinvoice {
	this := Mtdinvoice{}
	return &this
}

// NewMtdinvoiceWithDefaults instantiates a new Mtdinvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtdinvoiceWithDefaults() *Mtdinvoice {
	this := Mtdinvoice{}
	return &this
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *Mtdinvoice) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdinvoice) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *Mtdinvoice) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *Mtdinvoice) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetInvoiceID returns the InvoiceID field value if set, zero value otherwise.
func (o *Mtdinvoice) GetInvoiceID() string {
	if o == nil || o.InvoiceID == nil {
		var ret string
		return ret
	}
	return *o.InvoiceID
}

// GetInvoiceIDOk returns a tuple with the InvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdinvoice) GetInvoiceIDOk() (*string, bool) {
	if o == nil || o.InvoiceID == nil {
		return nil, false
	}
	return o.InvoiceID, true
}

// HasInvoiceID returns a boolean if a field has been set.
func (o *Mtdinvoice) HasInvoiceID() bool {
	if o != nil && o.InvoiceID != nil {
		return true
	}

	return false
}

// SetInvoiceID gets a reference to the given string and assigns it to the InvoiceID field.
func (o *Mtdinvoice) SetInvoiceID(v string) {
	o.InvoiceID = &v
}

// GetBillingStartDate returns the BillingStartDate field value if set, zero value otherwise.
func (o *Mtdinvoice) GetBillingStartDate() time.Time {
	if o == nil || o.BillingStartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BillingStartDate
}

// GetBillingStartDateOk returns a tuple with the BillingStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdinvoice) GetBillingStartDateOk() (*time.Time, bool) {
	if o == nil || o.BillingStartDate == nil {
		return nil, false
	}
	return o.BillingStartDate, true
}

// HasBillingStartDate returns a boolean if a field has been set.
func (o *Mtdinvoice) HasBillingStartDate() bool {
	if o != nil && o.BillingStartDate != nil {
		return true
	}

	return false
}

// SetBillingStartDate gets a reference to the given time.Time and assigns it to the BillingStartDate field.
func (o *Mtdinvoice) SetBillingStartDate(v time.Time) {
	o.BillingStartDate = &v
}

// GetBillingEndDate returns the BillingEndDate field value if set, zero value otherwise.
func (o *Mtdinvoice) GetBillingEndDate() time.Time {
	if o == nil || o.BillingEndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BillingEndDate
}

// GetBillingEndDateOk returns a tuple with the BillingEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdinvoice) GetBillingEndDateOk() (*time.Time, bool) {
	if o == nil || o.BillingEndDate == nil {
		return nil, false
	}
	return o.BillingEndDate, true
}

// HasBillingEndDate returns a boolean if a field has been set.
func (o *Mtdinvoice) HasBillingEndDate() bool {
	if o != nil && o.BillingEndDate != nil {
		return true
	}

	return false
}

// SetBillingEndDate gets a reference to the given time.Time and assigns it to the BillingEndDate field.
func (o *Mtdinvoice) SetBillingEndDate(v time.Time) {
	o.BillingEndDate = &v
}

// GetMonthlyTransactionAmount returns the MonthlyTransactionAmount field value if set, zero value otherwise.
func (o *Mtdinvoice) GetMonthlyTransactionAmount() string {
	if o == nil || o.MonthlyTransactionAmount == nil {
		var ret string
		return ret
	}
	return *o.MonthlyTransactionAmount
}

// GetMonthlyTransactionAmountOk returns a tuple with the MonthlyTransactionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdinvoice) GetMonthlyTransactionAmountOk() (*string, bool) {
	if o == nil || o.MonthlyTransactionAmount == nil {
		return nil, false
	}
	return o.MonthlyTransactionAmount, true
}

// HasMonthlyTransactionAmount returns a boolean if a field has been set.
func (o *Mtdinvoice) HasMonthlyTransactionAmount() bool {
	if o != nil && o.MonthlyTransactionAmount != nil {
		return true
	}

	return false
}

// SetMonthlyTransactionAmount gets a reference to the given string and assigns it to the MonthlyTransactionAmount field.
func (o *Mtdinvoice) SetMonthlyTransactionAmount(v string) {
	o.MonthlyTransactionAmount = &v
}

// GetTransactionLineItems returns the TransactionLineItems field value if set, zero value otherwise.
func (o *Mtdinvoice) GetTransactionLineItems() []Mtdlineitems {
	if o == nil || o.TransactionLineItems == nil {
		var ret []Mtdlineitems
		return ret
	}
	return o.TransactionLineItems
}

// GetTransactionLineItemsOk returns a tuple with the TransactionLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mtdinvoice) GetTransactionLineItemsOk() ([]Mtdlineitems, bool) {
	if o == nil || o.TransactionLineItems == nil {
		return nil, false
	}
	return o.TransactionLineItems, true
}

// HasTransactionLineItems returns a boolean if a field has been set.
func (o *Mtdinvoice) HasTransactionLineItems() bool {
	if o != nil && o.TransactionLineItems != nil {
		return true
	}

	return false
}

// SetTransactionLineItems gets a reference to the given []Mtdlineitems and assigns it to the TransactionLineItems field.
func (o *Mtdinvoice) SetTransactionLineItems(v []Mtdlineitems) {
	o.TransactionLineItems = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Mtdinvoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.InvoiceID != nil {
		toSerialize["invoice_id"] = o.InvoiceID
	}
	if o.BillingStartDate != nil {
		toSerialize["billing_start_date"] = o.BillingStartDate
	}
	if o.BillingEndDate != nil {
		toSerialize["billing_end_date"] = o.BillingEndDate
	}
	if o.MonthlyTransactionAmount != nil {
		toSerialize["monthly_transaction_amount"] = o.MonthlyTransactionAmount
	}
	if o.TransactionLineItems != nil {
		toSerialize["transaction_line_items"] = o.TransactionLineItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Mtdinvoice) UnmarshalJSON(bytes []byte) (err error) {
	varMtdinvoice := _Mtdinvoice{}

	if err = json.Unmarshal(bytes, &varMtdinvoice); err == nil {
		*o = Mtdinvoice(varMtdinvoice)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "billing_start_date")
		delete(additionalProperties, "billing_end_date")
		delete(additionalProperties, "monthly_transaction_amount")
		delete(additionalProperties, "transaction_line_items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableMtdinvoice is a helper abstraction for handling nullable mtdinvoice types. 
type NullableMtdinvoice struct {
	value *Mtdinvoice
	isSet bool
}

// Get returns the value.
func (v NullableMtdinvoice) Get() *Mtdinvoice {
	return v.value
}

// Set modifies the value.
func (v *NullableMtdinvoice) Set(val *Mtdinvoice) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableMtdinvoice) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableMtdinvoice) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMtdinvoice returns a pointer to a new instance of NullableMtdinvoice.
func NewNullableMtdinvoice(val *Mtdinvoice) *NullableMtdinvoice {
	return &NullableMtdinvoice{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableMtdinvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableMtdinvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
