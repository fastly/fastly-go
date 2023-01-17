# DictionaryInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdated** | Pointer to **string** | Timestamp (UTC) when the dictionary was last updated or an item was added or removed. | [optional] 
**ItemCount** | Pointer to **int32** | The number of items currently in the dictionary. | [optional] 
**Digest** | Pointer to **string** | A hash of all the dictionary content. | [optional] 

## Methods

### NewDictionaryInfoResponse

`func NewDictionaryInfoResponse() *DictionaryInfoResponse`

NewDictionaryInfoResponse instantiates a new DictionaryInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDictionaryInfoResponseWithDefaults

`func NewDictionaryInfoResponseWithDefaults() *DictionaryInfoResponse`

NewDictionaryInfoResponseWithDefaults instantiates a new DictionaryInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdated

`func (o *DictionaryInfoResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DictionaryInfoResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DictionaryInfoResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DictionaryInfoResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetItemCount

`func (o *DictionaryInfoResponse) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *DictionaryInfoResponse) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *DictionaryInfoResponse) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *DictionaryInfoResponse) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetDigest

`func (o *DictionaryInfoResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DictionaryInfoResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DictionaryInfoResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *DictionaryInfoResponse) HasDigest() bool`

HasDigest returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
