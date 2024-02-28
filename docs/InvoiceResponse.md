# InvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerID** | Pointer to **string** | Customer ID associated with the invoice. | [optional] 
**InvoiceID** | Pointer to **string** | Alphanumeric string identifying the invoice. | [optional] 
**InvoicePostedOn** | Pointer to **time.Time** | Date and time invoice was posted on, in ISO 8601 format. | [optional] 
**BillingStartDate** | Pointer to **time.Time** | Date and time (in ISO 8601 format) for initiation point of a billing cycle, signifying the start of charges for a service or subscription. | [optional] 
**BillingEndDate** | Pointer to **time.Time** | Date and time (in ISO 8601 format) for termination point of a billing cycle, signifying the end of charges for a service or subscription. | [optional] 
**StatementNumber** | Pointer to **string** | Alphanumeric string identifying the statement number. | [optional] 
**CurrencyCode** | Pointer to **string** | Three-letter code representing a specific currency used for financial transactions. | [optional] 
**MonthlyTransactionAmount** | Pointer to **float32** | Total billable amount for invoiced services charged within a single month. | [optional] 
**TransactionLineItems** | Pointer to [**[]Invoicelineitems**](Invoicelineitems.md) |  | [optional] 

## Methods

### NewInvoiceResponse

`func NewInvoiceResponse() *InvoiceResponse`

NewInvoiceResponse instantiates a new InvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceResponseWithDefaults

`func NewInvoiceResponseWithDefaults() *InvoiceResponse`

NewInvoiceResponseWithDefaults instantiates a new InvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerID

`func (o *InvoiceResponse) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *InvoiceResponse) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *InvoiceResponse) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *InvoiceResponse) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetInvoiceID

`func (o *InvoiceResponse) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *InvoiceResponse) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *InvoiceResponse) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceID

`func (o *InvoiceResponse) HasInvoiceID() bool`

HasInvoiceID returns a boolean if a field has been set.

### GetInvoicePostedOn

`func (o *InvoiceResponse) GetInvoicePostedOn() time.Time`

GetInvoicePostedOn returns the InvoicePostedOn field if non-nil, zero value otherwise.

### GetInvoicePostedOnOk

`func (o *InvoiceResponse) GetInvoicePostedOnOk() (*time.Time, bool)`

GetInvoicePostedOnOk returns a tuple with the InvoicePostedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePostedOn

`func (o *InvoiceResponse) SetInvoicePostedOn(v time.Time)`

SetInvoicePostedOn sets InvoicePostedOn field to given value.

### HasInvoicePostedOn

`func (o *InvoiceResponse) HasInvoicePostedOn() bool`

HasInvoicePostedOn returns a boolean if a field has been set.

### GetBillingStartDate

`func (o *InvoiceResponse) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *InvoiceResponse) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *InvoiceResponse) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.

### HasBillingStartDate

`func (o *InvoiceResponse) HasBillingStartDate() bool`

HasBillingStartDate returns a boolean if a field has been set.

### GetBillingEndDate

`func (o *InvoiceResponse) GetBillingEndDate() time.Time`

GetBillingEndDate returns the BillingEndDate field if non-nil, zero value otherwise.

### GetBillingEndDateOk

`func (o *InvoiceResponse) GetBillingEndDateOk() (*time.Time, bool)`

GetBillingEndDateOk returns a tuple with the BillingEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndDate

`func (o *InvoiceResponse) SetBillingEndDate(v time.Time)`

SetBillingEndDate sets BillingEndDate field to given value.

### HasBillingEndDate

`func (o *InvoiceResponse) HasBillingEndDate() bool`

HasBillingEndDate returns a boolean if a field has been set.

### GetStatementNumber

`func (o *InvoiceResponse) GetStatementNumber() string`

GetStatementNumber returns the StatementNumber field if non-nil, zero value otherwise.

### GetStatementNumberOk

`func (o *InvoiceResponse) GetStatementNumberOk() (*string, bool)`

GetStatementNumberOk returns a tuple with the StatementNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementNumber

`func (o *InvoiceResponse) SetStatementNumber(v string)`

SetStatementNumber sets StatementNumber field to given value.

### HasStatementNumber

`func (o *InvoiceResponse) HasStatementNumber() bool`

HasStatementNumber returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *InvoiceResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InvoiceResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InvoiceResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InvoiceResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetMonthlyTransactionAmount

`func (o *InvoiceResponse) GetMonthlyTransactionAmount() float32`

GetMonthlyTransactionAmount returns the MonthlyTransactionAmount field if non-nil, zero value otherwise.

### GetMonthlyTransactionAmountOk

`func (o *InvoiceResponse) GetMonthlyTransactionAmountOk() (*float32, bool)`

GetMonthlyTransactionAmountOk returns a tuple with the MonthlyTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyTransactionAmount

`func (o *InvoiceResponse) SetMonthlyTransactionAmount(v float32)`

SetMonthlyTransactionAmount sets MonthlyTransactionAmount field to given value.

### HasMonthlyTransactionAmount

`func (o *InvoiceResponse) HasMonthlyTransactionAmount() bool`

HasMonthlyTransactionAmount returns a boolean if a field has been set.

### GetTransactionLineItems

`func (o *InvoiceResponse) GetTransactionLineItems() []Invoicelineitems`

GetTransactionLineItems returns the TransactionLineItems field if non-nil, zero value otherwise.

### GetTransactionLineItemsOk

`func (o *InvoiceResponse) GetTransactionLineItemsOk() (*[]Invoicelineitems, bool)`

GetTransactionLineItemsOk returns a tuple with the TransactionLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLineItems

`func (o *InvoiceResponse) SetTransactionLineItems(v []Invoicelineitems)`

SetTransactionLineItems sets TransactionLineItems field to given value.

### HasTransactionLineItems

`func (o *InvoiceResponse) HasTransactionLineItems() bool`

HasTransactionLineItems returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
