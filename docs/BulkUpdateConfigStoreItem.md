# BulkUpdateConfigStoreItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemKey** | Pointer to **string** | Item key, maximum 256 characters. | [optional] 
**ItemValue** | Pointer to **string** | Item value, maximum 8000 characters. | [optional] 
**Op** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkUpdateConfigStoreItem

`func NewBulkUpdateConfigStoreItem() *BulkUpdateConfigStoreItem`

NewBulkUpdateConfigStoreItem instantiates a new BulkUpdateConfigStoreItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateConfigStoreItemWithDefaults

`func NewBulkUpdateConfigStoreItemWithDefaults() *BulkUpdateConfigStoreItem`

NewBulkUpdateConfigStoreItemWithDefaults instantiates a new BulkUpdateConfigStoreItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemKey

`func (o *BulkUpdateConfigStoreItem) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *BulkUpdateConfigStoreItem) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *BulkUpdateConfigStoreItem) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *BulkUpdateConfigStoreItem) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### GetItemValue

`func (o *BulkUpdateConfigStoreItem) GetItemValue() string`

GetItemValue returns the ItemValue field if non-nil, zero value otherwise.

### GetItemValueOk

`func (o *BulkUpdateConfigStoreItem) GetItemValueOk() (*string, bool)`

GetItemValueOk returns a tuple with the ItemValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemValue

`func (o *BulkUpdateConfigStoreItem) SetItemValue(v string)`

SetItemValue sets ItemValue field to given value.

### HasItemValue

`func (o *BulkUpdateConfigStoreItem) HasItemValue() bool`

HasItemValue returns a boolean if a field has been set.

### GetOp

`func (o *BulkUpdateConfigStoreItem) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *BulkUpdateConfigStoreItem) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *BulkUpdateConfigStoreItem) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *BulkUpdateConfigStoreItem) HasOp() bool`

HasOp returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


