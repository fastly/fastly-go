# NetworkProtectListMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Effective page limit after defaulting and clamping. | [default to 100]
**NextCursor** | Pointer to **string** | Continuation value to send as cursor when more resources remain. Omitted on the final page. | [optional] 

## Methods

### NewNetworkProtectListMeta

`func NewNetworkProtectListMeta(limit int32, ) *NetworkProtectListMeta`

NewNetworkProtectListMeta instantiates a new NetworkProtectListMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkProtectListMetaWithDefaults

`func NewNetworkProtectListMetaWithDefaults() *NetworkProtectListMeta`

NewNetworkProtectListMetaWithDefaults instantiates a new NetworkProtectListMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *NetworkProtectListMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NetworkProtectListMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NetworkProtectListMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNextCursor

`func (o *NetworkProtectListMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *NetworkProtectListMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *NetworkProtectListMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *NetworkProtectListMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
