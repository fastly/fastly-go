# DiscoveredOperationGetExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The current status of the operation. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp when the operation was last updated. | [optional] [readonly] 
**LastSeenAt** | Pointer to **time.Time** | The timestamp when the operation was last seen in traffic. | [optional] [readonly] 

## Methods

### NewDiscoveredOperationGetExtra

`func NewDiscoveredOperationGetExtra() *DiscoveredOperationGetExtra`

NewDiscoveredOperationGetExtra instantiates a new DiscoveredOperationGetExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveredOperationGetExtraWithDefaults

`func NewDiscoveredOperationGetExtraWithDefaults() *DiscoveredOperationGetExtra`

NewDiscoveredOperationGetExtraWithDefaults instantiates a new DiscoveredOperationGetExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DiscoveredOperationGetExtra) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscoveredOperationGetExtra) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscoveredOperationGetExtra) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiscoveredOperationGetExtra) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DiscoveredOperationGetExtra) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DiscoveredOperationGetExtra) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DiscoveredOperationGetExtra) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DiscoveredOperationGetExtra) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *DiscoveredOperationGetExtra) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *DiscoveredOperationGetExtra) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *DiscoveredOperationGetExtra) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *DiscoveredOperationGetExtra) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


