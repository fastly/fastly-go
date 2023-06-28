# InlineResponse2003Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **string** | Cursor for the next page. | [optional] 
**Limit** | Pointer to **int32** | Entries returned. | [optional] 

## Methods

### NewInlineResponse2003Meta

`func NewInlineResponse2003Meta() *InlineResponse2003Meta`

NewInlineResponse2003Meta instantiates a new InlineResponse2003Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003MetaWithDefaults

`func NewInlineResponse2003MetaWithDefaults() *InlineResponse2003Meta`

NewInlineResponse2003MetaWithDefaults instantiates a new InlineResponse2003Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *InlineResponse2003Meta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *InlineResponse2003Meta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *InlineResponse2003Meta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *InlineResponse2003Meta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetLimit

`func (o *InlineResponse2003Meta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineResponse2003Meta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineResponse2003Meta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineResponse2003Meta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
