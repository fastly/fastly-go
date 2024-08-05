# MtdInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerID** | Pointer to **string** | The Customer ID associated with the invoice. | [optional] 
**InvoiceID** | Pointer to **string** | An alphanumeric string identifying the invoice. | [optional] 
**BillingStartDate** | Pointer to **time.Time** | The date and time (in ISO 8601 format) for the initiation point of a billing cycle, signifying the start of charges for a service or subscription. | [optional] 
**BillingEndDate** | Pointer to **time.Time** | The date and time (in ISO 8601 format) for the termination point of a billing cycle, signifying the end of charges for a service or subscription. | [optional] 
**MonthlyTransactionAmount** | Pointer to **string** | The total billable amount for invoiced services charged within a single month. | [optional] 
**TransactionLineItems** | Pointer to [**[]Mtdlineitems**](Mtdlineitems.md) |  | [optional] 

## Methods

### NewMtdInvoiceResponse

`func NewMtdInvoiceResponse() *MtdInvoiceResponse`

NewMtdInvoiceResponse instantiates a new MtdInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMtdInvoiceResponseWithDefaults

`func NewMtdInvoiceResponseWithDefaults() *MtdInvoiceResponse`

NewMtdInvoiceResponseWithDefaults instantiates a new MtdInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerID

`func (o *MtdInvoiceResponse) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *MtdInvoiceResponse) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *MtdInvoiceResponse) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *MtdInvoiceResponse) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetInvoiceID

`func (o *MtdInvoiceResponse) GetInvoiceID() string`

GetInvoiceID returns the InvoiceID field if non-nil, zero value otherwise.

### GetInvoiceIDOk

`func (o *MtdInvoiceResponse) GetInvoiceIDOk() (*string, bool)`

GetInvoiceIDOk returns a tuple with the InvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceID

`func (o *MtdInvoiceResponse) SetInvoiceID(v string)`

SetInvoiceID sets InvoiceID field to given value.

### HasInvoiceID

`func (o *MtdInvoiceResponse) HasInvoiceID() bool`

HasInvoiceID returns a boolean if a field has been set.

### GetBillingStartDate

`func (o *MtdInvoiceResponse) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *MtdInvoiceResponse) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *MtdInvoiceResponse) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.

### HasBillingStartDate

`func (o *MtdInvoiceResponse) HasBillingStartDate() bool`

HasBillingStartDate returns a boolean if a field has been set.

### GetBillingEndDate

`func (o *MtdInvoiceResponse) GetBillingEndDate() time.Time`

GetBillingEndDate returns the BillingEndDate field if non-nil, zero value otherwise.

### GetBillingEndDateOk

`func (o *MtdInvoiceResponse) GetBillingEndDateOk() (*time.Time, bool)`

GetBillingEndDateOk returns a tuple with the BillingEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndDate

`func (o *MtdInvoiceResponse) SetBillingEndDate(v time.Time)`

SetBillingEndDate sets BillingEndDate field to given value.

### HasBillingEndDate

`func (o *MtdInvoiceResponse) HasBillingEndDate() bool`

HasBillingEndDate returns a boolean if a field has been set.

### GetMonthlyTransactionAmount

`func (o *MtdInvoiceResponse) GetMonthlyTransactionAmount() string`

GetMonthlyTransactionAmount returns the MonthlyTransactionAmount field if non-nil, zero value otherwise.

### GetMonthlyTransactionAmountOk

`func (o *MtdInvoiceResponse) GetMonthlyTransactionAmountOk() (*string, bool)`

GetMonthlyTransactionAmountOk returns a tuple with the MonthlyTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyTransactionAmount

`func (o *MtdInvoiceResponse) SetMonthlyTransactionAmount(v string)`

SetMonthlyTransactionAmount sets MonthlyTransactionAmount field to given value.

### HasMonthlyTransactionAmount

`func (o *MtdInvoiceResponse) HasMonthlyTransactionAmount() bool`

HasMonthlyTransactionAmount returns a boolean if a field has been set.

### GetTransactionLineItems

`func (o *MtdInvoiceResponse) GetTransactionLineItems() []Mtdlineitems`

GetTransactionLineItems returns the TransactionLineItems field if non-nil, zero value otherwise.

### GetTransactionLineItemsOk

`func (o *MtdInvoiceResponse) GetTransactionLineItemsOk() (*[]Mtdlineitems, bool)`

GetTransactionLineItemsOk returns a tuple with the TransactionLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLineItems

`func (o *MtdInvoiceResponse) SetTransactionLineItems(v []Mtdlineitems)`

SetTransactionLineItems sets TransactionLineItems field to given value.

### HasTransactionLineItems

`func (o *MtdInvoiceResponse) HasTransactionLineItems() bool`

HasTransactionLineItems returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
