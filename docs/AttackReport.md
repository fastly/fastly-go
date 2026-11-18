# AttackReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | ID of the workspace. | 
**Name** | **string** | Name of the workspace. | 
**TotalCount** | **int32** | Total request count | 
**BlockedCount** | **int32** | Blocked request count | 
**FlaggedCount** | **int32** | Flagged request count | 
**AttackCount** | **int32** | Attack request count | 
**AllFlaggedIPCount** | **int32** | Count of IPs that have been flagged | 
**FlaggedIPCount** | **int32** | Count of currently flagged IPs | 
**TopAttackSignals** | [**[]AttackSignal**](AttackSignal.md) |  | 
**TopAttackSources** | [**[]AttackSource**](AttackSource.md) |  | 

## Methods

### NewAttackReport

`func NewAttackReport(id string, name string, totalCount int32, blockedCount int32, flaggedCount int32, attackCount int32, allFlaggedIPCount int32, flaggedIPCount int32, topAttackSignals []AttackSignal, topAttackSources []AttackSource, ) *AttackReport`

NewAttackReport instantiates a new AttackReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackReportWithDefaults

`func NewAttackReportWithDefaults() *AttackReport`

NewAttackReportWithDefaults instantiates a new AttackReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *AttackReport) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *AttackReport) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *AttackReport) SetID(v string)`

SetID sets ID field to given value.


### GetName

`func (o *AttackReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttackReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttackReport) SetName(v string)`

SetName sets Name field to given value.


### GetTotalCount

`func (o *AttackReport) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AttackReport) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AttackReport) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetBlockedCount

`func (o *AttackReport) GetBlockedCount() int32`

GetBlockedCount returns the BlockedCount field if non-nil, zero value otherwise.

### GetBlockedCountOk

`func (o *AttackReport) GetBlockedCountOk() (*int32, bool)`

GetBlockedCountOk returns a tuple with the BlockedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedCount

`func (o *AttackReport) SetBlockedCount(v int32)`

SetBlockedCount sets BlockedCount field to given value.


### GetFlaggedCount

`func (o *AttackReport) GetFlaggedCount() int32`

GetFlaggedCount returns the FlaggedCount field if non-nil, zero value otherwise.

### GetFlaggedCountOk

`func (o *AttackReport) GetFlaggedCountOk() (*int32, bool)`

GetFlaggedCountOk returns a tuple with the FlaggedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedCount

`func (o *AttackReport) SetFlaggedCount(v int32)`

SetFlaggedCount sets FlaggedCount field to given value.


### GetAttackCount

`func (o *AttackReport) GetAttackCount() int32`

GetAttackCount returns the AttackCount field if non-nil, zero value otherwise.

### GetAttackCountOk

`func (o *AttackReport) GetAttackCountOk() (*int32, bool)`

GetAttackCountOk returns a tuple with the AttackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackCount

`func (o *AttackReport) SetAttackCount(v int32)`

SetAttackCount sets AttackCount field to given value.


### GetAllFlaggedIPCount

`func (o *AttackReport) GetAllFlaggedIPCount() int32`

GetAllFlaggedIPCount returns the AllFlaggedIPCount field if non-nil, zero value otherwise.

### GetAllFlaggedIPCountOk

`func (o *AttackReport) GetAllFlaggedIPCountOk() (*int32, bool)`

GetAllFlaggedIPCountOk returns a tuple with the AllFlaggedIPCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFlaggedIPCount

`func (o *AttackReport) SetAllFlaggedIPCount(v int32)`

SetAllFlaggedIPCount sets AllFlaggedIPCount field to given value.


### GetFlaggedIPCount

`func (o *AttackReport) GetFlaggedIPCount() int32`

GetFlaggedIPCount returns the FlaggedIPCount field if non-nil, zero value otherwise.

### GetFlaggedIPCountOk

`func (o *AttackReport) GetFlaggedIPCountOk() (*int32, bool)`

GetFlaggedIPCountOk returns a tuple with the FlaggedIPCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedIPCount

`func (o *AttackReport) SetFlaggedIPCount(v int32)`

SetFlaggedIPCount sets FlaggedIPCount field to given value.


### GetTopAttackSignals

`func (o *AttackReport) GetTopAttackSignals() []AttackSignal`

GetTopAttackSignals returns the TopAttackSignals field if non-nil, zero value otherwise.

### GetTopAttackSignalsOk

`func (o *AttackReport) GetTopAttackSignalsOk() (*[]AttackSignal, bool)`

GetTopAttackSignalsOk returns a tuple with the TopAttackSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAttackSignals

`func (o *AttackReport) SetTopAttackSignals(v []AttackSignal)`

SetTopAttackSignals sets TopAttackSignals field to given value.


### GetTopAttackSources

`func (o *AttackReport) GetTopAttackSources() []AttackSource`

GetTopAttackSources returns the TopAttackSources field if non-nil, zero value otherwise.

### GetTopAttackSourcesOk

`func (o *AttackReport) GetTopAttackSourcesOk() (*[]AttackSource, bool)`

GetTopAttackSourcesOk returns a tuple with the TopAttackSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAttackSources

`func (o *AttackReport) SetTopAttackSources(v []AttackSource)`

SetTopAttackSources sets TopAttackSources field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
