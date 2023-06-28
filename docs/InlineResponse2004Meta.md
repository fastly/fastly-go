# InlineResponse2004Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewInlineResponse2004Meta

`func NewInlineResponse2004Meta() *InlineResponse2004Meta`

NewInlineResponse2004Meta instantiates a new InlineResponse2004Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2004MetaWithDefaults

`func NewInlineResponse2004MetaWithDefaults() *InlineResponse2004Meta`

NewInlineResponse2004MetaWithDefaults instantiates a new InlineResponse2004Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *InlineResponse2004Meta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *InlineResponse2004Meta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *InlineResponse2004Meta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *InlineResponse2004Meta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetLimit

`func (o *InlineResponse2004Meta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineResponse2004Meta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineResponse2004Meta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineResponse2004Meta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
