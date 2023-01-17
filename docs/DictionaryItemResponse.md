# DictionaryItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemKey** | Pointer to **string** | Item key, maximum 256 characters. | [optional] 
**ItemValue** | Pointer to **string** | Item value, maximum 8000 characters. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DeletedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**UpdatedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**DictionaryID** | Pointer to **string** |  | [optional] [readonly] 
**ServiceID** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDictionaryItemResponse

`func NewDictionaryItemResponse() *DictionaryItemResponse`

NewDictionaryItemResponse instantiates a new DictionaryItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDictionaryItemResponseWithDefaults

`func NewDictionaryItemResponseWithDefaults() *DictionaryItemResponse`

NewDictionaryItemResponseWithDefaults instantiates a new DictionaryItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemKey

`func (o *DictionaryItemResponse) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *DictionaryItemResponse) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *DictionaryItemResponse) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *DictionaryItemResponse) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### GetItemValue

`func (o *DictionaryItemResponse) GetItemValue() string`

GetItemValue returns the ItemValue field if non-nil, zero value otherwise.

### GetItemValueOk

`func (o *DictionaryItemResponse) GetItemValueOk() (*string, bool)`

GetItemValueOk returns a tuple with the ItemValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemValue

`func (o *DictionaryItemResponse) SetItemValue(v string)`

SetItemValue sets ItemValue field to given value.

### HasItemValue

`func (o *DictionaryItemResponse) HasItemValue() bool`

HasItemValue returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DictionaryItemResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DictionaryItemResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DictionaryItemResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DictionaryItemResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DictionaryItemResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DictionaryItemResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDeletedAt

`func (o *DictionaryItemResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DictionaryItemResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DictionaryItemResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DictionaryItemResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *DictionaryItemResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DictionaryItemResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetUpdatedAt

`func (o *DictionaryItemResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DictionaryItemResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DictionaryItemResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DictionaryItemResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *DictionaryItemResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DictionaryItemResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDictionaryID

`func (o *DictionaryItemResponse) GetDictionaryID() string`

GetDictionaryID returns the DictionaryID field if non-nil, zero value otherwise.

### GetDictionaryIDOk

`func (o *DictionaryItemResponse) GetDictionaryIDOk() (*string, bool)`

GetDictionaryIDOk returns a tuple with the DictionaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaryID

`func (o *DictionaryItemResponse) SetDictionaryID(v string)`

SetDictionaryID sets DictionaryID field to given value.

### HasDictionaryID

`func (o *DictionaryItemResponse) HasDictionaryID() bool`

HasDictionaryID returns a boolean if a field has been set.

### GetServiceID

`func (o *DictionaryItemResponse) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *DictionaryItemResponse) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *DictionaryItemResponse) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *DictionaryItemResponse) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
