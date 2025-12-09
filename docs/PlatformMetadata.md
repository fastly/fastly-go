# PlatformMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** | An RFC-8339-formatted date and time indicating the inclusive start of the query time range. | [optional] 
**To** | Pointer to **time.Time** | An RFC-8339-formatted date and time indicating the exclusive end of the query time range. | [optional] 
**NextCursor** | Pointer to **string** | A string that can be used to request the next page of results, if any. | [optional] 
**GroupBy** | Pointer to **string** | A comma-separated list of fields used to group and order the results. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of results to return. | [optional] 

## Methods

### NewPlatformMetadata

`func NewPlatformMetadata() *PlatformMetadata`

NewPlatformMetadata instantiates a new PlatformMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformMetadataWithDefaults

`func NewPlatformMetadataWithDefaults() *PlatformMetadata`

NewPlatformMetadataWithDefaults instantiates a new PlatformMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *PlatformMetadata) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PlatformMetadata) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PlatformMetadata) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PlatformMetadata) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *PlatformMetadata) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PlatformMetadata) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PlatformMetadata) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *PlatformMetadata) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetNextCursor

`func (o *PlatformMetadata) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PlatformMetadata) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PlatformMetadata) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PlatformMetadata) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetGroupBy

`func (o *PlatformMetadata) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *PlatformMetadata) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *PlatformMetadata) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *PlatformMetadata) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetLimit

`func (o *PlatformMetadata) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PlatformMetadata) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PlatformMetadata) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PlatformMetadata) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


