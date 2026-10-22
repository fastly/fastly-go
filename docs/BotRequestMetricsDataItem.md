# BotRequestMetricsDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Timestamp for the time series being reported. | [optional] 
**Bot** | Pointer to **string** | Bot identifier for which data was recorded. Present only when grouped by bot. | [optional] 
**Datacenter** | Pointer to **string** | Datacenter (POP) for which data was recorded. Present only when grouped by datacenter. | [optional] 
**BotEdgeRequestsAllowedCount** | Pointer to **int64** | Count of edge requests where a bot was detected and the request was allowed by the configured policy. | [optional] 
**BotEdgeRequestsChallengedCount** | Pointer to **int64** | Count of edge requests where a bot was detected and a client challenge was issued for the request by the configured policy. | [optional] 
**BotEdgeRequestsBlockedCount** | Pointer to **int64** | Count of edge requests where a bot was detected and the request was blocked by the configured policy. | [optional] 
**BotEdgeRequestsChallengeTokenValidCount** | Pointer to **int64** | Count of edge requests where a bot was detected and the client making the request has already solved a challenge issued by the configured policy. | [optional] 

## Methods

### NewBotRequestMetricsDataItem

`func NewBotRequestMetricsDataItem() *BotRequestMetricsDataItem`

NewBotRequestMetricsDataItem instantiates a new BotRequestMetricsDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotRequestMetricsDataItemWithDefaults

`func NewBotRequestMetricsDataItemWithDefaults() *BotRequestMetricsDataItem`

NewBotRequestMetricsDataItemWithDefaults instantiates a new BotRequestMetricsDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *BotRequestMetricsDataItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BotRequestMetricsDataItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BotRequestMetricsDataItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BotRequestMetricsDataItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBot

`func (o *BotRequestMetricsDataItem) GetBot() string`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *BotRequestMetricsDataItem) GetBotOk() (*string, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *BotRequestMetricsDataItem) SetBot(v string)`

SetBot sets Bot field to given value.

### HasBot

`func (o *BotRequestMetricsDataItem) HasBot() bool`

HasBot returns a boolean if a field has been set.

### GetDatacenter

`func (o *BotRequestMetricsDataItem) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *BotRequestMetricsDataItem) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *BotRequestMetricsDataItem) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *BotRequestMetricsDataItem) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetBotEdgeRequestsAllowedCount

`func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsAllowedCount() int64`

GetBotEdgeRequestsAllowedCount returns the BotEdgeRequestsAllowedCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsAllowedCountOk

`func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsAllowedCountOk() (*int64, bool)`

GetBotEdgeRequestsAllowedCountOk returns a tuple with the BotEdgeRequestsAllowedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsAllowedCount

`func (o *BotRequestMetricsDataItem) SetBotEdgeRequestsAllowedCount(v int64)`

SetBotEdgeRequestsAllowedCount sets BotEdgeRequestsAllowedCount field to given value.

### HasBotEdgeRequestsAllowedCount

`func (o *BotRequestMetricsDataItem) HasBotEdgeRequestsAllowedCount() bool`

HasBotEdgeRequestsAllowedCount returns a boolean if a field has been set.

### GetBotEdgeRequestsChallengedCount

`func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsChallengedCount() int64`

GetBotEdgeRequestsChallengedCount returns the BotEdgeRequestsChallengedCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsChallengedCountOk

`func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsChallengedCountOk() (*int64, bool)`

GetBotEdgeRequestsChallengedCountOk returns a tuple with the BotEdgeRequestsChallengedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsChallengedCount

`func (o *BotRequestMetricsDataItem) SetBotEdgeRequestsChallengedCount(v int64)`

SetBotEdgeRequestsChallengedCount sets BotEdgeRequestsChallengedCount field to given value.

### HasBotEdgeRequestsChallengedCount

`func (o *BotRequestMetricsDataItem) HasBotEdgeRequestsChallengedCount() bool`

HasBotEdgeRequestsChallengedCount returns a boolean if a field has been set.

### GetBotEdgeRequestsBlockedCount

`func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsBlockedCount() int64`

GetBotEdgeRequestsBlockedCount returns the BotEdgeRequestsBlockedCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsBlockedCountOk

`func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsBlockedCountOk() (*int64, bool)`

GetBotEdgeRequestsBlockedCountOk returns a tuple with the BotEdgeRequestsBlockedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsBlockedCount

`func (o *BotRequestMetricsDataItem) SetBotEdgeRequestsBlockedCount(v int64)`

SetBotEdgeRequestsBlockedCount sets BotEdgeRequestsBlockedCount field to given value.

### HasBotEdgeRequestsBlockedCount

`func (o *BotRequestMetricsDataItem) HasBotEdgeRequestsBlockedCount() bool`

HasBotEdgeRequestsBlockedCount returns a boolean if a field has been set.

### GetBotEdgeRequestsChallengeTokenValidCount

`func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsChallengeTokenValidCount() int64`

GetBotEdgeRequestsChallengeTokenValidCount returns the BotEdgeRequestsChallengeTokenValidCount field if non-nil, zero value otherwise.

### GetBotEdgeRequestsChallengeTokenValidCountOk

`func (o *BotRequestMetricsDataItem) GetBotEdgeRequestsChallengeTokenValidCountOk() (*int64, bool)`

GetBotEdgeRequestsChallengeTokenValidCountOk returns a tuple with the BotEdgeRequestsChallengeTokenValidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotEdgeRequestsChallengeTokenValidCount

`func (o *BotRequestMetricsDataItem) SetBotEdgeRequestsChallengeTokenValidCount(v int64)`

SetBotEdgeRequestsChallengeTokenValidCount sets BotEdgeRequestsChallengeTokenValidCount field to given value.

### HasBotEdgeRequestsChallengeTokenValidCount

`func (o *BotRequestMetricsDataItem) HasBotEdgeRequestsChallengeTokenValidCount() bool`

HasBotEdgeRequestsChallengeTokenValidCount returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


