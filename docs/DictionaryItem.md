# DictionaryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemKey** | Pointer to **string** | Item key, maximum 256 characters. | [optional] 
**ItemValue** | Pointer to **string** | Item value, maximum 8000 characters. | [optional] 

## Methods

### NewDictionaryItem

`func NewDictionaryItem() *DictionaryItem`

NewDictionaryItem instantiates a new DictionaryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDictionaryItemWithDefaults

`func NewDictionaryItemWithDefaults() *DictionaryItem`

NewDictionaryItemWithDefaults instantiates a new DictionaryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemKey

`func (o *DictionaryItem) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *DictionaryItem) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *DictionaryItem) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *DictionaryItem) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### GetItemValue

`func (o *DictionaryItem) GetItemValue() string`

GetItemValue returns the ItemValue field if non-nil, zero value otherwise.

### GetItemValueOk

`func (o *DictionaryItem) GetItemValueOk() (*string, bool)`

GetItemValueOk returns a tuple with the ItemValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemValue

`func (o *DictionaryItem) SetItemValue(v string)`

SetItemValue sets ItemValue field to given value.

### HasItemValue

`func (o *DictionaryItem) HasItemValue() bool`

HasItemValue returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
