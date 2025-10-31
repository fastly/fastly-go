# ComputeAclListEntriesMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The maximum number of results shown in this response. | [optional] 
**NextCursor** | Pointer to **string** | Used for pagination, supply to the next request to get the next block of results. | [optional] 

## Methods

### NewComputeAclListEntriesMeta

`func NewComputeAclListEntriesMeta() *ComputeAclListEntriesMeta`

NewComputeAclListEntriesMeta instantiates a new ComputeAclListEntriesMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeAclListEntriesMetaWithDefaults

`func NewComputeAclListEntriesMetaWithDefaults() *ComputeAclListEntriesMeta`

NewComputeAclListEntriesMetaWithDefaults instantiates a new ComputeAclListEntriesMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ComputeAclListEntriesMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ComputeAclListEntriesMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ComputeAclListEntriesMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ComputeAclListEntriesMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextCursor

`func (o *ComputeAclListEntriesMeta) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ComputeAclListEntriesMeta) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ComputeAclListEntriesMeta) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ComputeAclListEntriesMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


