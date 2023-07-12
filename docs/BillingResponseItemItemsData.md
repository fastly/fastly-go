# BillingResponseItemItemsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItems** | Pointer to [**[]BillingResponseLineItem**](BillingResponseLineItem.md) |  | [optional] 

## Methods

### NewBillingResponseItemItemsData

`func NewBillingResponseItemItemsData() *BillingResponseItemItemsData`

NewBillingResponseItemItemsData instantiates a new BillingResponseItemItemsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingResponseItemItemsDataWithDefaults

`func NewBillingResponseItemItemsDataWithDefaults() *BillingResponseItemItemsData`

NewBillingResponseItemItemsDataWithDefaults instantiates a new BillingResponseItemItemsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineItems

`func (o *BillingResponseItemItemsData) GetLineItems() []BillingResponseLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *BillingResponseItemItemsData) GetLineItemsOk() (*[]BillingResponseLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *BillingResponseItemItemsData) SetLineItems(v []BillingResponseLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *BillingResponseItemItemsData) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
