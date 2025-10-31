# ConfigStoreInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemCount** | Pointer to **int32** | The number of items currently in the config store. | [optional] 

## Methods

### NewConfigStoreInfoResponse

`func NewConfigStoreInfoResponse() *ConfigStoreInfoResponse`

NewConfigStoreInfoResponse instantiates a new ConfigStoreInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigStoreInfoResponseWithDefaults

`func NewConfigStoreInfoResponseWithDefaults() *ConfigStoreInfoResponse`

NewConfigStoreInfoResponseWithDefaults instantiates a new ConfigStoreInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemCount

`func (o *ConfigStoreInfoResponse) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ConfigStoreInfoResponse) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ConfigStoreInfoResponse) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *ConfigStoreInfoResponse) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


