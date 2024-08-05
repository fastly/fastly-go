# EomInvoiceResponse

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

### NewEomInvoiceResponse

`func NewEomInvoiceResponse() *EomInvoiceResponse`

NewEomInvoiceResponse instantiates a new EomInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEomInvoiceResponseWithDefaults

`func NewEomInvoiceResponseWithDefaults() *EomInvoiceResponse`

NewEomInvoiceResponseWithDefaults instantiates a new EomInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerID

`func (o *EomInvoiceResponse) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *EomInvoiceResponse) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *EomInvoiceResponse) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *EomInvoiceResponse) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetInvoiceID

`func (o *EomInvoiceResponse) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *EomInvoiceResponse) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *EomInvoiceResponse) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceID

`func (o *EomInvoiceResponse) HasInvoiceID() bool`

HasInvoiceID returns a boolean if a field has been set.

### GetInvoicePostedOn

`func (o *EomInvoiceResponse) GetInvoicePostedOn() time.Time`

GetInvoicePostedOn returns the InvoicePostedOn field if non-nil, zero value otherwise.

### GetInvoicePostedOnOk

`func (o *EomInvoiceResponse) GetInvoicePostedOnOk() (*time.Time, bool)`

GetInvoicePostedOnOk returns a tuple with the InvoicePostedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePostedOn

`func (o *EomInvoiceResponse) SetInvoicePostedOn(v time.Time)`

SetInvoicePostedOn sets InvoicePostedOn field to given value.

### HasInvoicePostedOn

`func (o *EomInvoiceResponse) HasInvoicePostedOn() bool`

HasInvoicePostedOn returns a boolean if a field has been set.

### GetBillingStartDate

`func (o *EomInvoiceResponse) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *EomInvoiceResponse) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *EomInvoiceResponse) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.

### HasBillingStartDate

`func (o *EomInvoiceResponse) HasBillingStartDate() bool`

HasBillingStartDate returns a boolean if a field has been set.

### GetBillingEndDate

`func (o *EomInvoiceResponse) GetBillingEndDate() time.Time`

GetBillingEndDate returns the BillingEndDate field if non-nil, zero value otherwise.

### GetBillingEndDateOk

`func (o *EomInvoiceResponse) GetBillingEndDateOk() (*time.Time, bool)`

GetBillingEndDateOk returns a tuple with the BillingEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndDate

`func (o *EomInvoiceResponse) SetBillingEndDate(v time.Time)`

SetBillingEndDate sets BillingEndDate field to given value.

### HasBillingEndDate

`func (o *EomInvoiceResponse) HasBillingEndDate() bool`

HasBillingEndDate returns a boolean if a field has been set.

### GetStatementNumber

`func (o *EomInvoiceResponse) GetStatementNumber() string`

GetStatementNumber returns the StatementNumber field if non-nil, zero value otherwise.

### GetStatementNumberOk

`func (o *EomInvoiceResponse) GetStatementNumberOk() (*string, bool)`

GetStatementNumberOk returns a tuple with the StatementNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementNumber

`func (o *EomInvoiceResponse) SetStatementNumber(v string)`

SetStatementNumber sets StatementNumber field to given value.

### HasStatementNumber

`func (o *EomInvoiceResponse) HasStatementNumber() bool`

HasStatementNumber returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *EomInvoiceResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *EomInvoiceResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *EomInvoiceResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *EomInvoiceResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetMonthlyTransactionAmount

`func (o *EomInvoiceResponse) GetMonthlyTransactionAmount() float32`

GetMonthlyTransactionAmount returns the MonthlyTransactionAmount field if non-nil, zero value otherwise.

### GetMonthlyTransactionAmountOk

`func (o *EomInvoiceResponse) GetMonthlyTransactionAmountOk() (*float32, bool)`

GetMonthlyTransactionAmountOk returns a tuple with the MonthlyTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyTransactionAmount

`func (o *EomInvoiceResponse) SetMonthlyTransactionAmount(v float32)`

SetMonthlyTransactionAmount sets MonthlyTransactionAmount field to given value.

### HasMonthlyTransactionAmount

`func (o *EomInvoiceResponse) HasMonthlyTransactionAmount() bool`

HasMonthlyTransactionAmount returns a boolean if a field has been set.

### GetTransactionLineItems

`func (o *EomInvoiceResponse) GetTransactionLineItems() []Invoicelineitems`

GetTransactionLineItems returns the TransactionLineItems field if non-nil, zero value otherwise.

### GetTransactionLineItemsOk

`func (o *EomInvoiceResponse) GetTransactionLineItemsOk() (*[]Invoicelineitems, bool)`

GetTransactionLineItemsOk returns a tuple with the TransactionLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLineItems

`func (o *EomInvoiceResponse) SetTransactionLineItems(v []Invoicelineitems)`

SetTransactionLineItems sets TransactionLineItems field to given value.

### HasTransactionLineItems

`func (o *EomInvoiceResponse) HasTransactionLineItems() bool`

HasTransactionLineItems returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
