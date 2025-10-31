# BulkUpdateDictionaryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemKey** | Pointer to **string** | Item key, maximum 256 characters. | [optional] 
**ItemValue** | Pointer to **string** | Item value, maximum 8000 characters. | [optional] 
**Op** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkUpdateDictionaryItem

`func NewBulkUpdateDictionaryItem() *BulkUpdateDictionaryItem`

NewBulkUpdateDictionaryItem instantiates a new BulkUpdateDictionaryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateDictionaryItemWithDefaults

`func NewBulkUpdateDictionaryItemWithDefaults() *BulkUpdateDictionaryItem`

NewBulkUpdateDictionaryItemWithDefaults instantiates a new BulkUpdateDictionaryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemKey

`func (o *BulkUpdateDictionaryItem) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *BulkUpdateDictionaryItem) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *BulkUpdateDictionaryItem) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *BulkUpdateDictionaryItem) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### GetItemValue

`func (o *BulkUpdateDictionaryItem) GetItemValue() string`

GetItemValue returns the ItemValue field if non-nil, zero value otherwise.

### GetItemValueOk

`func (o *BulkUpdateDictionaryItem) GetItemValueOk() (*string, bool)`

GetItemValueOk returns a tuple with the ItemValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemValue

`func (o *BulkUpdateDictionaryItem) SetItemValue(v string)`

SetItemValue sets ItemValue field to given value.

### HasItemValue

`func (o *BulkUpdateDictionaryItem) HasItemValue() bool`

HasItemValue returns a boolean if a field has been set.

### GetOp

`func (o *BulkUpdateDictionaryItem) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *BulkUpdateDictionaryItem) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *BulkUpdateDictionaryItem) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *BulkUpdateDictionaryItem) HasOp() bool`

HasOp returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


