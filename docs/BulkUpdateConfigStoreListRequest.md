# BulkUpdateConfigStoreListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]BulkUpdateConfigStoreItem**](BulkUpdateConfigStoreItem.md) |  | [optional] 

## Methods

### NewBulkUpdateConfigStoreListRequest

`func NewBulkUpdateConfigStoreListRequest() *BulkUpdateConfigStoreListRequest`

NewBulkUpdateConfigStoreListRequest instantiates a new BulkUpdateConfigStoreListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateConfigStoreListRequestWithDefaults

`func NewBulkUpdateConfigStoreListRequestWithDefaults() *BulkUpdateConfigStoreListRequest`

NewBulkUpdateConfigStoreListRequestWithDefaults instantiates a new BulkUpdateConfigStoreListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BulkUpdateConfigStoreListRequest) GetItems() []BulkUpdateConfigStoreItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkUpdateConfigStoreListRequest) GetItemsOk() (*[]BulkUpdateConfigStoreItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkUpdateConfigStoreListRequest) SetItems(v []BulkUpdateConfigStoreItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *BulkUpdateConfigStoreListRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
