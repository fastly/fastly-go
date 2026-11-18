# BotRequestMetricsMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **time.Time** | An RFC-8339-formatted date and time string indicating the inclusive start of the query time range. | [optional] 
**To** | Pointer to **time.Time** | An RFC-8339-formatted date and time indicating the exclusive end of the query time range. | [optional] 
**GroupBy** | Pointer to **string** | List of dimensions (comma-separated) being summed over in the results. | [optional] 
**NextCursor** | Pointer to **string** | Optional string that can be used to request the next page of results, if any. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of results to return. Defaults to 1000. | [optional] 

## Methods

### NewBotRequestMetricsMetadata

`func NewBotRequestMetricsMetadata() *BotRequestMetricsMetadata`

NewBotRequestMetricsMetadata instantiates a new BotRequestMetricsMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotRequestMetricsMetadataWithDefaults

`func NewBotRequestMetricsMetadataWithDefaults() *BotRequestMetricsMetadata`

NewBotRequestMetricsMetadataWithDefaults instantiates a new BotRequestMetricsMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *BotRequestMetricsMetadata) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BotRequestMetricsMetadata) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BotRequestMetricsMetadata) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BotRequestMetricsMetadata) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *BotRequestMetricsMetadata) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BotRequestMetricsMetadata) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BotRequestMetricsMetadata) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *BotRequestMetricsMetadata) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetGroupBy

`func (o *BotRequestMetricsMetadata) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *BotRequestMetricsMetadata) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *BotRequestMetricsMetadata) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *BotRequestMetricsMetadata) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetNextCursor

`func (o *BotRequestMetricsMetadata) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *BotRequestMetricsMetadata) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *BotRequestMetricsMetadata) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *BotRequestMetricsMetadata) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetLimit

`func (o *BotRequestMetricsMetadata) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BotRequestMetricsMetadata) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BotRequestMetricsMetadata) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BotRequestMetricsMetadata) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
