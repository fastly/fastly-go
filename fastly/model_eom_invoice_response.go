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

// EomInvoiceResponse struct for EomInvoiceResponse
type EomInvoiceResponse struct {
	// Customer ID associated with the invoice.
	CustomerID *string `json:"customer_id,omitempty"`
	// Numeric string identifying the invoice.
	InvoiceID *string `json:"invoice_id,omitempty"`
	// Date and time invoice was posted on, in ISO 8601 format.
	InvoicePostedOn *time.Time `json:"invoice_posted_on,omitempty"`
	// Date and time (in ISO 8601 format) for initiation point of a billing cycle, signifying the start of charges for a service or subscription.
	BillingStartDate *time.Time `json:"billing_start_date,omitempty"`
	// Date and time (in ISO 8601 format) for termination point of a billing cycle, signifying the end of charges for a service or subscription.
	BillingEndDate *time.Time `json:"billing_end_date,omitempty"`
	// Alphanumeric string identifying the statement number.
	StatementNumber *string `json:"statement_number,omitempty"`
	// Three-letter code representing a specific currency used for financial transactions.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// Total billable amount for invoiced services charged within a single month.
	MonthlyTransactionAmount *float32           `json:"monthly_transaction_amount,omitempty"`
	TransactionLineItems     []Invoicelineitems `json:"transaction_line_items,omitempty"`
	AdditionalProperties     map[string]any
}

type _EomInvoiceResponse EomInvoiceResponse

// NewEomInvoiceResponse instantiates a new EomInvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEomInvoiceResponse() *EomInvoiceResponse {
	this := EomInvoiceResponse{}
	return &this
}

// NewEomInvoiceResponseWithDefaults instantiates a new EomInvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEomInvoiceResponseWithDefaults() *EomInvoiceResponse {
	this := EomInvoiceResponse{}
	return &this
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *EomInvoiceResponse) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EomInvoiceResponse) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *EomInvoiceResponse) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *EomInvoiceResponse) SetCustomerID(v string) {
	o.CustomerID = &v
}

// GetInvoiceID returns the InvoiceID field value if set, zero value otherwise.
func (o *EomInvoiceResponse) GetInvoiceID() string {
	if o == nil || o.InvoiceID == nil {
		var ret string
		return ret
	}
	return *o.InvoiceID
}

// GetInvoiceIDOk returns a tuple with the InvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EomInvoiceResponse) GetInvoiceIDOk() (*string, bool) {
	if o == nil || o.InvoiceID == nil {
		return nil, false
	}
	return o.InvoiceID, true
}

// HasInvoiceID returns a boolean if a field has been set.
func (o *EomInvoiceResponse) HasInvoiceID() bool {
	if o != nil && o.InvoiceID != nil {
		return true
	}

	return false
}

// SetInvoiceID gets a reference to the given string and assigns it to the InvoiceID field.
func (o *EomInvoiceResponse) SetInvoiceID(v string) {
	o.InvoiceID = &v
}

// GetInvoicePostedOn returns the InvoicePostedOn field value if set, zero value otherwise.
func (o *EomInvoiceResponse) GetInvoicePostedOn() time.Time {
	if o == nil || o.InvoicePostedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.InvoicePostedOn
}

// GetInvoicePostedOnOk returns a tuple with the InvoicePostedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EomInvoiceResponse) GetInvoicePostedOnOk() (*time.Time, bool) {
	if o == nil || o.InvoicePostedOn == nil {
		return nil, false
	}
	return o.InvoicePostedOn, true
}

// HasInvoicePostedOn returns a boolean if a field has been set.
func (o *EomInvoiceResponse) HasInvoicePostedOn() bool {
	if o != nil && o.InvoicePostedOn != nil {
		return true
	}

	return false
}

// SetInvoicePostedOn gets a reference to the given time.Time and assigns it to the InvoicePostedOn field.
func (o *EomInvoiceResponse) SetInvoicePostedOn(v time.Time) {
	o.InvoicePostedOn = &v
}

// GetBillingStartDate returns the BillingStartDate field value if set, zero value otherwise.
func (o *EomInvoiceResponse) GetBillingStartDate() time.Time {
	if o == nil || o.BillingStartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BillingStartDate
}

// GetBillingStartDateOk returns a tuple with the BillingStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EomInvoiceResponse) GetBillingStartDateOk() (*time.Time, bool) {
	if o == nil || o.BillingStartDate == nil {
		return nil, false
	}
	return o.BillingStartDate, true
}

// HasBillingStartDate returns a boolean if a field has been set.
func (o *EomInvoiceResponse) HasBillingStartDate() bool {
	if o != nil && o.BillingStartDate != nil {
		return true
	}

	return false
}

// SetBillingStartDate gets a reference to the given time.Time and assigns it to the BillingStartDate field.
func (o *EomInvoiceResponse) SetBillingStartDate(v time.Time) {
	o.BillingStartDate = &v
}

// GetBillingEndDate returns the BillingEndDate field value if set, zero value otherwise.
func (o *EomInvoiceResponse) GetBillingEndDate() time.Time {
	if o == nil || o.BillingEndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BillingEndDate
}

// GetBillingEndDateOk returns a tuple with the BillingEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EomInvoiceResponse) GetBillingEndDateOk() (*time.Time, bool) {
	if o == nil || o.BillingEndDate == nil {
		return nil, false
	}
	return o.BillingEndDate, true
}

// HasBillingEndDate returns a boolean if a field has been set.
func (o *EomInvoiceResponse) HasBillingEndDate() bool {
	if o != nil && o.BillingEndDate != nil {
		return true
	}

	return false
}

// SetBillingEndDate gets a reference to the given time.Time and assigns it to the BillingEndDate field.
func (o *EomInvoiceResponse) SetBillingEndDate(v time.Time) {
	o.BillingEndDate = &v
}

// GetStatementNumber returns the StatementNumber field value if set, zero value otherwise.
func (o *EomInvoiceResponse) GetStatementNumber() string {
	if o == nil || o.StatementNumber == nil {
		var ret string
		return ret
	}
	return *o.StatementNumber
}

// GetStatementNumberOk returns a tuple with the StatementNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EomInvoiceResponse) GetStatementNumberOk() (*string, bool) {
	if o == nil || o.StatementNumber == nil {
		return nil, false
	}
	return o.StatementNumber, true
}

// HasStatementNumber returns a boolean if a field has been set.
func (o *EomInvoiceResponse) HasStatementNumber() bool {
	if o != nil && o.StatementNumber != nil {
		return true
	}

	return false
}

// SetStatementNumber gets a reference to the given string and assigns it to the StatementNumber field.
func (o *EomInvoiceResponse) SetStatementNumber(v string) {
	o.StatementNumber = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *EomInvoiceResponse) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EomInvoiceResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *EomInvoiceResponse) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *EomInvoiceResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetMonthlyTransactionAmount returns the MonthlyTransactionAmount field value if set, zero value otherwise.
func (o *EomInvoiceResponse) GetMonthlyTransactionAmount() float32 {
	if o == nil || o.MonthlyTransactionAmount == nil {
		var ret float32
		return ret
	}
	return *o.MonthlyTransactionAmount
}

// GetMonthlyTransactionAmountOk returns a tuple with the MonthlyTransactionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EomInvoiceResponse) GetMonthlyTransactionAmountOk() (*float32, bool) {
	if o == nil || o.MonthlyTransactionAmount == nil {
		return nil, false
	}
	return o.MonthlyTransactionAmount, true
}

// HasMonthlyTransactionAmount returns a boolean if a field has been set.
func (o *EomInvoiceResponse) HasMonthlyTransactionAmount() bool {
	if o != nil && o.MonthlyTransactionAmount != nil {
		return true
	}

	return false
}

// SetMonthlyTransactionAmount gets a reference to the given float32 and assigns it to the MonthlyTransactionAmount field.
func (o *EomInvoiceResponse) SetMonthlyTransactionAmount(v float32) {
	o.MonthlyTransactionAmount = &v
}

// GetTransactionLineItems returns the TransactionLineItems field value if set, zero value otherwise.
func (o *EomInvoiceResponse) GetTransactionLineItems() []Invoicelineitems {
	if o == nil || o.TransactionLineItems == nil {
		var ret []Invoicelineitems
		return ret
	}
	return o.TransactionLineItems
}

// GetTransactionLineItemsOk returns a tuple with the TransactionLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EomInvoiceResponse) GetTransactionLineItemsOk() ([]Invoicelineitems, bool) {
	if o == nil || o.TransactionLineItems == nil {
		return nil, false
	}
	return o.TransactionLineItems, true
}

// HasTransactionLineItems returns a boolean if a field has been set.
func (o *EomInvoiceResponse) HasTransactionLineItems() bool {
	if o != nil && o.TransactionLineItems != nil {
		return true
	}

	return false
}

// SetTransactionLineItems gets a reference to the given []Invoicelineitems and assigns it to the TransactionLineItems field.
func (o *EomInvoiceResponse) SetTransactionLineItems(v []Invoicelineitems) {
	o.TransactionLineItems = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o EomInvoiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.InvoiceID != nil {
		toSerialize["invoice_id"] = o.InvoiceID
	}
	if o.InvoicePostedOn != nil {
		toSerialize["invoice_posted_on"] = o.InvoicePostedOn
	}
	if o.BillingStartDate != nil {
		toSerialize["billing_start_date"] = o.BillingStartDate
	}
	if o.BillingEndDate != nil {
		toSerialize["billing_end_date"] = o.BillingEndDate
	}
	if o.StatementNumber != nil {
		toSerialize["statement_number"] = o.StatementNumber
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
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
func (o *EomInvoiceResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEomInvoiceResponse := _EomInvoiceResponse{}

	if err = json.Unmarshal(bytes, &varEomInvoiceResponse); err == nil {
		*o = EomInvoiceResponse(varEomInvoiceResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "invoice_posted_on")
		delete(additionalProperties, "billing_start_date")
		delete(additionalProperties, "billing_end_date")
		delete(additionalProperties, "statement_number")
		delete(additionalProperties, "currency_code")
		delete(additionalProperties, "monthly_transaction_amount")
		delete(additionalProperties, "transaction_line_items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableEomInvoiceResponse is a helper abstraction for handling nullable eominvoiceresponse types.
type NullableEomInvoiceResponse struct {
	value *EomInvoiceResponse
	isSet bool
}

// Get returns the value.
func (v NullableEomInvoiceResponse) Get() *EomInvoiceResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableEomInvoiceResponse) Set(val *EomInvoiceResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableEomInvoiceResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableEomInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableEomInvoiceResponse returns a pointer to a new instance of NullableEomInvoiceResponse.
func NewNullableEomInvoiceResponse(val *EomInvoiceResponse) *NullableEomInvoiceResponse {
	return &NullableEomInvoiceResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableEomInvoiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableEomInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
