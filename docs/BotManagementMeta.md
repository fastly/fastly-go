# BotManagementMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The `limit` value used when making the request. | [optional] 
**Total** | Pointer to **int32** | The count of results matching the request. | [optional] 
**NextCursor** | Pointer to **NullableString** | The cursor to use to request the next page of results. Set to `null` when there are no more results. | [optional] 

## Methods

### NewBotManagementMeta

`func NewBotManagementMeta() *BotManagementMeta`

NewBotManagementMeta instantiates a new BotManagementMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotManagementMetaWithDefaults

`func NewBotManagementMetaWithDefaults() *BotManagementMeta`

NewBotManagementMetaWithDefaults instantiates a new BotManagementMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *BotManagementMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BotManagementMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BotManagementMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BotManagementMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *BotManagementMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BotManagementMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BotManagementMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BotManagementMeta) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetNextCursor

`func (o *BotManagementMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *BotManagementMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *BotManagementMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *BotManagementMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *BotManagementMeta) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *BotManagementMeta) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


